// Code generated by go-bindata.
// sources:
// www/index.html
// www/javascript/mapzen.whosonfirst.geojson.js
// www/javascript/mapzen.whosonfirst.inspector.init.js
// www/javascript/mapzen.whosonfirst.inspector.js
// www/javascript/mapzen.whosonfirst.render.js
// www/css/mapzen.whosonfirst.inspector.css
// www/css/mapzen.whosonfirst.render.css
// DO NOT EDIT!

package http

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/elazarl/go-bindata-assetfs"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _wwwIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x94\xcf\x92\x9b\x30\x0c\xc6\xcf\x9b\xa7\xd0\xe8\xd2\x5b\xfc\x02\x36\xa7\x76\x67\xf6\xb4\xc7\x1c\x3b\x0e\x16\xc5\x89\xb1\xa9\x25\x48\xe9\xd3\x77\xc0\xe4\xef\xf4\x96\x0d\x27\xf3\x49\xf3\xd3\x27\x21\xac\x5b\xe9\x42\xb5\x01\xd0\x2d\x59\x37\x1f\x00\xb4\x78\x09\x54\xed\xda\xf4\x8d\xe1\x33\xc2\xbb\xcf\x2c\xf0\x11\xb9\xa7\x5a\x52\xd6\xaa\xc4\x4b\x6e\xf0\xf1\x08\x32\xf5\x64\x50\xe8\x8f\xa8\x9a\x19\x21\x53\x30\xc8\x32\x05\xe2\x96\x48\x10\xda\x4c\x8d\xc1\x9a\x59\x75\xb6\xff\x4b\x71\x7b\xe0\xed\x92\xa9\x9e\xc3\x9c\xda\xc4\x29\x36\xb3\xc1\xad\x3f\x1b\xfc\x72\x72\xa6\xe8\xe8\x01\xcb\x75\xf6\xbd\xdc\x82\x0f\x76\xb4\x45\x45\xe0\x5c\x1b\xbc\x0a\x67\x68\xe7\xe7\xce\xb1\xd2\xaa\xe8\xcf\xb0\x6e\x0d\xfe\xa2\x74\xe0\xf4\x1a\xf6\xda\xfc\x2b\xd0\xd7\x2f\xf6\x5a\xba\x8f\x5e\xee\x4b\xc0\xfa\xcc\x8b\xaf\xce\x9b\xaf\xf7\xc9\x4d\x6b\x75\xe7\x47\xf0\xce\x60\x6d\xe3\x68\x19\x8b\x7a\xd1\xc1\xe0\xef\x81\xf2\x84\xd5\xe6\x4d\x37\x29\x77\xd5\xe6\x0d\x40\xfb\xd8\x0f\xb7\x8e\x71\x41\x9c\x52\xf3\xd3\x3b\x84\x68\x3b\xba\xbe\x8d\x36\x0c\x64\x10\xa1\x0f\xb6\xa6\x36\x05\x47\xd9\xe0\x8f\x28\x94\xc1\xc2\xee\xf3\x1d\x3e\xbe\x2f\xdb\x36\x83\xf7\x83\x48\x8a\x2b\x99\x87\x7d\xe7\x57\x76\x48\xe9\x38\xf4\x58\xad\x3f\xa7\x56\x25\x73\xb6\xa5\x8a\xaf\xd5\xb7\x72\x7e\x7c\x68\xc2\xa0\x23\xb1\x3e\x2c\x73\xf9\x5f\x38\x13\x0f\x41\xee\xc3\x97\xa3\x56\x65\x5a\x5a\x95\x1b\xe4\x5f\x00\x00\x00\xff\xff\xab\x2a\xf8\x40\x49\x04\x00\x00")

func wwwIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_wwwIndexHtml,
		"www/index.html",
	)
}

