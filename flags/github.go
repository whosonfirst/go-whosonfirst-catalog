package flags

import (
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-catalog/index"
	"github.com/whosonfirst/go-whosonfirst-catalog/utils"
	_ "log"
	"strings"
)

type GitHubFlags struct {
	catalog.Flags
	flags []string
}

func (fl *GitHubFlags) String() string {
	return strings.Join(fl.flags, "\n")
}

func (fl *GitHubFlags) Set(value string) error {
	fl.flags = append(fl.flags, value)
	return nil
}

func (fl GitHubFlags) ToIndexes() ([]catalog.Index, error) {

	indexes := make([]catalog.Index, 0)
	repos := make([]string, 0)

	for _, repo := range fl.flags {

		switch repo {
		case "*":

			r, err := utils.ListGitHubRepos()

			if err != nil {
				return nil, err
			}

			repos = r

		default:
			repos = strings.Split(repo, ",")
		}

		idx, err := index.NewGitHubIndex(repos...)

		if err != nil {
			return nil, err
		}

		indexes = append(indexes, idx)
	}

	return indexes, nil
}
