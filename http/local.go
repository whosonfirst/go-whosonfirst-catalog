package http

import (
	gohttp "net/http"
)

func LocalFileSystem(root string) gohttp.FileSystem {
	return gohttp.Dir(root)
}

func LocalHandler(fs gohttp.FileSystem) (gohttp.Handler, error) {
	return gohttp.FileServer(fs), nil
}