func wwwIndexHtml() (*asset, error) {
	bytes, err := wwwIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/index.html", size: 1097, mode: os.FileMode(420), modTime: time.Unix(1563546049, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwJavascriptMapzenWhosonfirstGeojsonJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x94\xcf\x6e\x9b\x40\x10\xc6\xcf\xf8\x29\xa6\x27\x83\x82\xb0\x9d\xaa\xad\x14\x42\xa4\xaa\x52\x6f\x7d\x02\x84\xa2\x8d\x19\x60\x5b\x98\x45\xb0\x24\x4e\x1d\xbf\x7b\xb5\xff\x00\x1b\x5b\xbd\xb1\xf3\x7d\xfb\xdb\x99\xd9\x59\x5e\x59\x07\x0d\x6b\xff\x22\x41\xe2\x3e\x3e\x3e\xe0\x78\x8a\x57\x66\x15\xbd\x55\xa2\x17\x54\xf0\xae\x97\xa3\xe5\x2c\x68\xed\x57\xfc\x51\x89\xe2\x77\x2f\x14\xda\x2f\x06\xda\x4b\x2e\xc8\x0f\x8e\xab\x95\xa7\x8e\xed\xb1\x2e\x20\x01\xb5\xf4\xd6\x39\x76\xfc\x15\x9f\x5f\x5e\xc4\x61\xfd\x00\xa3\xd9\x02\xf4\x1e\xcf\xe3\x05\xb8\x48\xba\xd6\xd6\x2c\x38\x2a\xc1\xeb\x50\x0e\x1d\xc1\x85\x18\x2b\xed\xb4\xdc\x2a\xdf\x5b\x5c\x67\x90\x24\xb0\xfe\x89\x4c\x0e\x1d\xfe\x10\x75\x8d\xfa\xc8\xb5\x3d\x4b\xa7\x58\x18\xb5\x87\x64\x42\xbb\x98\xc5\x6b\xdf\x5e\x0c\xa4\xba\xe3\xb4\xa8\x46\x2a\x65\x15\x4f\xa4\xfe\xad\x66\xca\x31\x50\x8e\x05\x27\xcc\xe3\xb9\xa4\x5b\x74\x45\x22\xbc\xb9\x8b\x70\xb1\x4b\x6b\x85\xe8\xc0\x57\x06\x9e\x6c\x63\xe0\xf0\x68\x92\x8b\x81\xdf\xdd\xb9\xd2\x34\x40\xb5\x08\x12\x7d\x0b\xd1\xac\xfb\xbe\xab\x21\xe5\x59\x10\xcf\xfc\xcf\xae\x04\xe5\x4a\xb7\xb6\xfa\x51\xd2\xc9\x68\x69\x77\x26\xb9\x12\xb4\x74\x7f\x29\x4d\xbb\x3e\x67\xee\x30\x75\x55\xfe\x27\xd3\xb1\x40\x4d\x97\x6f\x8f\x7e\xb4\x31\x7b\xe7\x9e\xe7\x32\x32\xba\x45\x9f\x16\x18\x41\x13\x46\x90\xc1\x08\x3a\xc3\xe8\x3c\x8c\x7e\x1d\xa3\xcb\xb0\x18\x53\xd2\x93\x8d\x8d\x18\x57\xa9\xd1\x6f\x62\xc6\x6c\x4c\xf9\x4f\x36\x36\xc7\x98\x6c\xf4\xc7\x88\x99\xc1\xec\xb0\xa7\xa6\x1d\xa1\x29\x27\x34\xe9\x84\x76\x30\xe6\xb3\x8f\x75\x8f\xf0\x9f\x07\x30\x8e\xfd\x66\x03\xdf\x73\xd6\x4a\xcc\xa1\xe8\x44\x03\x95\x94\xed\xc3\x66\x53\xf2\x3e\xea\x25\xdb\xff\xc1\xc3\xbe\x62\x54\x62\xb4\x17\xcd\x86\x6d\x76\xdf\xee\xbf\x7c\xdd\x8d\x53\x59\xa2\x68\xe6\x8f\x45\xad\x51\x76\xef\xe7\x8f\x45\x74\xb9\x7d\x52\x4d\xa4\x57\x9c\x98\xc4\x7e\xf6\x5c\x6a\x26\x95\x23\xcd\x42\x1d\x02\x00\xa8\xa9\x34\xa1\xc5\xa0\x43\x02\xe3\xa8\x2b\x74\xba\xcd\xdc\x03\xd4\x53\x0f\xb6\xb5\x0a\x1a\xb5\x43\x5f\xf9\xa3\x2f\xe5\x59\xba\x53\x73\x6e\x0c\x54\x5e\x33\x6c\x9d\xe1\x34\x25\xd8\x70\x32\xb7\xfd\x8b\xc9\x2a\x6a\x38\x45\xac\x6d\xeb\x77\x9f\x86\xba\x0e\x75\xfa\xc1\x94\x7a\xc3\x0e\x73\x33\x3b\x2c\xcd\xf1\x19\x99\xca\x9b\x64\x2a\x2f\xc9\x33\xf3\x25\x59\x99\xe3\x8b\xa1\x31\x99\x87\xf6\x1c\x83\xf2\x6c\x8a\xa1\x03\x2e\xc6\xe7\xa8\x26\x50\xad\xd5\x9f\xde\xc1\xd4\xbf\x23\x5e\xad\x4e\x81\x1f\xc4\xab\x7f\x01\x00\x00\xff\xff\x2a\x82\x1a\x4f\x4a\x06\x00\x00")

func wwwJavascriptMapzenWhosonfirstGeojsonJsBytes() ([]byte, error) {
	return bindataRead(
		_wwwJavascriptMapzenWhosonfirstGeojsonJs,
		"www/javascript/mapzen.whosonfirst.geojson.js",
	)
}

func wwwJavascriptMapzenWhosonfirstGeojsonJs() (*asset, error) {
	bytes, err := wwwJavascriptMapzenWhosonfirstGeojsonJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/javascript/mapzen.whosonfirst.geojson.js", size: 1610, mode: os.FileMode(420), modTime: time.Unix(1563546049, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwJavascriptMapzenWhosonfirstInspectorInitJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\x41\x8b\xdb\x30\x10\x85\xcf\xd6\xaf\x18\x74\x52\x20\x88\xde\x5d\x1f\xda\x12\x48\x20\xbd\x05\x7a\x2c\x5a\x69\x9c\x88\x28\x23\x23\xc9\x31\xd9\xe0\xff\xbe\x48\x31\xde\xec\x12\x7c\x91\xc7\xe3\xf7\xbd\xf1\x3c\x0d\x96\x8c\x1f\xa4\x32\x66\x73\x45\x4a\x7b\x1b\x13\x12\x06\xc1\x9d\x57\x86\xaf\xa1\xed\x49\x27\xeb\x09\xf2\xbb\xc0\xac\x59\xdd\x59\xc5\xaa\xab\x0a\xf0\xd6\xa7\xe4\x09\x1a\x30\x5e\xf7\x17\xa4\x24\x8f\x98\x36\x0e\x73\xf9\xfb\xb6\x33\xd9\xc5\x9f\xfb\x8e\xaf\xea\x8c\x3c\xe4\xd2\x93\x76\x56\x9f\xa1\x99\xcd\xc5\xea\xce\x58\x05\x00\x90\xc2\x0d\xee\x2c\x57\x55\x99\x60\xcd\x92\xfb\xe0\xdb\xff\xd6\x64\x77\x56\x55\x45\x6a\x8d\xbc\x2a\xd7\x63\xfd\xd4\x48\xc1\x5e\xc4\xa4\x69\x41\xe4\x76\x03\x9c\xe7\x35\xca\xcc\x80\xa9\x0f\x04\xad\x72\xb1\x70\xe3\xec\xd6\xa9\x10\x71\x47\x49\x58\xf3\x85\xff\x09\x3f\x96\xe9\x8b\xea\xde\x91\xe4\x70\xf2\xd1\x53\x6b\x43\x4c\xd2\x52\xec\x50\x27\x1f\xe4\x23\x13\x61\xcd\x67\xba\x22\xc4\x6e\x76\x9c\x1e\x79\xfd\x00\x0d\x2c\x7a\x05\x24\x83\xa1\xe0\xf5\x0b\x1c\x63\xef\x52\x5c\x8a\x70\x92\xf0\x19\x9f\x1a\xd2\x12\x61\xd8\x1e\xfe\xee\x21\x87\x55\xb6\x7f\xfe\xac\xba\x0e\xc9\xfc\x39\x59\x67\x44\x28\xf0\x98\xcf\x2c\x19\xd9\xf4\x1f\xf9\xd4\x2a\xe9\x13\x08\x2c\xdb\x69\x4f\xd1\x3b\x94\xce\x1f\x05\xff\xb7\xfd\x75\xe0\x6b\xc0\xef\x58\xb9\xfc\x17\xc9\x8e\x35\xcb\x23\x3e\x02\x00\x00\xff\xff\x40\x1b\x11\x0b\xb1\x02\x00\x00")

