package index

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-catalog/record"
	"github.com/whosonfirst/go-whosonfirst-catalog/utils"	
	"github.com/whosonfirst/go-whosonfirst-uri"
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

     	root := fmt.Sprintf("https://s3.amazonaws.com/%s/data/", e.bucket)

	url, err := uri.Id2AbsPath(root, id)

	if err != nil {
		return nil, err
	}

	body, err := utils.GetURL(url)
	
	if err != nil {
		return nil, err
	}

	return record.NewDefaultRecord("s3", id, url, body)
}
