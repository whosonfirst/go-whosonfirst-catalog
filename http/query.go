package http

import (
	gohttp "net/http"
)

func QueryHandler() (gohttp.Handler, error) {

	root := assetFS()
        return gohttp.FileServer(root), nil
}