func wwwJavascriptMapzenWhosonfirstInspectorInitJsBytes() ([]byte, error) {
	return bindataRead(
		_wwwJavascriptMapzenWhosonfirstInspectorInitJs,
		"www/javascript/mapzen.whosonfirst.inspector.init.js",
	)
}

func wwwJavascriptMapzenWhosonfirstInspectorInitJs() (*asset, error) {
	bytes, err := wwwJavascriptMapzenWhosonfirstInspectorInitJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/javascript/mapzen.whosonfirst.inspector.init.js", size: 689, mode: os.FileMode(420), modTime: time.Unix(1563546049, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwJavascriptMapzenWhosonfirstInspectorJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\x5b\x73\xdb\xb8\x15\x7e\x96\x7e\x05\x16\x9d\x59\x93\x63\x99\xce\xa5\x0f\x1d\xab\x7a\x48\x76\x9d\xbd\xd4\x4e\x3a\xb1\xdb\xd9\x19\x8e\xaa\x81\xc9\x23\x09\x6b\x0a\x60\x00\x28\x8a\xea\xf5\x7f\xef\xe0\x0e\x4a\xb4\x2d\xb5\x2f\xcd\x43\x6c\x02\xdf\xb9\xe0\x9c\x0f\xe7\x00\xf0\x57\x22\xd0\x8a\xb4\xff\x06\x86\x26\xfe\x97\x3f\xfe\x40\x0f\x8f\xe3\xa1\xfd\x2a\x36\x4b\x2e\x39\x9b\x53\x21\x55\x80\x74\x06\x1d\xbc\x07\x5f\x50\x26\x5b\xa8\x14\x17\x68\x82\xb2\xf9\x9a\x55\x8a\x72\x96\xe5\x0f\xc3\x21\x42\x08\x69\xdb\x15\xa9\x96\x80\x26\x56\x83\x1f\x94\xd0\xcc\xf5\xd8\x70\x30\x1c\x9c\x34\x9c\xdf\xaf\xdb\x93\x0b\x14\xe4\x69\x3d\x42\xd5\x5d\xfe\x30\x1c\x68\x01\xfb\xbf\x16\x6b\x05\x57\x1c\x4d\x50\xcd\xab\xf5\x0a\x98\x2a\x1a\x5e\x11\x2d\x51\x98\x99\x8a\x37\xe3\x08\x5e\x8b\x06\x4d\x9c\xc8\x29\x3a\x39\x3f\x3f\x41\xa7\x28\x08\x2c\xb9\x54\xe9\x77\x4b\xd4\x92\x91\x15\x68\x28\xad\x0d\x96\xd6\xe3\x5d\x07\x38\x6b\x38\xa9\xd1\x24\xba\x2a\x64\xab\xfd\x1c\x0c\x07\x03\x0d\x50\x44\x2c\x40\x87\x51\xc8\xb6\xb0\x1f\x63\x3b\x4b\xe7\x28\xb3\x03\x85\x00\x52\x6f\x6f\x14\x51\x80\xbe\x9b\xa0\x3f\x1b\x79\x6d\x41\x80\x5a\x0b\xa6\xf1\x8f\x51\xa3\x54\x44\xad\xe5\xac\xe2\xb5\x8e\xa2\xd5\x50\x9e\xd8\xd1\x93\xa9\x06\x77\x70\x0a\xbe\xa9\x3d\xdc\x2d\x7c\x53\x16\xeb\x94\x0a\xb2\x49\x40\x02\x64\xcb\x99\x84\x08\xd3\x98\x9a\x28\x82\x26\x88\xad\x9b\xc6\x49\x2a\xb1\x45\xde\x57\x37\xfb\xeb\xcd\xa7\x8f\x45\x4b\x84\x84\x4c\x90\x4d\x9e\xf8\x5e\x11\x55\x2d\x51\x06\x61\x75\x15\x67\x92\x37\x50\x34\x7c\x91\xe1\xcb\xcf\x9f\xf1\x48\xa7\x68\x84\x20\xef\x2c\x1f\xcd\x49\x23\x21\x55\x74\x97\x69\x63\xb9\x4b\xc6\xe3\x5e\x52\x04\x7c\xd1\x7e\xc2\x06\xfd\x76\x7d\xf5\xb3\x52\xed\x67\xf8\xb2\x06\xa9\x32\x2f\x22\xe0\x4b\x41\xea\xfa\xf2\x2b\x30\x75\x45\xa5\x02\x06\x22\xc3\x3a\x93\x78\xe4\x52\x9a\x77\xb4\x6a\x01\xde\x02\xcb\xf0\x4f\x97\xb7\xde\x4f\x25\xd6\x90\x6a\x94\xc0\x6a\x63\xe2\x71\x64\x68\x2c\x80\xd5\x20\x52\x1a\x3b\x6e\xec\x7a\x5b\x71\x51\x4b\x4b\x91\x12\xbb\x4f\x50\x78\x1a\x3e\xf0\x34\xa1\x71\xc5\xd7\xcc\x10\xca\xce\x15\x0d\xb0\x85\x5a\xee\x05\xe1\x6e\x3b\x53\xdb\xd6\x6f\xb4\x30\xac\xc7\xb4\xad\x32\x55\x29\xf9\x5a\x54\xdd\x61\xfb\xff\x9c\x0b\x94\x69\x04\x9d\xbc\x1a\x23\x8a\xfe\x6a\xad\x8f\x11\x3d\x3d\x4d\x59\x2e\xa2\x3f\x25\x0d\x84\x31\x5e\x96\x58\x9b\xc4\xd3\x84\xf5\xdf\x79\xe7\x4a\x35\x0d\x6c\x70\x3f\xe2\x8c\xf6\x06\x09\x64\x04\xf5\x8c\xf1\xbc\x68\xd7\x72\x99\xa9\x94\x57\xd0\x48\x08\x2c\x8c\xe2\x16\x29\x3c\xd2\x30\x25\x5d\x9a\x55\x27\xb9\x88\xb4\x48\xa2\x44\xee\x1a\x48\xeb\x4a\x25\x80\x28\xb8\x6c\x40\x7f\x65\xd8\xcc\xe3\x7d\x39\x1b\xc8\xd9\x12\x48\x0d\xe2\x39\xf9\x65\x10\xee\x88\x14\xa4\x6d\x81\xd5\x3f\x2c\x69\x53\x67\x3b\xc2\x7a\x33\x7e\xe4\x35\x64\xd8\x8a\xe0\xbc\xc7\xef\x6d\x7b\x9c\xf5\x44\xe0\x30\xdb\x26\x99\x3d\x96\x97\x44\x2e\x8f\xb2\x9c\x08\x1c\x66\x59\x0b\xf4\x59\x5e\x0b\x7a\x94\xe1\x88\x3f\xcc\xee\x5a\xd0\xde\x50\xd3\x15\x65\x8b\xe3\x82\x9d\x8a\x1c\x18\x6e\x23\xd2\x67\x5f\x2e\xf9\xe6\x38\xa2\x45\x81\xc3\x6c\x9f\xf5\x99\xad\x41\x11\xda\xc8\xa3\x2c\x77\x65\xfe\x07\xe3\x56\xc1\x4c\xf0\xcd\x73\x86\x45\xe4\x58\xc0\x77\x8c\x76\xb6\xdc\x0b\xd8\x64\x83\xbc\x80\x4c\x08\xfd\x02\x32\x32\xf0\x25\xe3\x29\x61\x5e\xc0\x76\x83\xdc\x8d\x9d\x29\x57\x5d\x67\x83\x92\xbc\xbf\xe0\x2b\x44\x99\xad\x91\xbe\xce\x9b\x53\x84\x1e\x28\x55\xac\xf0\x3a\x3a\xb1\x81\xc5\xea\x9b\x1c\x2d\x4c\xc7\x48\x40\xa9\x4c\x6c\x5e\x1e\x7d\xb7\x9d\xd9\xf4\xf8\xc6\x35\xe8\x69\x50\x1a\xfc\x44\x6b\xf2\x9a\x63\x8b\x4a\x9a\x4b\x72\x8a\x49\xbd\x30\x2d\x2b\x81\x04\x07\x34\xb6\xf4\xe5\x76\x3a\xd6\x27\xab\x04\xe6\xda\xab\x05\xc5\x26\x97\xd4\xa5\x30\xab\x6b\x48\x67\xd2\x66\x36\x4a\xdb\x6d\x3e\x45\xe7\xe8\xf5\x2b\xff\x6f\x3c\x38\x3f\x47\x8c\x30\x2e\xa1\xe2\xac\x96\x56\x3c\x51\xa2\x19\x17\x54\x98\xf2\x68\x7c\xec\xe0\x3a\x87\xac\x1f\xdf\xdd\xbe\x33\xa7\x17\x3a\x32\xc2\xf9\x18\xed\xaf\x7b\x56\x41\xd3\xbc\xbc\xb1\xd3\xe6\xa5\x25\x0e\xda\xd6\x16\x9f\x07\xf9\x9d\x68\xbe\x68\xba\x8e\xa6\x03\xfe\x20\xc3\x1a\xdd\x6b\xd6\xec\xda\x23\xcc\x06\xfc\x41\x66\x4d\x90\xf3\x0e\xb9\xf4\xe6\x3f\xc2\xde\xf9\x79\x90\x28\x24\xa8\x77\x4a\x09\x7a\xb7\x56\x90\x61\x9d\xf6\x33\x01\x2b\xae\xe0\x4c\xf3\xcb\xe4\x35\x1f\x6b\x31\x27\x1b\x04\x0f\x71\x55\x0b\xf7\xe6\xc5\xd6\xa0\x63\x32\x13\x25\x0e\xcb\x8d\xc1\x17\x8a\x7f\xa0\xdf\xa0\xce\xde\xe6\xe8\x14\x61\x89\x77\xe2\xe6\xeb\xdb\x11\x8e\xa4\x22\x3b\xc1\xab\x1a\x22\x25\x1e\x21\x5c\x35\xb4\xba\x3f\x5b\xc1\x41\x52\x2f\x87\xbc\x23\x7c\x50\xa3\x73\x12\xbb\xcb\x5d\x81\x22\x07\x35\x3a\x27\x63\x6e\x93\xa6\x20\x4d\x10\x06\x21\xb8\xc0\xa6\xfa\x0d\xbc\xa2\xa7\x02\xe0\xb0\x7e\xf5\x8f\x5e\x61\x90\xeb\xe9\x9c\x7a\x75\x41\xa2\x17\x18\x76\xe7\xf3\xb0\xb0\x9b\x9e\x87\x79\x22\xbf\x60\x33\xf2\xee\x79\x60\x9a\xa4\x5d\xca\x77\x12\xc8\x99\xa1\x47\x7a\xbd\xb7\xd7\x57\x5d\x63\x4d\x63\x02\xcd\x45\x48\x2e\xf6\x83\x58\xfa\xa1\x29\x16\xcf\xd2\x27\x77\xad\x6c\x20\xa1\x99\x17\xe6\x7c\xe6\xcc\x9b\xdd\x18\x52\x12\x92\x9c\xf6\xf1\xcc\x2f\x2d\x72\xc0\xbc\xac\x94\x6b\x41\xa7\xae\x2b\x8c\xed\xba\xfa\xae\x3f\xee\x6e\x6d\x34\xda\x2b\xeb\x70\x70\xb2\xa4\x35\x78\x17\xd2\x7b\x6b\xfe\x30\xdc\x3b\x01\xa6\xc4\x5c\x80\x72\xac\x7c\xbf\xfd\xa5\x4e\x38\xdd\x3d\x00\x16\x94\x31\x10\x3f\xdf\x5e\x5f\xa1\x09\xc2\x38\xde\x94\xd3\xa5\xa7\x76\x75\x14\xbc\x69\x13\xa3\xd4\xc1\xcc\x2c\x7c\xa7\xad\xc7\x10\x24\xd7\xdb\x3b\x5e\x6f\x43\x9f\xd4\x1f\x9d\xfb\x74\x6f\x1b\x8f\xd3\x1b\xa1\xe3\xfd\xdc\x49\xb7\xa6\x5f\xc3\x4a\x1d\x7a\x67\xb3\xd1\x5a\xef\x34\x7f\x6d\x4b\xb5\xcf\x81\xa8\xb5\xd0\xf6\xcd\xfa\xb4\x13\x33\xc5\x67\x6e\xb8\xf3\xce\xe1\xde\xbb\x5a\x1d\x79\x37\x5f\x62\x3d\x00\x42\x51\x90\xd1\x6b\x0f\xea\x79\x9d\xb3\x4f\x12\xee\xc7\x4c\x6b\xcf\x0c\x7a\xd7\xfd\x74\xb7\x78\x40\xf4\x62\x45\xda\x03\xc3\xb1\x22\x6d\x6f\x28\x56\xa4\xed\xc6\xe1\xb8\x0b\xb7\x91\x78\xf6\xf2\x21\xba\xea\x1b\x98\xab\x03\x3b\x47\x78\x96\xa1\x8b\xe5\x81\x32\x56\x28\xd8\x08\x1b\x94\xb4\xe1\x79\x28\x28\xf3\x93\x2e\xd2\x41\x5a\x75\x83\x1e\x94\x85\xbb\x63\x77\x3e\x2a\x8c\x1a\xf6\x0e\xf9\x4a\x74\xb7\xc8\x7f\xbf\x71\x3b\x5a\xb5\x9d\xa0\xd8\xd2\x56\x90\xcd\x6c\x45\xda\xcc\xf1\x32\x0f\x15\xc5\xcf\xa4\xbb\xda\x83\xd2\xa2\x72\x10\xb1\xed\xdb\x91\xe5\xa3\xae\xc2\xe1\x79\xd4\x37\xae\xb0\xe1\xef\xf8\xb7\xfe\x0d\xb0\x00\xfe\xbb\xe4\xac\xa8\x41\xd0\xaf\x30\xd3\xc0\xc4\xe9\xe4\x76\xad\x1b\xef\x55\xd1\x10\x75\xc5\x16\x99\x86\x95\xaf\xa7\x23\xa3\xb8\x7c\x35\x4d\x79\xc2\x60\x0f\xf9\xd6\x23\xdf\x4c\x83\xd2\xce\x39\xfc\xfd\xa7\x7f\x7c\xfc\xf1\x06\x8f\x90\xdc\x8c\x10\x0b\x8f\x88\x1d\xcc\xed\xe5\xcd\x2d\x1e\xa1\x4c\x7b\x32\xd1\xa0\xdc\x3d\x97\xfb\x7f\x57\xc5\xb5\x5d\x1f\x69\xe9\xdf\x60\x9b\xa6\x55\xd7\xb8\xbe\xde\x63\x03\x72\x46\x5a\x7a\x76\x0f\xdb\xf8\x7e\x95\xaa\xb5\x9b\x3b\x28\xd7\x59\x3d\xd1\x09\xcc\xd3\x2c\x64\x72\x53\xe2\x86\x28\x3c\xb5\xce\xb9\x8f\x1c\x7d\xff\x3d\xb2\x73\xe6\x36\xe3\xe7\xf4\x47\x6e\x52\xe6\x0a\xc2\x3f\x29\x6c\x32\xbd\xf8\xd7\x6f\xc2\x33\x6e\xda\x9d\xc2\x9b\x9e\xc6\xcf\xa9\x7a\xcf\xd7\xac\x96\x59\x69\xe3\x35\xcd\x7b\x72\xde\x90\xad\xa9\xd2\x57\x3a\xc7\xbf\xde\x7c\xfa\x98\x92\xd1\xec\x4f\x0d\x28\x48\x5d\xdf\x72\xbf\x35\x1d\x47\xbb\x45\x37\xa5\xaa\xa9\xbf\x7b\x2f\xb6\xcf\x74\x8c\xce\x31\xcc\x91\x0d\x27\x64\xed\x6b\x42\x9d\x95\x77\x14\xb4\x5c\xaa\x05\x95\x38\x7d\x6a\xed\x6f\x66\x03\xfd\x5b\x89\xaf\x41\x11\x1d\xf7\xf4\x19\x3e\x9d\xc9\x93\x9b\xf6\x02\xf8\x2a\x79\x94\xb5\xb0\x9f\x80\xaf\x70\x7c\x94\xd5\x98\x3e\x6d\x0e\xf6\xd4\x2b\xec\x53\x62\x3f\x00\x53\x82\xd3\xba\x2b\x7a\x7e\x8e\x6a\x68\x40\xf5\x28\xdf\x9d\xda\x51\xe0\x56\xe2\x8b\x87\xc6\x24\xc3\xb1\xb7\x3e\x24\x77\xc6\x0b\x84\x3f\xd8\x09\x3c\x4a\xdc\x05\x25\xb6\x17\xe6\x37\x3f\x1a\x2b\xd0\x85\xb5\x30\x8a\x2e\xfb\xbf\x4c\x58\x45\x3b\x6c\xec\x1e\xc5\x1b\x22\x15\xad\x24\x10\x51\x2d\xb1\xa9\x79\x1d\x9f\x77\xf3\x98\x94\x2f\x03\x29\x35\x8b\x56\x17\x7a\xc8\x25\xda\x4e\xea\x1f\x85\x6c\x1b\xaa\x32\x3c\x72\xf7\x00\x5b\xb8\x1a\xce\xdc\x7c\xf9\x2a\xa8\x94\x9b\x86\x28\x3f\xfc\x3a\x0c\x33\x48\xd0\x6f\xd2\xe1\x88\x7e\x3b\x0d\xca\x2b\xee\xde\x70\x4a\x17\xa3\xd2\x2a\x1e\x39\xb3\xd3\x51\x18\x37\x2a\x9e\x1e\xb7\x86\x93\x71\xa7\xe7\xa9\x71\xaf\x27\x9c\xd5\xa3\x53\x8e\x6c\x3e\xc7\x76\x43\x5e\x20\xfc\x77\xde\x6c\x17\x9c\x85\x2c\x63\xe3\x3d\x65\x44\x81\xc4\x17\xa8\xf4\xab\x99\xda\xbc\x3e\x49\x9a\xa0\x70\x97\x36\xd8\xf3\x06\x77\x89\x93\xf6\xae\x94\x3a\xc6\xe3\xc3\x99\xa3\x68\x03\x6f\xff\xd2\x43\x99\xf8\x2c\xe6\x97\xfe\x18\x82\x41\xeb\xc0\x29\xbd\x4b\x3c\xd0\x2c\x08\xff\x0b\x9b\x3f\x62\xea\xeb\xf5\x9f\x70\x10\xe9\x2d\x28\xe9\x03\xdb\x3d\xa2\xcc\xa0\xac\x2b\xb1\x13\xdf\x17\x2b\xa2\xaa\x65\x26\xc0\x16\xf7\x41\xc5\x99\xa2\x6c\x0d\x7b\x97\x49\x8d\xbf\x47\xdf\x4d\xbc\x75\x7d\x77\xc1\xf1\x26\xe5\xd6\x61\xf6\xf8\xfd\xb4\xc4\xfc\xee\x77\xa8\x94\x75\x7f\xb7\x06\xa4\xf3\xfb\xb7\x56\xbb\x67\xee\xa7\x51\xdb\xf8\xff\x26\xbb\x9d\xf6\xfe\xe1\x97\xdf\xd0\xf5\x25\x1e\xa1\xf4\x88\xdf\xdf\x23\x1e\x6d\xdf\x7f\x1c\x0f\x43\xb3\x76\x40\x7d\xf8\x72\xa3\x8f\x79\x96\x8f\x07\xc3\xff\x04\x00\x00\xff\xff\xf4\xba\x78\x77\xdd\x1f\x00\x00")

