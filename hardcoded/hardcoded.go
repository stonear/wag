// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../_hardcoded/doer.go (10.35kB)
// ../_hardcoded/middleware.go (1.685kB)
// ../_hardcoded/tracing.go (6.716kB)

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

var __hardcodedDoerGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5a\xff\x8f\xdb\x36\xb2\xff\xd9\xfe\x2b\x26\x06\xee\x22\x25\x5e\xd9\x9b\xbb\x6d\x0f\xee\xdb\x03\xda\xcd\xf6\xb5\x40\x9a\x04\xf5\xf6\x35\x40\x10\x5c\x69\x69\x6c\xb3\xa1\x48\x97\xa4\xd6\x6b\xec\xed\xff\xfe\x30\x43\x4a\x96\x6c\x79\x93\xbb\x16\x0f\xef\xf2\xc3\xda\x96\x66\x3e\x1c\x0e\xe7\x3b\xb3\x11\xf9\x47\xb1\x42\xc8\x95\x44\xed\x87\x43\x59\x6e\x8c\xf5\x90\x0c\x07\xa3\xc5\xce\xa3\x1b\x0d\x07\xa3\xdc\x68\x8f\x77\x9e\xbe\xa2\xce\x4d\x21\xf5\x6a\xf2\xab\x33\x9a\x1f\x58\x6b\x2c\x53\x49\x33\x91\xa6\xf2\x52\xd1\x0f\x65\x56\xf4\x51\x0a\xbf\x9e\x58\xa1\x0b\xfa\xa1\xd1\xc7\x8f\xc9\xda\xfb\x0d\x7d\x77\x3b\x9d\xd3\xa7\x97\x25\x8e\x86\xc3\xc1\x36\x57\x30\x5a\x49\xbf\xae\x16\x59\x6e\xca\xc9\x95\xc2\x5b\xb4\x93\xad\x58\x4d\x94\x59\xad\x68\xe1\xad\x58\x05\x59\xe9\x01\x5a\xe2\x6a\x73\x88\x25\xde\x4d\xd6\x3b\xe7\xad\xbc\x3b\x5b\x99\xfa\xeb\xa8\x4b\x55\x18\x6d\x6e\x85\x5e\xcb\x02\x27\x78\x8b\xda\x3b\x53\xd9\x1c\x47\xc3\x74\x38\x9c\x4c\xa0\x30\x68\x41\x3a\x10\x1a\xa4\xf6\x68\x97\x22\x47\x58\x1a\x0b\xa3\xc2\x48\xbd\x1a\x01\xc9\x0f\x16\x7f\xab\xd0\x79\x07\x1b\xe3\x9c\x5c\xa8\x1d\x6c\xa5\x5f\xc3\xd6\x8a\xcd\x46\xea\xd5\xd0\xef\x36\x18\xa1\x1a\x90\xfb\xe1\xe0\xa5\x49\x72\x78\x46\x08\xd9\x15\x6f\x64\x0c\x36\xfe\xfe\x31\x20\xa6\x90\xd4\xbf\xdd\xc6\x68\x87\x63\x60\x35\xa7\xc3\x87\x61\x40\x35\x9b\xd7\xa2\xc4\x2b\x7f\x07\xce\xdb\x2a\xf7\xf7\x0f\x2c\xf7\x42\x38\x8c\x18\xdf\x09\x5d\x28\xb4\xb0\x41\xbb\x34\xb6\x74\xe0\xd7\xc8\xef\x3b\xb2\x07\x34\x7a\xfc\x92\xe4\xdc\x83\x2d\x2b\x9d\x43\x52\x34\xaf\x52\xf8\x1d\x62\xd3\xae\x2d\xfa\xca\x6a\xc8\xb3\x97\x26\x09\x1b\x99\x4c\xc0\xa2\xb7\xbb\x5a\x52\xfa\x21\xd1\xc1\xc5\xf4\x5d\x57\xbf\x41\x48\xa6\x6d\x49\x49\xa0\x05\xec\xff\x91\xa2\x79\x19\xbb\x7b\x6b\x94\xcc\x77\xf0\xe3\xfe\x7b\x5c\xaf\xf5\x04\x0a\x5c\x4a\x8d\x0e\x44\x40\x86\x0d\x3f\xce\xc2\x62\x6d\xc2\xce\xe1\x4d\x26\xf0\x8d\xc8\x3f\x9a\xe5\xd2\x41\xd8\x52\xd0\xac\xae\xca\x05\x5a\x10\xba\x00\x2f\x4b\xa9\x57\x60\x96\x11\x58\x78\x8f\xe5\xc6\xbb\x6c\x38\xa8\x59\x93\x14\xde\x7f\x20\x9b\xcf\x5e\x56\x56\x78\x69\x34\x23\xf3\xaa\x60\x31\x47\x79\x8b\x01\xb7\xad\x88\x31\x08\x07\x5b\x54\x8a\x3e\xe9\xa5\x45\x57\x29\x0f\x66\xc9\xdc\xb5\x5f\xc5\xe3\x79\xea\xe0\x97\x97\xe6\x17\x28\xd1\xaf\x4d\x91\x0d\x07\x8c\x9e\x74\x4e\x6c\x0c\x27\x0e\x6c\x61\x8c\x8a\x3a\x9b\x4b\xbd\x52\xf8\x98\xe6\xfc\x5a\xf8\xe6\xf4\x44\x2d\x2d\x18\x9d\x63\xd0\xe6\x31\x44\xc7\x6a\x7b\x14\x2a\x3c\xec\x4c\x05\x6e\x6d\x2a\x55\x34\xcb\x60\x83\x7d\xee\x30\x37\xba\x00\xb1\xf4\xe4\x5e\x1e\x96\x42\x2a\x97\x45\xb3\x3d\x5a\x2f\x85\xd3\x9a\x6f\x19\xe7\xc1\x9b\xfb\x73\x78\x06\xfc\x64\xce\xab\x3d\xb4\xad\x08\xb6\x52\xa9\x28\x99\x36\xfa\xec\xed\x9b\xf9\xcd\x38\x7c\xfb\xfa\xe6\xea\xbb\x7d\x70\xe0\xcd\x5c\xbc\x7b\x97\x11\xeb\xcd\x9b\x97\x6f\x66\xf0\xbd\x27\x6b\x75\xa0\x8d\x87\xbc\xb2\x96\x82\xd9\xae\xb6\x16\xbd\x0b\x87\x50\xab\x03\x0b\x58\xec\x4e\x1c\xee\x23\x1b\x0e\xa7\x6d\xf1\x37\x38\x38\x71\x8b\x6e\xd3\x77\xec\xed\xa3\x27\x9d\xc8\x25\x3f\x7d\x72\x09\x5a\x2a\xf8\xe7\x3f\x69\x47\xd9\x0f\x6c\x4c\x70\x79\x09\x23\xda\xf0\xa8\xef\x39\x6d\x9f\x5e\x0c\x07\x03\x5a\x2b\x9b\x7b\xe1\x2b\x77\x65\x0a\x84\xff\x82\x8b\xe9\x94\xc0\x6b\x8d\x2f\x85\x72\x38\x1c\x3c\x34\x47\xe0\x6d\x85\x51\xcd\xd7\x77\x1b\xa3\x51\x7b\x29\x54\xaf\xf5\x69\xc0\x3d\x45\xc7\x89\x83\xd5\x9d\xe0\x7f\xdc\xf4\x96\xf2\x96\xe2\x61\x7c\xca\xf1\xbc\xb5\x8a\xa2\x68\x90\x5b\x14\x8e\x3c\x7c\x2b\xa4\x67\xfb\x70\x1c\x7b\xd1\x6f\x11\x75\x73\xf0\x33\x38\x9f\x4e\xc7\xf0\x82\xfe\xfc\x95\xfe\xfc\x8d\xfe\x50\x80\x38\xff\x62\x3a\x85\x52\x2a\x25\x83\x15\x3b\x78\x3e\x39\x83\x6a\x03\xde\xc0\xc5\x9f\xe0\x57\xe9\x3d\xda\xfa\x68\xfb\x77\xf1\x19\x06\x0d\xb3\x4b\x28\xc5\x47\x4c\x0e\x5e\x8f\xe1\x22\x1d\x0e\x34\xde\x31\xc5\xf9\x74\x5a\x1b\xf9\x0f\x7b\x89\x86\x03\xab\x0b\x7a\x4d\xa9\x3b\x7b\x8d\xdb\xa4\xfe\x32\xe7\x3c\x99\x30\xc3\x6b\xb3\x4d\xd2\xec\x27\x2d\xef\x5e\x0b\x6d\x92\x34\x4d\x87\x03\x24\xae\x69\x36\xbd\x80\xc9\x84\xb7\x75\x41\x19\x28\x47\xed\xe3\xbe\x86\x03\xca\xa3\x32\x82\xaf\x38\xae\xd7\x16\xf1\x5e\x7e\x80\x4b\x60\xd1\x9e\x43\x47\xe8\x24\x49\xac\x2e\xb2\x6f\x95\x11\xfe\x8b\xbf\x26\xe9\xb3\x17\xe9\xd9\x79\xfa\x0c\x9f\x2d\xe3\x13\x62\xa2\xf5\xc3\xc6\x9e\x5d\xc2\x8b\xb6\x55\x59\xf4\xff\x6f\x7d\xf7\xd4\x01\xff\xe7\x3b\xf0\x6b\xd3\x9f\x35\x82\x93\x52\x82\xd4\x54\xdd\x05\xe5\x91\x43\x35\xb9\x23\x66\xe1\x2e\xc0\xe3\x8e\x4b\xe1\xa0\xdc\xf8\x1d\x38\x25\x73\xac\x95\xdb\x41\xf8\x37\xb3\x40\x37\xec\x0b\xb5\x15\xbb\x56\xbc\xa0\xfd\x9f\x58\xed\x5f\x4d\xb8\x2d\x19\x82\x5a\xeb\x6a\x8f\x15\x74\x15\x8a\xf0\xae\x1a\x7e\x96\x7e\xdd\xd6\x51\xa3\x0d\xd0\xb8\x85\x58\xb7\x07\x33\x36\xb7\x68\xad\x2c\x62\x5d\x11\x4a\x68\x30\x8b\x5f\x31\xf7\x4f\x5d\x53\x8c\x35\x65\x10\xef\xe8\x00\x3e\xc9\xfd\x5d\x0d\x9a\x45\x79\xc6\x70\xa2\xe4\x4a\x0f\x29\xdb\x55\x60\x7c\x43\xf8\xff\x23\x54\x85\x84\x3c\xee\xec\xf3\xfe\xa1\x83\xcc\x15\x63\x5d\x94\x3e\x6b\x6a\xc1\x3f\xa2\x2c\xad\x97\x18\x83\xf9\xc8\x61\xa9\x16\x38\x49\xb3\x20\x5c\x57\xae\x34\x4b\xda\xdb\x64\x47\x7b\x62\x3e\xd6\x5e\xd1\xe8\xe2\x12\x8a\xac\xf5\x9b\x9d\xa4\xc9\x2c\xb4\xce\xfe\x5d\xb6\x37\xcc\xe1\xe0\x56\xd8\x3e\x1f\x0f\x2f\x1a\x3f\x1f\x72\xcd\x37\x17\xb7\xd8\x29\x8c\x16\xa6\xa0\x1c\x05\x39\x55\xfa\x5b\x84\x35\x13\x98\xb0\x58\x06\x6f\xfc\x1a\xed\x56\x86\x77\x1c\x05\x99\x40\x28\x8b\xa2\x20\xf3\x11\x05\xe3\x72\xb3\x50\x2d\x97\x68\xc1\xe8\x26\xba\x15\x9d\xa5\x98\x9d\xea\xae\x0c\xe6\x88\xcc\x46\x02\xcf\x26\x13\xe7\x69\x3b\xb7\x68\x97\xca\x6c\xb9\xdf\x62\x0e\x69\xb4\x9b\xbc\xf8\xcb\xf4\xcb\xe9\xdf\xbe\xfc\x62\x42\x6b\x49\xbd\x3a\x23\x89\xcf\xcc\xf2\x8c\x78\xcf\x22\xf6\x19\x25\x5e\x53\xf9\xb3\xd2\x14\x72\x49\xb1\xa1\x79\xe3\xbc\xf0\x51\x17\x8b\x6a\x09\xef\x3f\x50\x97\xca\x67\x60\xb3\x6f\x68\xf3\x31\xde\xd1\x69\x74\x15\x36\x18\x2c\xaa\x65\x88\x94\x97\x10\x3a\xd5\xec\x47\x14\xc5\xd7\x4a\x25\x81\x97\xf2\x47\x37\x6a\x12\x4a\x6d\xb4\x5a\x2a\xe6\x1e\x0e\xe8\x24\x1f\x86\x21\x93\xd5\x75\x2f\xa5\xbd\xaf\x38\xfa\x7d\x55\x3f\x7b\xfe\x9c\xf9\x7b\x45\x1b\xd8\xc2\x12\x4f\x94\xe3\xb5\xd9\x5c\x29\xe3\xd0\x26\xdc\x74\x53\xa6\xfd\x86\xd5\x9f\x2c\xaa\x25\xe7\xb5\x41\xc4\xb8\x04\x5b\x44\x11\x38\x3a\xd7\x1b\x2a\xb2\x82\x5a\xab\x7c\x0c\x36\x6e\xa3\x96\xec\xf2\x12\x14\xea\xa4\x36\xbd\x94\x42\xfe\x93\xb6\xf1\xc5\x4c\x13\x32\x0b\xe3\xa5\x41\xc6\x85\x45\xf1\x31\xae\x35\x99\x00\x4b\x58\x37\x1e\x6c\x93\xd1\xdc\x96\xfb\x07\x32\x64\x46\x2d\x55\x2d\x84\xdb\x1c\x28\xd3\x6d\x78\x27\x19\xc3\x91\xc1\x33\x7e\xa8\xb3\x15\xe2\xa6\x91\xf4\x7d\xdc\xc1\x87\xb4\x9b\xcb\xa3\x90\x31\x2c\xe7\xd2\xe6\x95\xf4\xdf\x90\xac\x68\xb9\x41\x94\xe5\x46\x61\x49\x5d\x7d\x88\x75\x81\x02\x16\x81\x04\x36\xd4\x91\x59\xcd\x49\x7d\x8e\x1e\x0a\x5c\x54\x2b\x72\x12\x3a\x3d\xfa\x34\x1b\xb4\xc2\x63\x87\x59\x6a\x10\x50\x52\x0e\xdc\xae\xd1\x22\xf5\x1c\xec\x00\x46\xab\x1d\x28\xb3\x6a\x89\x02\x6c\xa3\xa1\xd6\x63\x1a\xd2\xc8\x42\x99\xfc\x23\x57\x08\x4d\x75\xb1\xb4\xa6\x84\x95\xa1\xcc\xe7\xd7\xd6\x54\xab\x75\xcc\x7b\x3d\x5b\x7a\xac\xe7\x0d\xf2\x87\x7f\xdc\xb8\x0d\x22\xc0\x6b\x51\x22\x71\x4a\xbd\x1a\x0e\xc2\xb8\x24\x50\x6d\x73\x95\xfd\x2c\x56\x21\x60\xbe\xe2\x17\xa4\x4d\xf2\x97\xc8\x3a\x9f\x5f\xbf\xd1\x39\x82\xdb\xe9\x3c\xa3\x6f\xac\xeb\xef\xc2\x30\x65\x3e\xbf\xbe\xbe\xa5\xf4\x21\x1d\x60\x49\x65\x1d\xd7\x39\xfb\xa9\x0b\xdc\x4a\x01\x0e\xed\x2d\xda\x33\x47\x84\x61\xc6\x92\x71\xf5\x84\x2e\xb7\x72\x11\x8a\x66\xd2\x30\x6b\x8b\xaa\x01\x51\x2f\x1e\xd5\x70\xb8\xda\x5e\x07\x37\xf4\xfa\x91\x7f\x61\xcf\xf0\xcb\xaf\xce\xe8\xd9\x88\xc0\x46\xbf\x0c\x07\xac\x8e\xcf\xe6\xd2\xa2\x64\xae\x98\x45\xae\x4c\xa5\x7d\x1f\x97\x0c\x8f\x23\x97\x6d\x51\x13\xf7\x35\x85\x9e\x93\xbc\x87\xdc\xd8\x50\x37\xbc\x6f\x43\xf5\x2c\x56\xc7\xa2\xf7\xf0\xee\xa9\x09\xe0\x7b\x77\xd5\xb1\xa4\x37\x1b\xd4\x6d\x00\x2e\x3a\x1a\x00\xd9\x43\xcd\x0a\x30\x4a\x49\xbd\x62\xb1\xbe\x15\x52\x55\x16\x1f\x53\xc0\x31\xf5\x31\x88\x52\xe4\xe3\x6d\xb0\xc7\x40\x3a\xd4\xa7\xc0\xe6\x55\x9e\xa3\x73\x9f\x09\x16\xa9\x0f\xc1\xe6\x6b\x63\x7d\xd4\x02\x16\x9f\xda\x5e\x97\xfa\x08\xab\x25\xd0\xa7\x55\x75\x42\xa0\x9b\x35\xa5\xc9\xb7\xc6\xa8\x1f\x91\x2a\x35\x12\xea\x34\xc8\x31\xf5\x11\x9e\x2c\xd1\x54\xfe\x33\x85\x8a\xd4\x04\x72\x15\x5a\x9d\x2b\xa3\x63\xd3\x73\x7d\x87\x79\x45\x29\x3d\x18\x77\x17\x24\x7f\x9c\x9a\x00\x5f\x09\x8f\x3a\xdf\xdd\x18\x2f\xd4\x0f\x28\xf4\xe3\xc6\xad\x5a\xd4\xff\x28\x51\x90\x61\xd6\x05\xa1\x32\x2b\x8e\x0f\x89\xea\x8b\x6b\x63\x38\x0a\x24\x9c\xdd\x54\xf6\xca\xac\x12\x62\xf8\x5e\x2f\xcd\x18\x46\xa3\x31\x94\x62\xf3\x3e\x04\x81\x0f\xcd\xd0\xef\xfe\x81\xd2\x56\xd7\xaf\x67\x7d\x9e\x8c\x59\x3b\x52\x8c\x89\xa9\xe5\xce\xbd\x2c\xc4\xb4\x0f\x10\x7b\x96\x96\x17\x1f\xf3\x45\x96\x3d\x0d\xf3\xf5\x3a\xef\xac\xcb\xd7\x17\x0e\x98\xb9\xcf\x67\x67\x87\x9b\x3b\xa6\xe9\xe1\xed\xba\xea\xac\x97\xb7\x43\x73\x12\xa3\x76\x88\xc7\x30\x22\xcd\x11\xc6\x81\x63\xce\xfa\xf6\xd0\xa5\x39\x86\xe8\x2c\xdf\xaf\x86\x53\xcb\xf7\xb8\xe1\xec\x90\xf7\x98\xe6\x18\x26\x7a\xdf\x63\x22\x44\x1a\xe6\xfd\x94\xd3\xcd\x00\xb3\x4f\xb8\x31\xe3\xa8\x03\xcf\x3c\xb6\x41\xcc\x0e\xbd\x77\x3c\x1c\x3c\x74\x5b\xb4\xe3\x1a\x26\x05\xa9\xa5\x4f\xd2\x38\x46\x7f\x8b\x56\x9a\x42\xe6\x3c\x55\x53\x66\x15\xe6\xec\x5a\xa8\x6e\x19\x45\xf5\x98\x70\x4e\x3a\x2e\xc1\x1c\x7a\xcf\xd5\x4c\xab\xda\xf2\x6b\x8b\x6e\x6d\x54\xe1\xb8\xe2\xaa\x74\x81\xd6\x79\xa1\x0b\x1e\x29\x6c\x36\x4a\xe6\xa1\xc5\x5f\xe0\x5a\xdc\x4a\x63\x33\x66\xff\x49\x2f\x8d\xf5\x95\x16\x1e\xd5\x6e\xdc\x2e\x5e\x0a\x83\x4e\x3f\xf5\xb1\x35\x82\x95\x45\xe1\x61\x2b\x76\x24\x0a\xde\x6d\x42\x11\x2c\x1d\xa3\x14\xc2\x8b\x31\x38\x43\xfd\x94\x45\x67\xac\xe7\x42\xb2\xb2\x9a\xef\x00\x34\x48\xef\xc0\x55\x1b\xbe\x54\xa3\x7e\x61\x61\x8d\x28\x72\xe1\x8e\xb6\x51\x52\xb1\x9b\x3b\x2e\x9d\x78\xea\x7f\x5c\x3f\x41\x32\x9f\x5f\xa7\x41\xfc\x39\x86\xcb\x01\x37\x9b\x4c\x1e\xb9\xfe\x92\xce\x55\xe8\x26\x17\x5f\x64\x4d\x55\x18\x4b\x3b\xea\x16\xe8\xb4\xc2\x81\x0c\x22\xcf\xdc\x5b\x14\x65\x7d\x25\x33\xbb\xac\xf5\xc2\xb3\xbf\xf6\x3b\x2e\xdc\xfb\x98\xb2\xb9\x17\xd6\xf3\x6b\x25\x9d\x47\xcd\x31\xd8\x32\x98\x46\x9f\xbd\xe2\x87\xc9\xc8\xe7\x9b\xd1\x18\x46\xe7\x2f\xbe\xcc\xa6\xd9\x34\x3b\x9f\x4d\x47\x27\xfa\xaf\x8d\xd0\x32\x4f\xa8\x2d\xe1\x5e\x61\x38\x18\xac\x0c\xef\x3d\x9b\x93\x8a\x92\xfd\x32\x7d\xf2\xa4\x81\xbe\xb5\x55\x52\xdf\xcf\xc8\x36\x87\xb7\x68\x77\x70\x3e\x85\x7a\x12\xbb\xc0\x5c\x54\xb1\xcb\x21\x34\xb4\xf2\xae\x39\x1a\x61\x91\x8e\x54\xc0\xf9\xf4\xac\x1e\x94\x32\x5a\x74\x5b\xd8\x4a\x5d\x98\x6d\x06\x37\x6b\xe9\x40\xa1\x7f\xea\xa0\xa2\xe3\x2f\x79\xba\xdb\x07\xe9\x0d\xac\x90\xaf\x2f\xac\xda\x45\x34\x91\xe7\x55\xe8\x3f\xc8\xbd\xe2\xd5\x92\x83\x44\x1b\xee\x49\x84\x67\xf3\x03\xc9\x46\xba\x41\xbb\xc4\xdc\x37\x82\x9b\xca\xc6\x1b\xa8\x88\x56\x5b\x73\x24\x54\x3b\x28\x85\xcf\xd7\x8d\x28\x4f\x5d\xec\x54\x10\x84\x72\x86\xd5\xb2\x5d\xa3\xee\x6b\x9d\x22\xa4\xe3\x81\x20\xe4\x6b\xa1\x57\xe8\xb8\x3b\x55\x66\xf5\x2d\xe7\x48\x9d\xef\xc2\x5c\xb9\x7b\x77\xc2\x34\xc2\x79\x4e\xc0\x73\x44\x1d\xc6\xd3\x4d\xbe\x3d\x48\xd1\xf7\x0f\x1d\x86\x57\x66\x45\x91\xee\x80\x87\xe1\xe9\xf9\x3d\x1b\x05\xb7\xe3\xff\x80\x7a\xae\x1c\xdf\xe6\x1f\x93\x83\x6b\x9c\x14\xee\xa1\x19\x77\x49\xcd\x73\x49\x49\x41\x80\x30\x06\x95\x55\xb4\xcc\x28\x8e\x32\x46\xf0\x1c\x6a\xfb\xca\xbe\x2e\x0a\x9b\xa4\xd9\x9c\x97\x67\x0b\xa7\x7e\xf6\xb7\xc6\xbc\xd9\x26\x5f\xe3\x36\x16\x02\xc9\xe8\xbf\xaf\x6f\x46\x63\xa8\xac\x1a\x93\x35\x07\x86\x1e\x03\x67\xf5\x65\x6f\xad\xd4\x7e\x99\x84\x12\x00\x72\xa3\x35\xe6\x9e\x9b\x43\xd3\x9c\xc3\x7c\x7e\x4d\x6d\x0a\x8a\x72\x06\x7f\x72\xa3\xd0\xac\x07\x88\xdc\x68\x2f\x75\x85\xfc\x8b\xd5\x37\x08\x94\x8d\x78\xad\x8b\xee\x6c\x5e\x2d\x42\x1b\xf6\xb3\xf4\xeb\x84\x8a\x9f\x3f\xb7\x66\x68\x61\xf8\xf6\xdb\xff\xb9\xc0\x74\x82\x78\xbb\xbf\x1b\x08\x9c\xd9\x75\x88\x7e\x71\x69\x9e\xea\x1c\xd6\x74\xe1\x55\x14\x75\x76\x09\x54\x36\x66\x3f\xe9\x52\x58\xb7\x16\x2a\x09\x63\xa2\x04\x6f\xb3\x97\xc2\x8b\x24\x4d\xc7\xf0\x67\x4c\xbf\xea\xd9\x58\x57\xac\x28\x17\x03\x67\xdc\x75\x3e\xb9\x84\x51\x5c\xfb\xca\x94\xa5\xd0\xc5\xe8\x24\x6b\xf8\x9c\x4c\xe0\xc6\x14\x62\xc7\x1e\x66\x11\x72\xca\x27\x61\xdc\x1d\xf6\x07\x61\xdb\x14\x89\xe2\x58\x96\xdc\x10\xbd\x6b\x86\x0b\x75\xfc\x27\xc2\x90\x7a\x22\x70\xdd\x30\xc3\xf7\x9a\xff\xb3\x42\x81\x42\xc1\xd6\x58\xc5\xfe\xcc\xec\xbc\x1c\xe5\x31\xc7\x97\x84\xf5\x9a\xe4\xf2\x1b\xe1\x1c\xc8\x90\xb1\x9a\x99\x70\xd6\x80\xbf\x22\x11\xb6\xc6\xda\x1d\x88\x05\x75\x0d\x45\x9c\x55\x48\x17\x22\xc4\x16\xf7\xa3\x16\x06\xa3\xd7\xf9\x41\x8d\xc9\x63\xbf\x06\xd3\x1b\x28\xa4\x13\x0b\x85\x71\xec\x52\x9a\x02\x5b\x3a\xe6\x1e\xfd\xc9\x25\x14\x59\x7b\x86\xf1\xb8\x82\x29\x4c\x50\x48\xa9\x07\xb8\x9d\x38\xf3\x3e\x60\x7e\xd8\x93\x1e\xbd\x82\x4b\xc0\xfd\x59\x91\xba\xdb\xa1\x2f\x54\x21\x72\x09\x92\xc2\x39\x4a\xbf\x0e\x71\x90\x89\x93\xf3\x94\x89\x97\xd2\xba\x98\xa5\x61\x8b\x4f\x6f\x11\x1c\xc5\xb8\x70\x5e\x7b\x2c\xd3\xe2\x7c\x91\x76\x5f\x51\x6f\x1e\xd6\x5a\x8b\x3a\xb6\x16\x1d\x8e\xbf\xa4\xed\x4c\xc5\xd5\x09\x29\x1d\x0b\x3a\xda\x9c\x47\xba\x3c\xe4\x29\xc0\x99\x12\xfd\x9a\x4e\xe3\x40\x82\x46\xd3\xf5\xc0\x7a\xd0\x13\x66\x5b\x7a\xd9\xdf\xf3\xd5\xb4\x75\xb3\x55\x64\xaa\xee\xaf\xd2\x4f\x78\x4e\x7d\x3e\xbd\x0d\x08\x1d\x76\x7f\x6b\xf2\x47\xcb\xf7\x00\xa8\x1c\x9f\x64\xeb\xf6\x72\x5e\x2d\x92\x53\xf8\x29\xfc\x1d\x3a\x79\xed\x8f\x16\x68\x1f\xf9\xe8\xcf\x03\xf1\x7c\x46\x15\xfd\x3b\x6f\x3c\xe4\x12\x8a\x2c\xf8\xde\x7d\x67\x86\x3c\x3b\x1a\x22\xd7\xb5\xdf\x4b\x93\x74\xfc\x71\x5c\x57\x53\x21\xf2\xb3\x5a\xfa\xb2\x44\x9c\xd8\x86\x99\x79\xd8\x68\x1c\x09\xb7\xef\x14\xff\x7e\xd9\x5c\x2a\x0e\xc2\xff\x63\xc1\x22\x5e\x59\x10\xb7\xd0\x71\x19\x1e\xd0\x4a\x1e\x62\xb6\x9d\x26\x56\x3a\xad\xc5\x8c\xe5\x01\x7a\x32\xba\x78\xf7\x8e\x6b\xca\xb0\x74\x47\x9a\x87\x26\x23\x1f\x8d\x95\x39\xae\x7c\xe6\x25\xcc\x20\x0c\xdf\xff\x05\x45\x3d\x3e\xb3\x3f\x71\xf5\xd0\xdc\x3a\x7c\x4a\x7f\xff\xb6\xfa\x1e\xd1\xde\xc3\xb0\x2b\x47\x50\x1d\x4c\x9e\x81\x36\xb0\x8c\xad\x38\x6f\x34\xfc\xf7\x82\x1d\x7a\x78\x36\x49\x7b\x07\xf6\xff\x1b\x00\x00\xff\xff\x0b\xf5\x91\xd7\x6e\x28\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xeb, 0xd5, 0x50, 0x29, 0x13, 0x67, 0xba, 0xf5, 0xf1, 0x59, 0x97, 0x63, 0x87, 0xc0, 0x3f, 0xe5, 0xc6, 0xda, 0xc0, 0x7, 0x6c, 0x4e, 0x4c, 0xc2, 0xdc, 0x12, 0xef, 0x1f, 0xaf, 0x7f, 0x8e, 0x40}}
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

