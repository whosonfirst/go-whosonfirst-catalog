package http

import (
	gohttp "net/http"
)

func QueryHandler() (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		gohttp.Error(rsp, "Please implement me", gohttp.StatusInternalServerError)
		return
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
