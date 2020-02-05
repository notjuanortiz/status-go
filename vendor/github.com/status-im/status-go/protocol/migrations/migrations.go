// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 000001_init.down.db.sql (65B)
// 000001_init.up.db.sql (2.719kB)
// 000002_add_last_ens_clock_value.down.sql (0)
// 000002_add_last_ens_clock_value.up.sql (77B)
// doc.go (377B)

package migrations

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var __000001_initDownDbSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\xce\x48\x2c\x29\xb6\xe6\x42\x12\x29\x2d\x4e\x2d\x8a\xcf\x4d\x2d\x2e\x4e\x4c\x4f\x45\x95\x49\xce\xcf\x2b\x49\x4c\x06\x29\x07\x04\x00\x00\xff\xff\x61\x86\xbd\x5f\x41\x00\x00\x00")

func _000001_initDownDbSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_initDownDbSql,
		"000001_init.down.db.sql",
	)
}

func _000001_initDownDbSql() (*asset, error) {
	bytes, err := _000001_initDownDbSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_init.down.db.sql", size: 65, mode: os.FileMode(0644), modTime: time.Unix(1578682784, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5e, 0xbb, 0x3f, 0x1, 0x75, 0x19, 0x70, 0x86, 0xa7, 0x34, 0x40, 0x17, 0x34, 0x3e, 0x18, 0x51, 0x79, 0xd4, 0x22, 0xad, 0x8f, 0x80, 0xcc, 0xa6, 0xcc, 0x6, 0x2b, 0x62, 0x2, 0x47, 0xba, 0xf9}}
	return a, nil
}

