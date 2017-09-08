package index

import (
	"errors"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-catalog/record"
	"github.com/whosonfirst/go-whosonfirst-catalog/utils"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"strings"
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

	for _, repo := range e.repos {

		root := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/data/", e.org, repo, e.branch)

		url, err := uri.Id2AbsPath(root, id)

		if err != nil {
			return nil, err
		}

		rsp, err := utils.GetURLAsJSON(url)

		if err != nil {

			if strings.HasPrefix(err.Error(), "404") {
				continue
			}

			return nil, err
		}

		return record.NewDefaultRecord("geojson", "github", id, url, rsp)
	}

	msg := fmt.Sprintf("can't find record for ID %d", id)
	return nil, errors.New(msg)
}
