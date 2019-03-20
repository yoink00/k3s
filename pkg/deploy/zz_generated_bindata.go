// Code generated by go-bindata.
// sources:
// manifests/coredns.yaml
// manifests/traefik.yaml
// DO NOT EDIT!

package deploy

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

var _corednsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\xdd\x6e\x1b\xb7\x12\xbe\xd7\x53\x10\x7b\x90\xbb\xb3\xb2\x04\x23\x39\x3e\xbc\x4b\x24\x37\x35\x90\xb8\x82\x65\xe7\xa6\x28\x82\x11\x77\x24\xb1\xe6\x72\x58\x72\x56\xb1\x9a\xe6\xdd\x0b\xee\x9f\xb9\xf2\x3a\x48\x82\xac\x2e\x44\x72\x38\xdf\x90\xf3\xf3\x0d\xc1\xe9\x0f\xe8\x83\x26\x2b\xc5\x61\x3e\xb9\xd7\xb6\x90\x62\x8d\xfe\xa0\x15\xbe\x56\x8a\x2a\xcb\x93\x12\x19\x0a\x60\x90\x13\x21\x2c\x94\x28\x85\x22\x8f\x85\x0d\xed\x3c\x38\x50\x28\xc5\x7d\xb5\xc1\x3c\x1c\x03\x63\x39\xc9\xf3\x7c\x92\x42\xfb\x0d\xa8\x29\x54\xbc\x27\xaf\xff\x06\xd6\x64\xa7\xf7\x17\x61\xaa\xe9\xec\x30\xdf\x20\x43\x67\x79\x61\xaa\xc0\xe8\x6f\xc8\xe0\xc0\xac\x81\x0d\x9a\x10\x47\xa2\xb6\xe3\x2d\x32\xd6\xfa\x1b\x22\x0e\xec\xc1\x39\x6d\x77\x8d\xa1\xbc\xc0\x2d\x54\x86\x43\x7f\xde\xe6\x54\xb2\x3b\xb6\xaf\x0c\x06\x39\xc9\x05\x38\xfd\xd6\x53\xe5\x6a\xe4\x5c\x64\xd9\x44\x08\x8f\x81\x2a\xaf\xb0\x5d\x43\x5b\x38\xd2\xb6\x06\xcb\x45\x68\x3c\xd3\x4c\x1c\x15\xcd\xa0\x77\x42\x9c\x1e\xd0\x6f\x5a\x5d\xa3\x03\xd7\x83\x4f\xc0\x6a\xff\x6d\xf6\x2c\x15\xa7\x30\x3b\xe4\x9f\xe1\xd0\x37\xda\x16\xda\xee\x06\x7e\x05\x6b\x89\x6b\xf5\xd6\xb9\x63\xb8\x03\x7f\x43\xc5\x54\xb9\x02\x18\xa5\xc8\xd8\x57\x98\xfd\xfc\xf0\x90\xc1\x1b\xdc\xd6\xe7\x6b\x1d\xf6\x95\x0b\x4f\x84\x78\x9a\x3b\xcf\x20\x87\x6a\xf3\x27\x2a\xae\x63\x3f\x9a\xea\x3f\x9c\xe0\x7d\xed\x2c\xc8\x6e\xf5\xee\x3d\xb8\x1f\x29\x9b\x6e\xfb\x82\x3c\x6e\xb5\x41\x29\xfe\xa9\x7d\x3a\x95\x2f\xcf\xc5\xe7\x7a\x18\x3f\xf4\x9e\x7c\xe8\xa7\x7b\x04\xc3\xfb\x7e\xfa\x18\x00\xa1\x1a\x97\x4c\x0d\x29\x30\x42\xdb\x1c\x8a\xc2\x4f\xc1\x3b\x10\xda\xbd\x6a\x06\x8f\xb0\xa2\xce\x68\xa1\x6d\x40\x55\x79\x4c\xd6\x2b\x17\xd8\x23\x94\xc9\xd2\x16\x8c\xe1\xbd\xa7\x6a\xb7\x1f\x07\xee\xf7\x7e\xe9\x47\xce\x53\x89\xbc\xc7\x2a\x08\xf9\xff\xf9\xcb\xf3\x54\xf0\x70\x14\x53\x31\x9f\xd6\xbf\x7e\x5d\x81\xda\xa3\x38\x9f\xf5\x0b\x86\xc8\xf5\x13\x8f\x86\xa0\x48\x64\x50\x6c\xc0\x80\x55\xcd\xd1\xbf\x3c\x09\x12\x3e\x30\xda\x38\x0c\x27\x55\xb2\x44\x67\xe8\x58\xe2\x8f\x91\xdd\x49\xfe\x5f\x84\x1c\x9c\x6b\xb7\x34\x8a\xa7\x55\xd1\x00\x67\x31\xcc\xcb\xeb\x75\x36\x09\x0e\x55\xd4\xfe\x8f\x47\x67\xb4\x82\x20\x45\x74\x42\x2c\x1c\xc6\xdd\xb1\x01\xe6\xa3\x43\x29\x6e\xc8\x18\x6d\x77\x77\x75\x09\x36\x25\x9b\xae\xc8\xd6\x1d\x25\x3c\xdc\x59\x38\x80\x36\xb0\x89\x79\x54\xc3\xa1\x41\xc5\xe4\x9b\x3d\x65\xe4\xa4\x77\xc9\xc1\xc7\x8f\xce\x58\x3a\xd3\x03\xa7\xde\xa9\x7d\x3e\xd0\x7f\xee\xf2\xdd\xf5\xea\xf1\xa0\xe0\xae\x4f\x3c\x5c\xdf\x93\x0c\xfa\x94\x93\xe2\x97\x8b\x7b\x3c\x46\x97\x79\xcd\x5a\x81\x79\x5d\x14\x64\xc3\x6f\xd6\x1c\xb3\x24\x29\xc9\x45\x4d\xf2\x52\x64\x97\x0f\x3a\x70\xe8\x84\x91\x55\xd7\x83\xeb\xc7\x2f\xa6\xc0\x09\xbd\x51\x90\xc2\x68\x5b\x3d\xb4\x9b\x14\x59\x06\x6d\xd1\xf7\x67\xc9\x9f\xa4\x45\xf3\xe9\x12\x76\x8f\xcb\x67\xed\xbf\x9c\x4f\xcf\xa7\xb3\xe1\xa6\x55\x65\xcc\x8a\x8c\x56\x47\x29\xae\xb6\xd7\xc4\x2b\x8f\x01\x6b\xf6\xe9\x12\x3b\x69\x09\x7d\x7a\xeb\x52\xf3\x60\x25\x86\xa3\x24\x7f\x94\x62\xfe\xbf\xd9\x7b\x9d\x48\x3c\xfe\x55\x61\x38\xdd\xad\x5c\x25\xc5\x7c\x36\x2b\x47\x31\x06\x10\xe0\x77\x41\x8a\xdf\x45\x96\x2b\xb2\xdb\xec\xbf\x22\x3b\x43\x56\xdd\xa5\xce\x3a\x7e\xca\xc4\x1f\xbd\xca\x81\x4c\x55\xe2\xfb\x18\xd5\x41\xdc\x3a\x6f\x45\x5a\xcc\x9b\x4d\x89\xfd\x32\xee\x5f\x01\xef\xa5\x48\x2d\x0c\xee\x02\x45\x8c\xb3\x14\xb1\xdb\x3c\x52\x06\xf9\xa1\x9d\x3e\x52\x2b\xf2\x2c\x45\xc2\x2e\x5d\x21\x0f\x71\x9d\x27\x26\x45\x46\x8a\xbb\xe5\xea\x7b\x71\x72\x56\x6e\x14\xeb\x76\xf1\x15\xac\x01\xe7\x75\x68\x25\xb2\xd7\x6a\xfc\x64\x29\x5a\x4d\xca\x9a\x8f\x0b\xb2\x8c\x0f\x9c\x86\x16\x8c\xa1\x4f\x2b\xaf\x0f\xda\xe0\x0e\x2f\x83\x02\x53\xd7\x8f\x8c\x2c\x1d\x52\x77\x2b\x70\xb0\xd1\x46\xb3\xc6\x93\xe4\x80\xa2\x18\x2e\xe4\xe2\xfa\xf2\xf6\xe3\x9b\xab\xeb\xe5\xc7\xf5\xe5\xcd\x87\xab\xc5\xe5\x40\x5c\x78\x72\xa7\x0a\x60\xcc\x48\xe0\x6e\x88\xf8\x17\x6d\xb0\xed\xc5\xc3\x30\x1a\x7d\x40\x8b\x21\xac\x3c\x6d\x30\xc5\xdb\x33\xbb\xb7\xc8\x43\x13\xae\x49\x94\x93\x86\x27\xda\x74\x90\xe2\x62\x76\x31\x1b\x2c\x07\xb5\xc7\xe8\xe4\x5f\x6f\x6f\x57\x89\x40\x5b\xcd\x1a\xcc\x12\x0d\x1c\xd7\xa8\xc8\x16\x41\x8a\x57\xa9\x2a\xeb\x12\xa9\xe2\x5e\xf8\x32\x91\x85\x4a\x29\x0c\xe1\x76\xef\x31\xec\xc9\x14\x0d\xbb\x76\xdf\x16\xb4\xa9\x3c\x26\xd2\x4e\xb7\xb0\xa1\x2b\xfb\x65\xf3\x04\x6a\x05\x4d\x55\x7c\x47\xd5\xa8\xee\x91\x31\x74\xcf\x38\x31\xd5\x17\x66\x2c\xc3\x69\xb8\x6a\x46\xed\x4a\x79\x20\xeb\x3c\xdd\x0b\x9f\x7d\xee\xb4\xef\xa7\x91\xb6\x99\x74\x80\x67\xfb\xe6\x93\xe7\xe7\xe3\x0b\x21\x92\x71\x13\xd4\x2c\x96\x4d\x36\x22\x0e\xca\x83\x7b\xf6\x19\xfa\x0d\x6d\xb8\x7d\x1e\xe5\x6d\x4f\x4a\x90\xbe\xb5\x61\x0f\x5b\xea\x98\xcd\xd6\xc6\xd5\x4a\x8a\x17\x9f\x17\xef\xee\xd6\xb7\x97\x37\x1f\x97\xd7\xeb\x2f\x2f\x26\x09\x89\xe5\x27\x14\xe5\x52\xee\x39\x65\xaa\x7c\x84\x87\x9e\x51\x68\x08\x24\x1f\xa1\x1a\x37\x64\xa4\xa1\xca\xbf\x01\x00\x00\xff\xff\xfd\x32\xb3\x08\x16\x0e\x00\x00")

