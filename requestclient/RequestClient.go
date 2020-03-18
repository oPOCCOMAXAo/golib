package requestclient

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

type RequestClient struct {
	client         http.Client
	semaphore      *semaphore.Weighted
	defaultHeaders map[string]string
}

func NewRequestClient(parallel int64) *RequestClient {
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

func (rq *RequestClient) Get(url string, timeout time.Duration, headers map[string]string) (result string, status int, outHeaders http.Header) {
	r, _ := http.NewRequest("GET", url, nil)
	return rq.do(r, timeout, headers)
}

func (rq *RequestClient) Post(url string, data string, timeout time.Duration, headers map[string]string) (result string, status int, outHeaders http.Header) {
	r, _ := http.NewRequest("POST", url, strings.NewReader(data))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return rq.do(r, timeout, headers)
}

func (rq *RequestClient) Delete(url string, timeout time.Duration, headers map[string]string) (result string, status int, outHeaders http.Header) {
	r, _ := http.NewRequest("DELETE", url, nil)
	return rq.do(r, timeout, headers)
}

func (rq *RequestClient) printErr(err error) {
	fmt.Printf("%#v\n", err)
}

func (rq *RequestClient) do(req *http.Request, timeout time.Duration, headers map[string]string) (result string, status int, outHeaders http.Header) {
	defer rq.goLeave(timeout)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	for k, v := range rq.defaultHeaders {
		if _, prs := req.Header[k]; prs {
			req.Header.Set(k, v)
		}
	}
	if err := rq.semaphore.Acquire(context.Background(), 1); err != nil {
		status = -1
		rq.printErr(err)
		return
	}
	res, err := rq.client.Do(req)
	if err != nil {
		status = -1
		rq.printErr(err)
		return
	}
	status = res.StatusCode
	var buffer []byte
	if buffer, err = ioutil.ReadAll(res.Body); err != nil {
		status = 0
		rq.printErr(err)
		return
	}
	result = string(buffer)
	outHeaders = res.Header
	return
}

func (rq *RequestClient) GetCookie(host string, name string) string {
	a := rq.client.Jar.Cookies(&url.URL{Scheme: "http", Host: host, Path: "/"})
	for _, c := range a {
		if c.Name == name {
			return c.Value
		}
	}
	return ""
}

func (rq *RequestClient) SetCookie(host string, name string, value string) {
	rq.client.Jar.SetCookies(&url.URL{Scheme: "http", Host: host, Path: "/"}, []*http.Cookie{{Name: name, Value: value}})
}

func (rq *RequestClient) GetHeader(key string) string {
	return rq.defaultHeaders[key]
}

func (rq *RequestClient) SetHeader(key string, value string) {
	rq.defaultHeaders[key] = value
}
