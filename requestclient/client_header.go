package requestclient

func (rq *RequestClient) GetHeader(key string) string {
	return rq.defaultHeaders[key]
}

func (rq *RequestClient) SetHeader(key string, value string) {
	rq.defaultHeaders[key] = value
}
