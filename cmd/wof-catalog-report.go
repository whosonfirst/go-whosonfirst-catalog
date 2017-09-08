package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-catalog/flags"
	"github.com/whosonfirst/go-whosonfirst-catalog/probe"
	"github.com/whosonfirst/go-whosonfirst-log"
)

func main() {

	var github_flags flags.GitHubFlags
	var s3_flags flags.S3Flags
	var es_flags flags.ElasticsearchFlags
	var wof_flags flags.WOFFlags

	flag.Var(&github_flags, "github", "...")
	flag.Var(&s3_flags, "s3", "...")
	flag.Var(&es_flags, "es", "...")
	flag.Var(&wof_flags, "wof", "...")

	wofid := flag.Int64("id", -1, "...")
	flag.Parse()

	logger := log.SimpleWOFLogger()

	if *wofid == -1 {
		logger.Fatal("Invalid WOF ID")
	}

	indexes, err := flags.ToIndexes(github_flags, s3_flags, wof_flags, es_flags)

	if err != nil {
		logger.Fatal("Failed to create indexes because %v", err)
	}

	pr, err := probe.NewDefaultProbe(indexes...)

	if err != nil {
		logger.Fatal("Failed to create new probe because %v", err)
	}

	r, err := pr.GetById(*wofid)

	if err != nil {
		logger.Fatal("Failed to probe ID %d because %v", *wofid, err)
	}

	enc, err := json.Marshal(r)

	if err != nil {
		logger.Fatal("Failed to serialize results because %v", err)
	}

	fmt.Println(string(enc))
}
