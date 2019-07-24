package record

import (
	"encoding/json"
	"github.com/tidwall/gjson"
)

func LastModified(body interface{}) int64 {

	enc_body, err := json.Marshal(body)

	if err != nil {
		return -1
	}

	lastmod_rsp := gjson.GetBytes(enc_body, "properties.wof:lastmodified")

	if !lastmod_rsp.Exists() {
		return -1
	}

	return lastmod_rsp.Int()
}
