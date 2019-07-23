package http

import (
	go_http "net/http"
	"strings"
)

func StaticAssetsHandler() (go_http.Handler, error) {

	fs := assetFS()
	return go_http.FileServer(fs), nil
}

func AppendStaticAssetHandlers(mux *go_http.ServeMux) error {

	asset_handler, err := StaticAssetsHandler()

	if err != nil {
		return nil
	}

	for _, path := range AssetNames() {
		path := strings.Replace(path, "static", "", 1)
		mux.Handle(path, asset_handler)
	}

	return nil
}
