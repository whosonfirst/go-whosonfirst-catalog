package index

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-catalog/record"
	"github.com/whosonfirst/go-whosonfirst-catalog/utils"
	"net/url"
	"strconv"
)

type ElasticsearchIndex struct {
	catalog.Index
	endpoint string
	index    string
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

	str_url := url.String()

	rsp, err := utils.GetURLAsJSON(str_url)

	if err != nil {
		return nil, err
	}

	return record.NewDefaultRecord("elasticsearch", id, str_url, rsp)
}

func NewElasticsearchIndex(host string, port int, idx string) (catalog.Index, error) {

	endpoint := fmt.Sprintf("%s:%d", host, port)

	e := ElasticsearchIndex{
		endpoint: endpoint,
		index:    idx,
	}

	return &e, nil
}
