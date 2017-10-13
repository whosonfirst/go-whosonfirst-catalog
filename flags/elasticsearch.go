package flags

import (
	"github.com/whosonfirst/go-whosonfirst-inspector"
	"github.com/whosonfirst/go-whosonfirst-inspector/index"
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

	for _, addr := range fl.flags {

		var host string
		var port int
		var es_index string

		var tmp []string

		tmp = strings.Split(addr, "#")

		if len(tmp) == 1 {
			addr = tmp[0]
			es_index = "spelunker"
		} else {
			addr = tmp[0]
			es_index = tmp[1]
		}

		tmp = strings.Split(addr, ":")

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

		idx, err := index.NewElasticsearchIndex(host, port, es_index)

		if err != nil {
			return nil, err
		}

		indexes = append(indexes, idx)
	}

	return indexes, nil
}
