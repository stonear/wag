// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../_hardcoded/doer.go (10.327kB)
// ../_hardcoded/middleware.go (1.685kB)
// ../_hardcoded/tracing.go (1.824kB)

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

var __hardcodedDoerGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x3a\x7f\x6f\xdb\x46\xb2\x7f\x4b\x9f\x62\x22\xe0\x2e\x64\x22\x53\x72\xee\xdc\x1e\xd4\xe7\x03\x5a\xc7\x7d\x2d\x90\x26\x41\xe5\xbe\x06\x08\x82\xeb\x8a\x1c\x4a\xdb\x2c\x77\xd5\xdd\xa5\x65\xc1\xe7\xef\xfe\x30\xb3\x4b\x8a\xd4\x0f\x27\x77\x2d\x1e\xde\xf9\x0f\x49\xe4\xce\xcc\xce\xcc\xce\xef\xf5\x5a\xe4\x1f\xc5\x12\x21\x57\x12\xb5\x1f\x0e\x65\xb5\x36\xd6\x43\x32\x1c\x8c\x16\x5b\x8f\x6e\x34\x1c\x8c\x72\xa3\x3d\xde\x79\xfa\x89\x3a\x37\x85\xd4\xcb\xc9\xaf\xce\x68\x7e\x61\xad\xb1\x0c\x25\xcd\x44\x9a\xda\x4b\x45\x0f\xca\x2c\xe9\xab\x12\x7e\x35\xb1\x42\x17\xf4\xa0\xd1\xc7\xaf\xc9\xca\xfb\x35\xfd\x76\x5b\x9d\xd3\xb7\x97\x15\x8e\x86\xc3\xc1\x68\x29\xfd\xaa\x5e\x64\xb9\xa9\x26\xa2\xc4\xbb\xc9\x6a\xeb\xbc\x95\x77\x67\x4b\xd3\xfc\x1c\xf5\xa1\x0a\xa3\xcd\xad\xd0\x2b\x59\xe0\x04\x6f\x51\x7b\x67\x6a\x9b\xe3\x68\x98\x0e\x87\x93\x09\x14\x06\x2d\x48\x07\x42\x83\xd4\x1e\x6d\x29\x72\x84\xd2\x58\x18\x15\x46\xea\xe5\x08\x88\x13\xb0\xf8\x5b\x8d\xce\x3b\x58\x1b\xe7\xe4\x42\x6d\x61\x23\xfd\x0a\x36\x56\xac\xd7\x52\x2f\x87\x7e\xbb\xc6\x48\xaa\x25\x72\x3f\x1c\xbc\x34\x49\x0e\xcf\x88\x42\x76\xc5\xea\x1b\x83\x8d\xcf\x3f\x06\x8a\x29\x24\xcd\xb3\x5b\x1b\xed\x70\x0c\xac\xb0\x74\xf8\x30\x0c\x54\xcd\xfa\xb5\xa8\xf0\xca\xdf\x81\xf3\xb6\xce\xfd\xfd\x03\xf3\xbd\x10\x0e\x23\x8d\xef\x84\x2e\x14\x5a\x58\xa3\x2d\x8d\xad\x1c\xf8\x15\xf2\x7a\x8f\xf7\x40\x8d\x5e\xbf\x24\x3e\x77\xc4\xca\x5a\xe7\x90\x14\xed\x52\x0a\xbf\x83\x6d\x92\xda\xa2\xaf\xad\x86\x3c\x7b\x69\x12\x9b\x0e\x07\x93\x09\x34\xaf\xfc\x1d\x63\xd1\x42\x76\x15\x8c\x26\x49\xc7\x90\x8f\x21\x48\x1c\x40\xed\xb6\x11\x89\x1e\x24\x3a\xb8\x98\xbe\xeb\x1f\x44\x90\x86\x61\x3b\xe2\xd0\xee\x05\xec\xfe\xe8\x44\x98\x1f\xbb\x7d\x6b\x94\xcc\xb7\xf0\xe3\xee\x77\xdc\xaf\xf3\x06\x0a\x2c\xa5\x46\x07\x22\x50\x86\x35\xbf\xce\xc2\x66\x5d\xc0\xde\x29\x4f\x26\xf0\x8d\xc8\x3f\x9a\xb2\x74\x51\xd0\x70\x04\xba\xae\x16\x68\x41\xe8\x02\xbc\xac\xa4\x5e\x82\x29\x23\x61\xe1\x3d\x56\x6b\xef\xb2\xe1\xa0\x41\x4d\x52\x78\xff\x81\xcc\x3c\x7b\x59\x5b\xe1\xa5\xd1\x4c\x99\x77\x05\x8b\x39\xca\x5b\x0c\x74\xbb\x8a\x18\x83\x70\xb0\x41\xa5\xe8\x9b\x16\x2d\xba\x5a\x79\x30\x25\x63\x37\xae\x14\xcf\xf1\xa9\x83\x5f\x5e\x9a\x5f\xa0\x42\xbf\x32\x45\x36\x1c\x30\xf5\xa4\x77\xb4\x63\x38\x71\xb2\x0b\x63\x54\xd4\xd9\x5c\xea\xa5\xc2\xc7\x34\xe7\x57\xc2\xb7\xa7\x27\x1a\x6e\xc1\xe8\x1c\x83\x36\x0f\x49\xf4\xcc\xfb\x88\x42\x85\x87\xad\xa9\xc1\xad\x4c\xad\x8a\x76\x1b\x6c\x69\x9f\x3b\xcc\x8d\x2e\x40\x94\x9e\xfc\xd0\x43\x29\xa4\x72\x59\xb4\xef\x83\xfd\x52\x38\xad\xf9\x8e\x15\xef\xad\xdc\x9f\xc3\x33\xe0\x37\x73\xde\xed\xa1\x6b\x45\xb0\x91\x4a\x45\xce\xb4\xd1\x67\x6f\xdf\xcc\x6f\xc6\xe1\xd7\xd7\x37\x57\xdf\xed\xa2\x08\x0b\x73\xf1\xee\x5d\x46\xa8\x37\x6f\x5e\xbe\x99\xc1\xf7\x9e\xac\xd5\x81\x36\x1e\xf2\xda\x5a\xd4\x5e\x6d\x1b\x6b\xd1\xdb\x70\x08\x8d\x3a\xb0\x80\xc5\xf6\xc4\xe1\x3e\x22\x70\x38\x6d\x8b\xbf\xc1\xde\x89\x5b\x74\xeb\x63\xc7\xde\x3d\x7a\xd2\x89\x2c\xf9\xed\x93\x4b\xd0\x52\xc1\x3f\xff\x49\x12\x65\x3f\xb0\x31\xc1\xe5\x25\x8c\x48\xe0\xd1\xb1\xf7\x24\x3e\x2d\x0c\x07\x03\xda\x2b\x9b\x7b\xe1\x6b\x77\x65\x0a\x84\xff\x82\x8b\xe9\x94\x88\x37\x1a\x2f\x85\x72\x38\x1c\x3c\xb4\x47\xe0\x6d\x8d\x51\xcd\xd7\x77\x6b\xa3\x51\x7b\x29\xd4\x51\xeb\xd3\x80\x3b\x88\x9e\x13\x07\xab\x3b\x81\xff\xb8\xe9\x95\xf2\x96\x02\x67\x7c\xcb\x81\xbf\xb3\x8b\xa2\x68\x90\x5b\x14\x8e\x3c\x7c\x23\xa4\x67\xfb\x70\x1c\xa4\xd1\x6f\x10\x75\x7b\xf0\x33\x38\x9f\x4e\xc7\xf0\x82\x3e\xfe\x4a\x1f\x7f\xa3\x0f\x0a\x10\xe7\x5f\x4c\xa7\x50\x49\xa5\x64\xb0\x62\x07\xcf\x27\x67\x50\xaf\xc1\x1b\xb8\xf8\x13\xfc\x2a\xbd\x47\xdb\x1c\xed\x71\x29\x3e\xc3\xa0\x61\x76\x09\x95\xf8\x88\xc9\xde\xf2\x18\x2e\xd2\xe1\x40\xe3\x1d\x43\x9c\x4f\xa7\x8d\x91\xff\xb0\xe3\x68\x38\xb0\xba\xa0\x65\xca\xd6\xd9\x6b\xdc\x24\xcd\x8f\x39\x27\xd4\x84\x11\x5e\x9b\x4d\x92\x66\x3f\x69\x79\xf7\x5a\x68\x93\xa4\x69\x3a\x1c\x20\x61\x4d\xb3\xe9\x05\x4c\x26\x2c\xd6\x05\xa5\xaa\x1c\xb5\x8f\x72\x0d\x07\x94\x70\x65\x24\xbe\xe4\xb8\xde\x58\xc4\x7b\xf9\x01\x2e\x81\x59\x7b\x0e\x3d\xa6\x93\x24\xb1\xba\xc8\xbe\x55\x46\xf8\x2f\xfe\x9a\xa4\xcf\x5e\xa4\x67\xe7\xe9\x33\x7c\x56\xc6\x37\x84\x44\xfb\x07\xc1\x9e\x5d\xc2\x8b\xae\x55\x59\xf4\xff\x6f\x7d\xf7\xd4\x01\xff\xe7\x3b\xf0\x6b\x73\x3c\x6b\x04\x27\xa5\x04\xa9\xf1\x36\x26\xfe\x2d\x39\x54\x9b\x3b\x62\x16\xee\x13\x78\xdc\x71\x29\x1c\x54\x6b\xbf\x05\xa7\x64\x8e\x8d\x72\x7b\x14\xfe\xcd\x2c\xd0\x0f\xfb\x42\x6d\xc4\xb6\x13\x2f\x48\xfe\x13\xbb\xfd\xab\x09\xb7\xc3\x43\x50\x6b\x53\x16\xb2\x82\x62\x09\xd5\x57\xc3\xcf\xd2\xaf\xba\x3a\x6a\xb5\x01\x1a\x37\x10\x4b\xf5\x60\xc6\xe6\x16\xad\x95\x45\xac\x2b\x42\x85\x0f\x66\xf1\x2b\xe6\xfe\xa9\x6b\x8b\xb1\xb6\x0c\x62\x89\xf6\xc8\x27\xb9\xbf\x6b\x88\x36\x25\xdd\x18\x4e\x94\x5c\xe9\x3e\x64\xb7\x5c\x8c\x2b\x44\xff\x7f\x84\xaa\x91\x28\x8f\x7b\x72\xde\x3f\xf4\x28\x73\xc5\xd8\x54\xaf\xcf\xda\x5a\xf0\x8f\xa8\x5f\x9b\x2d\xc6\x60\x3e\x72\x58\xda\x55\xab\x59\x60\xae\xcf\x57\x9a\x25\x5d\x31\xd9\xd1\x9e\x98\x8f\x8d\x57\xb4\xba\xb8\x84\x22\xeb\x3c\xb3\x93\xb4\x99\x85\xf6\xd9\xad\x65\x3b\xc3\x1c\x0e\x6e\x85\x3d\xe6\xe3\x61\xa1\xf5\xf3\x21\xd7\x7c\x73\x71\x8b\xbd\xc2\x68\x61\x0a\xca\x51\x90\x53\x4b\xb0\x41\x58\x31\x80\x09\x9b\x65\xf0\xc6\xaf\xd0\x6e\x64\x58\xe3\x28\xc8\x00\x42\x59\x14\x05\x99\x8f\x28\x98\x2e\x77\x15\x75\x59\xa2\x05\xa3\xdb\xe8\x56\xf4\xb6\x62\x74\xaa\xbb\x32\x98\x23\x32\x1a\x31\x3c\x9b\x4c\x9c\x27\x71\x6e\xd1\x96\xca\x6c\xb8\x31\x63\x0c\x69\xb4\x9b\xbc\xf8\xcb\xf4\xcb\xe9\xdf\xbe\xfc\x62\x42\x7b\x49\xbd\x3c\x23\x8e\xcf\x4c\x79\x46\xb8\x67\x91\xf6\x19\x25\x5e\x53\xfb\xb3\xca\x14\xb2\xa4\xd8\xd0\xae\x38\x2f\x7c\xd4\xc5\xa2\x2e\xe1\xfd\x07\x6a\x4c\xf9\x0c\x6c\xf6\x0d\x09\x1f\xe3\x1d\x9d\x46\x5f\x61\x83\xc1\xa2\x2e\x43\xa4\xbc\x84\xd0\x9c\x66\x3f\xa2\x28\xbe\x56\x2a\x09\xb8\x94\x3f\xfa\x51\x93\xa8\x34\x46\xab\xa5\x62\xec\xe1\x80\x4e\xf2\x61\x18\x32\x59\x53\xf7\x52\xda\xfb\x8a\xa3\xdf\x57\xcd\xbb\xe7\xcf\x19\xff\x28\x6b\x03\x5b\x58\xc2\x89\x7c\xbc\x36\xeb\x2b\x65\x1c\xda\x84\xfb\x6c\xca\xb4\xdf\xb0\xfa\x93\x45\x5d\x72\x5e\x1b\x44\x1a\x97\x60\x8b\xc8\x02\x47\xe7\x46\xa0\x22\x2b\xa8\xd5\x0a\xdd\x55\xd8\x34\x72\x76\x79\x09\x0a\x75\xd2\x98\x5e\x4a\x21\xff\x49\xd7\xf8\x62\xa6\x09\x99\x85\xe9\xa5\x81\xc7\x85\x45\xf1\x31\xee\x35\x99\x00\x73\xd8\x34\x1e\x6c\x93\xd1\xdc\xca\xdd\x0b\x19\x32\xa3\x96\xaa\x61\xc2\xad\xf7\x94\xe9\xd6\x2c\x49\xc6\xe4\xc8\xe0\x99\x7e\xa8\xb3\x15\xe2\xba\xe5\xf4\x7d\x94\xe0\x43\xda\xcf\xe5\x91\xc9\x18\x96\x73\x69\xf3\x5a\xfa\x6f\x88\x57\xb4\xdc\x20\xca\x6a\xad\xb0\xa2\xf6\x3f\xc4\xba\x00\x01\x8b\x00\x02\x6b\xea\xc8\xac\xe6\xa4\x3e\x47\x0f\x05\x2e\xea\x25\x39\x09\x9d\x1e\x7d\x9b\x35\x5a\xe1\xb1\x87\x2c\x35\x08\xa8\x28\x07\x6e\x56\x68\x91\x7a\x0e\x76\x00\xa3\xd5\x16\x94\x59\x76\x58\x01\xb6\xd1\x50\xeb\x31\x0c\x69\x64\xa1\x4c\xfe\x91\x2b\x84\xb6\xba\x28\xad\xa9\x60\x69\x28\xf3\xf9\x95\x35\xf5\x72\x15\xf3\xde\x11\x91\x1e\xeb\x79\x03\xff\xe1\x8f\x1b\xb7\x41\x24\xf0\x5a\x54\x48\x98\x52\x2f\x87\x03\x65\x96\x4b\xb4\x01\xea\x67\xb1\x0c\xc1\xf2\x15\xbf\x24\x4d\x92\xaf\x44\xb4\xf9\xfc\xfa\x8d\xce\x11\xdc\x56\xe7\x19\xfd\x62\x3d\x7f\x17\x26\x2e\xf3\xf9\xf5\xf5\x2d\xa5\x0e\xe9\x00\x2b\x2a\xe9\xb8\xc6\xd9\x8d\x66\xe0\x56\x0a\x70\x68\x6f\xd1\x9e\x39\x02\x0c\x83\x98\x8c\x2b\x27\x74\xb9\x95\x8b\x50\x30\x93\x76\x59\x53\x54\x09\x88\x66\xf3\xa8\x82\xfd\xdd\x76\xf2\xdf\xd0\xf2\x23\x7f\x41\x5e\xf8\xe5\x57\x67\xf4\x6c\x44\xc4\x46\xbf\x0c\x07\xac\x8a\xcf\xc6\xd2\xa2\x62\xac\x98\x41\xae\x4c\xad\xfd\x31\x2c\x19\x5e\x47\x2c\xdb\x81\x26\xec\x6b\x0a\x3b\x27\x71\xf7\xb1\xb1\x85\x6e\x71\xdf\x86\xca\x59\x2c\x0f\x59\x3f\x82\xbb\x83\x26\x02\xdf\xbb\xab\x9e\x15\xbd\x59\xa3\xee\x12\xe0\x82\xa3\x25\x20\x8f\x40\xb3\x02\x8c\x52\x52\x2f\x99\xad\x6f\x85\x54\xb5\xc5\xc7\x14\x70\x08\x7d\x48\x44\x29\xf2\xef\x2e\xb1\xc7\x88\xf4\xa0\x4f\x11\x9b\xd7\x79\x8e\xce\x7d\x26\xb1\x08\xbd\x4f\x6c\xbe\x32\xd6\x47\x2d\x60\xf1\x29\xf1\xfa\xd0\x07\xb4\x3a\x0c\x7d\x5a\x55\x27\x18\xba\x59\x51\x8a\x7c\x6b\x8c\xfa\x11\xa9\x4a\x23\xa6\x4e\x13\x39\x84\x3e\xa0\x27\x2b\x34\xb5\xff\x4c\xa6\x22\x34\x11\xb9\x0a\x6d\xce\x95\xd1\xb1\xe1\xb9\xbe\xc3\xbc\xa6\x74\x1e\x8c\xbb\x4f\x24\x7f\x1c\x9a\x08\xbe\x12\x1e\x75\xbe\xbd\x31\x5e\xa8\x1f\x50\xe8\xc7\x8d\x5b\x75\xa0\xff\x51\xa1\x20\xc3\x6c\x8a\x41\x65\x96\x1c\x1f\x12\xb5\x1f\xd3\xc6\x70\x10\x44\x38\xab\xa9\xec\x95\x59\x26\x23\xa9\x4b\x33\x1a\xc3\x68\x34\x86\x4a\xac\xdf\x07\xf7\xff\xd0\x8e\xfa\xee\x1f\x28\x59\xf5\x3d\x7a\x76\xcc\x87\x31\xeb\xc6\x88\x31\x21\x75\x1c\xf9\x28\x0a\x21\xed\x42\xc3\x0e\xa5\xe3\xbf\x87\x78\x11\x65\x07\xc3\x78\x47\xdd\x76\xd6\xc7\x3b\x16\x08\x18\xf9\x98\xb7\xce\xf6\x85\x3b\x84\x39\x82\xdb\x77\xd2\xd9\x51\xdc\x1e\xcc\x49\x1a\x8d\x2b\x3c\x46\x23\xc2\x1c\xd0\xd8\x73\xc9\xd9\x31\x19\xfa\x30\x87\x24\x7a\xdb\x1f\x57\xc3\xa9\xed\x8f\x38\xe0\x6c\x1f\xf7\x10\xe6\x90\x4c\xf4\xbb\xc7\x58\x88\x30\x8c\xfb\x29\x77\x9b\x01\x66\x9f\x70\x60\xa6\xa3\xf6\x7c\xf2\xd0\x06\x31\xdb\xf7\xdb\xf1\x70\xf0\xd0\x6f\xcc\x0e\x2b\x97\x14\xa4\x96\x3e\x49\xe3\xf0\xfc\x2d\x5a\x69\x0a\x99\xf3\x2c\x4d\x99\x65\x98\xae\x6b\xa1\xfa\xc5\x13\x55\x61\xc2\x39\xe9\xb8\xf0\x72\xe8\x3d\xd7\x30\x9d\x1a\xcb\xaf\x2c\xba\x95\x51\x85\xe3\x3a\xab\xd6\x05\x5a\xe7\x85\x2e\x78\x90\xb0\x5e\x2b\x99\x87\xc6\x7e\x81\x2b\x71\x2b\x8d\xcd\x18\xfd\x27\x5d\x1a\xeb\x6b\x2d\x3c\xaa\xed\xb8\x5b\xb6\x14\x06\x9d\x7e\xea\x63\x43\x04\x4b\x8b\xc2\xc3\x46\x6c\x89\x15\xbc\x5b\x87\xd2\x57\x3a\xa6\x52\x08\x2f\xc6\xe0\x0c\x75\x51\x16\x9d\xb1\x9e\xcb\xc7\xda\x6a\x9e\xfc\x6b\x90\xde\x81\xab\xd7\x7c\x7b\x46\x5d\xc2\xc2\x1a\x51\xe4\xc2\x1d\x88\x51\x51\x89\x9b\x3b\x2e\x9a\x78\xd6\x7f\x58\x39\x41\x32\x9f\x5f\xa7\x81\xfd\x39\x86\x2b\x01\x37\x9b\x4c\x1e\xb9\x1d\x93\xce\xd5\xe8\x26\x17\x5f\x64\x6d\x2d\x18\x8b\x3a\xea\x11\xe8\xb4\xc2\x81\x0c\x22\xce\xdc\x5b\x14\x55\x73\x11\x33\xbb\x6c\xf4\xc2\x13\xbf\xee\x1a\x97\xeb\xc7\x90\xb2\xb9\x17\xd6\xf3\xb2\x92\xce\xa3\xe6\x08\x6c\x99\x98\x46\x9f\xbd\xe2\x97\xc9\xc8\xe7\x6b\x0a\xbb\xe7\x2f\xbe\xcc\xa6\xd9\x34\x3b\x9f\x4d\x47\x27\xba\xae\xb5\xd0\x32\x4f\xa8\x19\xe1\x0e\x61\x38\x18\x2c\x0d\xcb\x9e\xcd\x49\x45\xc9\x6e\x9b\x63\xfc\xa4\x01\xbe\x23\x2a\xa9\xef\x67\x64\x9b\xc3\x5b\xb4\x5b\x38\x9f\x42\x33\x7f\x5d\x60\x2e\xea\xd8\xdb\x10\x35\xb4\xf2\xae\x3d\x1a\x61\x91\x8e\x54\xc0\xf9\xf4\xac\x19\x8f\x32\xb5\xe8\xb6\xb0\x91\xba\x30\x9b\x0c\x6e\x56\xd2\x81\x42\xff\xd4\x41\x4d\xc7\x5f\xf1\x4c\xf7\x18\x49\x6f\x60\x89\x7c\x69\x61\xd5\x36\x52\x13\x79\x5e\x87\xae\x83\xdc\x2b\x5e\x28\x39\x48\xb4\xe1\x4e\x44\x78\x36\x3f\x90\x6c\xa4\x6b\xb4\x25\xe6\xbe\x65\xdc\xd4\x36\xde\x3b\x45\x6a\x8d\x35\x47\x40\xb5\x85\x4a\xf8\x7c\xd5\xb2\xf2\xd4\xc5\xfe\x84\x9a\x7f\x67\x58\x2d\x9b\x15\xea\x63\x0d\x53\x24\xe9\x78\x0c\x08\xf9\x4a\xe8\x25\x3a\xee\x49\x95\x59\x7e\xcb\x39\x52\xe7\xdb\x30\x4d\xee\xdf\x98\x30\x8c\x70\x9e\xd3\xef\x1c\x51\x87\xa1\x74\x9b\x6f\xf7\x12\xf4\xfd\x43\x0f\xe1\x95\x59\x52\xa4\xdb\xc3\x61\xf2\xf4\xfe\x9e\x8d\x82\x9b\xf0\x7f\x40\x33\x4d\x8e\xab\xf9\xc7\x64\xef\xf2\x26\x85\x7b\x68\x87\x5c\x52\xf3\x34\x52\x52\x10\x20\x1a\x83\xda\x2a\xda\x66\x14\x07\x18\x23\x78\x0e\x8d\x7d\x65\x5f\x17\x85\x4d\xd2\x6c\xce\xdb\xb3\x85\x53\x17\xfb\x5b\x6b\xde\x6c\x93\xaf\x71\x13\x0b\x81\x64\xf4\xdf\xd7\x37\xa3\x31\xd4\x56\x8d\xc9\x9a\x03\xc2\x11\x03\x67\xf5\x65\x6f\xad\xd4\xbe\x4c\x42\x09\x00\xb9\xd1\x1a\x73\xcf\x2d\xa1\x69\xcf\x61\x3e\xbf\xa6\x06\x05\x45\x35\x83\x3f\xb9\x51\x68\xd1\x03\x89\xdc\x68\x2f\x75\x8d\xfc\xc4\xea\x1b\x04\xc8\x96\xbd\xce\x3d\x78\x36\xaf\x17\xa1\x01\xfb\x59\xfa\x55\x42\xc5\xcf\x9f\x3b\x93\xb3\x30\x72\xfb\xed\xff\x9c\x61\x3a\x41\xbc\xdd\xdd\x08\x04\xcc\xec\x3a\x44\xbf\xb8\x35\xcf\x72\xf6\x2b\xba\xb0\x14\x59\x9d\x5d\x02\x15\x8c\xd9\x4f\xba\x12\xd6\xad\x84\x4a\xc2\x70\x28\xc1\xdb\xec\xa5\xf0\x22\x49\xd3\x31\xfc\x19\xd3\xaf\x8e\x08\xd6\x67\x2b\xf2\xc5\x84\x33\xee\x37\x9f\x5c\xc2\x28\xee\x7d\x65\xaa\x4a\xe8\x62\x74\x12\x35\x7c\x4f\x26\x70\x63\x0a\xb1\x65\x0f\xb3\x08\x39\xe5\x93\x30\xe4\x0e\xf2\x41\x10\x9b\x22\x51\x1c\xc6\x92\x1b\xa2\x77\xed\x48\xa1\x89\xff\x04\x18\x52\x4f\x24\xdc\xb4\xca\xf0\xbd\xe6\xff\x65\x28\x50\x28\xd8\x18\xab\xd8\x9f\x19\x9d\xb7\xa3\x3c\xe6\xf8\x6a\xb0\xd9\x93\x5c\x7e\x2d\x9c\x03\x19\x32\x56\x3b\x09\xce\x5a\xe2\xaf\x88\x85\x8d\xb1\x76\x0b\x62\x41\xfd\x42\x11\x27\x14\xd2\x85\x08\xb1\xc1\xdd\x80\x85\x89\xd1\x72\xbe\x57\x63\xf2\xb0\xaf\xa5\xe9\x0d\x14\xd2\x89\x85\xc2\x38\x6c\xa9\x4c\x81\x1d\x1d\x73\x77\xfe\xe4\x12\x8a\xac\x3b\xb9\x78\x5c\xc1\x14\x26\x28\xa4\x34\x63\xdb\x5e\x9c\x79\x1f\x68\x7e\xd8\x81\x1e\x2c\xc1\x25\xe0\xee\xac\x48\xdd\xdd\xd0\x17\xaa\x10\x59\x82\xa4\x70\x8e\xd2\xaf\x42\x1c\x64\xe0\xe4\x3c\x65\xe0\x52\x5a\x17\xb3\x34\x6c\xf0\xe9\x2d\x82\xa3\x18\x17\xce\x6b\x47\xcb\x74\x30\x5f\xa4\xfd\x25\xea\xca\xc3\x5e\x2b\xd1\xc4\xd6\xa2\x87\xf1\x97\xb4\x9b\xa9\xb8\x3a\x21\xa5\x63\x41\x47\x9b\xf3\x20\x97\x47\x3b\x05\x38\x53\xa1\x5f\xd1\x69\xec\x71\xd0\x6a\xba\x19\x53\x0f\x8e\x84\xd9\x8e\x5e\x76\xb7\x7b\x0d\x6c\xd3\x66\x15\x99\x6a\xba\xab\xf4\x13\x9e\xd3\x9c\xcf\xd1\x06\x84\x0e\xfb\x78\x6b\xf2\x47\xf3\xf7\x00\xa8\x1c\x9f\x64\xe7\xce\x72\x5e\x2f\x92\x53\xf4\x53\xf8\x3b\xf4\xf2\xda\x1f\xcd\xd0\x2e\xf2\xd1\xc7\x03\xe1\x7c\x46\x15\xfd\x3b\xef\x39\x64\x09\x45\x16\x7c\xef\xbe\x37\x39\x9e\x1d\x8c\x8e\x9b\xda\xef\xa5\x49\x7a\xfe\x38\x6e\xaa\xa9\x10\xf9\x59\x2d\xc7\xb2\x44\x9c\xd3\x86\x49\x79\x10\x34\x0e\x82\xbb\x37\x89\x7f\xbf\x6c\xaf\x12\x07\xe1\xbf\x57\xb0\x88\x17\x15\x84\x2d\x74\xdc\x86\xc7\xb2\x92\x47\x97\x5d\xa7\x89\x95\x4e\x67\x33\x63\x79\x6c\x9e\x8c\x2e\xde\xbd\xe3\x9a\x32\x6c\xdd\xe3\xe6\xa1\xcd\xc8\x07\xc3\x64\x8e\x2b\x9f\x79\xf5\x32\x08\x23\xf7\x7f\x41\x51\x8f\x4f\xea\x4f\x5c\x38\xb4\x77\x0d\x9f\xd2\xdf\xbf\xad\xbe\x47\xb4\xf7\x30\xec\xf3\x11\x54\x07\x93\x67\xa0\x0d\x94\xb1\x15\x67\x41\xc3\x3f\x15\x6c\xd1\xc3\xb3\x49\x7a\x74\x4c\xff\xbf\x01\x00\x00\xff\xff\xcb\xed\x54\xed\x57\x28\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x11, 0x37, 0x2d, 0xa4, 0x79, 0x55, 0xcb, 0xc, 0x2a, 0x2f, 0x3b, 0x89, 0x97, 0xe, 0x3f, 0x72, 0x4b, 0x77, 0xd5, 0xcf, 0xa5, 0xe4, 0xc1, 0xd6, 0x5c, 0x73, 0x44, 0xec, 0xd, 0xb, 0x3, 0x99}}
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

