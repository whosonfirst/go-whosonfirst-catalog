package record

import (
       "github.com/whosonfirst/go-whosonfirst-catalog"
)

type DefaultRecord struct {
	catalog.Record `json:",omitempty"`
	RecordSource   string      `json:"source"`	
	RecordID       int64       `json:"id"`
	RecordURI      string      `json:"uri"`
	RecordBody     []byte `json:"body"`
}

func (r *DefaultRecord) Id() int64 {
	return r.RecordID
}

func (r *DefaultRecord) Source() string {
	return r.RecordSource
}

func (r *DefaultRecord) URI() string {
	return r.RecordURI
}

func (r *DefaultRecord) Body() []byte {
	return r.RecordBody
}

func NewDefaultRecord(source string, id int64, uri string, body []byte) (catalog.Record, error){

	r := DefaultRecord{
		RecordSource: source,
		RecordID:     id,
		RecordURI:    uri,
		RecordBody:   body,
	}

	return &r, nil
}
