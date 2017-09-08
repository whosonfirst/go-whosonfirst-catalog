package utils

import (
       "encoding/json"
       "errors"
	"io/ioutil"
	"net/http"
)

func GetURL(url string) ([]byte, error) {

	client := &http.Client{}
	
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	rsp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer rsp.Body.Close()

	if rsp.StatusCode != 200 {
		return nil, errors.New(rsp.Status)
	}
	
	body, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func GetURLAsJSON(url string) (interface{}, error) {

	body, err := GetURL(url)

	if err != nil {
		return nil, err
	}

	var rsp interface{}
	
	err = json.Unmarshal(body, &rsp)

	if err != nil {
		return nil, err
	}

	return rsp, nil
}
