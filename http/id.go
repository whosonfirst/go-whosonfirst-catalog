package http

import (
       "encoding/json"
       "github.com/whosonfirst/go-whosonfirst-catalog"
	gohttp "net/http"
	"strconv"
)

func IDHandler(pr catalog.Probe) (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		query := req.URL.Query()

		str_id := query.Get("id")

		if str_id == "" {
			gohttp.Error(rsp, "Missing 'id' parameter", gohttp.StatusBadRequest)
		}
		
		id, err := strconv.ParseInt(str_id, 10, 64)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
		}

		results, err := pr.GetById(id)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)		
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
