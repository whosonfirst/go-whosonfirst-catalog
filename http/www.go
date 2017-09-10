package http

import (
	"bytes"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	_ "log"
	gohttp "net/http"
)

func WWWHandler() (gohttp.Handler, error) {

	fs := assetFS()
	return gohttp.FileServer(fs), nil
}

func APIKeyHandler(next gohttp.Handler, api_key string) (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		if req.URL.Path != "/" {
			next.ServeHTTP(rsp, req)
			return
		}

		fs := assetFS()

		fh, err := fs.Open("index.html")

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		body, err := ioutil.ReadAll(fh)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		reader := bytes.NewReader(body)
		doc, err := html.Parse(reader)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		var f func(node *html.Node, writer io.Writer)

		f = func(n *html.Node, w io.Writer) {

			if n.Type == html.ElementNode && n.Data == "body" {

				api_key_ns := ""
				api_key_key := "data-mapzen-api-key"
				api_key_value := api_key

				api_key_attr := html.Attribute{api_key_ns, api_key_key, api_key_value}
				n.Attr = append(n.Attr, api_key_attr)
			}

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c, w)
			}
		}

		f(doc, rsp)

		html.Render(rsp, doc)
		return
	}

	return gohttp.HandlerFunc(fn), nil
}
