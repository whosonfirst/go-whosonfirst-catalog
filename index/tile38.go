package index

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-inspector"
	"github.com/whosonfirst/go-whosonfirst-inspector/record"
	"github.com/whosonfirst/go-whosonfirst-tile38"
	"github.com/whosonfirst/go-whosonfirst-tile38/client"
	_ "log"
	"time"
)

type Tile38Index struct {
	catalog.Index
	endpoint   string
	collection string
	repos      []string
	client     tile38.Tile38Client
}

type Tile38Response map[string]interface{}

func (e *Tile38Index) GetById(id int64) (catalog.Record, error) {

	t1 := time.Now()
	var t2 time.Duration

	for _, r := range e.repos {

		geom_key := fmt.Sprintf("%d#%s", id, r)
		meta_key := fmt.Sprintf("%d#meta", id)

		uri := fmt.Sprintf("tile38://%s/%s/%s", e.endpoint, e.collection, geom_key)

		geom_rsp, err := e.client.Do("GET", e.collection, geom_key, "WITHFIELDS")

		t2 = time.Since(t1)

		if err != nil {
			return record.NewErrorRecord("tile38", id, uri, err, t2)
		}

		if !geom_rsp.(tile38.Tile38Response).Ok {
			continue
		}

		rsp := make(map[string]interface{})

		rsp[geom_key] = map[string]interface{}{
			"object": geom_rsp.(tile38.Tile38Response).Object,
			"fields": geom_rsp.(tile38.Tile38Response).Fields,
		}

		meta_rsp, err := e.client.Do("GET", e.collection, meta_key)

		if meta_rsp.(tile38.Tile38Response).Ok {

			var meta interface{}
			obj := meta_rsp.(tile38.Tile38Response).Object

			err = json.Unmarshal([]byte(obj.(string)), &meta)

			if err != nil {
				rsp[meta_key] = obj
			} else {
				rsp[meta_key] = meta
			}
		}

		t2 = time.Since(t1)

		return record.NewDefaultRecord("tile38", "tile38", id, uri, rsp, t2)
	}

	t2 = time.Since(t1)

	msg := fmt.Sprintf("can't find tile38 record for ID %d", id)
	err := errors.New(msg)

	uri := fmt.Sprintf("x-tile38-%d", id)
	return record.NewErrorRecord("tile38", id, uri, err, t2)
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
