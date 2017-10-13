package flags

import (
	"github.com/whosonfirst/go-whosonfirst-inspector"
	"github.com/whosonfirst/go-whosonfirst-inspector/index"
)

type WOFFlags struct {
	catalog.Flags
	flag bool
}

func (fl *WOFFlags) IsBoolFlag() bool {
	return true
}

func (fl *WOFFlags) String() string {
	return "true"
}

func (fl *WOFFlags) Set(value string) error {
	fl.flag = true
	return nil
}

func (fl WOFFlags) ToIndexes() ([]catalog.Index, error) {

	idx, err := index.NewWOFIndex()

	if err != nil {
		return nil, err
	}

	indexes := []catalog.Index{idx}
	return indexes, nil
}