var __000001_initUpDbSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\xd1\x6f\xe2\xb8\x13\x7e\xe7\xaf\x18\xe9\xf7\x40\x2b\xd1\x9f\xf6\xa4\xbd\xbd\x93\xfa\x44\xd9\xf4\x0e\x1d\x07\x2b\x9a\x9e\xba\x4f\xd6\xe0\x0c\xc4\xc2\xb1\x23\x7b\x02\x8b\xb4\x7f\xfc\xc9\x81\x40\x0c\x81\xee\xea\xfa\x50\xb5\x33\xe3\x19\xcf\x37\xf3\x7d\xce\x68\x9e\x0c\xd3\x04\xd2\xe1\xd3\x24\x81\xf1\x33\x4c\x67\x29\x24\x6f\xe3\x97\xf4\x05\x64\x8e\xec\xe1\xae\x07\xa0\x32\xf8\x67\x38\x1f\xfd\x39\x9c\xc3\x97\xf9\xf8\xef\xe1\xfc\x2b\xfc\x95\x7c\x85\xd9\x14\x46\xb3\xe9\xf3\x64\x3c\x4a\x61\x9e\x7c\x99\x0c\x47\xc9\xa0\x07\x60\xb0\xa0\x63\x7c\xc8\x37\x7d\x9d\x4c\x82\x43\x5a\x6d\xdd\x85\x07\x3e\x27\xcf\xc3\xd7\x49\x0a\xfd\xff\xe1\x2f\xbf\xff\x96\xfd\xda\x0f\xb1\xbc\x2b\x09\xc6\xd3\x34\x4a\x80\x92\xd5\x86\xe0\x69\x36\x9b\x24\xc3\xe9\x65\x86\x74\xfe\x5a\xdf\x80\x55\x41\x9e\xb1\x28\x2f\x32\x64\xa4\x89\x29\x13\xc8\x42\x6a\x2b\xd7\x62\x83\xba\x8a\x0b\x1d\xb3\x7d\x08\x07\xca\x6a\xa1\x95\x14\x6b\xda\xc1\xd3\x64\xf6\x14\x4c\x95\xd9\x28\xda\x52\x26\x0a\xf2\x1e\x57\x24\xa4\xad\x0c\xdf\xc8\xa1\xd1\xff\x68\xb9\x3a\xf4\x90\xf7\x58\xb0\xa0\x62\x41\xce\x9f\xff\x9f\xab\x52\x54\x65\x86\x4c\x7b\x57\xef\xfe\xb1\xd7\x8b\xe6\x29\xad\x61\x94\xa7\x21\xa6\xc9\x5b\xfa\x23\x13\xc4\x2c\x73\xe4\xfd\x3e\xbe\x0d\x5f\x3d\xda\x0b\x2b\x19\x2f\x36\xe4\xd4\x52\x51\x76\x1c\x4e\xd3\xd6\xf3\x70\xf2\x92\x9c\x47\x09\xbc\x85\x17\x6a\x85\x1d\xc5\x55\x46\x86\x95\xb4\xe6\xd2\x55\xe6\x96\xed\xa5\xb9\x46\x73\x0f\x51\x76\xa3\x9e\xdf\x79\xa6\x42\x30\xae\x4e\x18\x67\xb4\x51\x92\x84\x32\x4b\x7b\xb4\xb1\x53\x8b\x8a\x49\xb0\x15\x8c\x7a\x1d\xd7\xab\xd1\x7f\x78\x80\x31\xf7\x3d\xa8\xa2\xb4\x8e\xd1\x30\x70\x8e\xe1\x97\xf2\xc0\xb8\xd0\x04\x39\x7a\x70\x76\xab\x32\x40\x0f\x5b\x02\x47\x7a\x07\xd6\x80\xe2\x70\x78\x9b\x93\x09\x87\x35\x15\xa1\x57\xb3\x02\x65\x96\xca\x28\xa6\x07\x2f\x9d\xd5\xfa\xff\xbd\x1b\x84\xad\x3c\xb9\x66\x79\xf6\x33\xff\x59\xea\x02\x6c\x73\xe5\x4b\x72\x22\xa2\x50\xf2\x47\x12\x33\x19\xc0\xdb\xca\xc9\x8e\x5d\x08\xc8\x79\x56\x06\x59\x59\x73\x44\x0e\x80\xe9\x1b\x77\x8a\x02\xd4\x5b\x4a\x86\x45\x27\xe5\xa1\xee\xaa\x2d\x29\x87\x7c\x57\x29\x0e\xb5\x70\x89\x56\xe3\xb1\x57\x5b\x89\x5a\xdc\x8e\xc9\x55\x46\xd7\x37\x19\xc0\x91\x2f\xad\xf1\x61\x15\xe2\x6b\x35\x92\xd0\xf4\x72\xb8\xd0\x15\xee\x1f\xa0\x24\x32\xd7\x35\xad\x55\xd5\x56\xbc\xb2\xca\xac\x84\x67\xe4\xca\xc7\x95\x4b\x74\x9e\x32\x51\xe3\x7c\x82\xdd\xe1\x56\x94\xb8\xd3\x16\xb3\x96\xd5\xb3\x92\x6b\x72\xa2\x44\xb9\x3e\xdd\xb2\xb1\xe6\xe8\xf3\x38\xb7\xb4\x45\x81\x26\x6b\xe1\x15\xdb\xf7\x9d\x75\xba\x1a\x29\xe9\x74\x2e\x9d\x2d\xba\x3d\x61\x27\x1c\x4a\xee\xf6\xb2\x43\xe3\xc3\x63\x60\xcd\x8d\xdb\x7a\xb5\x32\xc8\x95\xa3\x56\xe7\x47\x1f\x23\xd7\xb3\x68\x8b\xe6\x78\xfa\x39\x79\x03\x95\x7d\x13\x87\xed\x9e\x4d\x63\x4e\xdd\xed\xed\xf7\x8f\x1d\x27\x08\x9d\xcc\xc5\x62\x77\xdc\xac\xd9\x14\xce\x4e\xef\x51\xae\x16\x9e\xdd\x5d\xff\xc3\x7f\xfc\xe9\xc3\xf7\xef\xed\xc5\x1a\xc0\xc3\xa7\x8f\x03\xf8\xf4\xf1\x3e\x38\x54\x36\x68\x68\x30\xa8\xb7\xf9\xf2\x71\x88\xb5\x23\x2c\x4a\x24\x1d\x3f\x27\x1c\xef\x93\xaa\xd6\x62\x4f\x67\x0f\x65\x2d\xbf\x54\x0f\xfc\xfc\x0d\x3d\xb8\xf8\x3a\x0f\x1d\xd5\x47\xb1\x62\x5b\x20\x2b\x89\x5a\xef\xae\x47\x77\x51\xd3\x91\x54\xa5\x22\xc3\x27\xe1\x6f\xb3\xe5\x1d\xcc\x42\x46\x32\xab\xa0\x96\xa7\x85\xf4\xe1\x79\xd8\xa0\x56\xe1\xd5\xa9\x91\x6c\x0a\xc7\xec\xb9\xe4\x54\xd4\xf9\xb5\x15\x6f\x0f\x63\xdf\x01\xbb\xdd\x09\xbd\x60\x5a\x2a\x57\x43\x4d\xa6\xb1\xc4\x4c\x88\xeb\xb4\x2e\x7b\x0e\x5d\xf3\x39\x75\xf6\x0d\xd4\xc9\x98\x4e\x2c\x22\x28\x02\x1f\xde\x47\xec\xae\xf5\xf7\xfd\x63\xef\xdf\x00\x00\x00\xff\xff\xd7\x95\x8b\x6b\x9f\x0a\x00\x00")

