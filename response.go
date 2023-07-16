package httphelper

import "net/http"

type HttpHelper struct {
	status int
	header *http.Header
}

const Hello = "Hello there"

func Status(status int) *HttpHelper {
	return &HttpHelper{
		status: status,
		header: &http.Header{},
	}
}
