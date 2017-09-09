package record

import (
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-catalog/utils"
)

type ErrorRecord struct {
	catalog.Record `json:",omitempty"`
	RecordType     string      `json:"type"`
	RecordSource   string      `json:"source"`
	RecordID       int64       `json:"id"`
	RecordURI      string      `json:"uri"`
	RecordBody     interface{} `json:"body"`
	RecordHash     string      `json:"hash"`
}

func (r *ErrorRecord) Id() int64 {
	return r.RecordID
}

func (r *ErrorRecord) Type() string {
	return r.RecordType
}

func (r *ErrorRecord) Source() string {
	return r.RecordSource
}

func (r *ErrorRecord) URI() string {
	return r.RecordURI
}

func (r *ErrorRecord) Body() interface{} {
	return r.RecordBody
}

func NewErrorRecord(source string, id int64, uri string, err error) (catalog.Record, error) {

	err_str := err.Error()

	hash, err := utils.HashInterface(err_str)

	if err != nil {
		return nil, err
	}

	r := ErrorRecord{
		RecordType:   "error",
		RecordSource: source,
		RecordID:     id,
		RecordURI:    uri,
		RecordBody:   err_str,
		RecordHash:   hash,
	}

	return &r, nil
}