func wwwJavascriptMapzenWhosonfirstInspectorJsBytes() ([]byte, error) {
	return bindataRead(
		_wwwJavascriptMapzenWhosonfirstInspectorJs,
		"www/javascript/mapzen.whosonfirst.inspector.js",
	)
}

func wwwJavascriptMapzenWhosonfirstInspectorJs() (*asset, error) {
	bytes, err := wwwJavascriptMapzenWhosonfirstInspectorJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/javascript/mapzen.whosonfirst.inspector.js", size: 8157, mode: os.FileMode(420), modTime: time.Unix(1563548247, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwJavascriptMapzenWhosonfirstRenderJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x56\xc1\x6e\xdb\x3c\x0c\x3e\xdb\x4f\x41\xf8\x3f\xd4\x46\xf2\xbb\xdd\xb5\xae\x31\x14\xc3\xae\x3b\xed\x36\x14\x85\x62\x31\x8b\x16\xc5\x0e\x24\xa6\x49\xd7\x06\xd8\x3b\xec\x0d\xf7\x24\x03\x25\xd9\xb1\x13\xc7\x3b\xb4\xb5\xad\xef\xfb\x48\xf1\x23\x89\xbe\x08\x03\x1b\xb1\xfd\x89\x35\x94\xed\xc3\xfb\x3b\xbc\x1d\x8b\xd8\xbf\xe5\xfb\x55\x63\x9b\x7a\xa9\x8c\xa5\x0e\x32\xf8\x18\xe0\x23\xf8\xdc\x60\x2d\xd1\x40\x09\xe9\x72\x57\x57\xa4\x9a\x3a\xcd\xde\xe2\x38\xe2\xa8\x16\xf5\x12\x4a\xe0\xd7\xe8\xc6\x03\x9f\xa5\x20\x71\x73\x0f\x1d\x58\xce\xa1\xa2\x83\xa3\x44\x8e\xb4\x37\x62\xbb\x75\x8a\xb2\xa9\x76\x1b\xac\x29\xaf\x0c\x0a\xc2\xcf\x1a\xf9\x2d\x4d\xa4\x7a\x49\xb2\x82\xf1\x01\x9b\x5b\xa4\x47\x22\xa3\x16\x3b\xc2\x34\xa9\xb4\xb0\x36\x99\x43\x22\x77\x9b\x2d\x1a\xc6\xb6\xe2\x3e\x09\xc7\xe5\x1f\xb5\x84\xf4\xd1\x18\xf1\x9a\x2b\xeb\xfe\xa6\x32\xcb\xde\xf8\x24\xea\xee\xc5\x97\x08\xb7\x7c\xd6\xca\x52\x9b\xb1\x13\x39\xc6\x11\x00\x00\x3f\xa2\xb6\x08\x2c\x48\xaf\x5b\x6c\x96\xa9\xcc\xa0\x2c\x21\x69\x16\x3f\xb0\xa2\x64\x4a\x55\xaa\xea\x5c\xb5\x13\xf4\x34\x4e\xbd\x6a\x6a\xc2\x9a\xce\xb8\x84\x87\x1e\x77\x18\xe3\x5a\xfd\xec\x56\xd4\x49\x56\x44\xa1\x08\x81\x90\x73\x29\x6b\xf9\x69\xa5\xb4\x4c\x43\xb0\x7e\x3a\x6d\xad\xfb\x30\xcf\xf4\x28\x83\xb4\x33\x75\x6b\x5f\xe1\xc5\x8f\xf3\x81\xf7\xaa\xa2\x49\xef\x49\x2c\x34\x4e\x64\xee\xce\x83\xf7\xee\xf9\x9a\xf3\x5e\xc8\xfd\xfe\x5f\x22\x09\xa5\xad\xa3\xb5\x66\x2d\x1b\x03\xe9\x1a\x54\x0d\xd2\x39\xd3\x55\xd9\x34\xfb\xa9\xf8\x26\x04\x77\x58\x2d\x16\xa8\x9d\x03\x50\xc2\xda\x7f\xef\x0e\x9f\x2b\x3a\xf0\x54\xf0\x05\xe1\x23\xdf\x13\x66\x90\xe4\x09\xcc\x60\x0d\xf7\x17\xf0\x15\x8a\x69\xd3\x68\x75\x11\xfa\x12\xfd\x15\x0f\xf4\xa5\x91\x98\x9e\x52\x0b\x24\xaf\x3f\x30\xcf\x61\xdc\xf1\x48\x8f\x5d\x4d\x43\xb6\x69\x74\xa4\x45\x23\x5f\xcf\x3b\x5a\x90\x48\xe5\xb7\xf5\xd3\xdc\xd5\x21\x30\x82\xfa\x20\x09\x26\xf7\x05\x4d\xb3\x1f\x9c\xfb\xbc\x03\xe2\xfc\x70\xd0\xa5\x9e\xef\xbb\x62\xd0\xa3\xcd\xbe\x9b\x55\x0f\x0a\x9d\xea\xa0\x85\x6f\xd2\xa8\xd7\xa5\x3c\xe4\xa3\x5d\xda\x75\x8f\x2f\xd5\xce\x17\x2a\xd7\x58\x7f\xa7\xd5\x60\xa7\x84\xc3\x12\xee\xba\xb9\x77\x21\x2f\x26\x37\xf9\xf3\xeb\x77\x32\xbe\x4f\x4e\x3a\x0f\x25\x7c\xb8\xae\x13\x6a\x7d\xf7\x34\x2e\xe3\xba\x45\xd9\x29\x4f\x77\xfa\x72\x36\x98\xa6\xca\xbb\x02\x14\x3c\xf8\xbb\x16\xa0\x66\xb3\x90\x46\xe7\xbd\x22\xdc\x4c\x28\x6b\xd5\x6f\xda\xeb\x8d\xa2\x7c\xf2\x3c\x21\xff\xf1\x84\xa8\xbe\xa5\x1c\xe3\x5a\xcf\xf0\xcd\x06\x67\x0c\x1e\xb7\x9b\xa1\x23\x6e\xb3\x0d\xe3\x6e\x7b\xb6\x5b\x4b\x7e\xc4\x65\x11\xdd\xde\xfa\xf4\x57\xb4\xd1\x76\x8b\x95\x12\xba\x5a\x09\x63\x53\x99\x15\x7d\x06\xaf\xd8\x7f\x6f\x60\x46\xf3\xe3\xd9\x16\x53\xb2\xdf\x12\x23\x00\x52\xa4\xf1\x84\x61\xe3\x4e\xa1\x27\xf7\xc2\x69\x23\x38\xd9\x7e\xe5\xc2\x2e\xe8\x97\x8c\x31\x6d\xc9\xa2\x38\xe2\xff\x00\xfa\x1d\x58\xc4\xf1\x31\x4b\xb3\x22\x8a\xff\x06\x00\x00\xff\xff\x3a\xb1\xda\x6e\x63\x08\x00\x00")

