// Code generated by go-bindata.
// sources:
// static/templates/index.html
// DO NOT EDIT!

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
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

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x95\xb1\x72\xdb\x30\x0c\x86\xe7\xe4\x29\x50\x2e\x6d\x07\x45\x2f\x40\x79\x4a\x73\x97\x29\x63\xc6\x1e\x25\x42\x11\x13\x0a\x64\x41\xc8\x89\xeb\xcb\xbb\xf7\x68\x4a\xb1\xe5\x6b\x97\xa4\xf1\x44\x19\xe0\xf7\xff\x04\x20\x71\xbf\x07\x8b\xbd\x23\x04\xe5\xc8\xe2\x8b\x82\xd7\xd7\x4b\x3d\xc8\xe8\x37\x97\x00\x7a\x40\x63\xf3\x02\x40\x8b\x13\x8f\x9b\xfb\x21\x7c\x4d\x70\x47\x70\xe3\x38\x09\xdc\x52\x8a\xd8\x49\x60\x5d\x97\x78\xc9\xf5\x8e\x9e\x40\x76\x11\x1b\x25\xf8\x22\x75\x97\x92\x02\x46\xdf\xa8\x24\x3b\x8f\x69\x40\x14\x05\x03\x63\xdf\xa8\x1c\xac\xdb\x10\x24\x09\x9b\x78\x35\x3a\xba\x3a\xa4\xd7\xef\x61\x3d\x9a\xad\x49\x1d\xbb\x28\x67\xc8\xc7\x42\x7c\xa7\x3d\xc2\x17\xf9\x8d\x99\xb2\xf6\x56\xa4\x4e\x51\x47\x03\x0a\x12\x77\x6b\x4b\x0b\x66\x36\xb4\xd1\x75\x09\xbc\xdb\xd7\x68\x62\xe6\x3d\x0f\x21\x05\xea\x73\x47\xae\xdc\xd2\x91\x0f\x54\xf1\x1f\x68\x46\xb2\xc8\x1f\xaa\xc0\x5f\xa8\x0f\x18\x1e\x53\x38\x2f\xc8\x7f\x82\xcf\x96\x3f\x85\x7d\xac\xf4\x27\xe3\x1d\x39\x59\x6b\xc0\xfc\xcb\xaf\x68\xbd\xbc\xa3\xba\x0d\x76\x37\xcb\x5b\xb7\x05\x67\x1b\xd5\x19\xda\x9a\xa4\xa0\xf3\x26\xa5\x46\x75\x81\xc4\x38\x42\x56\xf3\xc8\xcd\xa9\x4b\xd8\xb0\x55\x9b\xcb\x8b\xb7\xed\xbf\x26\xe4\x9d\x3a\x0d\x57\x59\xed\xb0\xff\x02\x40\xf7\x81\xc7\x25\x9a\xd7\x95\x23\xef\x08\x4b\x74\xcd\x3e\x84\x1f\x38\x4c\x71\x89\x02\xe8\x2f\x55\x05\xda\x9b\x16\x3d\xf4\x81\x1b\xf5\x1c\xfa\x9f\xce\xaa\xf3\xaf\xcc\xb7\xfb\xbb\x9b\xef\x70\x7b\xad\xeb\x43\xee\x06\xaa\xea\xc8\x70\x14\xa7\xd3\x42\xab\x83\xf3\x99\x04\x64\x46\x3c\x3e\x6d\x8d\x9f\xb0\x51\x0a\xa2\x37\x1d\x0e\xc1\x5b\xe4\x46\xfd\x20\x41\x06\x03\xf7\x77\x37\x70\x7b\xad\x56\x8e\x73\xc5\x38\x78\x55\x2f\x47\xaa\xad\xdb\x2e\xeb\x76\x12\x09\x34\x6b\xa7\xa9\x1d\xdd\xac\xee\x43\x78\x9a\xe2\x1b\xa9\x15\x82\x56\xa8\x8a\xec\x46\xc3\x3b\xb5\x99\x3f\x9b\xba\x2e\x84\x85\x97\xa2\xa1\x65\x4f\x8a\x8e\x08\xb9\x6a\x03\x5b\x64\x58\x3f\x56\x69\x54\xc0\xc1\x67\x59\x31\x32\xa5\x22\x3b\x27\x1d\x06\x25\x9a\x82\xd5\x75\x3e\x47\x6e\x6a\x71\x3e\x0f\xce\xb1\xc7\x16\xc5\x38\x9f\xd6\x5d\xce\xa3\x94\x31\xe5\xb0\x6f\xb9\x8c\x69\xf2\x72\x96\xdb\x87\x20\x45\xf4\x44\x60\xad\xb6\xfa\x47\xd7\x65\x50\x75\x5d\xae\x99\xfd\x1e\x90\x6c\xbe\x77\xfe\x04\x00\x00\xff\xff\xbc\xb4\xa6\x48\x8d\x06\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 1677, mode: os.FileMode(420), modTime: time.Unix(1563974987, 0)}
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
	"index.html": indexHtml,
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
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
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

