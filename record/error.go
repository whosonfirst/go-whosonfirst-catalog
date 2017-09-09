package record

import (
	"github.com/whosonfirst/go-whosonfirst-catalog"
)

func NewErrorRecord(source string, id int64, uri string, err error) (catalog.Record, error) {
	return NewDefaultRecord("error", source, id, uri, err.Error())
}
