package index

import (
	"errors"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-catalog/record"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"os"
	"path/filepath"
)

type FSIndex struct {
	catalog.Index
	root  string
	repos []string
}

func NewFSIndex(root string, repos ...string) (catalog.Index, error) {

	info, err := os.Stat(root)

	if os.IsNotExist(err) {
		return nil, err
	}

	if !info.IsDir() {
		return nil, errors.New("root is not a directory")
	}

	e := FSIndex{
		root:  root,
		repos: repos,
	}

	return &e, nil
}

func (e *FSIndex) GetById(id int64) (catalog.Record, error) {

	for _, repo := range e.repos {

		root := filepath.Join(e.root, repo)

		_, err := os.Stat(root)

		if os.IsNotExist(err) {
			continue
		}

		data := filepath.Join(root, "data")

		uri, err := uri.Id2AbsPath(data, id)

		if err != nil {
			return nil, err
		}

		_, err = os.Stat(uri)

		if os.IsNotExist(err) {
			continue
		}

		return nil, errors.New("Please finish me")
	}

	msg := fmt.Sprintf("can't find filesystem record for ID %d", id)
	err := errors.New(msg)

	uri := fmt.Sprintf("x-fs-%d", id)
	return record.NewErrorRecord("github", id, uri, err)

}
