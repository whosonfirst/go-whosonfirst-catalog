package flags

import (
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-catalog/index"
	"github.com/whosonfirst/go-whosonfirst-catalog/utils"
	"strings"
)

type FSFlags struct {
	catalog.Flags
	flags []string
}

func (fl *FSFlags) String() string {
	return strings.Join(fl.flags, "\n")
}

func (fl *FSFlags) Set(value string) error {
	fl.flags = append(fl.flags, value)
	return nil
}

func (fl FSFlags) ToIndexes() ([]catalog.Index, error) {

	indexes := make([]catalog.Index, 0)

	for _, fs := range fl.flags {

		var root string
		var repos []string

		parts := strings.Split(fs, "#")

		if len(parts) == 1 {
			root = parts[0]
			repos = []string{"whosonfirst-data"}
		} else {
			root = parts[0]
			repos = strings.Split(parts[1], ",")
		}

		if len(repos) == 1 && repos[0] == "*" {

			r, err := utils.ListGitHubRepos()

			if err != nil {
				return nil, err
			}

			repos = r
		}

		idx, err := index.NewFSIndex(root, repos...)

		if err != nil {
			return nil, err
		}

		indexes = append(indexes, idx)
	}

	return indexes, nil
}
