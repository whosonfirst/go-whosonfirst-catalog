package main

import (
	"flag"
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/whosonfirst/go-whosonfirst-catalog/flags"
	"github.com/whosonfirst/go-whosonfirst-catalog/http"
	"github.com/whosonfirst/go-whosonfirst-catalog/probe"
	"github.com/whosonfirst/go-whosonfirst-log"
	gohttp "net/http"
	"os"
)

func main() {

	var es_flags flags.ElasticsearchFlags
	var gh_flags flags.GitHubFlags
	var pg_flags flags.PgSQLFlags
	var s3_flags flags.S3Flags
	var t38_flags flags.Tile38Flags
	var wof_flags flags.WOFFlags

	flag.Var(&gh_flags, "gh", "...")
	flag.Var(&s3_flags, "s3", "...")
	flag.Var(&es_flags, "es", "...")
	flag.Var(&pg_flags, "pg", "...")
	flag.Var(&t38_flags, "t38", "...")
	flag.Var(&wof_flags, "wof", "...")

	var host = flag.String("host", "localhost", "The hostname to listen for requests on")
	var port = flag.Int("port", 8080, "The port number to listen for requests on")
	var api_key = flag.String("api-key", "mapzen-12345", "")

	flag.Parse()

	logger := log.SimpleWOFLogger()

	indexes, err := flags.ToIndexes(gh_flags, s3_flags, wof_flags, es_flags, t38_flags, pg_flags)

	pr, err := probe.NewDefaultProbe(indexes...)

	if err != nil {
		logger.Fatal("Failed to create new probe because %v", err)
	}

	id_handler, err := http.IDHandler(pr)

	if err != nil {
		logger.Fatal("failed to create ID handler because %s", err)
	}

	www_handler, err := http.WWWHandler()

	if err != nil {
		logger.Fatal("failed to create fs handler because %s", err)
	}

	key_handler, err := http.APIKeyHandler(www_handler, *api_key)

	if err != nil {
		logger.Fatal("failed to create query handler because %s", err)
	}

	ping_handler, err := http.PingHandler()

	if err != nil {
		logger.Fatal("failed to create ping handler because %s", err)
	}

	endpoint := fmt.Sprintf("%s:%d", *host, *port)
	logger.Status("listening on %s", endpoint)

	mux := gohttp.NewServeMux()

	mux.Handle("/id/", id_handler)
	mux.Handle("/ping", ping_handler)

	mux.Handle("/javascript/", www_handler)
	mux.Handle("/css/", www_handler)
	mux.Handle("/tangram/", www_handler)
	mux.Handle("/", key_handler)

	err = gracehttp.Serve(&gohttp.Server{Addr: endpoint, Handler: mux})

	if err != nil {
		logger.Fatal("failed to start server because %s", err)
	}

	os.Exit(0)
}
