package main

import (
	"flag"
	"fmt"
	"github.com/rs/cors"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/whosonfirst/go-http-nextzenjs"
	"github.com/whosonfirst/go-whosonfirst-inspector/flags"
	"github.com/whosonfirst/go-whosonfirst-inspector/http"
	"github.com/whosonfirst/go-whosonfirst-inspector/probe"
	"github.com/whosonfirst/go-whosonfirst-inspector/templates"
	"github.com/whosonfirst/go-whosonfirst-log"
	"html/template"
	go_http "net/http"
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

	local_templates := flag.String("templates", "", "...")

	var api_key = flag.String("nextzen-api-key", "nextzen-xxxxxxx", "A valid Nextzen API key (necessary for displaying maps).")

	flag.Parse()

	logger := log.SimpleWOFLogger()

	indexes, err := flags.ToIndexes(gh_flags, s3_flags, wof_flags, es_flags, t38_flags, pg_flags, fs_flags)

	pr, err := probe.NewDefaultProbe(indexes...)

	if err != nil {
		logger.Fatal("Failed to create new probe because %v", err)
	}

	var www_templates *template.Template

	if *local_templates != "" {

		t, err := template.ParseGlob(*local_templates)

		if err != nil {
			logger.Fatal(err)
		}

		www_templates = t

	} else {

		t := template.New("templates")

		for _, name := range templates.AssetNames() {

			body, err := templates.Asset(name)

			if err != nil {
				logger.Fatal(err)
			}

			t, err = t.Parse(string(body))

			if err != nil {
				logger.Fatal(err)
			}
		}

		www_templates = t
	}

	mux := go_http.NewServeMux()

	index_handler, err := http.TemplateHandler(www_templates, "index")

	if err != nil {
		logger.Fatal(err)
	}

	nextzenjs_opts := nextzenjs.DefaultNextzenJSOptions()
	nextzenjs_opts.APIKey = *api_key

	index_handler, _ = nextzenjs.NextzenJSHandler(index_handler, nextzenjs_opts)

	mux.Handle("/", index_handler)

	id_handler, err := http.IDHandler(pr)

	if err != nil {
		logger.Fatal("failed to create ID handler because %s", err)
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST"},
	})

	id_handler = c.Handler(id_handler)

	mux.Handle("/id/", id_handler)

	ping_handler, err := http.PingHandler()

	if err != nil {
		logger.Fatal("failed to create ping handler because %s", err)
	}

	mux.Handle("/ping", ping_handler)

	assets_handler, err := nextzenjs.NextzenJSAssetsHandler()

	if err != nil {
		logger.Fatal(err)
	}

	mux.Handle("/javascript/nextzen.js", assets_handler)
	mux.Handle("/javascript/nextzen.min.js", assets_handler)
	mux.Handle("/javascript/tangram.js", assets_handler)
	mux.Handle("/javascript/tangram.min.js", assets_handler)
	mux.Handle("/css/nextzen.js.css", assets_handler)
	mux.Handle("/tangram/refill-style.zip", assets_handler)

	err = http.AppendStaticAssetHandlers(mux)

	if err != nil {
		logger.Fatal(err)
	}

	err = bootstrap.AppendAssetHandlers(mux)

	if err != nil {
		logger.Fatal(err)
	}

	endpoint := fmt.Sprintf("%s:%d", *host, *port)
	logger.Status("listening on %s", endpoint)

	err = go_http.ListenAndServe(endpoint, mux)

	if err != nil {
		logger.Fatal("failed to start server because %s", err)
	}

	os.Exit(0)
}
