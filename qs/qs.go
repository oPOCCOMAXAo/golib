package qs

import "strings"

func Stringify(data map[string]string) string {
	t := make([]string, 0)
	for key, value := range data {
		t = append(t, key+"="+value)
	}
	return strings.Join(t, "&")
}