var __hardcodedTracingGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x58\xdd\x6f\xdc\xb8\x11\x7f\x96\xfe\x8a\x39\x01\x77\x95\x02\x59\x8a\x73\xb9\xdc\xc1\x07\x3f\xf8\x2b\xc9\xa2\xb6\x6f\xe1\xf5\xf5\x5a\x14\x45\xc0\x95\x66\x25\x36\x12\xa9\x92\xd4\x7e\x20\xf0\xff\x5e\x0c\x29\x69\xb5\xf6\xee\x36\x17\x34\x0f\x8e\x96\x9c\x19\xce\x17\x67\x7e\xc3\x86\x65\x9f\x59\x81\xa0\x51\x2d\x51\x19\xc5\x32\x2e\x0a\xdf\xe7\x75\x23\x95\x81\xd0\xf7\x82\x4c\x0a\x83\x6b\x13\xf8\x5e\x80\x22\x93\x39\x17\x45\x3a\xe7\x82\xa9\xcd\xce\x52\x89\x6b\xfa\xbd\xa8\x2d\xa5\x40\x93\x96\xc6\x34\xf4\x2d\x35\xfd\xd5\x46\x65\x52\x2c\xe9\xd3\xf0\x1a\x03\xdf\xf7\x82\x82\x9b\xb2\x9d\x27\x99\xac\xd3\xab\x0a\x97\xa8\xd2\xcf\x6c\xb3\x44\x3c\x29\x64\xba\xfc\x39\xad\x64\x51\xa0\x22\xca\x34\x85\x31\x71\xce\x96\x98\x15\x65\x5a\xc8\x13\xdd\xe0\x2a\xa5\x3f\x24\xb8\x90\x89\x6c\x50\x18\xac\xb0\x46\xa3\x36\x09\x97\x29\xa9\xaf\xf8\x3c\xe5\x42\x1b\xd5\xd6\x28\x0c\x33\x5c\x8a\x74\x24\xae\x90\x8a\x57\x15\x4b\xeb\x76\x9d\x4a\x83\x55\xdd\xae\x0f\x49\xa3\xed\x63\x7b\x29\x33\x74\x5a\x6b\xf0\x28\xd5\x9c\x15\x05\x2b\x8e\xd3\xe0\x9a\x42\x80\x4a\xa7\xd2\x54\x8d\xfd\x43\xe1\xf9\x36\xa6\xed\x57\xa1\x9a\xec\xa8\x88\x46\xc9\x86\x15\xd6\x49\x47\xe9\x74\xfe\x39\x55\xa8\x65\xab\xac\x4e\x3a\xff\x6c\xe5\xc3\x71\x8e\xff\x6d\xc2\x40\xe5\xfe\x1a\xd4\x94\x51\x1a\x6b\x4a\x9f\x63\xd2\x1d\x45\xba\x3c\x4d\x4e\xdf\x24\xaf\x8f\x9e\xd1\x69\x11\xf9\xfe\x92\x29\xc8\x71\xc1\xda\xca\x5c\xc9\xaa\xc2\xcc\x48\xf5\x51\x6a\x03\xda\x28\x2e\x0a\x38\x87\xa0\x92\x19\xab\x4a\x49\x6a\xec\xa3\x9e\xd2\x45\x69\xb9\x30\xa7\xef\xe0\x1c\xde\xfe\x78\xfa\xb3\xef\x2f\x5a\x91\xc1\x0c\x4d\xdb\x7c\xa8\xe4\x9c\x55\x8f\x74\xde\x54\xc9\x25\xcf\x51\x5d\x88\xfc\xa6\x0b\x53\x98\x99\x35\x74\x17\x2c\xb9\x72\xff\x47\x10\xf6\xbe\x4c\x66\x0d\x13\x3d\x6d\x0c\xaf\x86\x75\x2b\x4f\xf5\x02\x63\x40\xa5\xa4\x8a\xe0\x8b\xef\x7b\x2c\xcf\x15\x9c\x9d\xc3\xa2\x36\xc9\xac\x51\x5c\x98\x45\x18\x7c\xaf\xcf\xbe\xcf\x83\x78\xaf\xa5\x2f\x57\xc9\xa2\xc8\xdd\xba\x9b\x25\xaa\x0d\x9c\xfe\x04\x1a\x33\x29\x72\x0d\x2b\xfc\x4b\x55\x81\x51\x1b\x30\x92\x34\x17\x98\x19\xfa\xdc\xf1\x33\x64\xbd\x28\x60\xc6\xca\x31\x25\xf6\xc7\x00\xf9\x93\xb2\x0b\xe4\x02\x06\xdf\x9e\x59\xc7\x11\xe9\x1f\x25\x0a\x50\xad\x10\xe4\x7e\x2e\xa0\x51\x32\x6f\x33\xcb\x60\x4a\xae\x81\x6b\x60\xa0\x79\x8e\x19\x53\x31\x30\x91\xc3\x6a\xc4\x61\x45\x58\xa9\xd5\x66\x44\xdf\xaf\xf4\x72\x65\x83\x83\xb6\x27\x83\xb6\x89\xef\xd1\x3d\xb9\xaa\x38\x0a\x43\x4e\xdc\xb9\x35\xc9\x3d\xae\xdc\x56\xe8\x7b\xde\xee\xd6\x1f\xdc\x94\x0f\xd8\xf9\x83\x4b\x31\x45\xc5\x65\x1e\x9e\xfe\xf4\x8a\x8a\x5d\x32\xb3\xde\x8b\xe2\xbd\x7c\x37\x22\x6f\x24\x17\x26\xa4\xd0\x1d\xa0\x99\x08\x8d\x59\xab\x30\xa4\xfd\xc8\xf7\xf4\x4e\x62\xa0\x52\x3b\xda\x92\xa6\x94\x5a\x31\x6c\xcd\x89\xac\x67\xf6\xf2\xfd\x9b\x61\x81\xca\x32\x75\x9f\x74\xe4\x90\x0e\x83\x7e\xa3\xcd\x61\x2d\xa0\x2a\x7f\x96\xa6\xdb\x40\x9e\xbe\x7d\xf3\xee\x97\x94\x35\xdc\xdd\x32\x1d\x44\x11\x25\x13\x5f\xd8\xe3\xbe\x3b\x07\xc1\x2b\xf8\xe2\x7b\x9e\x42\xd3\x2a\x41\x3f\x63\xf7\x87\x52\xf6\x86\x32\x79\x11\x06\x36\xa3\x21\x53\xc8\x0c\xc5\xab\xaf\x6b\x67\xf0\xfd\x32\xb0\x8a\x47\xbe\xf7\xe4\xfb\x9e\x69\xc8\x02\x81\xab\xdd\x2b\x11\xee\x1a\x2a\x70\xf5\xd0\x95\xaa\x30\x8a\x28\xca\x58\x25\x33\x34\xcf\x98\x4c\x43\x9a\xba\xb6\x93\xbc\x57\xb2\xee\xae\x24\xf9\x32\x4a\x26\x62\x21\xaf\xc3\x40\x1b\xa6\x48\xa7\x13\x6b\x9e\x0a\x62\xe8\x18\xee\xc8\xa8\x80\x82\x88\x5a\x07\x67\x40\x5f\xb1\xef\x3d\x45\x7e\x6f\xea\xae\x52\xa6\xb1\x76\xfb\x4f\x5d\xbd\x78\x69\x44\x6f\x35\x1c\x28\x08\x7d\xf9\x85\x57\xfd\x57\xd2\x9b\x19\x1d\xac\x16\xe4\x7b\xcd\xea\xa6\xe2\xa2\x98\x2a\x39\x67\x73\x5e\x71\xb3\x21\x37\xbe\x4e\x5e\xff\xe4\x7b\x5c\xdf\x52\x30\x6d\x46\xe9\xe4\x03\x1a\x14\xcb\x30\xf8\x34\x99\x7d\xba\xfd\xed\xea\xe2\x36\x88\xe0\xfc\x1c\x02\xa3\x5a\xaa\xe3\x7c\x01\x3d\x3d\x99\xbf\x4f\xf0\x39\x9c\x26\xaf\x7d\xef\x09\xb0\xd2\x08\x7c\x01\xcb\x67\xa2\x1f\x1f\x2e\xae\x26\xf7\x1f\x3e\xcd\x2e\xee\xa6\xb7\xf4\x31\x7d\xf8\xed\xf2\xe2\x72\x72\x3b\x79\xfc\x47\x10\xfd\x0a\x4b\xca\x9a\x20\x38\x74\x00\x05\xea\x46\x2c\x87\x74\xee\x90\x46\x32\x65\x4a\xe3\xfb\x4a\x32\x13\x2e\x63\x78\xf7\x36\xf2\xbd\x3d\x59\x78\x44\x22\x69\xee\x7b\xde\xd3\x41\xbb\x0e\xb3\x52\x72\x7a\x94\xcf\x53\xaa\xc0\x95\x08\x83\x3d\xb4\x67\x41\xbc\x4f\x44\x34\xa4\xf5\x10\xc1\xfb\x17\xa9\xe1\x7b\xb6\x56\x22\xb4\x1a\x77\xaa\xeb\xe4\x1a\x0a\x14\xa8\x18\xd5\x33\x98\x08\x90\x8a\x82\xbe\x90\x6a\x38\x8b\xea\xf5\x4a\xaa\xcf\x10\x32\x03\x15\x32\x6d\x60\xc5\x4d\xe9\xca\xa5\x25\x42\xba\x5e\x7d\xdd\x1e\x4b\x84\xba\xd5\xa6\xff\x89\xe0\x3a\xfe\xe4\x5a\x43\x2b\xf8\x42\xaa\xba\xda\x00\x33\xa0\x98\xc8\x65\x0d\x0b\x25\x6b\x2b\x01\x85\xe1\x0a\x29\xfd\x33\xa4\xaa\x4f\xdd\xf2\xdd\xdb\xc4\x1d\xf1\x5e\x2a\xc0\xb5\x3d\x35\xde\xb1\x64\x7d\xa2\xd8\x66\xf7\xf4\x5c\xa2\x06\x21\x0d\xe4\xd2\x6a\xdb\x89\x18\xfc\x44\xa5\x69\xe6\x0c\x08\x77\xd3\x7f\x72\xfd\x40\x4d\xe7\x92\x69\xcc\xc3\xc8\x56\xd9\xe3\x5c\x53\xa6\x50\x18\x47\x7f\x44\xd4\xbe\xf8\x39\xf1\x69\x0a\x8f\x25\x6a\x84\x9a\xad\x79\xdd\xd6\x1a\x98\x42\x72\x7d\xd1\x32\x95\x03\x2b\x18\x41\x52\xd0\xb2\x46\x53\x52\x54\x0a\x49\x7f\x57\x4a\x8a\xc2\xb6\x35\x8d\x82\x90\x35\x30\x30\xae\x59\xe6\xcc\x30\x68\x05\xae\x1b\xcc\x0c\xe6\xd5\xe6\x85\x11\x0d\x13\xb7\xbc\xe6\x46\xef\x42\x08\xb7\x66\xb3\xfd\xa2\xc7\xa7\x57\xb2\x15\xc6\x6e\x9c\xc1\xe9\xeb\xd7\xa4\xb1\x77\xb3\x44\x61\xc6\x1b\xf4\xaf\xdf\xbc\xe5\xe2\xf3\xf3\xbd\x7e\xf3\x29\x8a\x7d\x67\xf2\x25\x33\x59\x89\x8a\xba\x6e\x2d\x15\x02\x2e\x16\x3c\xa3\xfe\x13\x83\x5e\x71\x93\x95\xe4\x01\x6e\x80\x2d\xa8\xae\x11\xb8\xb3\x4d\xfb\x65\x14\x37\x22\x1b\x55\xc0\x97\x01\xeb\x0e\x3a\x42\x31\x54\xfc\xbe\x38\x76\xbd\xf3\x48\xf1\xdf\xee\xe1\xda\xdc\xb1\x66\xda\x61\x61\xa9\xc2\x11\x2c\xb6\x40\x40\xd6\x8d\xd4\xdc\xe0\x71\x4a\x7b\x46\xd7\x45\xbe\x3c\xc5\x30\xde\xbb\x74\x43\xc0\x97\x27\xdb\x1f\xbb\x16\x61\x1a\xea\x07\x69\xfa\x35\xe8\xf1\xbd\x54\x8f\xa8\x8d\xf5\x35\x32\x61\x71\xd8\xdc\x56\x84\x9c\x80\x53\x2b\xb8\xe9\x3d\x1c\x93\x48\x4a\xaa\x9a\x53\x63\xd5\xf6\xa2\x69\x3a\x02\xd8\x5c\x2e\xd1\xd6\x07\xd9\x1a\xc2\x55\xf6\x86\x8e\x78\x13\x98\x18\x70\xea\x69\x60\x02\xb8\x38\xa9\xb1\x96\x6a\x43\x32\x87\x16\xb5\xe8\xae\x31\xb7\xf8\xaa\xaf\x10\xb9\xed\x79\x3a\xf9\x6a\x44\xdc\xd9\x14\x46\x10\xbe\x1a\xf0\x7f\x32\x11\x77\xf6\xc4\x3f\x85\x85\xbd\x41\xb7\xb3\x73\xd8\xca\xba\xc7\xd5\x73\x71\x61\x74\x18\x48\xe0\xb7\x80\x88\x3e\x9a\xb8\xaf\xd9\xa7\x29\xdc\xb5\xeb\x99\x1d\xb8\xef\x78\x9e\x57\xb8\xa2\xca\xd0\x7b\xb8\xde\x2e\x99\x92\x19\xd0\xa5\x6c\xab\x9c\x02\xcb\x8c\x61\x59\x89\x39\xc5\x99\xc1\x68\x6a\xed\xa6\xf7\x84\x44\x4f\x8c\xab\x92\x66\x65\x4b\xa4\x28\xf4\x19\x58\xc8\xa2\x5d\x28\x1c\x64\x66\x79\xee\x7e\x3b\x90\x06\x5c\x2c\x24\x89\xa5\xb4\x50\xf8\x9f\x16\xb5\xa1\xd1\x3a\xe3\x0b\x9e\xf5\xf8\x86\xa4\x3f\xf0\xa2\x34\x20\xe4\x0a\x56\x08\x52\x54\x1b\xd0\x6d\x63\x9f\x0a\x88\x88\x22\x4f\xcd\x80\x0b\x2b\x88\x1a\x02\x33\xce\x8a\x6b\x66\x58\x2e\x2d\x90\xc3\xcc\xf4\x09\xb1\xc7\x0f\x21\xd9\xc2\x33\xbc\x67\x35\x76\x43\x58\x04\x44\x1c\x12\xd2\x4c\x3e\x32\x91\x53\x7b\x82\xf1\x2f\x0a\xb6\x34\x34\xb7\x3b\x18\x6c\x47\xf8\x64\xbf\xcc\x78\xd8\xa7\x1a\xb1\xbd\xb6\x3a\xb4\xf1\xfc\xb0\xef\xee\x13\x82\xdd\xed\xe6\x17\xb9\x2d\xcd\x5b\xdf\x8f\xc2\x16\x6c\xe3\xef\xf4\x86\xe3\x9a\xf7\xc4\xce\x82\x1d\x33\xdf\x13\xbf\x15\xa2\x56\x8e\xef\x01\x75\x23\x85\xc6\x3f\x14\x77\x10\x10\x5e\x75\xeb\x36\x68\x91\xc3\x34\x34\xa7\x2a\x9e\x77\xfe\xa3\x95\x34\xed\xed\x86\x92\x69\xc0\x35\xc5\x9d\xee\xa7\x2d\x04\x0d\x13\x89\x8d\x6a\xd3\x1a\xaa\xce\x5c\x74\xb9\x90\x99\xf5\xf3\x3c\xe8\xa4\xd1\x11\x19\x53\x8a\xa3\xda\xa9\x6a\x1f\x91\xe5\xa8\xae\xba\x9d\xf3\xc3\x7b\x5f\x08\xbe\xff\x5f\x15\xb3\x68\x6e\xb8\xec\xb6\xf3\x8d\x41\xbc\xea\x27\x6c\x7b\x7b\x3d\x6f\xce\x0a\x22\xee\xde\x62\x92\x63\xa4\xe4\x4b\x4b\x99\xdc\x61\x3d\x47\x15\x06\x7f\x3f\xe9\x3c\x7e\x32\xb9\x0e\xa2\xe4\x6f\xac\x6a\x31\x8c\xac\x0a\x7c\x61\x7d\x7f\x3e\x60\xd6\x9e\x5f\x75\x0e\xa0\x24\x7b\x2e\xc1\x92\xed\xe1\xec\x58\xb5\xb5\x66\x50\xaa\x07\x22\x61\x94\xcc\x6c\x88\x43\x27\x80\xa0\xaa\xf7\x34\xd6\xe2\xbb\x91\x2c\x4d\xd5\x6a\x00\x00\x3a\x1c\xde\xaa\x7a\x21\xbb\x3a\xc5\x24\xc0\xd9\xef\xd5\xd6\xec\x01\x62\xf7\x3e\xbb\xc7\xd5\x5e\x87\x38\xd6\xde\xa6\xe7\x88\x9b\x34\xa1\x61\x59\xe5\x76\xda\x0b\xdd\x34\xd7\xab\x4f\x71\x71\x07\x39\x8f\xcf\xd0\x74\x67\x38\x25\x1c\x69\x66\xd6\x63\x45\x3a\xcf\x38\x60\x60\x97\xc6\x31\x8c\x89\xce\xf1\x29\x1b\x07\x37\xdf\x6e\x67\x3b\xdf\xed\xad\xba\x00\x85\x51\x72\x91\xe7\x87\x6d\xda\x33\x21\x8e\x33\x86\x98\xfb\x5f\x87\x64\x0c\x41\xd2\x99\x45\xfa\xbb\xf1\xfd\x15\x74\x96\x7c\x64\x7a\x08\x73\x1f\xc1\x86\x89\xc9\x75\xec\x12\x7c\x72\x6d\x39\x33\xcb\xba\x93\x0b\x31\xad\x1e\x48\x91\x34\xb5\x38\x92\x6a\x31\x4d\x49\x48\x8d\xa1\xc4\x75\x57\x2a\x34\x35\x01\x07\xcf\xa9\x92\xc7\xa0\x25\xdd\x2e\x30\xa5\xd4\x48\x3f\x6c\x31\xcf\xa4\x52\x58\xb9\xc9\xbc\x92\x85\xb6\x2d\xc5\x4d\xfb\x76\xb0\xd0\x7d\xe4\x2b\x14\x61\xa7\xab\x9d\x1a\x7f\x7c\x03\x3f\xfc\x60\x57\x9d\x25\x76\xf1\xf4\x1d\x7c\x01\x2a\x03\x3b\x4f\x48\xae\x79\x9c\xd1\xee\x7c\x63\x10\xc2\x1f\xdf\x9c\x64\x25\x53\xa4\x6c\x14\xc3\x2f\xdd\xea\xe9\xbb\xed\x6a\x37\x8e\x58\xec\xdc\x10\x50\xc9\xb5\x0b\xac\xd7\xe9\x70\xa9\x63\xf8\x44\x4e\x2b\x71\x9d\x5c\x63\x26\x73\xec\x7c\xd3\x2b\xe9\xc8\xff\x4c\x7c\xf3\x3c\xb1\xcc\x9f\x78\x1e\xc4\x8e\xdd\xdb\x7d\x7b\xcb\x83\x18\xdc\x83\x79\x72\xc9\x8b\x1b\x91\x73\x26\x92\xdf\xad\x8f\xc3\x41\xb1\x7f\xfe\x72\xf6\xaf\xa8\xbb\x6c\x5d\x98\x0f\x6b\xdb\x39\xef\x9b\x94\x25\xde\x6f\xd3\xb5\xd7\xaa\x57\xb3\x2b\x36\x9e\xe7\x95\x89\xed\xe2\x1f\x1f\x1f\xa7\xa1\x5a\xc5\x60\xef\x28\xc1\x5a\xef\xa9\x03\x3c\x23\xec\xb4\x85\x92\xdb\xf7\x8b\x1c\x75\xa6\xf8\xdc\xce\xa6\x34\x84\xb2\xa6\xa9\xb8\x7b\x25\xb4\xb0\xe3\x77\x82\xb4\x76\x84\x45\x63\xb3\xae\x6d\x5c\xb0\x6d\xef\xb1\xb8\x6b\x78\x43\xd9\x62\xb4\x3d\xcf\x22\x74\x8f\xa8\x73\xb1\xa6\x19\x01\x0c\xfb\x8c\x31\x7e\xeb\xb8\x98\x4e\x3f\xdd\x5f\xdc\xdd\x04\xd1\xa8\x80\xf6\x3c\xe7\x07\x48\xc7\x4f\x1c\x23\x8a\xaf\x97\xb5\x23\xca\xf7\x54\x17\xfe\xc1\x88\x3b\x54\x05\x86\x16\x30\x74\x2b\xd7\x6e\x54\xb6\xaf\x82\xdb\xd5\x7b\x5c\x51\x81\x1b\x55\x7a\xdb\x19\xdd\xfb\x78\x32\xcb\x4a\xac\xd9\xef\x0f\xb7\xf1\xce\xea\x16\x21\xfd\x15\x37\x7d\xc9\xe8\xd4\xb4\xe2\xbb\xe9\xa9\x03\x2b\xca\x7f\xf2\xff\x1b\x00\x00\xff\xff\x4f\x0c\x49\xd3\x3c\x1a\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x62, 0xcd, 0xd9, 0xe3, 0xc5, 0x70, 0x6d, 0xa0, 0x44, 0xd5, 0x1b, 0x3f, 0x93, 0x8a, 0xd7, 0xd7, 0x5f, 0xc, 0xf6, 0x74, 0x44, 0x8, 0xf3, 0xeb, 0x9b, 0x88, 0xab, 0xd0, 0xcf, 0x5d, 0xa8, 0x8b}}
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
