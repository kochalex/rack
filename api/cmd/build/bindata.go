// Code generated by go-bindata.
// sources:
// data/git-restore-mtime
// DO NOT EDIT!

package main

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

var _dataGitRestoreMtime = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x54\x4d\x8f\x1a\x39\x10\xbd\xf7\xaf\xa8\x15\x5a\xd1\x2d\x61\xf7\x30\x3b\xda\xdd\x41\x42\xd1\x24\xe1\x90\x43\xa4\x51\xa4\x9c\x08\x07\xd3\x5d\x80\x07\xb7\xdd\x63\x57\xc3\xf0\xef\x53\xee\x0f\x06\x82\x12\x5f\x6c\x5c\xaf\xde\xab\x7a\x65\x7a\xf4\x57\xde\x04\x9f\xaf\xb5\xcd\xd1\x1e\xa0\x3e\xd1\xce\xd9\x24\x19\xc1\x8e\xa8\x9e\xe5\x79\x20\x55\xec\xdd\x01\xfd\xc6\xb8\xa3\x2c\x5c\x95\xbf\x36\x18\x48\x3b\x1b\xf2\xe9\xe3\xbf\x0f\x0f\xff\xdd\xe5\xc7\x9d\xa2\x20\x68\x87\x02\x5f\x1b\x7d\x50\x06\x2d\x09\xb7\x11\x4d\x40\xc1\x19\x95\x26\x41\xba\xc2\x20\x36\xce\x8b\xad\xa6\x7c\xfa\xcf\xfd\xff\x0f\xf7\xf7\x8f\xa3\xe1\x10\x05\x3f\x2a\x8f\x62\xed\x2c\x06\x60\xb9\xc0\x0a\x12\x3e\x35\xde\x33\x19\x94\xda\x43\xd5\x04\x82\x35\x02\xb9\x5a\x18\x3c\xa0\x01\xb7\x81\xa3\xf3\x7b\x20\x8f\x28\x99\xe1\x7b\x50\x5b\x9c\x01\x2b\x08\xcf\x35\x3a\xe6\xab\xa2\xb0\x58\x33\x35\x2c\x6b\x45\xbb\x50\x63\x11\xa4\x94\xab\x28\x78\x82\x12\x37\xaa\x31\x04\x4d\x5d\x2a\x42\x50\xc6\xc0\x46\x1b\x0c\x1c\x5d\xbc\xa9\xaa\x36\x4c\x47\x0e\x9c\x35\xa7\x01\xd3\x9e\xb9\x57\xf8\xb6\x78\xfa\xfc\x75\x01\xca\x96\x5d\x0e\x68\x0b\x32\x2f\x5d\x31\xe3\xec\xdf\xd4\xd0\xe7\x30\x28\x49\x74\x55\x3b\x4f\x10\x9a\x75\xed\x5d\x81\x21\x4c\x20\xec\x0c\xbe\x9d\x03\x27\xbe\x71\x41\xc6\xb2\x93\x24\x4a\x18\xcd\x0e\xcc\x21\x20\xa5\x59\xc2\x5e\x42\x0c\x45\xd9\x94\xb1\x52\xf9\xed\x61\x39\x9d\xad\x80\x03\xcb\x3e\x4f\x16\x8d\x67\xef\x56\xd9\x2c\x01\x5e\x7a\x33\x10\x4a\x1d\x22\x63\x1a\xcf\x59\xcc\x78\xbf\x37\xda\xee\xbb\xfb\x2e\x29\xae\x41\x5d\xaa\xb2\x4c\x07\xa8\x47\x13\xf7\x0e\x9b\xb5\x58\x06\x5d\x4a\xb0\xf4\x0d\x13\x6b\x79\xe7\x68\x12\x1b\xe7\x38\xf7\x78\x76\x8f\x13\x8f\xca\xdc\x88\xf7\x95\x8f\x25\x9b\x3a\x8e\xb8\x3e\xf3\x1a\x11\x57\x1f\xe0\xca\x2a\x7e\xb2\x69\x97\x91\x5d\xc1\xa2\x7e\x14\x8c\x3c\xad\xf0\x2d\xcb\x1f\x9b\x1d\x7e\xbf\x38\x6d\xd3\xae\x91\x88\xcf\xd8\x80\xa4\x1d\x34\x0f\xe8\x2e\x61\x5d\xb7\x7e\x89\xb3\x3a\x4f\x57\x3e\xbb\x1a\x6d\xda\x8e\x58\x86\xda\x68\x4a\xc7\x0c\x83\xf8\xe7\x29\x76\xca\x6e\xb1\x04\x21\x6a\x8f\x44\xa7\xf9\xdf\x8a\xeb\x9e\xdc\x54\x76\xd1\x29\x95\xae\xa1\xf9\x25\xfd\x97\xe7\x45\xf7\x2a\x78\x80\x6d\x7b\x5d\x11\xb2\x83\x76\x6d\xb6\xa1\x79\xbb\xf1\xbd\xd7\x75\x9a\x0d\x0f\xc3\x3a\x6a\xef\x67\x50\x38\x4b\xda\x36\x98\x0c\xa1\x1e\xae\x3c\x85\xa3\x66\x0f\xc6\xb3\xf1\x2f\x6f\xe3\xcc\xd9\xf5\xf5\x83\xab\x5f\x8a\xe9\x2a\xb9\x18\xdf\xa5\xe9\xd1\xdd\x6b\xdf\xcf\x9e\xf7\xa3\x6b\x3d\xbd\x42\x8c\x6a\xaf\xf9\x4b\xd0\x7a\xdc\x79\x7e\x15\xe6\xb9\x34\x31\xd4\x66\x4e\x20\xed\x71\xed\x76\x7e\x9c\x01\xdf\x55\x87\x61\x19\x67\xb7\x69\x2c\x3e\xeb\xfa\x1d\xc1\xd3\xf0\x15\xe0\xff\xa9\xc5\x0f\x97\x06\xdd\x16\xbf\xf6\xa8\xf6\xc9\xcf\x00\x00\x00\xff\xff\xd5\xa9\x9e\x5b\x46\x05\x00\x00")

func dataGitRestoreMtimeBytes() ([]byte, error) {
	return bindataRead(
		_dataGitRestoreMtime,
		"data/git-restore-mtime",
	)
}

func dataGitRestoreMtime() (*asset, error) {
	bytes, err := dataGitRestoreMtimeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/git-restore-mtime", size: 1350, mode: os.FileMode(493), modTime: time.Unix(1459609522, 0)}
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

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"data/git-restore-mtime": dataGitRestoreMtime,
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
	"data": &bintree{nil, map[string]*bintree{
		"git-restore-mtime": &bintree{dataGitRestoreMtime, map[string]*bintree{}},
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
