package index

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-catalog/record"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"io/ioutil"
	"net/http"
)

type S3Index struct {
	catalog.Index
	bucket string
}

func NewS3Index() (catalog.Index, error) {

	e := S3Index{
		bucket: "whosonfirst.mapzen.com",
	}

	return &e, nil
}

func (e *S3Index) GetById(id int64) (catalog.Record, error) {

	root := fmt.Sprintf("http://%s.s3.amazonaws.com/data", e.bucket)

	url, err := uri.Id2AbsPath(root, id)

	if err != nil {
		return nil, err
	}

	rsp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return nil, err
	}

	return record.NewDefaultRecord("s3", id, url, body)
}
