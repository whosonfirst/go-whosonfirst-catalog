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

type Tile38Feature struct {
	Type       string      `json:"type"`
	Geometry   interface{} `json:"geometry"`
	Properties interface{} `json:"properties"`
}

func (e *Tile38Index) GetById(id int64) (catalog.Record, error) {

	for _, r := range e.repos {

		geom_key := fmt.Sprintf("%d#%s", id, r)
		meta_key := fmt.Sprintf("%d#meta", id)

		geom_rsp, err := e.client.Do("GET", e.collection, geom_key)

		if err != nil {
			return nil, err
		}

		if !geom_rsp.(tile38.Tile38Response).Ok {
			continue
		}

		meta_rsp, err := e.client.Do("GET", e.collection, meta_key)

		str_url := fmt.Sprintf("tile38://%s/%s/%s", e.endpoint, e.collection, geom_key)

		// not sure we don't just want to return an array of raw
		// T38 responses... (20170908/thisisaaronland)

		feature := Tile38Feature{
			Type:     "Feature",
			Geometry: geom_rsp.(tile38.Tile38Response).Object,
		}

		if meta_rsp.(tile38.Tile38Response).Ok {
			feature.Properties = meta_rsp.(tile38.Tile38Response).Object
		}

		return record.NewDefaultRecord("tile38", "tile38", id, str_url, feature)
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
