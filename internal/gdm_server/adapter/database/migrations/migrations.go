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

var __000001_init_schemeUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xc1\x4e\x84\x30\x10\x86\xef\x3c\xc5\x9f\x3d\x2d\x89\x6f\xe0\xa9\xe0\xb0\x36\xd6\xa2\xa5\x18\xf7\x44\x08\x54\x53\xc5\x42\xa0\xfa\xfc\xc6\xa5\xc2\x46\x5d\x63\xe2\x71\xf2\x7d\x99\x74\xbe\x34\x55\xc4\x34\x41\xb3\x44\x10\x78\x06\x99\x6b\xd0\x3d\x2f\x74\x81\xcd\x30\xf6\x4f\xa6\xf1\xd3\x06\xdb\x08\x00\xc2\x5c\xd9\x16\x5c\x6a\xda\x91\x3a\xe8\xb2\x14\xdd\xd9\x41\xe8\xac\x7b\xc6\x1d\x53\xe9\x25\x5b\x90\x98\x51\xd3\x3b\x5f\x5b\x67\xc6\xca\xd5\x2f\xe6\x9b\x84\x52\xf2\xdb\x92\x66\xb7\x35\x53\x33\xda\xc1\xdb\xde\x7d\x8a\x33\x78\x1d\x2a\xdf\x57\x6d\xed\x0d\x92\x3c\x17\xc4\xe4\xba\xe0\x82\x32\x56\x0a\x8d\x8c\x89\x62\xd9\xd3\x19\x6f\xda\x3f\xb9\x37\x8a\x5f\x33\xb5\xc7\x15\xed\xb7\xeb\x9d\x71\x14\x9f\x47\xd1\x6f\x8d\xde\xcc\x38\xd9\xde\x2d\x8d\xc2\xfc\x43\xa3\x10\xe2\x74\xc4\x20\x84\x0d\xa7\x3b\x3e\xd8\x47\x24\x22\x4f\xfe\x77\xe4\xfa\xd0\x78\x86\x59\xae\x88\xef\xe4\x07\xc4\x71\x02\x28\xca\x48\x91\x4c\xe9\xf8\x4b\x7c\x89\xf4\x1e\x00\x00\xff\xff\x93\xf0\x77\x2b\x47\x02\x00\x00")

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

	info := bindataFileInfo{name: "000001_init_scheme.up.sql", size: 583, mode: os.FileMode(420), modTime: time.Unix(1694102238, 0)}
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
