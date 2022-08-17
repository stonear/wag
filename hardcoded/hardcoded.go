// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../_hardcoded/doer.go (10.342kB)
// ../_hardcoded/middleware.go (1.685kB)
// ../_hardcoded/tracing.go (6.177kB)

package hardcoded

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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
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

var __hardcodedDoerGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5a\xff\x6f\xdb\x46\xb2\xff\x59\xfa\x2b\x26\x02\xee\x42\x26\x32\x25\xe7\xce\xed\x41\x7d\x3e\xa0\x75\xdc\xd7\x02\x69\x12\x54\xee\x6b\x80\x20\xb8\xae\xc8\x91\xb4\xcd\x72\x57\xdd\x5d\x5a\x16\x7c\xfe\xdf\x1f\x66\x76\x49\x91\x12\xe5\xe4\xae\xc5\xc3\xbb\x00\xb1\x25\x72\xe6\xb3\xb3\xb3\xf3\x7d\xbd\x11\xf9\x47\xb1\x42\xc8\x95\x44\xed\x87\x43\x59\x6e\x8c\xf5\x90\x0c\x07\xa3\xc5\xce\xa3\x1b\x0d\x07\xa3\xdc\x68\x8f\x77\x9e\x3e\xa2\xce\x4d\x21\xf5\x6a\xf2\xab\x33\x9a\x1f\x58\x6b\x2c\x53\x49\x33\x91\xa6\xf2\x52\xd1\x17\x65\x56\xf4\xab\x14\x7e\x3d\xb1\x42\x17\xf4\x45\xa3\x8f\xbf\x26\x6b\xef\x37\xf4\xd9\xed\x74\x4e\xbf\xbd\x2c\x71\x34\x1c\x0e\xb6\xb9\x82\xd1\x4a\xfa\x75\xb5\xc8\x72\x53\x4e\xae\x14\xde\xa2\x9d\x6c\xc5\x8a\xfe\x07\x19\x95\x59\xad\xd0\x12\x75\x9b\x52\x2c\xf1\x6e\xb2\xde\x39\x6f\xe5\xdd\xd9\xca\xd4\x1f\x47\x5d\xaa\xc2\x68\x73\x2b\xf4\x5a\x16\x38\xc1\x5b\xd4\xde\x99\xca\xe6\x38\x1a\xa6\xc3\xe1\x64\x02\x85\x41\x0b\xd2\x81\xd0\x20\xb5\x47\xbb\x14\x39\xc2\xd2\x58\x18\x15\x46\xea\xd5\x08\x48\x6e\xb0\xf8\x5b\x85\xce\x3b\xd8\x18\xe7\xe4\x42\xed\x60\x2b\xfd\x1a\xb6\x56\x6c\x36\x52\xaf\x86\x7e\xb7\xc1\x08\xd5\x80\xdc\x0f\x07\x2f\x4d\x92\xc3\x33\x42\xc8\xae\x78\x23\x63\xb0\xf1\xfb\x8f\x01\x31\x85\xa4\xfe\xee\x36\x46\x3b\x1c\x03\xab\x37\x1d\x3e\x0c\x03\xaa\xd9\xbc\x16\x25\x5e\xf9\x3b\x70\xde\x56\xb9\xbf\x7f\x60\xb9\x17\xc2\x61\xc4\xf8\x4e\xe8\x42\xa1\x85\x0d\xda\xa5\xb1\xa5\x03\xbf\x46\x7e\xdf\x91\x3d\xa0\xd1\xe3\x97\x24\xe7\x1e\x6c\x59\xe9\x1c\x92\xa2\x79\x95\xc2\xef\x10\x9b\x76\x6d\xd1\x57\x56\x43\x9e\xbd\x34\x49\xd8\xc8\x64\x02\x16\xbd\xdd\xd5\x92\xd2\x17\x89\x0e\x2e\xa6\xef\xba\xfa\x0d\x42\x32\x6d\x4b\x4a\x02\x2d\x60\xff\x8f\x14\xcd\xcb\xd8\xdd\x5b\xa3\x64\xbe\x83\x1f\xf7\x9f\xe3\x7a\xad\x27\x50\xe0\x52\x6a\x74\x20\x02\x32\x6c\xf8\x71\x16\x16\x6b\x13\x76\x0e\x6f\x32\x81\x6f\x44\xfe\xd1\x2c\x97\x0e\xc2\x96\x82\x66\x75\x55\x2e\xd0\x82\xd0\x05\x78\x59\x4a\xbd\x02\xb3\x8c\xc0\xc2\x7b\x2c\x37\xde\x65\xc3\x41\xcd\x9a\xa4\xf0\xfe\x03\xd9\x7a\xf6\xb2\xb2\xc2\x4b\xa3\x19\x99\x57\x05\x8b\x39\xca\x5b\x0c\xb8\x6d\x45\x8c\x41\x38\xd8\xa2\x52\xf4\x9b\x5e\x5a\x74\x95\xf2\x60\x96\xcc\x5d\xfb\x53\x3c\x9e\xa7\x0e\x7e\x79\x69\x7e\x81\x12\xfd\xda\x14\xd9\x70\xc0\xe8\x49\xe7\xc4\xc6\x70\xe2\xc0\x16\xc6\xa8\xa8\xb3\xb9\xd4\x2b\x85\x8f\x69\xce\xaf\x85\x6f\x4e\x4f\xd4\xd2\x82\xd1\x39\x06\x6d\x1e\x43\x74\xac\xb6\x47\xa1\xc2\xc3\xce\x54\xe0\xd6\xa6\x52\x45\xb3\x0c\x36\xd8\xe7\x0e\x73\xa3\x0b\x10\x4b\x4f\xee\xe5\x61\x29\xa4\x72\x59\x34\xdb\xa3\xf5\x52\x38\xad\xf9\x96\x71\x1e\xbc\xb9\x3f\x87\x67\xc0\x4f\xe6\xbc\xda\x43\xdb\x8a\x60\x2b\x95\x8a\x92\x69\xa3\xcf\xde\xbe\x99\xdf\x8c\xc3\xa7\xaf\x6f\xae\xbe\xdb\x07\x07\xde\xcc\xc5\xbb\x77\x19\xb1\xde\xbc\x79\xf9\x66\x06\xdf\x7b\xb2\x56\x07\xda\x78\xc8\x2b\x6b\x29\x98\xed\x6a\x6b\xd1\xbb\x70\x08\xb5\x3a\xb0\x80\xc5\xee\xc4\xe1\x3e\xb2\xe1\x70\xda\x16\x7f\x83\x83\x13\xb7\xe8\x36\x7d\xc7\xde\x3e\x7a\xd2\x89\x5c\xf2\xd3\x27\x97\xa0\xa5\x82\x7f\xfe\x93\x76\x94\xfd\xc0\xc6\x04\x97\x97\x30\xa2\x0d\x8f\xfa\x9e\xd3\xf6\xe9\xc5\x70\x30\xa0\xb5\xb2\xb9\x17\xbe\x72\x57\xa6\x40\xf8\x2f\xb8\x98\x4e\x09\xbc\xd6\xf8\x52\x28\x87\xc3\xc1\x43\x73\x04\xde\x56\x18\xd5\x7c\x7d\xb7\x31\x1a\xb5\x97\x42\xf5\x5a\x9f\x06\xdc\x53\x74\x9c\x38\x58\xdd\x09\xfe\xc7\x4d\x6f\x29\x6f\x29\x1e\xc6\xa7\x1c\xcf\x5b\xab\x28\x8a\x06\xb9\x45\xe1\xc8\xc3\xb7\x42\x7a\xb6\x0f\xc7\xb1\x17\xfd\x16\x51\x37\x07\x3f\x83\xf3\xe9\x74\x0c\x2f\xe8\xc7\x5f\xe9\xc7\xdf\xe8\x07\x05\x88\xf3\x2f\xa6\x53\x28\xa5\x52\x32\x58\xb1\x83\xe7\x93\x33\xa8\x36\xe0\x0d\x5c\xfc\x09\x7e\x95\xde\xa3\xad\x8f\xb6\x7f\x17\x9f\x61\xd0\x30\xbb\x84\x52\x7c\xc4\xe4\xe0\xf5\x18\x2e\xd2\xe1\x40\xe3\x1d\x53\x9c\x4f\xa7\xb5\x91\xff\xb0\x97\x68\x38\xb0\xba\xa0\xd7\x94\xb2\xb3\xd7\xb8\x4d\xea\x0f\x73\xce\x93\x09\x33\xbc\x36\xdb\x24\xcd\x7e\xd2\xf2\xee\xb5\xd0\x26\x49\xd3\x74\x38\x40\xe2\x9a\x66\xd3\x0b\x98\x4c\x78\x5b\x17\x94\x81\x72\xd4\x3e\xee\x6b\x38\xa0\x3c\x2a\x23\xf8\x8a\xe3\x7a\x6d\x11\xef\xe5\x07\xb8\x04\x16\xed\x39\x74\x84\x4e\x92\xc4\xea\x22\xfb\x56\x19\xe1\xbf\xf8\x6b\x92\x3e\x7b\x91\x9e\x9d\xa7\xcf\xf0\xd9\x32\x3e\x21\x26\x5a\x3f\x6c\xec\xd9\x25\xbc\x68\x5b\x95\x45\xff\xff\xd6\x77\x4f\x1d\xf0\x7f\xbe\x03\xbf\x36\xfd\x59\x23\x38\x29\x25\x48\x4d\x55\x5d\x50\x1e\x39\x54\x93\x3b\x62\x16\xee\x02\x3c\xee\xb8\x14\x0e\xca\x8d\xdf\x81\x53\x32\xc7\x5a\xb9\x1d\x84\x7f\x33\x0b\x74\xc3\xbe\x50\x5b\xb1\x6b\xc5\x0b\xda\xff\x89\xd5\xfe\xd5\x84\xdb\x92\x21\xa8\xb5\xae\xf6\x58\x41\x57\xa1\xf8\xee\xaa\xe1\x67\xe9\xd7\x6d\x1d\x35\xda\x00\x8d\x5b\x88\xf5\x7a\x30\x63\x73\x8b\xd6\xca\x22\xd6\x15\xa1\x84\x06\xb3\xf8\x15\x73\xff\xd4\x35\xc5\x58\x53\x06\xf1\x8e\x0e\xe0\x93\xdc\xdf\xd5\xa0\x59\x94\x67\x0c\x27\x4a\xae\xf4\x90\xb2\x5d\x05\xc6\x37\x84\xff\x3f\x42\x55\x48\xc8\xe3\xce\x3e\xef\x1f\x3a\xc8\x5c\x31\xd6\x45\xe9\xb3\xa6\x16\xfc\x23\xca\xd2\x7a\x89\x31\x98\x8f\x1c\x96\x6a\x81\x93\x34\x0b\xc2\x75\xe5\x4a\xb3\xa4\xbd\x4d\x76\xb4\x27\xe6\x63\xed\x15\x8d\x2e\x2e\xa1\xc8\x5a\xdf\xd9\x49\x9a\xcc\x42\xeb\xec\xdf\x65\x7b\xc3\x1c\x0e\x6e\x85\xed\xf3\xf1\xf0\xa2\xf1\xf3\x21\xd7\x7c\x73\x71\x8b\x9d\xc2\x68\x61\x0a\xca\x51\x90\x53\xa5\xbf\x45\x58\x33\x81\x09\x8b\x65\xf0\xc6\xaf\xd1\x6e\x65\x78\xc7\x51\x90\x09\x84\xb2\x28\x0a\x32\x1f\x51\x30\x2e\x37\x0b\xd5\x72\x89\x16\x8c\x6e\xa2\x5b\xd1\x59\x8a\xd9\xa9\xee\xca\x60\x8e\xc8\x6c\x24\xf0\x6c\x32\x71\x9e\xb6\x73\x8b\x76\xa9\xcc\x96\xfb\x2d\xe6\x90\x46\xbb\xc9\x8b\xbf\x4c\xbf\x9c\xfe\xed\xcb\x2f\x26\xb4\x96\xd4\xab\x33\x92\xf8\xcc\x2c\xcf\x88\xf7\x2c\x62\x9f\x51\xe2\x35\x95\x3f\x2b\x4d\x21\x97\x14\x1b\x9a\x37\xce\x0b\x1f\x75\xb1\xa8\x96\xf0\xfe\x03\x75\xa7\x7c\x06\x36\xfb\x86\x36\x1f\xe3\x1d\x9d\x46\x57\x61\x83\xc1\xa2\x5a\x86\x48\x79\x09\xa1\x43\xcd\x7e\x44\x51\x7c\xad\x54\x12\x78\x29\x7f\x74\xa3\x26\xa1\xd4\x46\xab\xa5\x62\xee\xe1\x80\x4e\xf2\x61\x18\x32\x59\x5d\xf7\x52\xda\xfb\x8a\xa3\xdf\x57\xf5\xb3\xe7\xcf\x99\xbf\x57\xb4\x81\x2d\x2c\xf1\x44\x39\x5e\x9b\xcd\x95\x32\x0e\x6d\xc2\xcd\x36\x65\xda\x6f\x58\xfd\xc9\xa2\x5a\x72\x5e\x1b\x44\x8c\x4b\xb0\x45\x14\x81\xa3\x73\xbd\xa1\x22\x2b\xa8\xb5\xca\xc7\x60\xe3\x36\x6a\xc9\x2e\x2f\x41\xa1\x4e\x6a\xd3\x4b\x29\xe4\x3f\x69\x1b\x5f\xcc\x34\x21\xb3\x30\x5e\x1a\x64\x5c\x58\x14\x1f\xe3\x5a\x93\x09\xb0\x84\x75\xe3\xc1\x36\x19\xcd\x6d\xb9\x7f\x20\x43\x66\xd4\x52\xd5\x42\xb8\xcd\x81\x32\xdd\x86\x77\x92\x31\x1c\x19\x3c\xe3\x87\x3a\x5b\x21\x6e\x1a\x49\xdf\xc7\x1d\x7c\x48\xbb\xb9\x3c\x0a\x19\xc3\x72\x2e\x6d\x5e\x49\xff\x0d\xc9\x8a\x96\x1b\x44\x59\x6e\x14\x96\xd4\xd5\x87\x58\x17\x28\x60\x11\x48\x60\x43\x1d\x99\xd5\x9c\xd4\xe7\xe8\xa1\xc0\x45\xb5\x22\x27\xa1\xd3\xa3\xdf\x66\x83\x56\x78\xec\x30\x4b\x0d\x02\x4a\xca\x81\xdb\x35\x5a\xa4\x9e\x83\x1d\xc0\x68\xb5\x03\x65\x56\x2d\x51\x80\x6d\x34\xd4\x7a\x4c\x43\x1a\x59\x28\x93\x7f\xe4\x0a\xa1\xa9\x2e\x96\xd6\x94\xb0\x32\x94\xf9\xfc\xda\x9a\x6a\xb5\x8e\x79\xaf\x67\x4b\x8f\xf5\xbc\x41\xfe\xf0\x8f\x1b\xb7\x41\x04\x78\x2d\x4a\x24\x4e\xa9\x57\xc3\x41\x18\x97\x04\xaa\x6d\xae\xb2\x9f\xc5\x2a\x04\xcc\x57\xfc\x82\xb4\x49\xfe\x12\x59\xe7\xf3\xeb\x37\x3a\x47\x70\x3b\x9d\x67\xf4\x89\x75\xfd\x5d\x18\xa6\xcc\xe7\xd7\xd7\xb7\x94\x3e\xa4\x03\x2c\xa9\xac\xe3\x3a\x67\x3f\x75\x81\x5b\x29\xc0\xa1\xbd\x45\x7b\xe6\x88\x30\xcc\x58\x32\xae\x9e\xd0\xe5\x56\x2e\x42\xd1\x4c\x1a\x66\x6d\x51\x35\x20\xea\xc5\xa3\x1a\x0e\x57\xdb\xeb\xe0\x86\x5e\x3f\xf2\x2f\xec\x19\x7e\xf9\xd5\x19\x3d\x1b\x11\xd8\xe8\x97\xe1\x80\xd5\xf1\xd9\x5c\x5a\x94\xcc\x15\xb3\xc8\x95\xa9\xb4\xef\xe3\x92\xe1\x71\xe4\xb2\x2d\x6a\xe2\xbe\xa6\xd0\x73\x92\xf7\x90\x1b\x1b\xea\x86\xf7\x6d\xa8\x9e\xc5\xea\x58\xf4\x1e\xde\x3d\x35\x01\x7c\xef\xae\x3a\x96\xf4\x66\x83\xba\x0d\xc0\x45\x47\x03\x20\x7b\xa8\x59\x01\x46\x29\xa9\x57\x2c\xd6\xb7\x42\xaa\xca\xe2\x63\x0a\x38\xa6\x3e\x06\x51\x8a\x7c\xbc\x0d\xf6\x18\x48\x87\xfa\x14\xd8\xbc\xca\x73\x74\xee\x33\xc1\x22\xf5\x21\xd8\x7c\x6d\xac\x8f\x5a\xc0\xe2\x53\xdb\xeb\x52\x1f\x61\xb5\x04\xfa\xb4\xaa\x4e\x08\x74\xb3\xa6\x34\xf9\xd6\x18\xf5\x23\x52\xa5\x46\x42\x9d\x06\x39\xa6\x3e\xc2\x93\x25\x9a\xca\x7f\xa6\x50\x91\x9a\x40\xae\x42\xab\x73\x65\x74\x6c\x7a\xae\xef\x30\xaf\x28\xa5\x07\xe3\xee\x82\xe4\x8f\x53\x13\xe0\x2b\xe1\x51\xe7\xbb\x1b\xe3\x85\xfa\x01\x85\x7e\xdc\xb8\x55\x8b\xfa\x1f\x25\x0a\x32\xcc\xba\x20\x54\x66\xc5\xf1\x21\x51\x7d\x71\x6d\x0c\x47\x81\x84\xb3\x9b\xca\x5e\x99\x55\x42\x0c\xdf\xeb\xa5\x19\xc3\x68\x34\x86\x52\x6c\xde\x87\x20\xf0\xa1\x19\xfa\xdd\x3f\x50\xda\xea\xfa\xf5\xac\xcf\x93\x31\x6b\x47\x8a\x31\x31\xb5\xdc\xb9\x97\x85\x98\xf6\x01\x62\xcf\xd2\xf2\xe2\x63\xbe\xc8\xb2\xa7\x61\xbe\x5e\xe7\x9d\x75\xf9\xfa\xc2\x01\x33\xf7\xf9\xec\xec\x70\x73\xc7\x34\x3d\xbc\x5d\x57\x9d\xf5\xf2\x76\x68\x4e\x62\xd4\x0e\xf1\x18\x46\xa4\x39\xc2\x38\x70\xcc\x59\xdf\x1e\xba\x34\xc7\x10\x9d\xe5\xfb\xd5\x70\x6a\xf9\x1e\x37\x9c\x1d\xf2\x1e\xd3\x1c\xc3\x44\xef\x7b\x4c\x84\x48\xc3\xbc\x9f\x72\xba\x19\x60\xf6\x09\x37\x66\x1c\x75\xe0\x99\xc7\x36\x88\xd9\xa1\xf7\x8e\x87\x83\x87\x6e\x8b\x76\x5c\xc3\xa4\x20\xb5\xf4\x49\x1a\xc7\xe8\x6f\xd1\x4a\x53\xc8\x9c\xa7\x6a\xca\xac\xc2\x9c\x5d\x0b\xd5\x2d\xa3\xa8\x1e\x13\xce\x49\xc7\x25\x98\x43\xef\xb9\x9a\x69\x55\x5b\x7e\x6d\xd1\xad\x8d\x2a\x1c\x57\x5c\x95\x2e\xd0\x3a\x2f\x74\xc1\x23\x85\xcd\x46\xc9\x3c\xb4\xf8\x0b\x5c\x8b\x5b\x69\x6c\xc6\xec\x3f\xe9\xa5\xb1\xbe\xd2\xc2\xa3\xda\x8d\xdb\xc5\x4b\x61\xd0\xe9\xa7\x3e\xb6\x46\xb0\xb2\x28\x3c\x6c\xc5\x8e\x44\xc1\xbb\x4d\x28\x82\xa5\x63\x94\x42\x78\x31\x06\x67\xa8\x9f\xb2\xe8\x8c\xf5\x5c\x48\x56\x56\xf3\x1d\x80\x06\xe9\x1d\xb8\x6a\xc3\x97\x69\xd4\x2f\x2c\xac\x11\x45\x2e\xdc\xd1\x36\x4a\x2a\x76\x73\xc7\xa5\x13\x4f\xfd\x8f\xeb\x27\x48\xe6\xf3\xeb\x34\x88\x3f\xc7\x70\x39\xe0\x66\x93\xc9\x23\xd7\x5f\xd2\xb9\x0a\xdd\xe4\xe2\x8b\xac\xa9\x0a\x63\x69\x47\xdd\x02\x9d\x56\x38\x90\x41\xe4\x99\x7b\x8b\xa2\xac\xaf\x64\x66\x97\xb5\x5e\x78\xf6\xd7\x7e\xc7\x85\x7b\x1f\x53\x36\xf7\xc2\x7a\x7e\xad\xa4\xf3\xa8\x39\x06\x5b\x06\xd3\xe8\xb3\x57\xfc\x30\x19\xf9\x7c\x33\x1a\xc3\xe8\xfc\xc5\x97\xd9\x34\x9b\x66\xe7\xb3\xe9\xe8\x44\xff\xb5\x11\x5a\xe6\x09\xb5\x25\xdc\x2b\x0c\x07\x83\x95\xe1\xbd\x67\x73\x52\x51\xb2\x5f\xa6\x4f\x9e\x34\xd0\xb7\xb6\x4a\xea\xfb\x19\xd9\xe6\xf0\x16\xed\x0e\xce\xa7\x50\x4f\x62\x17\x98\x8b\x2a\x76\x39\x84\x86\x56\xde\x35\x47\x23\x2c\xd2\x91\x0a\x38\x9f\x9e\xd5\x83\x52\x46\x8b\x6e\x0b\x5b\xa9\x0b\xb3\xcd\xe0\x66\x2d\x1d\x28\xf4\x4f\x1d\x54\x74\xfc\x25\x4f\x77\xfb\x20\xbd\x81\x15\xf2\xf5\x85\x55\xbb\x88\x26\xf2\xbc\x0a\xfd\x07\xb9\x57\xbc\x5a\x72\x90\x68\xc3\x3d\x89\xf0\x6c\x7e\x20\xd9\x48\x37\x68\x97\x98\xfb\x46\x70\x53\xd9\x78\x03\x15\xd1\x6a\x6b\x8e\x84\x6a\x07\xa5\xf0\xf9\xba\x11\xe5\xa9\x8b\x9d\x0a\x82\x50\xce\xb0\x5a\xb6\x6b\xd4\x7d\xad\x53\x84\x74\x3c\x10\x84\x7c\x2d\xf4\x0a\x1d\x77\xa7\xca\xac\xbe\xe5\x1c\xa9\xf3\x5d\x98\x2b\x77\xef\x4e\x98\x46\x38\xcf\x09\x78\x8e\xa8\xc3\x78\xba\xc9\xb7\x07\x29\xfa\xfe\xa1\xc3\xf0\xca\xac\x28\xd2\x1d\xf0\x30\x3c\x3d\xbf\x67\xa3\xe0\x76\xfc\x1f\x50\xcf\x95\xe3\xdb\xfc\x63\x72\x70\x8d\x93\xc2\x3d\x34\xe3\x2e\xa9\x79\x2e\x29\x29\x08\x10\xc6\xa0\xb2\x8a\x96\x19\xc5\x51\xc6\x08\x9e\x43\x6d\x5f\xd9\xd7\x45\x61\x93\x34\x9b\xf3\xf2\x6c\xe1\xd4\xcf\xfe\xd6\x98\x37\xdb\xe4\x6b\xdc\xc6\x42\x20\x19\xfd\xf7\xf5\xcd\x68\x0c\x95\x55\x63\xb2\xe6\xc0\xd0\x63\xe0\xac\xbe\xec\xad\x95\xda\x2f\x93\x50\x02\x40\x6e\xb4\xc6\xdc\x73\x73\x68\x9a\x73\x98\xcf\xaf\xa9\x4d\x41\x51\xce\xe0\x4f\x6e\x14\x9a\xf5\x00\x91\x1b\xed\xa5\xae\x90\xbf\xb1\xfa\x06\x81\xb2\x11\xaf\x75\xd1\x9d\xcd\xab\x45\x68\xc3\x7e\x96\x7e\x9d\x50\xf1\xf3\xe7\xd6\x0c\x2d\x0c\xdf\x7e\xfb\x3f\x17\x98\x4e\x10\x6f\xf7\x77\x03\x81\x33\xbb\x0e\xd1\x2f\x2e\xcd\x53\x9d\xc3\x9a\x2e\xbc\x8a\xa2\xce\x2e\x81\xca\xc6\xec\x27\x5d\x0a\xeb\xd6\x42\x25\x61\x4c\x94\xe0\x6d\xf6\x52\x78\x91\xa4\xe9\x18\xfe\x8c\xe9\x57\x3d\x1b\xeb\x8a\x15\xe5\x62\xe0\x8c\xbb\xce\x27\x97\x30\x8a\x6b\x5f\x99\xb2\x14\xba\x18\x9d\x64\x0d\xbf\x27\x13\xb8\x31\x85\xd8\xb1\x87\x59\x84\x9c\xf2\x49\x18\x77\x87\xfd\x41\xd8\x36\x45\xa2\x38\x96\x25\x37\x44\xef\x9a\xe1\x42\x1d\xff\x89\x30\xa4\x9e\x08\x5c\x37\xcc\xf0\xbd\xe6\x3f\x56\x28\x50\x28\xd8\x1a\xab\xd8\x9f\x99\x9d\x97\xa3\x3c\xe6\xf8\x92\xb0\x5e\x93\x5c\x7e\x23\x9c\x03\x19\x32\x56\x33\x13\xce\x1a\xf0\x57\x24\xc2\xd6\x58\xbb\x03\xb1\xa0\xae\xa1\x88\xb3\x0a\xe9\x42\x84\xd8\xe2\x7e\xd4\xc2\x60\xf4\x3a\x3f\xa8\x31\x79\xec\xd7\x60\x7a\x03\x85\x74\x62\xa1\x30\x8e\x5d\x4a\x53\x60\x4b\xc7\xdc\xa3\x3f\xb9\x84\x22\x6b\xcf\x30\x1e\x57\x30\x85\x09\x0a\x29\xf5\x00\xb7\x13\x67\xde\x07\xcc\x0f\x7b\xd2\xa3\x57\x70\x09\xb8\x3f\x2b\x52\x77\x3b\xf4\x85\x2a\x44\x2e\x41\x52\x38\x47\xe9\xd7\x21\x0e\x32\x71\x72\x9e\x32\xf1\x52\x5a\x17\xb3\x34\x6c\xf1\xe9\x2d\x82\xa3\x18\x17\xce\x6b\x8f\x65\x5a\x9c\x2f\xd2\xee\x2b\xea\xcd\xc3\x5a\x6b\x51\xc7\xd6\xa2\xc3\xf1\x97\xb4\x9d\xa9\xb8\x3a\x21\xa5\x63\x41\x47\x9b\xf3\x48\x97\x87\x3c\x05\x38\x53\xa2\x5f\xd3\x69\x1c\x48\xd0\x68\xba\x1e\x58\x0f\x7a\xc2\x6c\x4b\x2f\xfb\x7b\xbe\x9a\xb6\x6e\xb6\x8a\x4c\xd5\xfd\x55\xfa\x09\xcf\xa9\xcf\xa7\xb7\x01\xa1\xc3\xee\x6f\x4d\xfe\x68\xf9\x1e\x00\x95\xe3\x93\x6c\xdd\x5e\xce\xab\x45\x72\x0a\x3f\x85\xbf\x43\x27\xaf\xfd\xd1\x02\xed\x23\x1f\xfd\x78\x20\x9e\xcf\xa8\xa2\x7f\xe7\x8d\x87\x5c\x42\x91\x05\xdf\xbb\xef\xcc\x90\x67\x47\x43\xe4\xba\xf6\x7b\x69\x92\x8e\x3f\x8e\xeb\x6a\x2a\x44\x7e\x56\x4b\x5f\x96\x88\x13\xdb\x30\x33\x0f\x1b\x8d\x23\xe1\xf6\x9d\xe2\xdf\x2f\x9b\x4b\xc5\x41\xf8\x3b\x16\x2c\xe2\x95\x05\x71\x0b\x1d\x97\xe1\x01\xad\xe4\x21\x66\xdb\x69\x62\xa5\xd3\x5a\xcc\x58\x1e\xa0\x27\xa3\x8b\x77\xef\xb8\xa6\x0c\x4b\x77\xa4\x79\x68\x32\xf2\xd1\x58\x99\xe3\xca\x67\x5e\xc2\x0c\xc2\xf0\xfd\x5f\x50\xd4\xe3\x33\xfb\x13\x57\x0f\xcd\xad\xc3\xa7\xf4\xf7\x6f\xab\xef\x11\xed\x3d\x0c\xbb\x72\x04\xd5\xc1\xe4\x19\x68\x03\xcb\xd8\x8a\xf3\x46\xc3\x9f\x17\xec\xd0\xc3\xb3\x49\xda\x3b\xb0\xff\xdf\x00\x00\x00\xff\xff\xbb\xa6\x2f\x90\x66\x28\x00\x00")

