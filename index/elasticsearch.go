package index

import (
	"errors"
	"github.com/whosonfirst/go-whosonfirst-catalog"
)

type ElasticsearchIndex struct {
	catalog.Index
}

func (e *ElasticsearchIndex) GetById(id in64) (catalog.Record, error) {

	return nil, errors.New("Please write me")
}
