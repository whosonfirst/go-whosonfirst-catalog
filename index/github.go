package index

import (
	"errors"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-inspector"
	"github.com/whosonfirst/go-whosonfirst-inspector/record"
	"github.com/whosonfirst/go-whosonfirst-inspector/utils"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"strings"
	"time"
)

type GitHubIndex struct {
	catalog.Index
	org    string
	repos  []string
	branch string
}

func NewGitHubIndex(repos ...string) (catalog.Index, error) {

	e := GitHubIndex{
		org:    "whosonfirst-data",
		branch: "master",
		repos:  repos,
	}

	return &e, nil
}

func (e *GitHubIndex) GetById(id int64) (catalog.Record, error) {

	t1 := time.Now()
	var t2 time.Duration

	for _, repo := range e.repos {

		root := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/data/", e.org, repo, e.branch)

		uri, err := uri.Id2AbsPath(root, id)

		if err != nil {
			return nil, err
		}

		rsp, err := utils.GetURLAsJSON(uri)

		t2 = time.Since(t1)

		if err != nil {

			if strings.HasPrefix(err.Error(), "404") {
				continue
			}

			return record.NewErrorRecord("github", id, uri, err, t2)
		}

		return record.NewDefaultRecord("geojson", "github", id, uri, rsp, t2)
	}

	t2 = time.Since(t1)

	msg := fmt.Sprintf("can't find github record for ID %d", id)
	err := errors.New(msg)

	uri := fmt.Sprintf("x-github-%d", id)
	return record.NewErrorRecord("github", id, uri, err, t2)
}
