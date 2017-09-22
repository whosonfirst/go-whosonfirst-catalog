package http

import (
	"fmt"
	"log"
	gohttp "net/http"
	"regexp"
)

type RewriteRule struct {
	Regexp  *regexp.Regexp
	Replace string
	Last    bool
}

func RemovePrefixRewriteRule(path string) RewriteRule {

	pat := fmt.Sprintf("^%s(.*)", path)

	re := regexp.MustCompile(pat)

	rule := RewriteRule{
		Regexp:  re,
		Replace: "$1",
		Last:    false,
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

			log.Println("BEFORE", path)
			path = re.ReplaceAllString(path, rw.Replace)
			log.Println("AFTER", path)

			if rw.Last {
				break
			}
		}

		req.URL.Path = path

		next.ServeHTTP(rsp, req)
		return
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
