package index

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"io/ioutil"
	"net/http"
)

type S3Index struct {
	catalog.Index
	bucket string
}

type S3Record struct {
	catalog.Record `json:",omitempty"`
	id             int64       `json:"id"`
	uri            string      `json:"uri"`
	address        string      `json:"address"`
	body           interface{} `json:"body"`
}

func (r *S3Record) Id() int64 {
	return r.id
}

func (r *S3Record) Source() string {
	return r.source
}

func (r *S3Record) URI() string {
	return r.uri
}

func (r *S3Record) Body() interface{} {
	return r.body
}

func NewS3Index() (catalog.Index, error) {

	e := S3Index{
		bucket: "whosonfirst.mapzen.com",
	}

	return &e, nil
}

func (e *S3Index) GetById(id int64) (catalog.Record, error) {

	root := fmt.Sprintf("http://%s.s3.amazonaws.com/data", e.bucket)

	abs_root, err := uri.IdToAbsPath(root, id)

	if err != nil {
		return nil, err
	}

	rsp, err := http.Get(abs_root)

	if err != nil {
		return nil, err
	}

	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp)

	if err != nil {
		return nil, err
	}

	r := S3Record{
		id:       id,
		"source": "S3",
		"uri":    url,
		"body":   body,
	}

	return &r, nil
}
