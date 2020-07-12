package requestclient

import (
	"context"
	"github.com/opoccomaxao/golib/console"
	"golang.org/x/sync/semaphore"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"
)

type RequestClient struct {
	client         http.Client
	semaphore      *semaphore.Weighted
	defaultHeaders map[string]string
}

func New(parallel int64) *RequestClient {
	if parallel < 1 {
		parallel = 1
	}
	jar, _ := cookiejar.New(nil)
	client := http.Client{Jar: jar}
	sem := semaphore.NewWeighted(parallel)
	return &RequestClient{client, sem, make(map[string]string)}
}

func (rq *RequestClient) leave(timeout time.Duration) {
	time.Sleep(timeout)
	rq.semaphore.Release(1)
}

func (rq *RequestClient) goLeave(timeout time.Duration) {
	go rq.leave(timeout)
}

func (rq *RequestClient) Get(url string, timeout time.Duration, headers map[string]string) Response {
	r, _ := http.NewRequest("GET", url, nil)
	return rq.do(r, timeout, headers)
}

func (rq *RequestClient) Post(url string, data string, timeout time.Duration, headers map[string]string) Response {
	r, _ := http.NewRequest("POST", url, strings.NewReader(data))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return rq.do(r, timeout, headers)
}

func (rq *RequestClient) Put(url string, data string, timeout time.Duration, headers map[string]string) Response {
	r, _ := http.NewRequest("PUT", url, strings.NewReader(data))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return rq.do(r, timeout, headers)
}

func (rq *RequestClient) Delete(url string, timeout time.Duration, headers map[string]string) Response {
	r, _ := http.NewRequest("DELETE", url, nil)
	return rq.do(r, timeout, headers)
}

func (rq *RequestClient) do(req *http.Request, timeout time.Duration, headers map[string]string) (response Response) {
	defer rq.goLeave(timeout)

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	for k, v := range rq.defaultHeaders {
		if _, prs := req.Header[k]; !prs {
			req.Header.Set(k, v)
		}
	}

	if err := rq.semaphore.Acquire(context.Background(), 1); err != nil {
		response.Status = -1
		response.Error = err
		console.Error("%s", err)
		return
	}

	res, err := rq.client.Do(req)
	if err != nil {
		response.Status = -1
		response.Error = err
		console.Error("%s", err)
		return
	}

	response.Status = res.StatusCode
	var buffer []byte
	if buffer, err = ioutil.ReadAll(res.Body); err != nil {
		response.Status = 0
		response.Error = err
		console.Error("%s", err)
		return
	}

	response.Response = buffer
	response.Headers = res.Header
	return
}
