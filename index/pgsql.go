package index

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-catalog/record"
	"github.com/whosonfirst/go-whosonfirst-pgis/client"
)

type PgSQLIndex struct {
	catalog.Index
	endpoint string
	db       string
	client   *pgis.PgisClient
}

func (e *PgSQLIndex) GetById(id int64) (catalog.Record, error) {

	uri := fmt.Sprintf("postgis://%s/%s#%d", e.endpoint, e.db, id)

	row, err := e.client.GetById(id)

	if err != nil {
		return record.NewErrorRecord("postgis", id, uri, err)
	}

	return record.NewDefaultRecord("postgis", "postgis", id, uri, row)
}

func NewPgSQLIndex(host string, port int, user string, password string, dbname string, maxconns int) (catalog.Index, error) {

	client, err := pgis.NewPgisClient(host, port, user, password, dbname, maxconns)

	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s:%d", host, port)

	e := PgSQLIndex{
		client:   client,
		endpoint: endpoint,
		db:       dbname,
	}

	return &e, nil
}