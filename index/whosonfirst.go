package index

import (
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-catalog/record"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"io/ioutil"
	"net/http"
)

type WOFIndex struct {
	catalog.Index
}

func NewWOFIndex() (catalog.Index, error) {

	e := WOFIndex{}

	return &e, nil
}

func (e *WOFIndex) GetById(id int64) (catalog.Record, error) {

	root := "https://whosonfirst.mapzen.com/data/"

	url, err := uri.Id2AbsPath(root, id)

	if err != nil {
		return nil, err
	}

	rsp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return nil, err
	}

	return record.NewDefaultRecord("whosonfirst", id, url, body)
}
