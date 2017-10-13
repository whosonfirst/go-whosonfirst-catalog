package record

import (
	"github.com/whosonfirst/go-whosonfirst-inspector"
	"time"
)

func NewErrorRecord(source string, id int64, uri string, err error, t time.Duration) (catalog.Record, error) {
	return NewDefaultRecord("error", source, id, uri, err.Error(), t)
}
