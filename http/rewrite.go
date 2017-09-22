package http

import (
	"fmt"
	_ "log"
	gohttp "net/http"
	"regexp"
)

type RewriteRule struct {
	Regexp  *regexp.Regexp
	Replace string
}

func RemovePrefixRewriteRule(path string) RewriteRule {

	pat := fmt.Sprintf("^%s(.*)", path)

	re := regexp.MustCompile(pat)

	rule := RewriteRule{
		Regexp:  re,
		Replace: "$1",
	}

	return rule
}

func RewriteHandler(rules []RewriteRule, next gohttp.Handler) (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		path := req.URL.Path

		for _, rw := range rules {

			re := rw.Regexp

			if !re.MatchString(path) {
				continue
			}

			path = re.ReplaceAllString(path, rw.Replace)
		}

		req.URL.Path = path

		next.ServeHTTP(rsp, req)
		return
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
