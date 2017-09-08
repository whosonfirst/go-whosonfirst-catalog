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

	var github_flags flags.GitHubFlags
	var s3_flags flags.S3Flags
	var wof_flags flags.WOFFlags

	flag.Var(&github_flags, "github", "...")
	flag.Var(&s3_flags, "s3", "...")
	flag.Var(&wof_flags, "wof", "...")

	var host = flag.String("host", "localhost", "The hostname to listen for requests on")
	var port = flag.Int("port", 8080, "The port number to listen for requests on")

	flag.Parse()

	logger := log.SimpleWOFLogger()

	indexes, err := flags.ToIndexes(github_flags, s3_flags, wof_flags)

	pr, err := probe.NewDefaultProbe(indexes...)

	if err != nil {
		logger.Fatal("Failed to create new probe because %v", err)
	}

	idhandler, err := http.IDHandler(pr)

	if err != nil {
		logger.Fatal("failed to create ID handler because %s", err)
	}

	pinghandler, err := http.PingHandler()

	if err != nil {
		logger.Fatal("failed to create Ping handler because %s", err)
	}

	endpoint := fmt.Sprintf("%s:%d", *host, *port)
	logger.Status("listening on %s", endpoint)

	mux := gohttp.NewServeMux()
	mux.Handle("/", idhandler)
	mux.Handle("/ping", pinghandler)

	err = gracehttp.Serve(&gohttp.Server{Addr: endpoint, Handler: mux})

	if err != nil {
		logger.Fatal("failed to start server because %s", err)
	}

	os.Exit(0)
}