func wwwJavascriptMapzenWhosonfirstRenderJsBytes() ([]byte, error) {
	return bindataRead(
		_wwwJavascriptMapzenWhosonfirstRenderJs,
		"www/javascript/mapzen.whosonfirst.render.js",
	)
}

func wwwJavascriptMapzenWhosonfirstRenderJs() (*asset, error) {
	bytes, err := wwwJavascriptMapzenWhosonfirstRenderJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/javascript/mapzen.whosonfirst.render.js", size: 2147, mode: os.FileMode(420), modTime: time.Unix(1563547821, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwCssMapzenWhosonfirstInspectorCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\xdd\x8e\x9b\x30\x10\x85\xef\x79\x8a\xa9\xa2\xbd\x24\x62\xd3\x46\xaa\xcc\xd3\x38\xf6\x00\x56\xfc\x43\xc7\x43\x36\xdb\x2a\xef\x5e\x81\x21\x05\xc2\x6e\x7a\xb1\xca\x55\x66\x6c\x7f\x67\xe6\x1c\x76\x4a\xfa\x8b\x8c\xf0\x27\x03\x00\x70\x92\x6a\xe3\x05\xbc\xa2\x2b\xb3\x5b\x96\xed\x9c\x6c\xc7\x56\x83\xa6\x6e\x58\xc0\xf7\xa2\x68\xaf\xa9\xf9\xab\x43\x7a\x1f\xdb\xa7\x40\x1a\x49\xc0\xa1\xbd\x42\x0c\xd6\x68\xd8\x15\x45\x51\xce\xaf\x1e\x87\x8b\x09\x73\xcd\x37\x8a\x3d\x3b\x3f\x05\xe6\xe0\xc4\x5d\x41\x82\x54\x81\xdc\x48\x6a\xa5\xd6\xc6\xd7\x02\xf6\xc7\xfe\xcc\xac\x94\x5b\xac\x78\x26\x5e\x23\x4b\x63\xe3\x7f\x28\x0c\x17\xa4\xca\x86\x37\x01\x51\x51\xb0\x16\xbe\x19\xd7\x06\x62\xe9\x79\x31\xc2\x38\xfc\x6a\x86\x45\x75\x73\x08\xc2\xd8\x59\xfe\x4a\x25\xb0\x2d\x05\x9e\x6b\x89\xa1\x23\x85\x77\xc3\x57\x57\x3f\xe7\x3d\x74\x3f\x95\x7b\xcb\x32\x96\x27\x3b\xb1\xac\xf1\x38\xc1\x0e\x93\x75\x55\xf0\x9c\x57\xd2\x19\xfb\x2e\x20\x4a\x1f\xf3\x88\x64\xaa\xd4\x7c\x33\x9a\x1b\xf1\x5a\x14\x2f\xe5\x6c\x71\x79\x6c\xa5\x1a\x12\xd0\xf9\x88\x23\x87\x46\xc8\x05\x89\x8d\x92\x36\x97\xd6\xd4\x5e\x00\x87\xb6\x5c\xc6\x66\xb5\x9f\x54\xe8\xff\x0f\xef\x34\x93\x47\x52\x9d\x6b\x0a\x9d\xd7\xb9\x0a\x36\x90\xd8\xa1\xee\x7f\xe5\xc2\xc1\xe4\x1e\x37\xc6\xa7\x3a\xe3\x95\x47\x72\x9f\xc5\x47\xce\x86\x31\xfb\xc3\x2a\xc6\x02\xf6\x87\x23\xba\x41\x8e\x1e\xe5\x6c\x3e\xbc\x3c\xff\x01\x6c\x94\xfa\x4f\xe9\xda\xa2\xfb\x67\x92\xbc\x62\x82\x3b\x35\xed\xff\xc7\xcf\x97\xa7\xbe\x0f\x2f\xb5\x34\x59\x3d\xb8\x1a\xcd\x6f\x14\xd1\xc9\xbe\x3b\x4f\xd2\x71\x33\x49\xf3\x24\x3e\x9e\xb8\x65\xd9\x5e\x59\xa3\xce\xb9\x9b\x18\xc9\x16\x38\xd9\x0e\x67\xcb\xd7\xa8\x02\x49\x36\xc1\x8b\xce\x6b\xa4\x3e\x74\xe9\x3a\x12\x05\xfa\xc8\x5d\x68\x8d\x3f\x2f\x89\x7f\x03\x00\x00\xff\xff\x08\x5c\xb8\xab\x1a\x05\x00\x00")