func _hardcodedDoerGoBytes() ([]byte, error) {
	return bindataRead(
		__hardcodedDoerGo,
		"../_hardcoded/doer.go",
	)
}

func _hardcodedDoerGo() (*asset, error) {
	bytes, err := _hardcodedDoerGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../_hardcoded/doer.go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc, 0xd5, 0x2c, 0x20, 0x46, 0x76, 0xdb, 0xea, 0xde, 0x4f, 0xd6, 0x45, 0xe4, 0x2b, 0x7, 0x75, 0x30, 0x43, 0x84, 0x99, 0xba, 0xea, 0x61, 0xad, 0xfe, 0x22, 0x44, 0x21, 0x79, 0x43, 0x1c, 0xcd}}
	return a, nil
}

var __hardcodedMiddlewareGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x5f\x6b\xe3\x46\x10\x7f\xb6\x3e\xc5\x74\x4b\xc8\xea\x70\xe4\x7b\x28\x14\x02\x7e\x28\xe9\xa5\x79\x39\x38\x2e\xa1\x2d\x94\xc2\x6d\x56\x23\x69\xb1\xbc\xab\xce\x8e\xac\x9a\xe0\xef\x5e\x66\x25\x39\x51\x48\xd3\xa7\x3e\x58\xde\x9d\x99\xdf\x6f\xfe\x6f\x67\xec\xce\xd4\x08\x11\xe9\x80\x94\x65\x6e\xdf\x05\x62\xd0\xd9\x4a\x55\x7b\x56\xd9\x4a\x79\xe4\x4d\xc3\xdc\xc9\x99\x7a\xcf\x6e\x8f\x9b\x12\x1f\xfb\x5a\x65\xd9\x4a\xd5\x8e\x9b\xfe\xb1\xb0\x61\xbf\xb9\x69\xf1\x80\xb4\xd9\x99\xe3\x01\xf1\xaa\x0e\x9b\xc3\x8f\x9b\x36\xd4\x35\x92\xca\xf2\x2c\xdb\x6c\xe0\x8b\xf1\xce\x7e\x76\x65\xd9\xe2\x60\x08\xa1\x0d\x75\x04\xe3\x8f\xd0\x89\x22\x16\x70\x1b\x08\x7c\x18\xd6\x30\xe0\x25\x21\xd8\xe0\xd9\xf9\x1e\x81\x1b\x0a\x83\xf3\x35\x70\x83\xa3\x31\xf4\x9d\x30\xca\x3d\xb2\xb1\x3b\x88\x01\xb8\x71\x11\xf6\xe6\x08\x96\x4c\x6c\x46\x5b\x0a\x16\x63\x2c\xb2\xaa\xf7\xf6\xb5\x7f\xdd\x80\x24\x56\xdc\x19\x5f\xb6\x48\xf9\xe2\x06\x4f\xd9\x8a\x90\x7b\xf2\x0b\xf1\x6d\xef\xad\x16\x32\x3d\x8c\xf2\xaf\x18\xbb\xe0\x23\xfe\x46\x8e\x91\xd6\x40\xf0\x61\x92\xff\xd5\x63\xe4\x5c\x78\x56\x25\x56\x48\x90\x60\xa3\x60\x95\x92\xf8\x44\x04\xd7\x5b\x20\xb4\xe1\x80\xa4\x73\x51\xb8\x0a\xce\xba\xed\x16\xbc\x6b\x47\xc0\x14\x8c\x1c\x4f\xf2\x39\x18\x02\xa4\xf4\x0b\x94\x89\x24\x0e\x8e\x6d\x03\x2f\x99\xe7\x73\xa1\xf9\xd8\xe1\xe4\xda\x9a\x28\x45\x23\xe7\xeb\xeb\xc4\x2c\x34\x5b\xa8\xf6\x5c\x7c\x12\xb2\x4a\xcf\xb0\xfc\x6c\x9e\xbc\xbc\xb4\x9e\x4d\x44\x54\x62\x65\xfa\x96\xff\x85\x4c\xf5\x7e\xe7\xc3\xe0\xa7\xbe\x5d\x7c\x7f\x80\x50\x81\xc4\x03\x17\x0f\x6a\x7d\x66\x7a\x3e\xe5\x63\x92\xf2\x1d\xe7\xa7\xb8\xa5\xb0\xbf\x09\x9e\xf1\x6f\xd6\x54\xcc\xa7\x3c\x1f\x7d\xfc\xac\x55\x82\xaa\x75\x8a\x60\xc2\x7c\x7e\x52\x48\xa4\xae\x25\xf6\x35\xa8\x34\x25\x4c\xc6\xa2\xba\x9e\xb2\xd7\x69\x8c\x8b\x7b\xd1\xe8\x3c\x3f\xe5\xe7\xc6\x2c\x2a\x70\x4a\x8d\x69\x8a\x7b\x59\x91\xbb\x87\x87\x2f\x7a\x58\x83\x68\x4e\x79\x76\x4a\x73\x1d\xd9\x70\x1f\x97\x93\x00\x03\x99\x2e\x82\x01\x9a\xc4\x30\x24\x79\x96\x32\x7f\x13\x11\x99\x7a\xcb\xd2\xa4\x37\x26\x2b\x5b\x8d\x18\x70\x9e\xc5\x6d\x1a\x68\x1d\xe1\xc3\x5b\x54\x39\xa4\xff\x3b\x34\x25\x92\xb6\xa1\x44\x81\xa5\xfe\xc7\x62\xe2\xd9\x82\xc8\x45\xb0\xc4\x16\xaf\xa1\x73\x96\xbf\x22\x45\x17\xfc\x57\xe3\x6b\x84\x12\xad\x2b\x31\xc2\xd0\x20\x37\x48\xc0\x01\x8c\xb5\xd8\x31\x18\x38\x8c\x86\xc5\x98\xea\x02\x96\x56\x60\xd2\x4f\x5d\xc8\xe1\x31\x84\x36\x79\xb8\x69\x1d\x7a\x9e\x00\x37\x0d\xda\xdd\x8b\xc7\xc2\xca\x3d\xa6\xad\xb6\xc9\xee\xd9\x4f\xaa\xc5\x7b\xe0\x57\x9b\xbe\x06\xf2\xf5\x22\xb0\xff\x71\xf7\xe7\x6c\x65\xd1\x8b\xb1\xac\xc5\x2f\xc8\x5a\xfd\x7e\x35\x86\x7c\x35\x05\xa2\x64\xcc\xfe\x6b\xe0\x7f\x2a\xcb\xf9\xa6\xc6\x32\x5c\x4d\x0e\xd4\x7a\x2e\x88\xf0\xb8\x0a\xbe\x23\x5f\xcf\xb5\x9e\x56\x7f\x58\x34\xf7\x87\x8f\x1f\xf3\x17\x52\xfd\xc7\x9f\x8f\x47\x46\x2d\xbb\x7b\xdf\x91\xf3\x5c\xe9\x6f\x4f\x6a\x8f\x31\x9a\x5a\xb6\x46\x2d\xeb\x0e\x97\x17\xf1\x12\x7c\xe0\xa9\xf3\x58\xae\xa1\x6b\x51\x9e\x8b\xbe\xab\xc9\x94\xa8\x4e\xdf\x9e\x83\xca\x93\xaf\xf3\x2b\x76\x7a\x67\xa7\xfe\x09\x00\x00\xff\xff\xa1\x73\xce\xd9\x95\x06\x00\x00")

