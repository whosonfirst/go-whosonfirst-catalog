package flags

import (
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-catalog/index"
	"github.com/whosonfirst/go-whosonfirst-catalog/utils"
	"strconv"
	"strings"
)

type Tile38Flags struct {
	catalog.Flags
	flags []string
}

func (fl *Tile38Flags) String() string {
	return strings.Join(fl.flags, "\n")
}

func (fl *Tile38Flags) Set(value string) error {
	fl.flags = append(fl.flags, value)
	return nil
}

func (fl Tile38Flags) ToIndexes() ([]catalog.Index, error) {

	indexes := make([]catalog.Index, 0)

	for _, addr := range fl.flags {

		var host string
		var port int
		var collection string
		var repos []string

		var tmp []string

		tmp = strings.Split(addr, "#")

		if len(tmp) == 1 {
			addr = tmp[0]
			repos = []string{"whosonfirst-data"}
		} else {
			addr = tmp[0]
			repos = strings.Split(tmp[1], ",")
		}

		tmp = strings.Split(addr, "/")

		if len(tmp) == 1 {
			addr = tmp[0]
			collection = "whosonfirst"
		} else {
			addr = tmp[0]
			collection = tmp[1]
		}

		tmp = strings.Split(addr, ":")

		if len(tmp) == 1 {
			host = tmp[0]
			port = 9851
		} else {

			host = tmp[0]

			p, err := strconv.Atoi(tmp[1])

			if err != nil {
				return nil, err
			}

			port = p
		}

		if len(repos) == 1 && repos[0] == "*" {

			r, err := utils.ListGitHubRepos()

			if err != nil {
				return nil, err
			}

			repos = r
		}

		idx, err := index.NewTile38Index(host, port, collection, repos...)

		if err != nil {
			return nil, err
		}

		indexes = append(indexes, idx)
	}

	return indexes, nil
}
