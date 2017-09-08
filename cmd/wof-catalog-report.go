package main

import (
        "encoding/json"
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-catalog/index"
	"github.com/whosonfirst/go-whosonfirst-catalog/probe"
	"github.com/whosonfirst/go-whosonfirst-log"
)

func main() {

	wofid := flag.Int64("id", -1, "...")
	flag.Parse()

	logger := log.SimpleWOFLogger()

	if *wofid == -1 {
		logger.Fatal("Invalid WOF ID")
	}

	indexes := make([]catalog.Index, 0)

	gh, err := index.NewGitHubIndex("whosonfirst-data")

	if err != nil {
		logger.Fatal("Failed to create new GitHub index because %v", err)
	}

	indexes = append(indexes, gh)

	s3, err := index.NewS3Index()

	if err != nil {
		logger.Fatal("Failed to create new GitHub index because %v", err)
	}

	indexes = append(indexes, s3)
	
	pr, err := probe.NewProbe(indexes...)

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
