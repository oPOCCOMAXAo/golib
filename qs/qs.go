package qs

import (
	"net/url"
)

func Stringify(data map[string]string) string {
	values := url.Values{}
	for key, value := range data {
		values.Add(key, value)
	}
	return values.Encode()
}