func _000001_initUpDbSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_initUpDbSql,
		"000001_init.up.db.sql",
	)
}

func _000001_initUpDbSql() (*asset, error) {
	bytes, err := _000001_initUpDbSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_init.up.db.sql", size: 2719, mode: os.FileMode(0644), modTime: time.Unix(1578682784, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x60, 0xdc, 0xeb, 0xe, 0xc2, 0x4f, 0x75, 0xa, 0xf6, 0x3e, 0xc7, 0xc4, 0x4, 0xe2, 0xe1, 0xa4, 0x73, 0x2f, 0x4a, 0xad, 0x1a, 0x0, 0xc3, 0x93, 0x9d, 0x77, 0x3e, 0x31, 0x91, 0x77, 0x2e, 0xc8}}
	return a, nil
}

var __000002_add_last_ens_clock_valueDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _000002_add_last_ens_clock_valueDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_add_last_ens_clock_valueDownSql,
		"000002_add_last_ens_clock_value.down.sql",
	)
}

func _000002_add_last_ens_clock_valueDownSql() (*asset, error) {
	bytes, err := _000002_add_last_ens_clock_valueDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_add_last_ens_clock_value.down.sql", size: 0, mode: os.FileMode(0644), modTime: time.Unix(1580459788, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4, 0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95, 0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55}}
	return a, nil
}

var __000002_add_last_ens_clock_valueUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x41\x0a\x85\x20\x10\x06\xe0\xfd\x3b\xc5\x7f\x84\xb7\x6f\x35\xa5\x41\x30\x8d\x10\xe3\x5a\x64\x70\x95\xe8\x42\xeb\xfc\x7d\xc4\xea\x2f\x28\xad\xec\x61\xbd\xcd\x6c\x73\x80\x9c\xc3\x16\x38\x9e\x82\x9a\xc7\x4c\xa5\x8d\x64\xb5\xdb\x9d\xde\x5c\x9f\x82\x43\x14\x12\x14\x12\x99\xe1\xfc\x4e\x91\x15\xff\xe5\xf7\x05\x00\x00\xff\xff\xd0\x66\x8a\xf7\x4d\x00\x00\x00")

func _000002_add_last_ens_clock_valueUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_add_last_ens_clock_valueUpSql,
		"000002_add_last_ens_clock_value.up.sql",
	)
}

func _000002_add_last_ens_clock_valueUpSql() (*asset, error) {
	bytes, err := _000002_add_last_ens_clock_valueUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_add_last_ens_clock_value.up.sql", size: 77, mode: os.FileMode(0644), modTime: time.Unix(1580459768, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x4d, 0x3, 0x8f, 0xd5, 0x85, 0x83, 0x47, 0xbe, 0xf9, 0x82, 0x7e, 0x81, 0xa4, 0xbd, 0xaa, 0xd5, 0x98, 0x18, 0x5, 0x2d, 0x82, 0x42, 0x3b, 0x3, 0x50, 0xc3, 0x1e, 0x84, 0x35, 0xf, 0xb6, 0x2b}}
	return a, nil
}

