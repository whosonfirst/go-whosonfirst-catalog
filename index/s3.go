package index

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-inspector"
	"github.com/whosonfirst/go-whosonfirst-inspector/record"
	"github.com/whosonfirst/go-whosonfirst-inspector/utils"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"time"
)

type S3Index struct {
	catalog.Index
	bucket string
}

func NewS3Index(bucket string) (catalog.Index, error) {

	e := S3Index{
		bucket: bucket,
	}

	return &e, nil
}

func (e *S3Index) GetById(id int64) (catalog.Record, error) {

	root := fmt.Sprintf("https://s3.amazonaws.com/%s/", e.bucket)

	uri, err := uri.Id2AbsPath(root, id)

	if err != nil {
		return nil, err
	}

	t1 := time.Now()

	rsp, err := utils.GetURLAsJSON(uri)

	t2 := time.Since(t1)

	if err != nil {
		return record.NewErrorRecord("s3", id, uri, err, t2)
	}

	return record.NewDefaultRecord("geojson", "s3", id, uri, rsp, t2)
}
