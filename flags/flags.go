package flags

import (
	"github.com/whosonfirst/go-whosonfirst-inspector"
)

func ToIndexes(possible_flags ...catalog.Flags) ([]catalog.Index, error) {

	indexes := make([]catalog.Index, 0)

	for _, fl := range possible_flags {

		possible_indexes, err := fl.ToIndexes()

		if err != nil {
			return nil, err
		}

		for _, idx := range possible_indexes {
			indexes = append(indexes, idx)
		}
	}

	return indexes, nil
}
