package flags

import (
	"github.com/whosonfirst/go-whosonfirst-inspector"
	"github.com/whosonfirst/go-whosonfirst-inspector/index"
	"strings"
)

type WOFFlags struct {
	catalog.Flags
	flags []string
}

func (fl *WOFFlags) String() string {
	return strings.Join(fl.flags, "\n")
}

func (fl *WOFFlags) Set(value string) error {
	fl.flags = append(fl.flags, value)
	return nil
}

func (fl WOFFlags) ToIndexes() ([]catalog.Index, error) {

	indexes := make([]catalog.Index, 0)

	for range fl.flags {

		idx, err := index.NewWOFIndex()

		if err != nil {
			return nil, err
		}

		indexes = append(indexes, idx)
	}

	return indexes, nil
}
