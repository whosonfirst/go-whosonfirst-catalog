package index

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"io/ioutil"
	"net/http"
)

type GitHubIndex struct {
	catalog.Index
	org    string
	repo   string
	branch string
}

type GitHubRecord struct {
	catalog.Record `json:",omitempty"`
	RecordID       int64       `json:"id"`
	RecordURI      string      `json:"uri"`
	RecordSource   string      `json:"source"`
	RecordBody     interface{} `json:"body"`
}

func (r *GitHubRecord) Id() int64 {
	return r.RecordID
}

func (r *GitHubRecord) Source() string {
	return r.RecordSource
}

func (r *GitHubRecord) URI() string {
	return r.RecordURI
}

func (r *GitHubRecord) Body() interface{} {
	return r.RecordBody
}

func NewGitHubIndex(repo string) (catalog.Index, error) {

	e := GitHubIndex{
		org:    "whosonfirst-data",
		branch: "master",
		repo:   "repo",
	}

	return &e, nil
}

func (e *GitHubIndex) GetById(id int64) (catalog.Record, error) {

	root := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/data/", e.org, e.repo, e.branch)

	url, err := uri.IdToAbsPath(root, id)

	if err != nil {
		return nil, err
	}

	rsp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp)

	if err != nil {
		return nil, err
	}

	r := GitHubRecord{
		RecordID:     id,
		RecordSource: "github",
		RecordURI:    url,
		RecordBody:   body,
	}

	return &r, nil
}
