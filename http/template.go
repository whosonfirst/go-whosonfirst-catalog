package http

import (
       "errors"
       go_http "net/http"
       "html/template"
)

func TemplateHandler(templates *template.Template, name string) (go_http.Handler, error) {

     t := templates.Lookup(name)

     if t == nil {
     	return nil, errors.New("Missing template")
     }

	fn := func(rsp go_http.ResponseWriter, req *go_http.Request) {

		err := t.Execute(rsp, nil)

		if err != nil {
			go_http.Error(rsp, err.Error(), go_http.StatusInternalServerError)
			return
		}

		return
	}

	return go_http.HandlerFunc(fn), nil
}