var _docGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8f\xbb\x6e\xc3\x30\x0c\x45\x77\x7f\xc5\x45\x96\x2c\xb5\xb4\x74\xea\xd6\xb1\x7b\x7f\x80\x91\x68\x89\x88\x1e\xae\x48\xe7\xf1\xf7\x85\xd3\x02\xcd\xd6\xf5\x00\xe7\xf0\xd2\x7b\x7c\x66\x51\x2c\x52\x18\xa2\x68\x1c\x58\x95\xc6\x1d\x27\x0e\xb4\x29\xe3\x90\xc4\xf2\x76\x72\xa1\x57\xaf\x46\xb6\xe9\x2c\xd5\x57\x49\x83\x8c\xfd\xe5\xf5\x30\x79\x8f\x40\xed\x68\xc8\xd4\x62\xe1\x47\x4b\xa1\x46\xc3\xa4\x25\x5c\xc5\x32\x08\xeb\xe0\x45\x6e\x0e\xef\x86\xc2\xa4\x06\xcb\x64\x47\x85\x65\x46\x20\xe5\x3d\xb3\xf4\x81\xd4\xe7\x93\xb4\x48\x46\x6e\x47\x1f\xcb\x13\xd9\x17\x06\x2a\x85\x23\x96\xd1\xeb\xc3\x55\xaa\x8c\x28\x83\x83\xf5\x71\x7f\x01\xa9\xb2\xa1\x51\x65\xdd\xfd\x4c\x17\x46\xeb\xbf\xe7\x41\x2d\xfe\xff\x11\xae\x7d\x9c\x15\xa4\xe0\xdb\xca\xc1\x38\xba\x69\x5a\x29\x9c\x29\x31\xf4\xab\x88\xf1\x34\x79\x9f\xfa\x5b\xe2\xc6\xbb\xf5\xbc\x71\x5e\xcf\x09\x3f\x35\xe9\x4d\x31\x77\x38\xe7\xff\x80\x4b\x1d\x6e\xfa\x0e\x00\x00\xff\xff\x9d\x60\x3d\x88\x79\x01\x00\x00")

func docGoBytes() ([]byte, error) {
	return bindataRead(
		_docGo,
		"doc.go",
	)
}

func docGo() (*asset, error) {
	bytes, err := docGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "doc.go", size: 377, mode: os.FileMode(0644), modTime: time.Unix(1574354941, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xef, 0xaf, 0xdf, 0xcf, 0x65, 0xae, 0x19, 0xfc, 0x9d, 0x29, 0xc1, 0x91, 0xaf, 0xb5, 0xd5, 0xb1, 0x56, 0xf3, 0xee, 0xa8, 0xba, 0x13, 0x65, 0xdb, 0xab, 0xcf, 0x4e, 0xac, 0x92, 0xe9, 0x60, 0xf1}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"000001_init.down.db.sql": _000001_initDownDbSql,

	"000001_init.up.db.sql": _000001_initUpDbSql,

	"000002_add_last_ens_clock_value.down.sql": _000002_add_last_ens_clock_valueDownSql,

	"000002_add_last_ens_clock_value.up.sql": _000002_add_last_ens_clock_valueUpSql,

	"doc.go": docGo,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"000001_init.down.db.sql":                  &bintree{_000001_initDownDbSql, map[string]*bintree{}},
	"000001_init.up.db.sql":                    &bintree{_000001_initUpDbSql, map[string]*bintree{}},
	"000002_add_last_ens_clock_value.down.sql": &bintree{_000002_add_last_ens_clock_valueDownSql, map[string]*bintree{}},
	"000002_add_last_ens_clock_value.up.sql":   &bintree{_000002_add_last_ens_clock_valueUpSql, map[string]*bintree{}},
	"doc.go":                                   &bintree{docGo, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
