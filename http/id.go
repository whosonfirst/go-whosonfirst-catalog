package http

import (
	"encoding/json"
	"github.com/whosonfirst/go-whosonfirst-inspector"
	go_http "net/http"
	"path/filepath"
	"strconv"
)

func IDHandler(pr catalog.Probe) (go_http.Handler, error) {

	fn := func(rsp go_http.ResponseWriter, req *go_http.Request) {

		req_path := req.URL.Path

		str_id := filepath.Base(req_path)

		if str_id == "" {
			go_http.Error(rsp, "Missing ID", go_http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseInt(str_id, 10, 64)

		if err != nil {
			go_http.Error(rsp, err.Error(), go_http.StatusBadRequest)
			return
		}

		results, err := pr.GetById(id)

		if err != nil {
			go_http.Error(rsp, err.Error(), go_http.StatusInternalServerError)
			return
		}

		js, err := json.Marshal(results)

		if err != nil {
			go_http.Error(rsp, err.Error(), go_http.StatusInternalServerError)
			return
		}

		rsp.Header().Set("Content-Type", "application/json")
		rsp.Write(js)
	}

	h := go_http.HandlerFunc(fn)
	return h, nil
}
