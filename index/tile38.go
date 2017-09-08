package index

import (
	"errors"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-catalog/record"
	"github.com/whosonfirst/go-whosonfirst-tile38"
	"github.com/whosonfirst/go-whosonfirst-tile38/client"
)

type Tile38Index struct {
	catalog.Index
	endpoint   string
	collection string
	repos      []string
	client     tile38.Tile38Client
}

func (e *Tile38Index) GetById(id int64) (catalog.Record, error) {

	for _, r := range e.repos {

		geom_key := fmt.Sprintf("%d#%s", id, r)
		// meta_key := fmt.Sprintf("%d#meta", id)

		rsp, err := e.client.Do("GET", e.collection, geom_key)

		if err != nil {
			return nil, err
		}

		str_url := fmt.Sprintf("tile38://%s/%s#%d", e.endpoint, e.collection, id)

		r, err := record.NewDefaultRecord("tile38", "tile38", id, str_url, rsp)

		if err == nil {
			return r, nil
		}
	}

	msg := fmt.Sprintf("can't find record for ID %d", id)
	return nil, errors.New(msg)
}

func NewTile38Index(host string, port int, collection string, repos ...string) (catalog.Index, error) {

	t38_client, err := client.NewRESPClient(host, port)

	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s:%d", host, port)

	e := Tile38Index{
		client:     t38_client,
		endpoint:   endpoint,
		collection: collection,
		repos:      repos,
	}

	return &e, nil
}
