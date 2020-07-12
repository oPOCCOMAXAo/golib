package requestclient

import (
	"net/http"
	"net/url"
)

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
