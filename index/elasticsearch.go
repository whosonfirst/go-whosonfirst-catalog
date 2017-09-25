package index

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-catalog/record"
	"github.com/whosonfirst/go-whosonfirst-catalog/utils"
	"net/url"
	"strconv"
	"time"
)

type ElasticsearchIndex struct {
	catalog.Index
	endpoint string
	index    string
	timing	 time.Duration
}

func (e *ElasticsearchIndex) GetById(id int64) (catalog.Record, error) {

	str_id := strconv.FormatInt(id, 10)

	path := fmt.Sprintf("%s/_search", e.index)

	query := url.Values{}
	query.Set("wof:id", str_id)

	url := &url.URL{
		RawQuery: query.Encode(),
		Host:     e.endpoint,
		Path:     path,
		Scheme:   "http",
	}

	uri := url.String()

	t1 := time.Now()

	rsp, err := utils.GetURLAsJSON(uri)

	t2 := time.Since(t1)

	if err != nil {
		return record.NewErrorRecord("elasticsearch", id, uri, err, t2)
	}

	return record.NewDefaultRecord("elasticsearch", "elasticsearch", id, uri, rsp, t2)
}

func NewElasticsearchIndex(host string, port int, idx string) (catalog.Index, error) {

	endpoint := fmt.Sprintf("%s:%d", host, port)

	e := ElasticsearchIndex{
		endpoint: endpoint,
		index:    idx,
	}

	return &e, nil
}
