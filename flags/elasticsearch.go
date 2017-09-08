package flags

import (
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-catalog/index"
	"strconv"
	"strings"
)

type ElasticsearchFlags struct {
	catalog.Flags
	flags []string
}

func (fl *ElasticsearchFlags) String() string {
	return strings.Join(fl.flags, "\n")
}

func (fl *ElasticsearchFlags) Set(value string) error {
	fl.flags = append(fl.flags, value)
	return nil
}

func (fl ElasticsearchFlags) ToIndexes() ([]catalog.Index, error) {

	indexes := make([]catalog.Index, 0)

	for _, pointer := range fl.flags {

		var endpoint string
		var host string
		var port int
		var es_idx string

		var tmp []string

		tmp = strings.Split(pointer, "#")

		if len(tmp) == 1 {
			endpoint = tmp[0]
			es_idx = "whosonfirst"
		} else {
			endpoint = tmp[0]
			es_idx = tmp[1]
		}

		tmp = strings.Split(endpoint, ":")

		if len(tmp) == 1 {
			host = tmp[0]
			port = 9200
		} else {

			host = tmp[0]

			p, err := strconv.Atoi(tmp[1])

			if err != nil {
				return nil, err
			}

			port = p
		}

		idx, err := index.NewElasticsearchIndex(host, port, es_idx)

		if err != nil {
			return nil, err
		}

		indexes = append(indexes, idx)
	}

	return indexes, nil
}
