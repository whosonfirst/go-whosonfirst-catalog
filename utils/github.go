package utils

import (
	"github.com/whosonfirst/go-whosonfirst-github/organizations"
)

func ListGitHubRepos() ([]string, error) {

	data, found := CacheGet("repos")

	if found {
		return data.([]string), nil
	}

	opts := organizations.NewDefaultListOptions()
	repos, err := organizations.ListRepos("whosonfirst-data", opts)

	if err != nil {
		return nil, err
	}

	CacheSet("repos", repos)

	return repos, nil
}
