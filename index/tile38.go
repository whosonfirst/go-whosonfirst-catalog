package index

import (
	"errors"
	"github.com/whosonfirst/go-whosonfirst-catalog"
)

type Tile38Index struct {
	catalog.Index
}

func (e *Tile38Index) GetById(id in64) (catalog.Record, error) {

	return nil, errors.New("Please write me")
}
