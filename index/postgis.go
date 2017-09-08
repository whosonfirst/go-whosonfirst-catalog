package index

import (
	"errors"
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-pgis"
)

type PostGISIndex struct {
	catalog.Index
}

func (e *PostGISIndex) GetById(id in64) (catalog.Record, error) {

	return nil, errors.New("Please write me")
}
