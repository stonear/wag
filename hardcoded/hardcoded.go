// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../_hardcoded/doer.go (10.35kB)
// ../_hardcoded/middleware.go (1.685kB)
// ../_hardcoded/tracing.go (7.082kB)

package hardcoded

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
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

var __hardcodedTracingGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x59\x5b\x6f\xdb\xb8\xf2\x7f\x96\x3e\xc5\xfc\x05\xec\x56\x2a\x1c\xa9\xd9\x4b\x77\x91\x45\x1e\x72\x6b\x6b\xfc\x93\x6c\x90\xa4\xe8\x39\x28\x16\x01\x23\x8d\x2d\xa2\x12\xa9\x43\x52\xbe\xa0\xc8\x77\x3f\x98\xa1\x24\xcb\x89\x93\x93\xdd\x3e\xb8\x36\x39\x33\x1c\xce\xf5\x37\x4c\x23\xf2\x6f\x62\x8e\x60\xd1\x2c\xd0\x38\x23\x72\xa9\xe6\x61\x28\xeb\x46\x1b\x07\x71\x18\x44\xb9\x56\x0e\x57\x2e\x0a\x83\x08\x55\xae\x0b\xa9\xe6\xd9\xbd\x54\xc2\xac\xb7\x96\x4a\x5c\xd1\xef\x59\xcd\x94\x0a\x5d\x56\x3a\xd7\xd0\x77\x6d\xe9\xd3\x3a\x93\x6b\xb5\xa0\xaf\x4e\xd6\x18\x85\x61\x10\xcd\xa5\x2b\xdb\xfb\x34\xd7\x75\x76\x52\xe1\x02\x4d\xf6\x4d\xac\x17\x88\x7b\x73\x9d\x2d\x7e\xcb\x2a\x3d\x9f\xa3\xf1\x94\x3a\xd5\x0d\x2a\x87\x15\xd6\xe8\xcc\x3a\x95\x3a\x23\xc5\x8c\xbc\xcf\xa4\xb2\xce\xb4\x35\x2a\x27\x9c\xd4\x2a\x1b\x49\x9d\x6b\x23\xab\x4a\x64\x75\xbb\xca\xb4\xc3\xaa\x6e\x59\xc9\x5d\xd2\x68\xfb\xa5\xbd\x0c\x57\x64\x12\x34\x36\xd3\xae\x6a\xf8\x83\xcc\x85\xff\x88\x69\xf3\x6d\x6e\x9a\xfc\x45\x11\x8d\xd1\x8d\x98\xf3\xd5\x5e\xa4\xb3\xc5\xb7\xcc\xa0\xd5\xad\x61\x9d\x6c\xf1\x8d\xe5\xc3\xcb\x1c\xff\xfb\x0a\x03\x95\xff\x74\x68\xc9\xc3\x16\x6b\x72\xe7\x4b\xd2\x3d\x45\xb6\xd8\x4f\xf7\xdf\xa5\xef\x5e\x3c\xa3\xd3\x22\x09\xc3\x85\x30\x50\xe0\x4c\xb4\x95\x3b\xd1\x55\x85\xb9\xd3\xe6\x93\xb6\x0e\xac\x33\x52\xcd\xe1\x10\xa2\x4a\xe7\xa2\x2a\x35\xa9\xb1\x8b\xfa\x8a\x02\xb7\x95\xca\xed\xbf\x87\x43\xf8\xe5\xe7\xfd\xdf\xc2\x30\xcb\xe0\x46\xc3\x12\xdf\x2c\x10\xe6\xda\x81\x6a\x6b\x34\xba\xb5\x10\x19\xfc\x4f\x8b\xd6\xed\xc9\x22\xb2\xe0\x34\x14\x28\x2a\x58\x4a\x57\xa6\xc4\xf4\xaf\xbd\xeb\x6e\x7f\x7a\x0a\xd2\x82\x45\x07\xf7\x6b\x40\xb5\xd0\xeb\x14\xbe\x20\x2c\x85\x72\xc4\xd6\x3b\x09\x41\x3a\x28\xa4\xc1\xdc\x55\x6b\x10\xaa\x80\xa5\xac\x2a\xa8\x85\x6a\x45\x45\x2b\x45\x41\x14\x4e\x43\xa5\xe7\x16\x84\x85\x25\x56\x15\x9f\xf5\x05\xa1\xd0\xea\x8d\x1b\x64\xe6\xa2\xaa\x88\x78\x5b\x09\x05\xae\x44\xcf\x7d\x8f\xb9\x68\x2d\x82\x76\x25\x1a\xce\x60\x99\xa3\x85\x52\x2c\x90\x88\xa4\x01\xbd\x54\x5b\xec\x74\x8e\x2b\x85\xa3\xed\x35\x09\x81\x99\x36\xb0\x10\x46\x92\x31\x0c\x0a\xab\x95\x65\x6d\x72\xce\xc7\xbd\x8d\x79\xe8\xfa\xba\x35\xe0\x4a\x72\xc3\x1c\x15\x1a\xe1\xb0\x20\x73\x90\x0b\x27\x1b\x13\xf0\x1a\x29\x49\xeb\xf0\x28\x3b\x21\x16\x16\xa4\x7b\x63\x59\xb7\x25\x4a\x53\x40\x89\xa2\x40\x63\x13\x3a\x97\x4c\xe6\x4a\x54\x1b\x93\x8d\xe4\x3e\x31\x1b\xf9\xbf\xe3\xbe\xd5\x57\x83\x0f\x0e\xe1\xeb\x5f\x3e\x5e\xbe\x47\xe3\xeb\x47\x13\x88\x9e\x5c\xcc\x2f\x4a\x54\x6e\x4f\x36\x7b\xa2\x28\x0c\x5a\x1b\x3d\x6c\xcb\x3e\xd7\xf3\x23\x0b\x87\x50\x8b\xe6\xab\x97\xbc\xfb\x80\x03\x88\x38\x38\x9e\xc8\x7f\x7c\xe8\xc1\xab\x35\x39\x18\x2d\x46\x0f\x3e\x96\xd1\xb5\xcd\xc7\x4a\xdf\x8b\xea\x96\x32\xe7\xca\xe8\x85\x2c\xd0\x1c\xa9\xe2\xac\x2b\x38\x14\xaa\x16\xda\x86\x1d\x31\x67\x52\xf0\xf5\xa0\xe9\x88\xd9\xd4\x7d\x7d\x4a\xc3\x59\xab\xf2\xd7\x08\x8e\x73\xb7\x82\xae\x27\xa4\x27\xfe\xff\x04\xe2\xbe\xdc\xa4\x37\x8d\x50\x3d\xed\x04\xde\x0e\xeb\x2c\xcf\xf4\x02\x27\x80\xc6\x68\x93\xc0\xf7\x30\x0c\xb2\x0c\xce\x16\x68\xd6\xb0\xff\x2b\x58\xcc\xb5\x2a\xc8\xbf\x6f\x2a\xd2\x78\xcd\x99\xa0\x95\xc2\x9c\x93\x62\xab\x7e\x40\xde\x27\x3d\x08\xc7\x72\xe8\xb6\x5d\x41\x00\xaa\x13\x1c\x72\x7a\x06\x43\xcd\x38\xe0\x82\x40\xa4\x5f\x28\xca\x4c\xab\x14\xc5\xb3\x54\x64\x97\xa2\xcd\x99\xc1\x95\xd2\x52\xbc\x0b\xb0\xb2\xc0\x5c\x98\x89\x4f\xe5\x11\x07\x8b\x60\xa9\xd5\x7a\x44\xdf\xaf\xf4\x72\x75\x83\x83\xb6\x7b\x83\xb6\x69\x18\x50\x74\xd9\x91\xa9\x60\xa7\x01\xc3\x80\xa2\x00\x0e\x0e\x61\x56\xbb\xf4\xa6\x31\x52\xb9\x59\x1c\xfd\x60\x0f\x7e\xa0\x68\xd9\x55\x28\x9f\xae\x52\x41\x4c\xc2\x00\x0d\x0b\x62\xbb\xc7\x4a\x56\x49\x18\xc8\x19\xc4\xda\xa6\x1f\xd1\xa1\x5a\xc4\xd1\xdd\xed\xf5\xd1\xc9\xf4\xf2\xe3\xdd\xd9\xe5\xd1\xf1\xf9\xd9\x69\x94\x24\x70\x78\x08\x91\x33\x2d\x46\xec\xa9\x80\xda\xd6\x09\x47\x23\xc9\xda\x6a\x62\xe9\x25\x2e\xfd\x56\x1c\x06\x9e\x72\xb3\xf7\x45\xba\xf2\x1a\x3b\x3f\x4a\xad\xae\xd0\x48\x5d\xc4\xfb\xbf\xbe\x25\x30\x90\xde\xb0\xd7\x93\xc9\x6e\xc6\x33\x55\x34\x5a\x2a\x17\x93\x31\x9e\x23\x9a\x2a\x8b\x79\x6b\x30\x66\x82\x24\x0c\x02\xbb\x15\x89\x74\xfd\x91\xc6\xa4\x2d\x85\xf2\x04\x36\x57\x22\x26\x39\x63\xca\xff\x3b\x04\x25\x2b\xf8\x4e\x67\x19\x74\xad\x51\xf4\x7b\xe2\x3f\xc8\x17\x67\x64\xc6\x59\x1c\xb1\x39\x21\x37\x28\x1c\xf9\xbb\xcf\xa7\x03\xf8\x61\x11\xf1\xa1\x24\xf4\x21\x0c\x1e\x00\x2b\x8b\x2c\x70\xcb\xed\x87\x30\x74\x55\x52\xe9\x52\xeb\x66\xc8\xb4\x24\x0c\x1e\xc2\x30\x70\x0d\xd9\x5a\xe1\x72\x3b\x85\xe2\xed\xfb\x29\x5c\x5e\x77\xdd\x3f\x4e\x92\x30\xa0\xe2\x9b\xde\xa0\x7b\xc4\xe4\x9a\x24\x0c\x03\x8f\xac\xd2\x0f\x46\xd7\x5d\x0a\x93\x2d\x92\x74\xaa\x66\xfa\x34\x8e\xac\x13\x86\xae\xb3\xc7\xba\x99\x68\x02\x1d\xc3\x05\xe9\x1f\x6d\x2a\x13\x7d\x9b\x84\xc1\x43\x12\xf6\x56\xda\x56\xca\x35\x6c\xb2\xf0\x21\xf4\xf5\xe5\xe9\x25\xf0\xc5\xf8\x9f\x40\x8f\x68\xe0\x6d\xff\x2d\xed\xaf\x99\x3c\x5b\x5d\xc8\xcc\x56\xd4\x4d\x25\xd5\xfc\xca\xe8\x7b\x71\x2f\x2b\xe9\xd6\x64\xc6\x77\xe9\xbb\x5f\xc3\x40\xda\x73\x4a\x54\x8e\xe1\x51\xfc\x4f\x6f\xee\xce\xff\x3c\x39\x3a\x8f\x46\x71\xcf\x59\xd2\xd3\xb3\xfb\x76\x08\x3e\x84\xfd\xf4\xdd\xe0\x64\x39\x83\xc5\x23\xd1\x7d\x66\xdd\x1c\x5d\x5c\x9d\xd3\x97\xab\xeb\x3f\x8f\x8f\x8e\xa7\xe7\xd3\xdb\x7f\x47\xc9\x1f\xb0\xa0\x88\x8b\xa2\xe7\x0e\x20\x47\x9d\xa9\x85\x8f\xe2\x83\x43\xe8\xc0\x74\x7a\x25\x8c\xc5\x0f\x95\x16\x2e\x5e\x4c\xe0\xfd\x2f\xcf\x44\xf0\xf3\x12\x49\x73\x1f\xa0\xcf\xdc\xeb\x79\xd6\x71\x70\x0e\x7e\xb8\x7c\xe2\xe0\x30\x08\x3c\xb8\x21\xa4\x32\x2e\xcf\xd3\xd3\x1e\x45\x68\x93\xc2\x54\x81\x36\xe4\x3a\x02\x24\xfd\xa1\x54\xf0\x97\xda\x7c\x83\x58\x38\xa8\x50\x58\xc7\xd0\xcc\xd7\x5b\x26\x42\xce\xaf\xae\xf0\x8f\x25\x42\xdd\x5a\x37\xc0\x94\xae\xf5\x4d\x4f\x2d\xb4\x4a\xce\xb4\xa9\x09\x8b\x39\x30\x42\x15\xba\x86\x99\xd1\x35\x4b\x40\xe5\xa4\x41\x0a\xe2\x1c\xa9\x6d\x10\x8c\x7c\xff\x4b\xea\x8f\xf8\xa0\x0d\xe0\x8a\x4f\x9d\x6c\xdd\x64\xb5\x67\xc4\x7a\xfb\xf4\x42\xa3\x05\xa5\x1d\x14\x9a\xb5\xed\x44\x0c\x76\xa2\x8a\x75\xe3\x2f\x10\x6f\x07\xf1\xf4\xf4\x9a\xba\xd6\xb1\xb0\x58\xc4\x09\x17\xb3\x97\xb9\xae\x84\x41\xe5\x3c\xfd\x0b\xa2\x76\x38\x32\xf1\xe2\xb3\x0c\x6e\x4b\xb4\x08\xb5\x58\xc9\xba\xad\x2d\x08\x83\x64\xfa\x79\x2b\x4c\x01\x62\x2e\x08\xc3\x81\xd5\x35\x76\xe0\x4f\xd3\xe7\xd2\x68\x35\xe7\xbe\x68\x51\xd1\x08\x08\x02\x9c\xef\xb6\x85\x70\x02\x5a\x85\xab\x06\x73\x87\x45\xb5\x7e\x72\x89\x46\xa8\x73\x59\x4b\x67\xb7\x81\x83\x5f\xe3\x98\x3d\x72\x34\xdc\xb5\x0e\x4f\x74\xab\x1c\x6f\x1c\xc0\xfe\xbb\x77\x5c\xfe\xcf\x16\xa8\xdc\x78\x83\xfe\xf5\x9b\xe7\x52\x7d\x7b\xbc\xd7\x6f\x3e\x24\x93\xf0\xb1\x32\xc7\xc2\xe5\xe5\xa8\x14\x3d\xb5\xf9\x50\x58\xfb\x1a\x44\x24\x2f\xd7\xd8\xcd\x1e\xae\xdc\x85\x68\x7a\x70\xaa\x4d\x3c\x1a\xe8\xb8\x67\xea\xba\xd1\x56\x3a\x7c\x99\x92\xcf\xe8\x8a\xf5\xf7\x87\x0d\xdc\xa6\xbd\x63\x31\x9f\x8b\x39\x7e\x7f\x48\xa8\xb8\x77\x95\xd8\x35\xe1\xab\xd1\xe2\x07\x6d\x6e\xd1\x3a\xc2\x30\x35\x76\xe3\xc7\x3d\xa7\x6c\x41\xd0\xa8\x55\x34\xb3\xa0\xa5\x9e\x30\xe9\x51\x7a\x2d\xa9\xf5\x59\xce\x04\x4b\x47\x80\xb8\xd7\x0b\xe4\x04\xd6\xad\x23\xe4\xc4\x29\x34\xe2\x4d\x61\xea\xc0\xab\x67\x41\x28\x90\x6a\xaf\xc6\x5a\x9b\x35\xc9\x1c\x3a\xc1\xac\xcb\x33\xa9\xb6\x27\x0d\x6a\x2d\xf6\xf5\x40\xb5\xbb\x53\x9c\x40\xfc\x76\xd3\x63\xa7\xea\x82\x4f\xfc\x5b\x10\x35\x18\x74\x3b\x78\xd4\xaf\x1f\x8b\xa3\x9e\xfd\x5c\xbf\xc6\x7f\xd2\xab\x7b\x6f\xe2\xae\x9e\x9a\x65\x70\xd1\xae\x6e\xf8\xe9\xe6\x42\x16\x45\x85\x4b\x4a\xdd\xde\xc2\xf5\x66\x89\x87\x3e\x5b\xea\xb6\x2a\xc8\xb1\xc2\x39\x91\x97\x7e\xa2\x12\x30\x7a\x25\xe9\xde\x81\x78\x04\x9c\x3a\x5f\xc6\xdc\x52\xfb\xa9\xcf\x1e\x00\x23\x03\xeb\x5d\xe1\x41\xb1\x28\x0a\xff\xdb\x0f\xf1\x20\xd5\x4c\x93\x58\x0a\x8b\x7e\xb4\xb1\x0d\xe6\x72\x26\xf3\x1e\x46\x90\xf4\x6b\x39\x2f\x1d\x28\xbd\x84\x25\x82\x56\xd5\x1a\x6c\xdb\xf0\xa3\x13\x11\x91\xe7\xa9\x5a\x77\xe3\x2e\x55\x6c\x9e\x5a\x85\x83\x53\xe1\x44\xa1\x19\x6a\x61\xee\xfa\x80\xd8\x61\x87\xb8\x9b\x88\x2f\x45\x8d\xdd\xf3\x41\x02\x44\x1c\x97\xce\x35\xe9\x27\xa1\x0a\xea\x1f\x30\xfe\x45\xce\xd6\xae\x22\x4b\x30\xbe\xe5\x27\xa3\x74\xb7\xcc\xc9\xb0\x4f\x35\x62\x93\xb6\x36\x66\x7f\x7e\xdc\x95\xfb\x09\xb9\x3b\xcb\x18\x42\x5e\x11\x9a\xaf\x54\x1c\x1d\x15\x5c\x3e\x37\xe6\x1f\x79\x2e\xda\x84\x80\x57\x1d\x5e\x56\xbe\x27\xf6\x97\xd8\xba\xe9\x07\xe2\x67\x21\x66\xe9\xf9\xae\xd1\x36\x5a\x59\xfc\x62\xa4\x07\x5b\xf0\xb6\x5b\x67\xbf\x25\x1e\x3d\xc8\x19\x98\x7e\xe9\xf3\xf5\x94\xa1\x51\x76\x57\xa2\xa8\x5c\xe9\x11\x4b\x10\x94\x29\x5b\xff\xd3\xed\xed\x55\x6c\x96\x13\xe0\xb6\xdc\x2b\x43\x5f\x09\x2b\x04\x3c\xf5\xe4\x46\x16\x9d\x3b\x78\xcd\x0e\x49\xc5\x2d\x60\x8c\x49\x4d\x3f\x60\x72\x96\x04\x01\xb3\x1e\x82\x65\xc2\x61\xab\x6f\x76\x71\x92\xde\xb0\xd8\x38\x61\xc1\xd4\x6f\xd1\xed\x78\xc9\xf0\x53\x3d\x0f\x96\x46\x16\x0c\xf1\xd3\x4f\xbc\x46\x49\x18\xef\x1c\xcc\x89\xd2\xeb\xe0\x38\x34\x46\xaa\xd1\x2a\xd5\xac\xbb\x49\x2f\x99\xf6\x85\x9a\xe3\xe3\xf7\x03\x6f\x2b\xb6\x67\x77\xde\x47\x74\xb1\x27\x4a\x46\x08\x30\x08\x9e\x45\xe8\x47\x45\xd1\xff\xde\x16\xfe\xd5\xff\xfc\x6b\xb2\x53\xb8\x77\xc7\xc3\xc6\x13\x59\x06\x47\x05\x3f\xb5\xf4\x4a\x6e\x92\xd6\x07\xc5\x6b\xae\xb5\x79\x72\xf9\x1b\x57\x33\xcb\x8e\x84\x1c\x36\x10\xbd\x5e\xef\x33\x95\xeb\xc2\x63\x49\xff\x2e\x49\xd5\x07\x64\x31\x54\x8b\xd3\xd3\xae\x60\x74\xd1\x6b\x73\xc6\xa8\xdb\x51\xf3\x07\xd8\x3c\xfd\x24\xec\x10\x3c\x9d\x7a\x24\x6c\x7a\x3a\xf1\xa2\xa7\xa7\xcc\x99\x33\xeb\x56\x84\x4d\x68\x75\x47\xe0\x05\x5e\xc7\xa2\x2b\x52\x84\xd2\x91\x2a\x66\x89\xab\x2e\xe8\xd9\xd0\x1e\x58\x52\x89\x9b\x80\xe5\x57\x2d\x70\xa5\xb6\x48\x3f\xb8\xca\xe5\xda\x18\xac\xfc\x50\xe9\xdf\xbc\x54\xe1\x75\xb2\x0c\x89\x6d\x6f\xef\x0a\x55\xdc\xe9\xca\x53\xcb\xcf\x3f\xc1\x8f\x3f\xf2\xaa\xbf\x09\x2f\xee\xbf\x87\xef\x90\x65\x8f\x5e\x4f\xbc\x91\x0e\x68\xf7\x7e\xed\x10\xe2\x9f\x7f\xda\xcb\x4b\x7e\xf6\x5a\x25\x13\xf8\xbd\x5b\xdd\x7f\xbf\x59\xed\x80\x34\xa3\xbe\xce\xea\xa1\xf7\x6a\xa7\xc3\xb1\x9d\xc0\x1d\x19\xad\xc4\x55\x7a\x8a\xe4\xa9\xce\x36\xbd\x92\xaf\x8e\xef\xa8\x28\x52\x66\xba\xa3\x0c\xf4\x6c\xc1\xf6\x3b\x08\x65\xa6\xff\x5b\x44\x7a\x2c\xe7\x67\xaa\x90\x42\xa5\x9f\xd9\xb6\xf1\xa0\xd0\xd7\xdf\x0f\xfe\x4a\xba\x48\xea\xdc\xfb\xbc\x96\x9d\xd1\xfe\x96\x92\xc4\xf3\xcf\x74\xec\xb5\x49\x9e\x06\xba\x01\x2a\x32\xd4\x59\xc6\x47\x87\xcf\xd5\x5a\x82\x7e\xc1\x43\x07\x0a\x46\xf8\x62\x03\xb7\x36\xa3\x74\x81\x36\x37\xf2\x9e\x07\x2c\x9a\xa4\x44\xd3\x54\xd2\xbf\x95\x71\x6b\xfe\x4c\xb0\x8f\xe7\x30\x74\x1c\x80\x6d\xe3\xfd\x6e\x86\xc7\xc3\x61\x9c\xdf\xe0\x98\x1d\x13\x3a\xa5\x14\x95\x7c\xd1\x34\xa3\x26\xcc\x13\xf5\x78\xec\x3e\xba\xba\xba\xbb\x3c\xba\x38\x8b\xc6\x85\xa2\xe7\x39\x7c\x86\x74\x3c\x6d\x8f\x28\x5e\x2f\x6b\x4b\x54\x18\x98\x2e\x22\x86\x4b\x5c\xa0\x99\x63\xcc\x1d\xb5\x5b\x39\xf5\xf3\x9e\x7f\x61\x1a\x56\x2f\x71\x49\x5e\x1a\xe6\x15\xcb\xcf\x5f\xdd\x5f\x3f\xd2\x9b\xbc\xc4\x5a\x7c\xbe\x3e\x9f\x6c\xad\x6e\x50\xc4\xff\xe3\xba\xaf\x1e\x9d\x9a\xfe\x01\xcb\x4f\x18\x5d\x37\x37\xe1\x43\xf8\xdf\x00\x00\x00\xff\xff\x63\x5e\x0f\xdf\xaa\x1b\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb7, 0x1b, 0x4c, 0x9b, 0x44, 0x92, 0xc4, 0xe, 0xd7, 0x6c, 0x9e, 0xef, 0xfb, 0x6, 0xbe, 0xed, 0x3d, 0xc, 0x84, 0xd3, 0x79, 0x5f, 0x99, 0x92, 0x47, 0xff, 0x9, 0x4, 0xa0, 0x3e, 0xa, 0x7c}}
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
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	err = os.WriteFile(_filePath(dir, name), data, info.Mode())
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
