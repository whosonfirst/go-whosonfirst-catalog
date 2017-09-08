package index

import (
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"io/ioutil"
	"net/http"
)

type WOFIndex struct {
	catalog.Index
}

type WOFRecord struct {
	catalog.Record `json:",omitempty"`
	id             int64       `json:"id"`
	uri            string      `json:"uri"`
	address        string      `json:"address"`
	body           interface{} `json:"body"`
}

func (r *WOFRecord) Id() int64 {
	return r.id
}

func (r *WOFRecord) Source() string {
	return r.source
}

func (r *WOFRecord) URI() string {
	return r.uri
}

func (r *WOFRecord) Body() interface{} {
	return r.body
}

func NewWOFIndex() (catalog.Index, error) {

	e := WOFIndex{}

	return &e, nil
}

func (e *WOFIndex) GetById(id int64) ([]byte, error) {

	root := "https://whosonfirst.mapzen.com/data/"

	abs_root, err := uri.IdToAbsPath(root, id)

	if err != nil {
		return nil, err
	}

	rsp, err := http.Get(abs_root)

	if err != nil {
		return nil, err
	}

	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp)

	if err != nil {
		return nil, err
	}

	r := WOFRecord{
		id:       id,
		"source": "WOF",
		"uri":    url,
		"body":   body,
	}

	return &r, nil
}
