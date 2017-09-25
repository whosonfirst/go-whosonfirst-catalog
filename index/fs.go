package index

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-catalog/record"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
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

	t1 := time.Now()
	var t2 time.Duration

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

		fh, err := os.Open(uri)

		t2 = time.Since(t1)

		if err != nil {
			return record.NewErrorRecord("fs", id, uri, err, t2)
		}

		body, err := ioutil.ReadAll(fh)

		t2 = time.Since(t1)

		if err != nil {
			return record.NewErrorRecord("fs", id, uri, err, t2)
		}

		var stub interface{}

		err = json.Unmarshal(body, &stub)

		t2 = time.Since(t1)

		if err != nil {
			return record.NewErrorRecord("fs", id, uri, err, t2)
		}

		return record.NewDefaultRecord("fs", "fs", id, uri, stub, t2)
	}

	t2 = time.Since(t1)

	msg := fmt.Sprintf("can't find filesystem record for ID %d", id)
	err := errors.New(msg)

	uri := fmt.Sprintf("x-fs-%d", id)
	return record.NewErrorRecord("fs", id, uri, err, t2)
}
