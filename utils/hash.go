package utils

import (
	"encoding/json"
	"github.com/whosonfirst/go-whosonfirst-hash"
)

func HashInterface(i interface{}) (string, error) {

	body, err := json.Marshal(i)

	h, err := hash.NewWOFHash()

	if err != nil {
		return "", err
	}

	return h.HashBytes(body)
}
