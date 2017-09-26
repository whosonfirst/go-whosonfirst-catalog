package record

import (
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-catalog/utils"
	"time"
)

type DefaultRecord struct {
	catalog.Record `json:",omitempty"`
	RecordType     string        `json:"type"`
	RecordSource   string        `json:"source"`
	RecordID       int64         `json:"id"`
	RecordURI      string        `json:"uri"`
	RecordBody     interface{}   `json:"body"`
	RecordHash     string        `json:"hash"`
	RecordTiming   time.Duration `json:"timing"`
}

func (r *DefaultRecord) Id() int64 {
	return r.RecordID
}

func (r *DefaultRecord) Type() string {
	return r.RecordType
}

func (r *DefaultRecord) Source() string {
	return r.RecordSource
}

func (r *DefaultRecord) URI() string {
	return r.RecordURI
}

func (r *DefaultRecord) Body() interface{} {
	return r.RecordBody
}

func (r *DefaultRecord) Timing() time.Duration {
	return r.RecordTiming
}

func NewDefaultRecord(rtype string, source string, id int64, uri string, body interface{}, t time.Duration) (catalog.Record, error) {

	hash, err := utils.HashInterface(body)

	if err != nil {
		return nil, err
	}

	r := DefaultRecord{
		RecordType:   rtype,
		RecordSource: source,
		RecordID:     id,
		RecordURI:    uri,
		RecordBody:   body,
		RecordHash:   hash,
		RecordTiming: t,
	}

	return &r, nil
}
