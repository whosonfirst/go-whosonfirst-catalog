package record

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/whosonfirst/go-whosonfirst-inspector"
	"github.com/whosonfirst/go-whosonfirst-inspector/utils"
	_ "log"
	"time"
)

type ElasticsearchRecord struct {
	catalog.Record     `json:",omitempty"`
	RecordType         string        `json:"type"`
	RecordSource       string        `json:"source"`
	RecordID           int64         `json:"id"`
	RecordURI          string        `json:"uri"`
	RecordBody         interface{}   `json:"body"`
	RecordHash         string        `json:"hash"`
	RecordTiming       time.Duration `json:"timing"`
	RecordLastModified int64         `json:"lastmodified"`
}

func (r *ElasticsearchRecord) Id() int64 {
	return r.RecordID
}

func (r *ElasticsearchRecord) Type() string {
	return r.RecordType
}

func (r *ElasticsearchRecord) Source() string {
	return r.RecordSource
}

func (r *ElasticsearchRecord) URI() string {
	return r.RecordURI
}

func (r *ElasticsearchRecord) Body() interface{} {
	return r.RecordBody
}

func (r *ElasticsearchRecord) Timing() time.Duration {
	return r.RecordTiming
}

func (r *ElasticsearchRecord) LastModified() int64 {
	return r.RecordLastModified
}

func NewElasticsearchRecord(id int64, uri string, body interface{}, t time.Duration) (catalog.Record, error) {

	b, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	// because the cruft surrending this will be different for each
	// host we call and by extension the hashes will never line up
	// (20170928/thisisaaronland)

	path := "hits.hits.0._source"
	source_rsp := gjson.GetBytes(b, path)

	if !source_rsp.Exists() {
		msg := fmt.Sprintf("Unable to find result for path '%s'", path)
		return nil, errors.New(msg)
	}

	body = source_rsp.Value()

	hash, err := utils.HashInterface(body)

	if err != nil {
		return nil, err
	}

	lastmod := LastModified(body)

	r := ElasticsearchRecord{
		RecordType:         "elasticsearch",
		RecordSource:       "elasticsearch",
		RecordID:           id,
		RecordURI:          uri,
		RecordBody:         body,
		RecordHash:         hash,
		RecordTiming:       t,
		RecordLastModified: lastmod,
	}

	return &r, nil
}
