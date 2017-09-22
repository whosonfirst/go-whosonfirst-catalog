package http

import (
	"encoding/json"
	"github.com/whosonfirst/go-whosonfirst-catalog"
	gohttp "net/http"
	"path/filepath"
	"strconv"
	_ "strings"
)

func IDHandler(pr catalog.Probe) (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		req_path := req.URL.Path

		str_id := filepath.Base(req_path)

		if str_id == "" {
			gohttp.Error(rsp, "Missing ID", gohttp.StatusBadRequest)
			return
		}

		id, err := strconv.ParseInt(str_id, 10, 64)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}

		results, err := pr.GetById(id)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		js, err := json.Marshal(results)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		rsp.Header().Set("Content-Type", "application/json")
		rsp.Header().Set("Access-Control-Allow-Origin", "*")

		rsp.Write(js)
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
