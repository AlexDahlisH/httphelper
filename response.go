package httphelper

import "net/http"

type HttpHelper struct {
	status int
	header *http.Header
}

func Status(status int) *HttpHelper {
	return &HttpHelper{
		status: status,
		header: &http.Header{},
	}
}