func _hardcodedMiddlewareGoBytes() ([]byte, error) {
	return bindataRead(
		__hardcodedMiddlewareGo,
		"../_hardcoded/middleware.go",
	)
}

func _hardcodedMiddlewareGo() (*asset, error) {
	bytes, err := _hardcodedMiddlewareGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../_hardcoded/middleware.go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcb, 0x1, 0xb6, 0x9d, 0x16, 0xe7, 0x46, 0xa2, 0x49, 0x7d, 0xf6, 0xc, 0x20, 0x3a, 0xd4, 0x73, 0x64, 0xea, 0xcc, 0xe3, 0x1f, 0x62, 0xec, 0xc7, 0x5b, 0x68, 0x93, 0x2d, 0xaf, 0x96, 0x19, 0x61}}
	return a, nil
}

var __hardcodedTracingGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x58\x6d\x6f\xdb\xc6\x93\x7f\x2d\x7e\x8a\x29\x81\x34\x64\x40\x93\x56\x93\xa6\x85\x0a\xbf\x70\x6c\x27\x11\xce\x76\x04\xdb\x45\x71\x38\x1c\x8c\x35\x39\x24\x17\x21\x77\x79\xbb\x4b\x3d\x20\xf0\x77\x3f\xcc\xf0\x41\x92\x2d\xfb\xdf\x56\x2f\x08\x71\x77\x76\x76\x1e\x7f\x33\xc3\x46\xa4\xdf\x45\x81\x60\xd1\x2c\xd1\x38\x23\x52\xa9\x0a\xcf\x93\x75\xa3\x8d\x83\xc0\x9b\xf8\xa9\x56\x0e\xd7\xce\xf7\x26\x3e\xaa\x54\x67\x52\x15\xc9\x83\x54\xc2\x6c\xf6\x96\x4a\x5c\xd3\x7b\x5e\x33\xa5\x42\x97\x94\xce\x35\xf4\x5f\x5b\x7a\x5a\x67\x52\xad\x96\xf4\xd7\xc9\x1a\x7d\xcf\x9b\xf8\x85\x74\x65\xfb\x10\xa7\xba\x4e\xce\x2a\x5c\xa2\x49\xbe\x8b\xcd\x12\xf1\xa8\xd0\xc9\xf2\xb7\xa4\xd2\x45\x81\xc6\xdf\x27\xcc\xc4\x12\xd3\xa2\x4c\x0a\x7d\x64\x1b\x5c\x25\xf4\x60\x12\x1d\xeb\x06\x95\xc3\x0a\x6b\x74\x66\x13\x4b\x9d\x90\xe8\x46\x3e\x24\x52\x59\x67\xda\x1a\x95\x13\x4e\x6a\x95\xec\xb0\x2b\xb4\x91\x55\x25\x92\xba\x5d\x27\xda\x61\x55\xb7\xeb\x97\xb8\xd1\xf6\x6b\x7b\x09\xae\xc9\x68\x68\x6c\xa2\x5d\xd5\xf0\x83\x0c\x8a\xff\xea\xd0\xf6\x5f\x61\x9a\xf4\x55\x16\x8d\xd1\x8d\x28\x58\xb5\x57\xe9\x6c\xf6\x3d\x31\x68\x75\x6b\x58\x26\x9b\x7d\x67\xfe\xf0\xfa\x89\xff\xac\xc2\x48\xd5\x3d\x1d\x5a\x8a\x01\x8b\x35\x39\xfc\x35\xee\x1d\x45\xb2\x9c\xc6\xd3\xe3\xf8\xf8\xd5\x3b\x7a\x29\x42\xcf\x5b\x0a\x03\x19\xe6\xa2\xad\xdc\x99\xae\x2a\x4c\x9d\x36\x5f\xb5\x75\x60\x9d\x91\xaa\x80\x13\xf0\x2b\x9d\x8a\xaa\xd4\x24\xc6\x21\xea\x05\x85\x76\x2b\x95\x9b\x7e\x84\x13\xf8\xf0\x7e\xfa\x9b\xe7\x25\x09\x0c\x66\xd4\x06\x9c\x86\xd6\x62\xcc\xa7\x77\x96\x77\x0c\x1d\xdf\xe1\xda\x5d\x89\x66\xb1\xdd\x3d\xd9\xdf\x27\x81\xcf\xba\xdc\xf9\xf1\x08\x49\x02\xac\x42\x23\x0c\x2a\x07\x25\x8a\x0c\x8d\x97\xb7\x2a\x85\x5b\x74\x6d\xf3\xa5\xd2\x0f\xa2\xe2\x43\x0b\xa3\x97\x32\x43\x73\xaa\xb2\x8b\x3e\x38\x82\xd4\xad\xa1\x4f\xc4\xb8\x67\x1a\x42\x30\x78\x30\xbe\x6d\x84\x1a\x68\x23\x78\x37\xae\x33\x3f\x33\x30\x8c\x00\x8d\xd1\x26\x84\x1f\xde\x24\x49\x60\x1a\x83\x45\x07\x6d\x03\x43\x0c\xf2\xf2\x2f\xe3\x32\xf3\x60\xad\xf9\x38\xef\xbe\x8f\x41\x58\x2b\x0b\x05\x05\x4b\xfc\x8c\x88\xa9\xe6\x39\xac\xf0\xad\x41\x30\xad\x52\xe4\x14\xf6\x48\xb5\x89\xc0\x95\xa8\xc0\xb5\x46\x81\xce\x73\xb0\xa2\x6e\x2a\xa9\x8a\x18\xbe\xb9\x12\xcd\x4a\x5a\xec\xd6\xb0\x93\xf0\x0d\x68\x03\xab\x52\x38\x02\x07\xb8\xbb\x39\x3d\x9b\x5f\x7f\xb9\xbf\x3d\xbd\x5a\x5c\xd2\x9f\xc5\xcd\xb7\x4f\xa7\x9f\xe6\x97\xf3\xbb\xff\x06\xdb\x60\x2a\x73\x89\x36\xf6\x26\x03\xdb\x85\xd1\x0f\xe2\x41\x56\xd2\x6d\x60\x76\x02\xc7\xf1\xf1\xd4\x9b\x48\x7b\x49\xc2\xd0\x82\xb6\xf1\x17\x74\xa8\x96\x81\x7f\x3f\xbf\xbd\xbf\xfc\x76\x76\x7a\xe9\x87\x70\x72\x02\xbe\x33\x2d\x05\xbd\xcc\x61\xa0\xff\xe1\x4d\x0e\x32\x3e\x81\x69\x7c\xec\x4d\x1e\x01\x2b\x8b\x20\x73\x58\x3e\x61\xfd\x9a\xd8\x7e\xf8\x07\x2c\xe1\xa7\x13\xf0\xfd\x97\x2e\xf8\x6c\x74\x7d\xa1\x96\xec\x3c\xe2\xdc\x03\x69\xbc\x10\xc6\xe2\xe7\x4a\x0b\x17\x2c\x23\xf8\xf8\x21\xf4\x26\x24\x2e\x51\xfd\x74\x02\x4a\x76\x12\x4f\x0c\xb2\xb5\x95\xac\xa2\xee\x91\xd7\x2e\xbe\xa0\x38\xc8\x03\x3f\xd5\x6d\x95\x81\xd2\x0e\x1a\xe2\x06\x6f\xdf\xd8\xb7\x14\xfc\x39\xb1\xf5\x23\x58\x12\xd3\xc7\x17\x15\x7f\x59\x5a\x8f\x4e\xd1\x4d\x0b\x23\x95\xab\x54\xe0\xf7\x66\x9c\xf9\xd1\x60\xd1\xd0\xf3\x26\x22\xcb\x58\x29\x22\xbd\x6d\x88\x36\x0f\xfc\x37\x76\xf6\x26\xf3\xa3\x83\x59\xfe\x7c\x95\xb2\x39\xec\xc2\xee\x62\x89\x66\x03\xd3\x5f\xc1\x62\xaa\x55\x66\x29\x08\x2b\x0a\xd1\x0d\x29\x95\x6a\xa5\x30\x75\xf4\x77\x0f\x63\x20\x1d\x58\x81\x70\xcc\xc7\x95\x38\x5c\xc3\x91\x4b\x09\x0d\x3a\x87\x11\x57\x66\x0c\x1a\x44\xfa\x17\xc5\xf3\x10\xe5\x52\x51\x1e\x64\x6d\xca\x07\x5c\x29\x2d\x48\x0b\x02\xac\xcc\x30\x15\x26\x02\xa1\x32\x58\xed\x9c\x60\x16\x7d\x6e\xec\xd0\x0f\x2b\x03\x5f\xdd\xe0\x28\xed\xd1\x28\x6d\xec\x4d\xa8\x46\x9c\x55\x92\x10\x85\x62\x6e\xb7\x62\xc4\xd7\xb8\xea\xb6\x02\x6f\x32\xd9\xdf\xfa\x4b\xba\xf2\x06\x7b\x7b\x48\xad\x16\x68\xa4\xce\x82\xe9\xaf\xef\xa8\x34\xc7\xb7\x6c\xbd\x30\x3a\x78\xee\x42\x65\x8d\x96\xca\x05\xe4\xba\x17\x68\xe6\xca\x62\xda\x1a\x0c\x68\x3f\xf4\x26\x76\x0f\x9e\xfa\x38\x1e\x8f\x91\xa4\x04\x70\x11\x6c\xd5\x09\x9f\x44\xcf\x70\x1a\xce\x0c\x0a\x87\x99\x1f\x7a\x07\x82\xfd\xd5\x58\x67\xec\x83\x94\xce\x93\x4d\x07\xcc\x9b\xc1\x9b\xa5\xcf\x42\x85\x14\xb4\xde\xc4\x35\x24\x9d\xc2\xd5\x3e\x78\x06\xfb\x4a\x1c\x08\xfd\x88\x0e\xdd\xf4\xf5\x35\x08\x43\x72\x0f\x56\xf1\x2d\xba\x27\x9c\x5c\xf3\x54\xbd\x8e\x00\x06\x8a\x1d\x2d\xbd\x49\xd7\x05\xc5\x94\x58\x3d\xf2\x93\xb1\xc2\x78\xae\x72\x7d\x1e\xf8\xd6\x09\x43\x0a\x1d\x75\x40\xec\x47\xd0\x1f\xb8\x22\x8b\xf8\xe4\x25\xb4\xd6\x9f\x41\xf7\xa3\x77\x72\x9a\x3f\x28\x70\x64\x84\x43\x7f\x76\x50\x21\x6f\xf2\x18\x7a\x83\x55\xf7\xf5\x77\x0d\x9b\xd8\x7b\xf4\xba\x22\xf6\xdc\x5e\x83\x81\xe1\x85\x2a\x75\x08\x52\x18\x77\x3e\x7e\x88\x60\x68\x53\xe0\xdd\xf0\x2f\x1e\x2c\x1b\xbe\x58\xdf\x28\x06\x06\x61\x07\x8a\xeb\x67\x72\x79\x13\xce\x5b\xa4\x0a\xbf\x97\xe9\xf3\x73\x28\x50\xa1\xa1\x62\x1e\xc3\x5c\x81\x36\xc4\x34\xd7\x66\x94\x95\xb0\x63\xa5\xcd\x77\x08\x84\x83\x0a\x85\x75\xb0\x92\xae\xec\x52\xb7\x2b\x5c\x14\x46\x03\x86\xec\x72\x84\xba\xb5\x6e\x78\xc5\xae\x6c\xc2\xfc\xdc\x42\xab\x64\xae\x4d\x5d\x6d\x40\x38\x30\x42\x65\xba\x86\xdc\xe8\x9a\x39\xa0\x72\xd2\x20\xd9\x3e\x45\x42\x20\xea\x5a\x3e\x7e\x88\xbb\x2b\x3e\x6b\x03\xb8\xe6\x5b\xa3\x3d\x4d\xd6\x47\x46\x6c\xf6\x6f\xcf\x34\x5a\x06\xfa\x4c\xb3\xb4\xc4\x62\x34\x12\xe5\xec\x6d\x27\xfd\xb6\xa7\x38\xad\x56\x62\x63\xbb\xe5\x20\xe4\x4c\x4f\x12\x78\xfd\xcc\x82\xbb\x9b\x4f\xc2\x62\x16\xec\x3b\x69\x7e\x7e\x43\x30\xda\x6f\x3d\x77\x7d\x38\xde\x70\x57\xa2\x45\xa8\xc5\x5a\xd6\x6d\x6d\x41\x18\x24\xab\x17\xad\x30\x19\x88\x42\x50\x2f\x0f\x56\xd7\xe8\x4a\x72\x48\xa1\xe9\xb9\x32\x5a\x15\x8c\xae\x16\x15\x8d\x23\x20\xc0\x75\x98\x9d\x09\x27\xa0\x55\xb8\x6e\x30\x75\x98\x55\x9b\x67\x8a\x37\x42\x5d\xca\x5a\x3a\xbb\xdf\x4f\x75\x6b\x5c\x43\x4f\x1d\x8d\x11\xad\xc3\x33\xdd\x2a\xc7\x1b\x33\x98\x1e\x1f\x93\xc4\x93\x8b\x25\x2a\xb7\xbb\x41\xbf\x61\xf3\x52\xaa\xef\x4f\xf7\x86\xcd\xc7\x5e\xe3\x4f\xc2\xa5\x25\x1a\x82\xfe\x5a\x1b\x04\xcc\x73\x99\x12\x08\x46\x60\x57\xd2\xa5\x25\xe9\x2f\x1d\x88\x9c\xb2\x89\xba\x6b\xae\x1c\x4f\xb4\xd8\xa8\x74\x27\xe9\x7a\xd6\x7b\x24\xfd\x3d\x4f\x69\xbe\x8a\x25\x5b\x38\x97\x45\x6b\x10\x74\xeb\xa0\xd4\x2b\x98\xbf\xad\x7b\xe3\x92\xf5\xc7\xb0\xa5\x40\x1f\x93\x33\x97\xc6\xba\x67\x91\x34\x42\xe0\x40\xd7\x55\x81\x47\x6e\xb2\xff\x46\xb3\xfb\x59\x9b\x3b\xb4\x8e\xed\x81\x42\x71\xc1\x7e\xe0\x74\xcd\xa8\xc2\xb6\x4a\xba\xc1\x0a\x11\xb1\x24\xb7\xd7\x92\xd0\xdd\x72\x16\x58\xba\x02\xc4\x83\x5e\x22\x27\xaf\x6e\x1d\x15\x60\x4e\x9f\x9d\xb3\x31\xcc\x1d\x74\x70\x61\x41\x28\x90\xea\xa8\xc6\x5a\x9b\x0d\xf1\x1c\xc1\x2b\xef\x73\x4c\x72\x21\x1e\xec\x90\x31\x1a\xda\xf8\x6f\x37\xf0\xbd\x4e\x41\x08\xc1\xbb\x71\x48\x8a\xe7\xea\x8a\x6f\xfc\x67\xad\xfb\x28\xdb\xec\x04\xb6\xbc\xae\x71\xf5\x94\x5d\x10\xbe\x5c\xcd\x70\xbc\x72\x1a\x1f\xff\xb3\xca\xd5\x43\x2c\x1e\xaa\x05\x49\x02\x57\xed\xfa\x96\x3f\x26\x5c\xc9\x2c\xab\x70\x45\x09\x3c\x98\xb9\xde\x2e\xb9\x52\x38\xb0\x25\x37\xa0\x0f\x08\xc2\x39\x91\x96\x98\x91\xb3\x05\xec\x4c\xe5\xfd\x97\x89\x98\x58\xcf\x5d\x87\x63\x6e\xc5\x20\xa6\x0a\x3b\x03\xae\x7f\xb6\xf3\x47\xd7\x60\x89\x2c\xeb\xde\xbb\xa1\x11\xa4\xca\x35\xb1\xa5\xd8\x30\xf8\x7f\x2d\x5a\x77\xd4\x4f\x0b\xe9\x50\x2c\x89\xfb\x8d\x2c\x4a\x07\x4a\xaf\x60\x85\xa0\x55\xb5\x01\xdb\x36\xfc\x19\x84\x88\xc8\xfd\x04\xd7\x52\x31\x23\x82\x6c\xe1\x3a\x2d\xce\x85\x13\x99\xe6\x96\x02\x53\x37\x44\xc5\x01\x3b\x04\xa4\x8b\x4c\xf1\x5a\xd4\xd8\x8f\xab\x21\x10\x71\x50\x3a\xd7\xc4\x5f\x85\xca\xa8\x80\xc0\xee\x1b\x79\x5c\xbb\x8a\x2c\xc1\x4d\x13\x7f\xa2\x88\x0f\xf3\x8c\xc6\x7d\xca\xc3\xed\x68\x6a\x83\xed\x10\x1b\x3e\xed\x3d\x4e\x33\x06\xcc\xad\xa9\x77\xbc\xe4\x6f\xdd\xdd\x89\x09\xaf\x0b\x3a\x10\x77\x02\xef\x69\xf5\x99\xce\x33\x13\xb3\xea\xce\xdd\xa0\x6d\xb4\xb2\xf8\x97\x91\x1c\x46\x06\xde\xf5\xeb\xec\xa3\xb0\x9b\x5f\x92\x64\x50\x0a\x4a\x61\x01\xd7\xe4\x54\xca\x40\x4e\xf5\x46\xa8\x98\x5d\xd6\xb4\x8e\x30\x52\xaa\xde\xd1\xa9\x5b\x3f\x75\xb2\x47\xec\xec\x98\x35\x0c\xf2\xbb\x5d\x95\x19\x26\x6b\xce\x80\xc9\xbe\x95\xbe\x68\xd7\x57\xf9\x46\xa8\x99\xcf\x04\xb6\xc1\x55\x7c\xde\xd6\x4d\x60\xf9\x5d\xe6\x60\x53\x1e\xd1\x98\xf7\xc8\xed\x0f\xb0\x69\xfc\x55\xd8\xbe\x10\x06\xbd\x66\xdc\x16\xcf\xcf\xa3\x4e\x9c\xf9\x39\x9f\x4c\xf9\x28\x11\xc5\xb7\x1c\x20\x41\x18\xd1\xea\x78\x76\x5c\x66\x16\x49\xc2\x05\x8e\xa2\x8f\x86\x42\xa4\x54\x28\x71\xdd\x07\x97\xe5\x0f\x19\xdc\x32\x50\xec\x46\x60\x35\xd9\x02\x5c\xa9\x69\xd0\xd6\x5d\xf8\xa6\xda\x18\xac\xba\xae\xb8\xd2\x85\xe5\x24\x62\x99\x2c\x37\x3b\x96\x6f\x92\x39\x54\xa8\x82\x5e\x56\x1e\x92\xdf\xff\x02\x3f\xff\xcc\xab\x9d\x26\xbc\x38\xfd\x08\x3f\x80\x9c\xb6\x37\x62\x75\xe9\x32\xa3\xdd\x87\x8d\x43\x08\xde\xff\x72\x94\x96\xc2\x90\xb0\x61\x04\xbf\xf7\xab\xd3\x8f\xdb\xd5\xbe\x45\xe2\xa2\xde\x10\x3e\x67\x9d\x20\x93\x5e\x84\x4f\x36\x82\x7b\xb2\x59\x89\xeb\xf8\x1c\x53\x9d\x61\x6f\x9a\x41\xc6\x8e\xfc\x40\xff\xbc\xeb\xe9\xf8\x34\xcb\x86\x37\x9f\x4f\xde\xcb\xcc\x8f\xba\xb3\x93\xfd\xb9\x94\x86\xd2\xee\xd3\x67\xfc\x49\x16\x17\x2a\x93\x42\xc5\x7f\xb2\x7d\x83\x51\xaa\xff\xf9\x7d\xf6\xbf\x61\xd8\x5f\xde\x19\xe6\x65\x51\x7b\xc3\xfd\x73\x49\xe9\xe0\xbf\x13\x74\x10\x69\x90\x91\xc6\xfb\xee\x51\xc6\x0c\x58\x5f\xef\xee\x16\x81\x59\x45\xc0\xad\xec\x63\xc8\x73\x51\x87\xed\x3b\x65\x62\x5b\x3a\xb7\xfd\x40\x86\x36\x35\xf2\x81\x9b\x06\xca\x15\xd1\x34\x95\xec\xc6\x67\x46\xd8\x3f\xa9\x84\x73\x3f\x8d\x8e\xc3\xed\xc0\x47\xa6\x61\x9a\xd8\x96\xa3\x03\x33\x00\x25\xd0\x52\x18\xba\x60\x07\x4b\x9f\x00\xdb\xd9\x30\xea\x0d\xa7\xfa\xa1\x71\xf7\xab\xcf\xe2\xdb\xf9\xfd\xfc\xdc\x0f\x77\x3e\xc1\x0c\x3c\xf7\x3f\x0f\x9d\x2e\x16\xf7\xd7\xa7\x57\x17\xc4\x63\xfb\xa9\x67\x87\xe2\xef\x72\xda\x63\xe4\x4d\x4c\x1f\x19\xa3\x8a\x57\x68\x0a\x0c\x18\x4a\xfb\x95\xf3\xae\xab\xe7\x61\x7a\xbb\x7a\x8d\x2b\x02\xf9\xb1\x35\xb5\x01\x43\x52\xf7\x49\x35\xbe\x4d\x4b\xac\xc5\x9f\x37\x97\xd1\xde\xea\xb6\x54\xfc\x17\x6e\x06\x24\xe9\xc5\x64\xf6\xfd\xc0\xde\xc3\xb8\xf1\x1e\xbd\xff\x0f\x00\x00\xff\xff\x0f\x03\x1b\xbe\x21\x18\x00\x00")

