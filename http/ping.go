package http

import (
	go_http "net/http"
)

func PingHandler() (go_http.Handler, error) {

	fn := func(rsp go_http.ResponseWriter, req *go_http.Request) {
		rsp.Header().Set("Content-Type", "text/plain")
		rsp.Write([]byte("PONG"))
	}

	h := go_http.HandlerFunc(fn)
	return h, nil
}
