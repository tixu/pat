// Code generated by go-bindata. DO NOT EDIT.
// sources:
// schema/schema.json
// schema/schema.yaml
package pkg

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

var _schemaSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x55\x4d\x6f\xdb\x30\x0c\xbd\xfb\x57\x08\x6a\x0f\x1b\x96\xcc\x3b\xec\x94\x5b\x81\xad\xa7\x01\x2b\x86\xdd\xd6\x0e\xd0\x62\x3a\x56\x61\x49\x1e\x45\x03\x0d\xe6\xfe\xf7\x41\xfe\x8a\x24\xcb\x45\x9a\x01\x3b\x86\xa2\xdf\xe3\xe3\x23\x99\x3f\x19\x63\xfc\xda\xee\x2b\x50\x82\xef\x18\xaf\x88\x9a\x5d\x9e\x3f\x5a\xa3\xb7\x43\xf4\xbd\xc1\x43\x5e\xa0\x28\x69\xfb\xe1\x63\x3e\xc4\xae\xf8\xc6\x7d\x27\x8a\x42\x92\x34\x5a\xd4\x77\x68\x1a\x40\x92\x60\xf9\x8e\x95\xa2\xb6\xd0\x27\x14\x50\x4a\xdd\xa7\xb8\xb8\xe3\x62\x8c\xdf\x58\xeb\x52\x8d\x9e\x43\x8c\xf1\xc6\x07\x98\xa2\x8e\x82\x82\xdf\xae\x58\x84\xd2\x55\x7a\x95\x7b\xe8\xf9\xa7\x16\x45\x8f\x39\xa7\x3e\x6f\x4e\x28\xf0\xd4\xc0\x9e\xa0\x88\xb1\x24\x81\xb2\x51\x70\x95\xe2\x73\x0f\x12\xb1\x04\x3c\x8c\x71\x3a\x36\xe0\x3e\x15\x88\xe2\xe8\x15\x93\x45\xc9\x73\xa2\xf9\xf5\x08\x7b\x1a\x32\xc7\x57\x3e\x8b\xf1\x1b\x24\x88\x00\x5d\x88\xff\xbc\xbf\x2f\xde\xbd\xd1\xb6\x6b\x6d\xa7\x6c\x67\x3b\xd5\x55\x6f\xaf\xf9\x02\xda\x12\x4a\x7d\x08\xa1\x7d\x11\x1e\xfa\x4b\xd5\xdc\xca\x27\x6a\x11\xce\x71\xab\x58\x16\x7e\xb1\x67\x0a\x08\xe5\x3e\x76\x67\xcd\xb2\x94\xe6\x7f\xf2\x07\xe1\x77\x2b\xb1\x1f\x99\x1f\x09\x81\x89\x42\xc7\xc8\xc3\x59\x16\x7f\x6b\xeb\xa0\x79\xdc\x68\xf8\x5a\x06\x64\x81\xec\x95\x8e\xf7\x6f\x25\x1a\x75\x2b\x6b\x58\xbc\xbc\xd8\x97\x93\xe8\x65\x97\x92\xe2\x43\x2e\x2f\xfc\x90\xf2\xef\x55\xd5\x7f\x91\x04\x28\xea\xff\x25\x60\xa2\x4b\x6b\x38\xcf\xc7\x6c\xa4\x4c\x69\xe3\x62\x3a\x71\x81\xc5\xcb\xd1\x5d\x59\x8c\xd3\x85\x5c\x3d\x1b\xde\xfc\x4e\x23\x55\x0e\x7b\x7a\x19\xe5\xb4\xe4\xaf\x21\xd4\x42\x41\xea\x88\xa4\xee\x0e\xc6\xf3\x9e\x2e\x63\x58\x8b\xa8\xc3\x0b\x33\x07\xe6\x00\x79\xb3\x68\x7c\x36\x5a\xc8\x49\x52\xbf\x1b\xee\x84\xb2\x2d\xbb\x43\xa3\x80\x2a\x68\x2d\xbb\xa9\x01\x89\x7d\x07\x4b\x52\x1f\x58\x4b\xb2\x96\x74\x1c\xfe\xda\x62\xcf\x9f\xb3\xbf\x01\x00\x00\xff\xff\xa4\x45\xaf\x81\x2a\x07\x00\x00")

func schemaSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaSchemaJson,
		"schema/schema.json",
	)
}