func _hardcodedTracingGoBytes() ([]byte, error) {
	return bindataRead(
		__hardcodedTracingGo,
		"../_hardcoded/tracing.go",
	)
}

func _hardcodedTracingGo() (*asset, error) {
	bytes, err := _hardcodedTracingGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../_hardcoded/tracing.go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x63, 0xda, 0xa, 0x4c, 0xfd, 0x33, 0xfb, 0xb4, 0xde, 0x32, 0x4, 0xf0, 0x61, 0xa9, 0xd7, 0x25, 0x87, 0xd4, 0x7a, 0x99, 0xf9, 0x3a, 0xa6, 0xdb, 0x5d, 0xb1, 0xb4, 0x3, 0x80, 0x6d, 0x9d, 0x3d}}
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
	"../_hardcoded/doer.go":       _hardcodedDoerGo,
	"../_hardcoded/middleware.go": _hardcodedMiddlewareGo,
	"../_hardcoded/tracing.go":    _hardcodedTracingGo,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

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
	"..": {nil, map[string]*bintree{
		"_hardcoded": {nil, map[string]*bintree{
			"doer.go":       {_hardcodedDoerGo, map[string]*bintree{}},
			"middleware.go": {_hardcodedMiddlewareGo, map[string]*bintree{}},
			"tracing.go":    {_hardcodedTracingGo, map[string]*bintree{}},
		}},
	}},
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
