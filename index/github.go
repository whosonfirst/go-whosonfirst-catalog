package index

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-catalog/record"
	"github.com/whosonfirst/go-whosonfirst-catalog/utils"	
	"github.com/whosonfirst/go-whosonfirst-uri"
)

type GitHubIndex struct {
	catalog.Index
	org    string
	repo   string
	branch string
}

func NewGitHubIndex(repo string) (catalog.Index, error) {

	e := GitHubIndex{
		org:    "whosonfirst-data",
		branch: "master",
		repo:   repo,
	}

	return &e, nil
}

func (e *GitHubIndex) GetById(id int64) (catalog.Record, error) {

	root := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/data/", e.org, e.repo, e.branch)

	url, err := uri.Id2AbsPath(root, id)

	if err != nil {
		return nil, err
	}

	body, err := utils.GetURL(url)
	
	if err != nil {
		return nil, err
	}

	return record.NewDefaultRecord("github", id, url, body)
}
