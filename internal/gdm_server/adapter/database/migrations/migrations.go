// Code generated for package migrations by go-bindata DO NOT EDIT. (@generated)
// sources:
// migrations/000001_init_scheme.down.sql
// migrations/000001_init_scheme.up.sql
package migrations

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __000001_init_schemeDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x4b\x2d\x2a\xce\xcc\xcf\x2b\xb6\xe6\xc2\x2a\x5d\x50\x94\x9f\x95\x9a\x5c\x12\x9f\x99\x97\x96\x8f\x5f\x49\xb1\x35\x20\x00\x00\xff\xff\x5d\x24\x06\xef\x60\x00\x00\x00")

func _000001_init_schemeDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_init_schemeDownSql,
		"000001_init_scheme.down.sql",
	)
}

func _000001_init_schemeDownSql() (*asset, error) {
	bytes, err := _000001_init_schemeDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_init_scheme.down.sql", size: 96, mode: os.FileMode(420), modTime: time.Unix(1693840088, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000001_init_schemeUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\x51\x4f\x83\x30\x10\x80\xdf\xf9\x15\x97\x3d\x8d\xc4\x7f\xe0\x53\x21\xc7\x6c\xac\xad\x76\xc5\xb8\x27\x42\xe8\x69\xaa\xd8\x12\xa8\xfe\x7e\xe3\xa8\xdb\x32\xc7\x1e\x2f\xdf\x17\x8e\xfb\x5a\x6a\x64\x06\xc1\xb0\x42\x20\xf0\x0a\xa4\x32\x80\x2f\x7c\x6b\xb6\xb0\x1a\xc6\xf0\x4e\x5d\x9c\x56\xb0\xce\x00\x00\xd2\xdc\x38\x0b\x5c\x1a\xdc\xa0\xde\xeb\xb2\x16\xfd\xcd\x5e\xe8\x9d\xff\x80\x67\xa6\xcb\x3b\x76\x40\x62\x46\x5d\xf0\xb1\x75\x9e\xc6\xc6\xb7\x9f\xf4\x4f\x82\x5a\xf2\xa7\x1a\x67\xd7\xd2\xd4\x8d\x6e\x88\x2e\xf8\x3f\x71\x06\x5f\x43\x13\x43\x63\xdb\x48\x50\x28\x25\x90\xc9\xb3\x2d\x96\x7a\x8a\x64\x17\xe8\xa3\xe6\x0f\x4c\xef\xe0\x1e\x77\xeb\xe3\x2d\x79\x96\xdf\x66\xd9\xb5\x0e\xdf\x34\x4e\x2e\xf8\x43\x87\x34\x5f\xe8\x90\x16\x2d\x87\x4a\x42\xfa\xc2\x72\xab\x57\xf7\x06\x85\x50\xc5\x95\x0b\x8e\x7f\x91\xcf\xb0\x52\x1a\xf9\x46\xfe\x42\x38\xbd\x0f\x34\x56\xa8\x51\x96\x78\xfa\xa6\x67\x05\x7e\x02\x00\x00\xff\xff\x7f\xe3\x10\x8b\x08\x02\x00\x00")

func _000001_init_schemeUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_init_schemeUpSql,
		"000001_init_scheme.up.sql",
	)
}

func _000001_init_schemeUpSql() (*asset, error) {
	bytes, err := _000001_init_schemeUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_init_scheme.up.sql", size: 520, mode: os.FileMode(420), modTime: time.Unix(1694004928, 0)}
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
	"000001_init_scheme.down.sql": _000001_init_schemeDownSql,
	"000001_init_scheme.up.sql":   _000001_init_schemeUpSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"000001_init_scheme.down.sql": &bintree{_000001_init_schemeDownSql, map[string]*bintree{}},
	"000001_init_scheme.up.sql":   &bintree{_000001_init_schemeUpSql, map[string]*bintree{}},
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
