package index

import (
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-catalog/record"
	"github.com/whosonfirst/go-whosonfirst-catalog/utils"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"time"
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

	uri, err := uri.Id2AbsPath(root, id)

	if err != nil {
		return nil, err
	}

	t1 := time.Now()

	rsp, err := utils.GetURLAsJSON(uri)

	t2 := time.Since(t1)

	if err != nil {
		return record.NewErrorRecord("whosonfirst", id, uri, err, t2)
	}

	return record.NewDefaultRecord("geojson", "whosonfirst", id, uri, rsp, t2)
}