func schemaSchemaJson() (*asset, error) {
	bytes, err := schemaSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/schema.json", size: 1834, mode: os.FileMode(420), modTime: time.Unix(1529467027, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaSchemaYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xcb\xca\xdb\x30\x10\x85\xf7\x7a\x8a\xc1\x7f\x16\x2d\xc5\x75\x17\x5d\x79\x17\x48\xb3\x2a\x34\x84\x2e\x43\x41\xb5\xc7\xb1\x82\x2e\xee\x68\x0c\x09\xf8\xe1\x8b\x7c\x4f\xec\x40\xda\x5f\x9b\x28\xa3\xe3\x6f\xe6\x1c\x69\xe3\xb3\x12\x8d\x4c\xa1\x64\xae\xd2\x24\xb9\x78\x67\xe3\xae\xf6\xd9\xd1\x39\xc9\x49\x16\x1c\x7f\xf9\x9a\x74\xb5\x37\xc1\x8a\x35\xa6\x50\x49\x86\x18\x0e\xe4\x0c\x72\x89\xb5\x87\xad\x46\x62\xf8\x89\x9e\x95\x3d\x43\xcd\x4a\x2b\xbe\x09\xbe\x55\x98\x82\xfb\x7d\xc1\x8c\x45\x45\xae\x42\x62\x85\x3e\x15\x00\x56\x1a\x0c\xbf\x00\x9d\xc8\x33\x29\x7b\x16\x00\x54\xeb\x4e\x01\xb0\x21\x2c\x52\x88\xde\x92\x1c\x0b\x65\x15\x2b\x67\x7d\x72\x0c\xe7\x91\x00\x28\xd4\x95\x6b\x1a\xb4\x1d\x45\x12\xc9\x5b\xfb\x5f\x31\x9a\xfe\xe8\x09\x68\xdf\x7d\x1f\x50\xd2\xfb\x30\x99\xb3\xff\x0d\xdb\x0e\x84\x48\x10\xfe\xa9\x15\x61\x1e\xf4\x71\xeb\xb3\xdd\xb4\xbe\xda\xdd\xd4\x4d\xc8\x3c\x6f\x01\x52\x1f\xa6\x74\xa0\x90\xda\xa3\x98\xe1\x03\xea\x38\xe5\x72\x17\x6b\x28\x38\x8b\x3f\x8a\x61\xbe\x18\xee\x93\x1e\x56\x41\xce\xec\x95\xc6\x79\x6d\x91\x7e\xb7\xe6\x16\x86\x15\x8f\x80\x17\xfa\x7c\x57\x8c\x24\xf5\xbb\x5a\xf5\x0c\x01\xd0\xdf\xd4\x13\xf3\xcb\x29\xf2\x9a\x64\xc8\x6d\x62\xae\xde\xd9\xae\x97\x45\xbd\xcc\x20\x93\xca\x66\x5e\x1e\x9f\xc1\xe2\x29\xac\xba\x7a\x74\x14\x8f\xf3\x8c\x85\xbe\x93\x00\x18\xdf\xcd\xcb\xe6\x24\xff\xa3\x2d\xbc\x56\x98\xf1\x3c\xe1\x97\x7c\xad\xb2\xbf\xb5\xac\x11\xbf\xbb\x0b\x7a\x91\x44\x25\x99\x91\x6c\x0a\xd1\xaf\xd3\x29\xff\xf4\xc1\xfa\xa6\xf6\x8d\xf1\x8d\x6f\x4c\x53\x7e\xdc\x04\xc4\x8c\xb8\x12\xc1\xdf\x00\x00\x00\xff\xff\x8a\x99\xea\xa2\xa0\x04\x00\x00")

func schemaSchemaYamlBytes() ([]byte, error) {
	return bindataRead(
		_schemaSchemaYaml,
		"schema/schema.yaml",
	)
}

func schemaSchemaYaml() (*asset, error) {
	bytes, err := schemaSchemaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/schema.yaml", size: 1184, mode: os.FileMode(420), modTime: time.Unix(1529466865, 0)}
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
	"schema/schema.json": schemaSchemaJson,
	"schema/schema.yaml": schemaSchemaYaml,
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
	"schema": &bintree{nil, map[string]*bintree{
		"schema.json": &bintree{schemaSchemaJson, map[string]*bintree{}},
		"schema.yaml": &bintree{schemaSchemaYaml, map[string]*bintree{}},
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