func wwwCssMapzenWhosonfirstInspectorCssBytes() ([]byte, error) {
	return bindataRead(
		_wwwCssMapzenWhosonfirstInspectorCss,
		"www/css/mapzen.whosonfirst.inspector.css",
	)
}

func wwwCssMapzenWhosonfirstInspectorCss() (*asset, error) {
	bytes, err := wwwCssMapzenWhosonfirstInspectorCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/css/mapzen.whosonfirst.inspector.css", size: 1306, mode: os.FileMode(420), modTime: time.Unix(1563546049, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwCssMapzenWhosonfirstRenderCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xdf\x6a\xf3\x30\x0c\xc5\xef\xf3\x14\xfa\xf8\xd8\xa5\x4b\x57\xd8\x8d\xf7\x34\x4e\xac\x24\x62\xfe\x87\xac\x6c\x2d\xa3\xef\x3e\x12\xbb\xa1\x66\x2d\x8c\x5c\xe9\x28\x3a\xd2\xef\x78\x71\xf0\xdd\x01\x00\x38\xca\xa2\xb2\x5c\x1c\x2a\xb9\x24\xd4\x10\x62\xc0\xf7\xad\xe5\x0d\x4f\x14\xf4\x31\x9d\x4b\x9d\x8c\xb5\x14\xa6\x22\x5c\xbb\xee\x60\x17\x9f\x90\x41\x4c\xef\x70\xb7\x0b\xa8\x66\xa4\x69\x16\x7d\x42\x5f\x06\xc7\x18\x44\x8d\xc6\x93\xbb\x68\xc8\x26\x64\x95\x91\x69\x2c\xcd\x2f\xb2\x32\xeb\xd7\xe3\xf1\xa5\xd4\x7d\x64\x8b\xac\x72\x32\xc3\xba\x0c\x96\x90\x51\x1e\xec\x13\xae\x2b\x3f\x91\x85\x06\xe3\x94\x71\x34\x05\x0d\x12\x53\x7b\x2f\xec\x04\x95\x68\x13\xd6\xba\x71\x9d\xab\x5f\x6f\x86\x8f\x89\xe3\x12\xac\x1a\xa2\x8b\xac\xff\xa3\x5d\xbf\xe6\x3c\x89\xe9\x3e\xaa\xaa\xf6\x51\x24\xfa\xda\xd8\x16\xdc\x35\x1d\x8e\xa2\x21\x47\x47\x16\x64\xa6\x00\xff\xc8\xa7\xc8\x62\x82\x34\x26\xbc\x85\xf7\xe4\xc7\xdd\x54\xf0\x2c\x95\x78\x35\xfe\xcd\xd7\x46\x70\x38\xbd\xa1\x6f\x71\xf5\x48\x9c\x45\x0d\x33\x39\x7b\x43\xbf\x83\x7b\x72\x68\x6b\xe1\xcc\x63\x87\x5b\x10\x7f\x30\x29\x8f\x79\x9b\x7f\x88\xd5\x52\x34\x5a\x4d\xf5\xb0\xeb\x4d\x04\xd7\xee\x27\x00\x00\xff\xff\xcd\x95\x93\x60\xe8\x02\x00\x00")

