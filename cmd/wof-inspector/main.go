package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-inspector"
	"github.com/whosonfirst/go-whosonfirst-inspector/flags"
	"github.com/whosonfirst/go-whosonfirst-inspector/probe"
	"github.com/whosonfirst/go-whosonfirst-log"
	"strconv"
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

	flag.Parse()

	logger := log.SimpleWOFLogger()

	indexes, err := flags.ToIndexes(gh_flags, s3_flags, wof_flags, es_flags, t38_flags, pg_flags)

	if err != nil {
		logger.Fatal("Failed to create indexes because %v", err)
	}

	pr, err := probe.NewDefaultProbe(indexes...)

	if err != nil {
		logger.Fatal("Failed to create new probe because %v", err)
	}

	results := make(map[string]catalog.ProbeResults)

	for _, str_id := range flag.Args() {

		wofid, err := strconv.ParseInt(str_id, 10, 64)

		if err != nil {
			logger.Fatal("Failed to parse ID (%s) %s", str_id, err)
		}

		r, err := pr.GetById(wofid)

		if err != nil {
			logger.Fatal("Failed to probe ID %d because %v", wofid, err)
		}

		results[str_id] = r
	}

	report, err := json.Marshal(results)

	if err != nil {
		logger.Fatal("Failed to serialize results because %v", err)
	}

	fmt.Println(string(report))

}
