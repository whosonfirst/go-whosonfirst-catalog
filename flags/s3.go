package flags

import (
	"github.com/whosonfirst/go-whosonfirst-inspector"
	"github.com/whosonfirst/go-whosonfirst-inspector/index"
	"strings"
)

type S3Flags struct {
	catalog.Flags
	flags []string
}

func (fl *S3Flags) String() string {
	return strings.Join(fl.flags, "\n")
}

func (fl *S3Flags) Set(value string) error {
	fl.flags = append(fl.flags, value)
	return nil
}

func (fl S3Flags) ToIndexes() ([]catalog.Index, error) {

	indexes := make([]catalog.Index, 0)

	for _, bucket := range fl.flags {

		idx, err := index.NewS3Index(bucket)

		if err != nil {
			return nil, err
		}

		indexes = append(indexes, idx)
	}

	return indexes, nil
}