func wwwCssMapzenWhosonfirstRenderCssBytes() ([]byte, error) {
	return bindataRead(
		_wwwCssMapzenWhosonfirstRenderCss,
		"www/css/mapzen.whosonfirst.render.css",
	)
}

func wwwCssMapzenWhosonfirstRenderCss() (*asset, error) {
	bytes, err := wwwCssMapzenWhosonfirstRenderCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/css/mapzen.whosonfirst.render.css", size: 744, mode: os.FileMode(420), modTime: time.Unix(1563546049, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"www/index.html": wwwIndexHtml,
	"www/javascript/mapzen.whosonfirst.geojson.js":        wwwJavascriptMapzenWhosonfirstGeojsonJs,
	"www/javascript/mapzen.whosonfirst.inspector.init.js": wwwJavascriptMapzenWhosonfirstInspectorInitJs,
	"www/javascript/mapzen.whosonfirst.inspector.js":      wwwJavascriptMapzenWhosonfirstInspectorJs,
	"www/javascript/mapzen.whosonfirst.render.js":         wwwJavascriptMapzenWhosonfirstRenderJs,
	"www/css/mapzen.whosonfirst.inspector.css":            wwwCssMapzenWhosonfirstInspectorCss,
	"www/css/mapzen.whosonfirst.render.css":               wwwCssMapzenWhosonfirstRenderCss,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"www": &bintree{nil, map[string]*bintree{
		"css": &bintree{nil, map[string]*bintree{
			"mapzen.whosonfirst.inspector.css": &bintree{wwwCssMapzenWhosonfirstInspectorCss, map[string]*bintree{}},
			"mapzen.whosonfirst.render.css":    &bintree{wwwCssMapzenWhosonfirstRenderCss, map[string]*bintree{}},
		}},
		"index.html": &bintree{wwwIndexHtml, map[string]*bintree{}},
		"javascript": &bintree{nil, map[string]*bintree{
			"mapzen.whosonfirst.geojson.js":        &bintree{wwwJavascriptMapzenWhosonfirstGeojsonJs, map[string]*bintree{}},
			"mapzen.whosonfirst.inspector.init.js": &bintree{wwwJavascriptMapzenWhosonfirstInspectorInitJs, map[string]*bintree{}},
			"mapzen.whosonfirst.inspector.js":      &bintree{wwwJavascriptMapzenWhosonfirstInspectorJs, map[string]*bintree{}},
			"mapzen.whosonfirst.render.js":         &bintree{wwwJavascriptMapzenWhosonfirstRenderJs, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