var __hardcodedTracingGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\xdf\x6f\xdb\x36\x10\x7e\xb6\xfe\x8a\x83\x81\x15\xd2\xa0\x90\x68\x3b\x64\x45\x86\x3c\xac\xf5\x8a\xf8\x21\x43\xd0\x64\xe8\xc3\x30\x14\x34\x79\x96\x88\x48\xa4\x46\x9e\x6c\x19\x81\xff\xf7\xe1\x28\xf9\x47\x86\x26\xc0\xf6\x62\x8b\xc7\xbb\xef\x3e\xde\xf7\x91\x9d\xd2\x8f\xaa\x42\x88\x18\x36\x18\x28\x28\x6d\x5d\x95\x65\xb6\xed\x7c\x20\xc8\xb3\xd9\x1c\x9d\xf6\xc6\xba\x4a\xae\xac\x53\x61\x37\x3f\x0f\xd5\x38\xf0\x7a\xdd\x12\xff\x39\x24\x59\x13\x75\xf3\x2c\x9b\xcd\x2b\x4b\x75\xbf\x12\xda\xb7\xf2\x53\x83\x1b\x0c\xf2\x51\xed\x36\x88\x17\x95\x97\x9b\x9f\x65\xe3\xab\x0a\x03\x57\x55\x5e\xf8\x0e\x1d\x61\x83\x2d\x52\xd8\x09\xeb\xa5\xf6\x8e\x82\x5d\x49\xeb\x22\x85\xbe\x45\x47\x8a\xac\x77\xf2\x0c\xb4\xf2\xc1\x36\x8d\x92\x6d\x3f\x48\x4f\xd8\xb4\xfd\xf0\x12\x1a\x6f\xcb\x2e\xf8\x4e\x55\x09\xe6\xd5\x3c\x1e\x01\xce\xb3\x22\xcb\xa4\x84\x43\x91\x0f\x40\x1e\xfa\x88\x22\xdb\xa8\x70\x1e\x3e\x83\x15\x0f\x38\xd0\xad\xea\xee\x4e\xbb\xd7\xcf\xf7\x19\xfa\x93\x77\x84\x03\x3d\xed\x41\x4a\x48\xcd\x3a\x15\xd0\x11\xd4\xa8\x0c\x86\xd4\xf6\xb6\x1f\xee\x93\x1e\xb7\xd6\x98\x06\xb7\x2a\x20\x04\xa4\x3e\xb8\x08\xed\x29\x44\xb5\x22\x88\xb5\xef\x1b\x03\x2b\x04\x45\xa4\x74\x8d\x86\xa9\x2a\x38\x9b\xcf\x24\xae\x60\xe8\x25\x81\xf1\x18\x81\xb6\x1e\xa8\xb6\xae\x8a\x57\x10\x49\x05\x8a\x10\x3b\xe5\x62\x09\xca\x19\x50\xc6\x8c\xeb\x71\x1c\x60\xdd\xda\x33\x2c\xd5\x4c\xe4\xef\x1e\x23\x5d\xc4\x0e\xb5\x5d\x5b\x0d\xa3\x94\x09\xfd\x8b\xad\x6a\x02\xe7\xb7\xb0\x45\xf0\xae\xd9\x41\xec\xbb\xe4\x24\x4e\xb2\xae\x82\xe5\x22\x82\x75\x09\x68\xed\x43\xab\x68\x3c\xc5\x42\x91\x32\xbe\x02\x1c\x3a\xd4\x14\x45\xb6\xee\x9d\xfe\xde\x1c\x72\x3e\x8b\xd5\xf8\xbb\x6a\x11\x22\x05\xeb\xaa\x02\x38\x39\x67\xe7\x89\x1b\xe5\x4c\x83\xa1\x80\xf3\x15\x3c\x65\x33\x4f\xec\x10\xb8\xba\x86\xc9\x2c\xe2\xfb\x98\xe5\x71\xff\xab\xa5\xfa\x24\x65\xcc\x4f\xa2\x17\x45\x36\x1b\xe5\x98\x3a\xc3\xeb\xbd\x0f\xc9\x23\x87\x67\x44\x3f\x73\x7d\x02\x09\xdb\xb1\xee\x0b\xc6\xce\xbb\x88\x5f\x83\x25\x0c\x25\x04\xf8\x71\x8a\xa7\xb1\x17\x09\x70\x26\xe5\x81\x27\xd4\x2a\x02\x0e\xac\x13\xb1\xf6\x35\x26\xe1\x44\x52\xa1\xeb\x09\x2c\x81\x75\x93\x76\x9a\x86\x7f\xeb\xc6\x68\x91\xe7\x92\x94\x16\xf7\x9d\x72\x9f\x83\x6f\x27\x9b\xe6\x41\x1c\xbe\xf8\xd4\xb3\xd9\xba\x25\x71\x17\xac\xa3\xc6\xe5\x31\x45\xec\x1a\xa2\x66\x80\x98\x8a\x8f\xe9\xbf\x40\xd4\xe2\x46\xc5\xe4\xfa\xe5\x22\x9f\x98\xcf\x98\xdd\x72\x51\x8e\xfd\x96\x8b\x54\xa9\x53\x29\x27\x89\xfb\xa4\x69\x5e\x94\x1c\x3d\xd6\x1e\xc3\x09\x42\x4a\x30\x93\x61\xb4\x77\xfc\x6e\x45\xa8\x71\x98\xfc\x10\xd3\x5d\xb5\x8e\x2e\x7f\x62\xbb\x95\x10\x3d\x9f\x15\xa8\xf6\x11\x79\x91\x1c\xa7\x7d\x08\xd8\x28\x62\x53\x36\xbe\x8a\xc9\xf7\x89\x53\x84\xad\x0f\x8f\x31\x75\xb2\x6b\x68\xd0\xe5\x13\xd7\x02\xae\xaf\xe1\xfd\x3b\x78\xf3\x26\x45\xc7\x93\xa4\xe0\xdb\x4b\x78\xe2\x0b\xfd\xec\x4d\x99\x1c\x7e\xc5\xbb\xab\x1d\x21\xe4\xef\xdf\x5d\xe8\x5a\x05\x26\x5b\x94\xf0\x61\x8a\xbe\xbd\x3c\x45\x47\x0a\x89\x0c\xc3\x83\x35\x23\x91\xd9\x44\xe1\x63\x2c\xe1\x1b\xcf\xac\xc6\x41\x2c\x50\x7b\x83\xd3\x68\x0e\x1c\xc7\xf4\xe9\x52\xbe\x24\xa5\xf8\xd5\x98\xc3\x6a\x9e\x2a\xbf\x59\x33\x2f\xc7\xda\x24\xf2\x7d\xc7\x2a\xaf\xf3\xf9\x0f\x66\x5e\xc2\xf8\xf6\x8b\x8f\xb6\xfa\xcd\x19\xab\x9c\xf8\x23\xcd\x37\x3f\xb2\xfa\xf3\xc3\xd5\x5f\x45\x31\x35\x1f\x07\xf3\x32\xd5\x69\x70\xff\x9d\x29\x17\xfe\x3f\xa2\x07\x4a\x07\x8e\xfb\xec\xf0\x53\x8b\xf4\xc6\xdc\x3c\x3c\xdc\xe5\x61\x5b\x42\xe0\x84\x3d\xa7\xed\xb3\x7d\xf6\x4f\x00\x00\x00\xff\xff\x17\x0a\x14\xd8\x20\x07\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x4f, 0xa8, 0xd2, 0xb5, 0x8e, 0x1b, 0x3e, 0x4a, 0x4c, 0xb1, 0x93, 0x2a, 0x36, 0x47, 0x9c, 0x64, 0xef, 0xfa, 0xca, 0x3f, 0x4, 0xc7, 0xa5, 0xb0, 0x7e, 0x87, 0x2d, 0x5b, 0x28, 0x38, 0xa4, 0x60}}
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
