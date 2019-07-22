package main

import (
	"flag"
	"fmt"
	"github.com/whosonfirst/go-http-nextjs"
	"github.com/whosonfirst/go-http-rewrite"
	"github.com/whosonfirst/go-whosonfirst-inspector/flags"
	"github.com/whosonfirst/go-whosonfirst-inspector/http"
	"github.com/whosonfirst/go-whosonfirst-inspector/probe"
	"github.com/whosonfirst/go-whosonfirst-log"
	gohttp "net/http"
	"os"
)

func main() {

	var es_flags flags.ElasticsearchFlags
	var fs_flags flags.FSFlags
	var gh_flags flags.GitHubFlags
	var pg_flags flags.PgSQLFlags
	var s3_flags flags.S3Flags
	var t38_flags flags.Tile38Flags
	var wof_flags flags.WOFFlags

	flag.Var(&es_flags, "es", "The endpoint of an Elasticsearch database to inspect. Endpoints are defined as HOST + ':' + PORT + '#' + INDEX")
	flag.Var(&fs_flags, "fs", "The path of a Who's On First data directory to inspect. Paths are defined as ROOT + '#' + COMMA-SEPARATED LIST OF REPOSITORIES. If the value of list of repositories is '*' then all the repos in the 'whosonfirst-data' origanization will be used.")
	flag.Var(&gh_flags, "gh", "The name of a GitHub repos to inspect. If '*' then all the repos in the 'whosonfirst-data' organization will be used.")
	flag.Var(&pg_flags, "pg", "The DSN of a PostgreSQL endpoint to inspect.")
	flag.Var(&s3_flags, "s3", "The name of an AWS S3 buckets to inspect.")
	flag.Var(&t38_flags, "t38", "The endpoint of a Tile38 endpoints to inspect. Endpoints are defined as HOST + ':' + PORT + '#' + COMMA-SEPARATED LIST OF REPOSITORIES. If the value of list of repositories is '*' then all the repos in the 'whosonfirst-data' origanization will be used.")
	flag.Var(&wof_flags, "wof", "Inspect records hosted on whosonfirst.mapzen.com/data.")

	var host = flag.String("host", "localhost", "The hostname to listen for requests on.")
	var port = flag.Int("port", 8080, "The port number to listen for requests on.")

	var local = flag.String("local", "", "The path to a local directory containing HTML (and CSS/Javascript) assets that the wof-inspectord daemon should serve. The default behaviour is to served bundled assets.")
	var root = flag.String("root", "/", "The root path to host the wof-inspectord daemon on.")

	var api_key = flag.String("api-key", "nextzen-xxxxxxx", "A valid Nextzen API key (necessary for displaying maps).")

	flag.Parse()

	logger := log.SimpleWOFLogger()

	indexes, err := flags.ToIndexes(gh_flags, s3_flags, wof_flags, es_flags, t38_flags, pg_flags, fs_flags)

	pr, err := probe.NewDefaultProbe(indexes...)

	if err != nil {
		logger.Fatal("Failed to create new probe because %v", err)
	}

	id_handler, err := http.IDHandler(pr)

	if err != nil {
		logger.Fatal("failed to create ID handler because %s", err)
	}

	var www_handler gohttp.Handler
	var www_fs gohttp.FileSystem

	if *local == "" {

		asset_handler, err := http.WWWHandler()

		if err != nil {
			logger.Fatal("failed to create www handler because %s", err)
		}

		asset_fs := http.WWWFileSystem()

		www_handler = asset_handler
		www_fs = asset_fs

	} else {

		local_fs := http.LocalFileSystem(*local)
		local_handler, err := http.LocalHandler(local_fs)

		if err != nil {
			logger.Fatal("failed to create www handler because %s", err)
		}

		www_handler = local_handler
		www_fs = local_fs
	}

	root_handler, err := nextzenjs.NextzenAPIKeyHandler(www_handler, www_fs, *api_key)

	if err != nil {
		logger.Fatal("failed to create query handler because %s", err)
	}

	nextzenjs_handler, err := nextzenjs.NextzenJSHandler()

	if err != nil {
		logger.Fatal("failed to create nextzen.js handler because %s", err)
	}

	ping_handler, err := http.PingHandler()

	if err != nil {
		logger.Fatal("failed to create ping handler because %s", err)
	}

	handlers := map[string]gohttp.Handler{
		"/":                          root_handler,
		"/id/":                       id_handler,
		"/ping/":                     ping_handler,
		"/javascript/nextzen.min.js":  nextzenjs_handler,
		"/javascript/tangram.min.js": nextzenjs_handler,
		"/css/nextzen.js.css":         nextzenjs_handler,
		"/tangram/refill-style.zip":  nextzenjs_handler,
	}

	if *root != "/" {

		opts := rewrite.DefaultRewriteRuleOptions()

		rule := rewrite.RemovePrefixRewriteRule(*root, opts)
		rules := []rewrite.RewriteRule{rule}

		for path, handler := range handlers {

			rw_path := *root + path
			rw_handler, err := rewrite.RewriteHandler(rules, handler)

			if err != nil {
				logger.Fatal("failed to create rewrite handler for %s (%s) because %s", rw_path, path, err)
			}

			delete(handlers, path)
			handlers[rw_path] = rw_handler
		}
	}

	mux := gohttp.NewServeMux()

	for path, handler := range handlers {
		logger.Status("configure %s handler", path)
		mux.Handle(path, handler)
	}

	endpoint := fmt.Sprintf("%s:%d", *host, *port)
	logger.Status("listening on %s", endpoint)

	err = gohttp.ListenAndServe(endpoint, mux)

	if err != nil {
		logger.Fatal("failed to start server because %s", err)
	}

	os.Exit(0)
}
