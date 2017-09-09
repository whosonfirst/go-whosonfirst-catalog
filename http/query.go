package http

import (
	gohttp "net/http"
)

func QueryHandler(docroot string) (gohttp.Handler, error) {

	root := gohttp.Dir(docroot)
        return gohttp.FileServer(root), nil
}