func corednsYamlBytes() ([]byte, error) {
	return bindataRead(
		_corednsYaml,
		"coredns.yaml",
	)
}

func corednsYaml() (*asset, error) {
	bytes, err := corednsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "coredns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _traefikYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xcc\x4d\x4b\xc4\x30\x10\xc6\xf1\x7b\x3e\xc5\xb0\xb0\xc7\x4d\x5c\x14\x0f\x73\x53\x29\x28\x82\x88\x6f\x57\x99\xa6\xa3\x0d\x79\x69\xc8\x4c\x05\x15\xbf\xbb\xb4\xf4\xb6\xc7\x99\xe7\xcf\x8f\x6a\x78\xe3\x26\x61\x2a\x08\xf1\x5c\xac\x27\xd5\xc4\x36\x4c\xee\xeb\x68\x62\x28\x03\xc2\x2d\xa7\x7c\x33\x52\x53\x93\x59\x69\x20\x25\x34\x00\x85\x32\x23\x68\x23\xfe\x08\x71\xbb\xa5\x92\x67\x84\x38\xf7\x7c\x90\x6f\x51\xce\x46\x2a\xfb\x25\xf7\x0b\x80\x30\xaa\x56\x41\xe7\xf6\xbf\xf7\xaf\xd7\xdd\xd3\x43\xf7\xd2\x3d\xbf\x5f\x3d\xde\xfd\xed\x9d\x28\x69\xf0\x6e\x0d\xc5\x6d\xf0\xe1\x68\x2f\x2f\xec\x99\xd5\xcf\x1f\x03\x20\xac\x8b\x05\xd0\x7a\xf2\x96\x0b\xf5\x89\x07\x84\x9d\xb6\x99\x77\xeb\x20\x92\x4e\xfe\xff\x01\x00\x00\xff\xff\xf0\x93\x36\xe7\xe3\x00\x00\x00")

func traefikYamlBytes() ([]byte, error) {
	return bindataRead(
		_traefikYaml,
		"traefik.yaml",
	)
}

func traefikYaml() (*asset, error) {
	bytes, err := traefikYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "traefik.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"coredns.yaml": corednsYaml,
	"traefik.yaml": traefikYaml,
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
	"coredns.yaml": &bintree{corednsYaml, map[string]*bintree{}},
	"traefik.yaml": &bintree{traefikYaml, map[string]*bintree{}},
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
