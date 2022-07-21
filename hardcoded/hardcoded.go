// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../_hardcoded/doer.go (10.391kB)
// ../_hardcoded/middleware.go (1.683kB)
// ../_hardcoded/tracing.go (7.417kB)

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

var __hardcodedDoerGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x3a\x7f\x6f\xdb\x46\xb2\x7f\x4b\x9f\x62\x22\xe0\x2e\x64\x22\x53\x72\xee\xdc\x1e\xd4\xe7\x03\x5a\xc7\x7d\x2d\x90\x26\x41\xe5\xbe\x06\x08\x82\xeb\x8a\x1c\x4a\xdb\x2c\x77\xd5\xdd\xa5\x65\xc1\xe7\xef\xfe\x30\xb3\x4b\x8a\xd4\x0f\x27\x77\x2d\x1e\xde\xf9\x0f\x89\xe2\xce\xcc\xce\xcc\xce\xef\xf5\x5a\xe4\x1f\xc5\x12\x21\x57\x12\xb5\x1f\x0e\x65\xb5\x36\xd6\x43\x32\x1c\x8c\x16\x5b\x8f\x6e\x34\x1c\x8c\x72\xa3\x3d\xde\x79\x7a\x44\x9d\x9b\x42\xea\xe5\xe4\x57\x67\x34\xbf\xb0\xd6\x58\x86\x92\x66\x22\x4d\xed\xa5\xa2\x1f\xca\x2c\xe9\xab\x12\x7e\x35\xb1\x42\x17\xf4\x43\xa3\x8f\x5f\x93\x95\xf7\x6b\x7a\x76\x5b\x9d\xd3\xb7\x97\x15\x8e\x86\xc3\x81\x32\xcb\x25\x5a\x18\x2d\xa5\x5f\xd5\x8b\x2c\x37\xd5\xe4\x4a\xe1\x2d\xda\xc9\x46\x2c\x27\x61\xd5\xd1\x73\x78\x24\xd4\x0e\xa8\x28\xf1\x6e\xb2\xda\x3a\x6f\xe5\xdd\xd9\xd2\x34\x8f\x7b\x50\x85\xd1\xe6\x56\xe8\x95\x2c\x70\x82\xb7\xa8\xbd\x33\xb5\xcd\x71\x34\x4c\x87\xc3\xc9\x04\x0a\x83\x16\xa4\x03\xa1\x41\x6a\x8f\xb6\x14\x39\x42\x69\x2c\x8c\x0a\x23\xf5\x72\x04\xc4\x3b\x58\xfc\xad\x46\xe7\x1d\xac\x8d\x73\x72\xa1\xb6\xb0\x91\x7e\x05\x1b\x2b\xd6\x6b\xa9\x97\x43\xbf\x5d\x63\x24\xd5\x12\xb9\x1f\x0e\x5e\x9a\x24\x87\x67\x44\x21\xbb\x62\x85\x8f\xc1\xc6\xdf\x3f\x06\x8a\x29\x24\xcd\x6f\xb7\x36\xda\xe1\x18\x58\xc5\xe9\xf0\x61\x18\xa8\x9a\xf5\x6b\x51\xe1\x95\xbf\x03\xe7\x6d\x9d\xfb\xfb\x07\xe6\x7b\x21\x1c\x46\x1a\xdf\x09\x5d\x28\xb4\xb0\x46\x5b\x1a\x5b\x39\xf0\x2b\xe4\xf5\x1e\xef\x81\x1a\xbd\x7e\x49\x7c\xee\x88\x95\xb5\xce\x21\x29\xda\xa5\x14\x7e\x07\xdb\x24\xb5\x45\x5f\x5b\x0d\x79\xf6\xd2\x24\x36\x1d\x0e\x26\x13\x68\x5e\xf9\x3b\xc6\xa2\x85\xec\x2a\x98\x59\x92\x8e\x21\x1f\x43\x90\x38\x80\xda\x6d\x23\x12\xfd\x90\xe8\xe0\x62\xfa\xae\x7f\x10\x41\x1a\x86\xed\x88\x43\xbb\x17\xb0\xfb\xa3\x13\x61\x7e\xec\xf6\xad\x51\x32\xdf\xc2\x8f\xbb\xe7\xb8\x5f\xe7\x0d\x14\x58\x4a\x8d\x0e\x44\xa0\x0c\x6b\x7e\x9d\x85\xcd\xba\x80\xbd\x53\x9e\x4c\xe0\x1b\x91\x7f\x34\x65\xe9\xa2\xa0\xe1\x08\x74\x5d\x2d\xd0\x82\xd0\x05\x78\x59\x49\xbd\x04\x53\x46\xc2\xc2\x7b\xac\xd6\xde\x65\xc3\x41\x83\x9a\xa4\xf0\xfe\x03\x39\x46\xf6\xb2\xb6\xc2\x4b\xa3\x99\x32\xef\x0a\x16\x73\x94\xb7\x18\xe8\x76\x15\x31\x06\xe1\x60\x83\x4a\xd1\x37\x2d\x5a\x74\xb5\xf2\x60\x4a\xc6\x6e\x9c\x2f\x9e\xe3\x53\x07\xbf\xbc\x34\xbf\x40\x85\x7e\x65\x8a\x6c\x38\x60\xea\x49\xef\x68\xc7\x70\xe2\x64\x17\xc6\xa8\xa8\xb3\xb9\xd4\x4b\x85\x8f\x69\xce\xaf\x84\x6f\x4f\x4f\x34\xdc\x82\xd1\x39\x06\x6d\x1e\x92\xe8\x99\xf7\x11\x85\x0a\x0f\x5b\x53\x83\x5b\x99\x5a\x15\xed\x36\xd8\xd2\x3e\x77\x98\x1b\x5d\x80\x28\x3d\xf9\xa1\x87\x52\x48\xe5\xb2\x68\xdf\x07\xfb\xa5\x70\x5a\xf3\x1d\x2b\xde\x5b\xb9\x3f\x87\x67\xc0\x6f\xe6\xbc\xdb\x43\xd7\x8a\x60\x23\x95\x8a\x9c\x69\xa3\xcf\xde\xbe\x99\xdf\x8c\xc3\xd3\xd7\x37\x57\xdf\xed\xa2\x08\x0b\x73\xf1\xee\x5d\x46\xa8\x37\x6f\x5e\xbe\x99\xc1\xf7\x9e\xac\xd5\x81\x36\x1e\xf2\xda\x5a\xd4\x5e\x6d\x1b\x6b\xd1\xdb\x70\x08\x8d\x3a\xb0\x80\xc5\xf6\xc4\xe1\x3e\x22\x70\x38\x6d\x8b\xbf\xc1\xde\x89\x5b\x74\xeb\x63\xc7\xde\x3d\x7a\xd2\x89\x2c\xf9\xed\x93\x4b\xd0\x52\xc1\x3f\xff\x49\x12\x65\x3f\xb0\x31\xc1\xe5\x25\x8c\x48\xe0\xd1\xb1\xf7\x24\x3e\x2d\x0c\x07\x03\xda\x2b\x9b\x7b\xe1\x6b\x77\x65\x0a\x84\xff\x82\x8b\xe9\x94\x88\x37\x1a\x2f\x85\x72\x38\x1c\x3c\xb4\x47\xe0\x6d\x8d\x51\xcd\xd7\x77\x6b\xa3\x51\x7b\x29\xd4\x51\xeb\xd3\x80\x3b\x88\x9e\x13\x07\xab\x3b\x81\xff\xb8\xe9\x95\xf2\x96\x02\x67\x7c\xcb\x81\xbf\xb3\x8b\xa2\x68\x90\x5b\x14\x8e\x3c\x7c\x23\xa4\x67\xfb\x70\x1c\xa4\xd1\x6f\x10\x75\x7b\xf0\x33\x38\x9f\x4e\xc7\xf0\x82\x3e\xfe\x4a\x1f\x7f\xa3\x0f\x0a\x10\xe7\x5f\x4c\xa7\x50\x49\xa5\x64\xb0\x62\x07\xcf\x27\x67\x50\xaf\xc1\x1b\xb8\xf8\x13\xfc\x2a\xbd\x47\xdb\x1c\xed\x71\x29\x3e\xc3\xa0\x61\x76\x09\x95\xf8\x88\xc9\xde\xf2\x18\x2e\xd2\xe1\x40\xe3\x1d\x43\x9c\x4f\xa7\x8d\x91\xff\xb0\xe3\x68\x38\xb0\xba\xa0\x65\xca\xef\xd9\x6b\xdc\x24\xcd\xc3\x9c\x13\x6a\xc2\x08\xaf\xcd\x26\x49\xb3\x9f\xb4\xbc\x7b\x2d\xb4\x49\xd2\x34\x1d\x0e\x90\xb0\xa6\xd9\xf4\x02\x26\x13\x16\xeb\x82\x52\x55\x8e\xda\x47\xb9\x86\x03\x4a\xb8\x32\x12\x5f\x72\x5c\x6f\x2c\xe2\xbd\xfc\x00\x97\xc0\xac\x3d\x87\x1e\xd3\x49\x92\x58\x5d\x64\xdf\x2a\x23\xfc\x17\x7f\x4d\xd2\x67\x2f\xd2\xb3\xf3\xf4\x19\x3e\x2b\xe3\x1b\x42\xa2\xfd\x83\x60\xcf\x2e\xe1\x45\xd7\xaa\x2c\xfa\xff\xb7\xbe\x7b\xea\x80\xff\xf3\x1d\xf8\xb5\x39\x9e\x35\x82\x93\x52\x82\xd4\x54\xff\x05\xe5\x91\x43\xb5\xb9\x23\x66\xe1\x3e\x81\xc7\x1d\x97\xc2\x41\xb5\xf6\x5b\x70\x4a\xe6\xd8\x28\xb7\x47\xe1\xdf\xcc\x02\xfd\xb0\x2f\xd4\x46\x6c\x3b\xf1\x82\xe4\x3f\xb1\xdb\xbf\x9a\x70\x3b\x3c\x04\xb5\x36\x65\x21\x2b\x28\x96\x50\x7d\x35\xfc\x2c\xfd\xaa\xab\xa3\x56\x1b\xa0\x71\x03\xb1\xb8\x0f\x66\x6c\x6e\xd1\x5a\x59\xc4\xba\x22\xf4\x04\x60\x16\xbf\x62\xee\x9f\xba\xb6\x18\x6b\xcb\x20\x96\x68\x8f\x7c\x92\xfb\xbb\x86\x68\x53\xd2\x8d\xe1\x44\xc9\x95\xee\x43\x76\xcb\xc5\xb8\x42\xf4\xff\x47\xa8\x1a\x89\xf2\xb8\x27\xe7\xfd\x43\x8f\x32\x57\x8c\x4d\xf5\xfa\xac\xad\x05\xff\x88\xfa\xb5\xd9\x62\x0c\xe6\x23\x87\xa5\x5d\xb5\x9a\x05\xe6\xfa\x7c\xa5\x59\xd2\x15\x93\x1d\xed\x89\xf9\xd8\x78\x45\xab\x8b\x4b\x28\xb2\xce\x6f\x76\x92\x36\xb3\xd0\x3e\xbb\xb5\x6c\x67\x98\xc3\xc1\xad\xb0\xc7\x7c\x3c\x2c\xb4\x7e\x3e\xe4\x9a\x6f\x2e\x6e\xb1\x57\x18\x2d\x4c\x41\x39\x0a\x72\x6a\x09\x36\x08\x2b\x06\x30\x61\xb3\x0c\xde\xf8\x15\xda\x8d\x0c\x6b\x1c\x05\x19\x40\x28\x8b\xa2\x20\xf3\x11\x05\xd3\xe5\xae\xa2\x2e\x4b\xb4\x60\x74\x1b\xdd\x8a\xde\x56\x8c\x4e\x75\x57\x06\x73\x44\x46\x23\x86\x67\x93\x89\xf3\x24\xce\x2d\xda\x52\x99\x0d\x37\x66\x8c\x21\x8d\x76\x93\x17\x7f\x99\x7e\x39\xfd\xdb\x97\x5f\x4c\x68\x2f\xa9\x97\x67\xc4\xf1\x99\x29\xcf\x08\xf7\x2c\xd2\x3e\xa3\xc4\x6b\x6a\x7f\x56\x99\x42\x96\x14\x1b\xda\x15\xe7\x85\x8f\xba\x58\xd4\x25\xbc\xff\x40\xad\x2c\x9f\x81\xcd\xbe\x21\xe1\x63\xbc\xa3\xd3\xe8\x2b\x6c\x30\x58\xd4\x65\x88\x94\x97\x10\xda\xd9\xec\x47\x14\xc5\xd7\x4a\x25\x01\x97\xf2\x47\x3f\x6a\x12\x95\xc6\x68\xb5\x54\x8c\x3d\x1c\xd0\x49\x3e\x0c\x43\x26\x6b\xea\x5e\x4a\x7b\x5f\x71\xf4\xfb\xaa\x79\xf7\xfc\x39\xe3\x1f\x65\x6d\x60\x0b\x4b\x38\x91\x8f\xd7\x66\x7d\xa5\x8c\x43\x9b\x70\x67\x4e\x99\xf6\x1b\x56\x7f\xb2\xa8\x4b\xce\x6b\x83\x48\xe3\x12\x6c\x11\x59\xe0\xe8\xdc\x08\x54\x64\x05\xb5\x5a\xa1\xbb\x0a\x9b\x46\xce\x2e\x2f\x41\xa1\x4e\x1a\xd3\x4b\x29\xe4\x3f\xe9\x1a\x5f\xcc\x34\x21\xb3\x30\xbd\x34\xf0\xb8\xb0\x28\x3e\xc6\xbd\x26\x13\x60\x0e\x9b\xc6\x83\x6d\x32\x9a\x5b\xb9\x7b\x21\x43\x66\xd4\x52\x35\x4c\xb8\xf5\x9e\x32\xdd\x9a\x25\xc9\x98\x1c\x19\x3c\xd3\x0f\x75\xb6\x42\x5c\xb7\x9c\xbe\x8f\x12\x7c\x48\xfb\xb9\x3c\x32\x19\xc3\x72\x2e\x6d\x5e\x4b\xff\x0d\xf1\x8a\x96\x1b\x44\x59\xad\x15\x56\xd4\xfe\x87\x58\x17\x20\x60\x11\x40\x60\x4d\x1d\x99\xd5\x9c\xd4\xe7\xe8\xa1\xc0\x45\xbd\x24\x27\xa1\xd3\xa3\x6f\xb3\x46\x2b\x3c\xf6\x90\xa5\x06\x01\x15\xe5\xc0\xcd\x0a\x2d\x52\xcf\xc1\x0e\x60\xb4\xda\x82\x32\xcb\x0e\x2b\xc0\x36\x1a\x6a\x3d\x86\x21\x8d\x2c\x94\xc9\x3f\x72\x85\xd0\x56\x17\xa5\x35\x15\x2c\x0d\x65\x3e\xbf\xb2\xa6\x5e\xae\x62\xde\x3b\x22\xd2\x63\x3d\x6f\xe0\x3f\xfc\x71\xe3\x36\x88\x04\x5e\x8b\x0a\x09\x53\xea\x65\x3b\x7f\xe1\xbf\xf0\x9c\xfd\x2c\x96\x21\x66\xbe\xe2\xdf\xa4\x50\x72\x99\x88\x3d\x9f\x5f\xbf\xd1\x39\x82\xdb\xea\x3c\xa3\x27\x56\xf7\x77\x61\xf0\x32\x9f\x5f\x5f\xdf\x52\x06\x91\x0e\xb0\xa2\xca\x8e\x4b\x9d\xdd\x84\x06\x6e\xa5\x00\x87\xf6\x16\xed\x99\x23\xc0\x30\x8f\xc9\xb8\x80\x42\x97\x5b\xb9\x08\x75\x33\x29\x99\x15\x46\x05\x81\x68\x36\x8f\x9a\xd8\xdf\x6d\xa7\x86\x1b\x5a\x7e\xe4\x2f\x88\x0d\xbf\xfc\xea\x8c\x9e\x8d\x88\xd8\xe8\x97\xe1\x80\x35\xf2\xd9\x58\x5a\x54\x8c\x15\x13\xc9\x95\xa9\xb5\x3f\x86\x25\xc3\xeb\x88\x65\x3b\xd0\x84\x7d\x4d\xd1\xe7\x24\xee\x3e\x36\xb6\xd0\x2d\xee\xdb\x50\x40\x8b\xe5\x21\xeb\x47\x70\x77\xd0\x44\xe0\x7b\x77\xd5\x33\xa6\x37\x6b\xd4\x5d\x02\x5c\x77\xb4\x04\xe4\x11\x68\x56\x80\x51\x4a\xea\x25\xb3\xf5\xad\x90\xaa\xb6\xf8\x98\x02\x0e\xa1\x0f\x89\x28\x45\x6e\xde\x25\xf6\x18\x91\x1e\xf4\x29\x62\xf3\x3a\xcf\xd1\xb9\xcf\x24\x16\xa1\xf7\x89\xcd\x57\xc6\xfa\xa8\x05\x2c\x3e\x25\x5e\x1f\xfa\x80\x56\x87\xa1\x4f\xab\xea\x04\x43\x37\x2b\xca\x94\x6f\x8d\x51\x3f\x22\x15\x6b\xc4\xd4\x69\x22\x87\xd0\x07\xf4\x64\x85\xa6\xf6\x9f\xc9\x54\x84\x26\x22\x57\xa1\xdb\xb9\x32\x3a\xf6\x3d\xd7\x77\x98\xd7\x94\xd5\x83\x71\xf7\x89\xe4\x8f\x43\x13\xc1\x57\xc2\xa3\xce\xb7\x37\xc6\x0b\xf5\x03\x0a\xfd\xb8\x71\xab\x0e\xf4\x3f\x2a\x14\x64\x98\x4d\x4d\xa8\xcc\x92\xe3\x43\xa2\x4e\x84\xb6\x31\x1c\xc4\x12\xce\x71\x2a\x7b\x65\x96\xc9\x48\xea\xd2\x8c\xc6\x30\x1a\x8d\xa1\x12\xeb\xf7\x21\x0a\x7c\x68\x07\x7f\xf7\x0f\x94\xba\xfa\x8e\x3d\x3b\xe6\xca\x98\x75\x43\xc5\x98\x90\x3a\xfe\x7c\x14\x85\x90\x76\x11\x62\x87\xd2\x71\xe3\x43\xbc\x88\xb2\x83\x61\xbc\xa3\xde\x3b\xeb\xe3\x1d\x8b\x07\x8c\x7c\xcc\x69\x67\xfb\xc2\x1d\xc2\x1c\xc1\xed\xfb\xea\xec\x28\x6e\x0f\xe6\x24\x8d\xc6\x23\x1e\xa3\x11\x61\x0e\x68\xec\x79\xe6\xec\x98\x0c\x7d\x98\x43\x12\xbd\xed\x8f\xab\xe1\xd4\xf6\x47\xfc\x70\xb6\x8f\x7b\x08\x73\x48\x26\xba\xdf\x63\x2c\x44\x18\xc6\xfd\x94\xd7\xcd\x00\xb3\x4f\xf8\x31\xd3\x51\x7b\xae\x79\x68\x83\x98\xed\xbb\xef\x78\x38\x78\xe8\xb7\x69\x87\x75\x4c\x0a\x52\x4b\x9f\xa4\x71\x94\xfe\x16\xad\x34\x85\xcc\x79\xb2\xa6\xcc\x32\xcc\xda\xb5\x50\xfd\x52\x8a\x6a\x32\xe1\x9c\x74\x5c\x86\x39\xf4\x9e\x2b\x9a\x4e\xc5\xe5\x57\x16\xdd\xca\xa8\xc2\x71\xd5\x55\xeb\x02\xad\xf3\x42\x17\x3c\x56\x58\xaf\x95\xcc\x43\x9b\xbf\xc0\x95\xb8\x95\xc6\x66\x8c\xfe\x93\x2e\x8d\xf5\xb5\x16\x1e\xd5\x76\xdc\xad\x5e\x0a\x83\x4e\x3f\xf5\xb1\x3d\x82\xa5\x45\xe1\x61\x23\xb6\xc4\x0a\xde\xad\x43\x21\x2c\x1d\x53\x29\x84\x17\x63\x70\x86\x7a\x2a\x8b\xce\x58\xcf\xc5\x64\x6d\x35\xdf\x03\x68\x90\xde\x81\xab\xd7\x7c\xfb\x46\x3d\xc3\xc2\x1a\x51\xe4\xc2\x1d\x88\x51\x51\xc1\x9b\x3b\xae\x9d\x78\xf2\x7f\x58\x40\x41\x32\x9f\x5f\xa7\x81\xfd\x39\x86\x0b\x02\x37\x9b\x4c\x1e\xb9\x2b\x93\xce\xd5\xe8\x26\x17\x5f\x64\x6d\x65\x18\x6b\x3b\xea\x18\xe8\xb4\xc2\x81\x0c\x22\xce\xdc\x5b\x14\x55\x73\x2d\x33\xbb\x6c\xf4\xc2\xf3\xbf\xee\x1a\x17\xef\xc7\x90\xb2\xb9\x17\xd6\xf3\xb2\x92\xce\xa3\xe6\x08\x6c\x99\x98\x46\x9f\xbd\xe2\x97\xc9\xc8\xe7\x6b\x0a\xbb\xe7\x2f\xbe\xcc\xa6\xd9\x34\x3b\x9f\x4d\x47\x27\x7a\xb0\xb5\xd0\x32\x4f\xa8\x35\xe1\x7e\x61\x38\x18\x2c\x0d\xcb\x9e\xcd\x49\x45\xc9\x6e\x9b\x63\xfc\xa4\x01\xbe\x23\x2a\xa9\xef\x67\x64\x9b\xc3\x5b\xb4\x5b\x38\x9f\x42\x33\x8d\x5d\x60\x2e\xea\xd8\xe9\x10\x35\xb4\xf2\xae\x3d\x1a\x61\x91\x8e\x54\xc0\xf9\xf4\xac\x19\x96\x32\xb5\xe8\xb6\xb0\x91\xba\x30\x9b\x0c\x6e\x56\xd2\x81\x42\xff\xd4\x41\x4d\xc7\x5f\xf1\x84\xf7\x18\x49\x6f\x60\x89\x7c\x85\x61\xd5\x36\x52\x13\x79\x5e\x87\x1e\x84\xdc\x2b\x5e\x2f\x39\x48\xb4\xe1\xbe\x44\x78\x36\x3f\x90\x6c\xa4\x6b\xb4\x25\xe6\xbe\x65\xdc\xd4\x36\xde\x42\x45\x6a\x8d\x35\x47\x40\xb5\x85\x4a\xf8\x7c\xd5\xb2\xf2\xd4\xc5\x6e\x05\x41\x28\x67\x58\x2d\x9b\x15\xea\x63\xed\x53\x24\xe9\x78\x28\x08\xf9\x4a\xe8\x25\x3a\xee\x50\x95\x59\x7e\xcb\x39\x52\xe7\xdb\x30\x5b\xee\xdf\x9f\x30\x8c\x70\x9e\xd3\xef\x1c\x51\x87\x11\x75\x9b\x6f\xf7\x12\xf4\xfd\x43\x0f\xe1\x95\x59\x52\xa4\xdb\xc3\x61\xf2\xf4\xfe\x9e\x8d\x82\x5b\xf2\x7f\x40\x33\x5b\x8e\xab\xf9\xc7\x64\xef\x2a\x27\x85\x7b\x68\x47\x5e\x52\xf3\x6c\x52\x52\x10\x20\x1a\x83\xda\x2a\xda\x66\x14\xc7\x19\x23\x78\x0e\x8d\x7d\x65\x5f\x17\x85\x4d\xd2\x6c\xce\xdb\xb3\x85\x53\x4f\xfb\x5b\x6b\xde\x6c\x93\xaf\x71\x13\x0b\x81\x64\xf4\xdf\xd7\x37\xa3\x31\xd4\x56\x8d\xc9\x9a\x03\xc2\x11\x03\x67\xf5\x65\x6f\xad\xd4\xbe\x4c\x42\x09\x00\xb9\xd1\x1a\x73\xcf\x0d\xa2\x69\xcf\x61\x3e\xbf\xa6\x3e\x05\x45\x35\x83\x3f\xb9\x51\x68\xd8\x03\x89\xdc\x68\x2f\x75\x8d\xfc\x8b\xd5\x37\x08\x90\x2d\x7b\x9d\x5b\xf1\x6c\x5e\x2f\x42\x1f\xf6\xb3\xf4\xab\x84\x8a\x9f\x3f\x77\xe6\x68\x61\x00\xf7\xdb\xff\x39\xc3\x74\x82\x78\xbb\xbb\x1f\x08\x98\xd9\x75\x88\x7e\x71\x6b\x9e\xec\xec\x57\x74\x61\x29\xb2\x3a\xbb\x04\xaa\x1b\xb3\x9f\x74\x25\xac\x5b\x09\x95\x84\x51\x51\x82\xb7\xd9\x4b\xe1\x45\x92\xa6\x63\xf8\x33\xa6\x5f\x1d\x11\xac\xcf\x56\xe4\x8b\x09\x67\xdc\x76\x3e\xb9\x84\x51\xdc\xfb\xca\x54\x95\xd0\xc5\xe8\x24\x6a\xf8\x9e\x4c\xe0\xc6\x14\x62\xcb\x1e\x66\x11\x72\xca\x27\x61\xe4\x1d\xe4\x83\x20\x36\x45\xa2\x38\x9a\x25\x37\x44\xef\xda\x01\x43\x13\xff\x09\x30\xa4\x9e\x48\xb8\xe9\x98\xe1\x7b\xcd\xff\xd9\x50\xa0\x50\xb0\x31\x56\xb1\x3f\x33\x3a\x6f\x47\x79\xcc\xf1\x45\x61\xb3\x27\xb9\xfc\x5a\x38\x07\x32\x64\xac\x76\x2e\x9c\xb5\xc4\x5f\x11\x0b\x1b\x63\xed\x16\xc4\x82\xda\x86\x22\xce\x2b\xa4\x0b\x11\x62\x83\xbb\x71\x0b\x13\xa3\xe5\x7c\xaf\xc6\xe4\xd1\x5f\x4b\xd3\x1b\x28\xa4\x13\x0b\x85\x71\xf4\x52\x99\x02\x3b\x3a\xe6\x26\xfd\xc9\x25\x14\x59\x77\x8e\xf1\xb8\x82\x29\x4c\x50\x48\x69\x86\xb8\xbd\x38\xf3\x3e\xd0\xfc\xb0\x03\x3d\x58\x82\x4b\xc0\xdd\x59\x91\xba\xbb\xa1\x2f\x54\x21\xb2\x04\x49\xe1\x1c\xa5\x5f\x85\x38\xc8\xc0\xc9\x79\xca\xc0\xa5\xb4\x2e\x66\x69\xd8\xe0\xd3\x5b\x04\x47\x31\x2e\x9c\xd7\x8e\x96\xe9\x60\xbe\x48\xfb\x4b\xd4\x9c\x87\xbd\x56\xa2\x89\xad\x45\x0f\xe3\x2f\x69\x37\x53\x71\x75\x42\x4a\xc7\x82\x8e\x36\xe7\xb1\x2e\x77\x40\x05\x38\x53\xa1\x5f\xd1\x69\xec\x71\xd0\x6a\xba\x19\x5a\x0f\x8e\x84\xd9\x8e\x5e\x76\x77\x7d\x0d\x6c\xd3\x6d\x15\x99\x6a\xba\xab\xf4\x13\x9e\xd3\x9c\xcf\xd1\x06\x84\x0e\xfb\x78\x6b\xf2\x47\xf3\xf7\x00\xa8\x1c\x9f\x64\xe7\x06\x73\x5e\x2f\x92\x53\xf4\x53\xf8\x3b\xf4\xf2\xda\x1f\xcd\xd0\x2e\xf2\xd1\xc7\x03\xe1\x7c\x46\x15\xfd\x3b\x6f\x3d\x64\x09\x45\x16\x7c\xef\xbe\x37\x47\x9e\x1d\x0c\x92\x9b\xda\xef\xa5\x49\x7a\xfe\x38\x6e\xaa\xa9\x10\xf9\x59\x2d\xc7\xb2\x44\x9c\xda\x86\xb9\x79\x10\x34\x8e\x85\xbb\xf7\x8a\x7f\xbf\x6c\x2f\x16\x07\xe1\x7f\x59\xb0\x88\xd7\x16\x84\x2d\x74\xdc\x86\x87\xb4\x92\x07\x99\x5d\xa7\x89\x95\x4e\x67\x33\x63\x79\x88\x9e\x8c\x2e\xde\xbd\xe3\x9a\x32\x6c\xdd\xe3\xe6\xa1\xcd\xc8\x07\xa3\x65\x8e\x2b\x9f\x79\x11\x33\x08\x03\xf8\x7f\x41\x51\x8f\xcf\xed\x4f\x5c\x3f\xb4\x37\x0f\x9f\xd2\xdf\xbf\xad\xbe\x47\xb4\xf7\x30\xec\xf3\x11\x54\x07\x93\x67\xa0\x0d\x94\xb1\x15\x67\x41\xc3\xbf\x18\x6c\xd1\xc3\xb3\x49\x7a\x74\x68\xff\xbf\x01\x00\x00\xff\xff\x61\xa5\xf0\x68\x97\x28\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x80, 0x4, 0x11, 0x39, 0xa1, 0x94, 0xbe, 0xc3, 0xda, 0xe7, 0x62, 0x5e, 0x21, 0x98, 0xaa, 0xf9, 0x45, 0x4d, 0xa2, 0x73, 0x57, 0x2d, 0xb1, 0x46, 0x3b, 0xad, 0xdb, 0x36, 0xd1, 0x12, 0xa2, 0x24}}
	return a, nil
}

var __hardcodedMiddlewareGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\xcd\x6e\xe3\x36\x10\x3e\x5b\x4f\x31\x65\x11\x84\x5a\x38\xf2\x1e\x8a\x1e\x02\xf8\x50\xa4\x9b\xe6\xb2\xc0\x62\x13\xb4\x05\x8a\x02\xcb\x50\x23\x89\xb0\x4c\xaa\xc3\x91\x55\x23\xf0\xbb\x17\x43\x49\x4e\x14\xa4\xe9\xa9\x07\xcb\xe4\xcc\x7c\xdf\xfc\xb3\x33\x76\x67\x6a\x84\x88\x74\x40\xca\x32\xb7\xef\x02\x31\xe8\x6c\xa5\xaa\x3d\xab\x6c\xa5\x3c\xf2\xa6\x61\xee\xe4\x4c\xbd\x67\xb7\xc7\x4d\x89\x8f\x7d\xad\xb2\x6c\xa5\xea\xd0\xed\xea\xc2\xf9\xcd\x4d\x8b\x07\xa4\xcd\xce\x1c\x0f\x88\x57\x75\x28\x0e\x3f\x6e\xda\x50\xd7\x48\x2a\xcb\xb3\x6c\xb3\x81\x2f\xc6\x3b\xfb\xd9\x95\x65\x8b\x83\x21\x84\x36\xd4\x11\x8c\x3f\x42\x27\x8a\x58\xc0\x6d\x20\xf0\x61\x58\xc3\x80\x97\x84\x60\x83\x67\xe7\x7b\x04\x6e\x28\x0c\xce\xd7\xc0\x0d\x8e\xc6\xd0\x77\xc2\x28\xf7\xc8\xc6\xee\x20\x06\xe0\xc6\x45\xd8\x9b\x23\x58\x32\xb1\x19\x6d\x29\x58\x8c\xb1\xc8\xaa\xde\xdb\xd7\xfe\x75\x03\x92\x56\x71\x67\x7c\xd9\x22\xe5\x8b\x1b\x3c\x65\x2b\x42\xee\xc9\x2f\xc4\xb7\xbd\xb7\x5a\xc8\xf4\x30\xca\xbf\x62\xec\x82\x8f\xf8\x1b\x39\x46\x5a\x03\xc1\x87\x49\xfe\x57\x8f\x91\x73\xe1\x59\x95\x58\x21\x41\x82\x8d\x82\x55\x4a\xe2\x13\x11\x5c\x6f\x81\xd0\x86\x03\x92\xce\x45\xe1\x2a\x38\xeb\xb6\x5b\xf0\xae\x1d\x01\x53\x30\x72\x3c\xc9\xe7\x60\x08\x90\xd2\x2f\x50\x26\x92\x38\x38\xb6\x0d\xbc\x64\x9e\xcf\x85\xe6\x63\x87\x93\x6b\x6b\xa2\x14\x8d\x9c\xaf\xaf\x13\xb3\xd0\x6c\xa1\xda\x73\xf1\x49\xc8\x2a\x3d\xc3\xf2\xb3\x79\xf2\xf2\xd2\x7a\x36\x11\x51\x89\x95\xe9\x5b\xfe\x17\x32\xd5\xfb\x9d\x0f\x83\x9f\xfa\x76\xf1\xfd\x01\x42\x05\x12\x0f\x5c\x3c\xa8\xf5\x99\xe9\xf9\x94\x8f\x49\xca\x77\x9c\x9f\xe2\x96\xc2\xfe\x26\x78\xc6\xbf\x59\x53\x31\x9f\xf2\x7c\xf4\xf1\xb3\x56\x09\xaa\xd6\x29\x82\x09\xf3\xf9\x49\x21\x91\xba\x96\xd8\xd7\xa0\xd2\x94\x30\x19\x8b\xea\x7a\xca\x5e\xa7\x21\x2e\xee\x45\xa3\xf3\xfc\x94\x9f\x1b\xb3\xa8\xc0\x29\x35\xa6\x29\xee\x65\x41\xee\x1e\x1e\xbe\xe8\x61\x0d\xa2\x39\xe5\xd9\x29\xcd\x75\x64\xc3\x7d\x5c\x4e\x02\x0c\x64\xba\x08\x06\x68\x12\xc3\x90\xe4\x59\xca\xfc\x4d\x44\x64\xea\x2d\x4b\x93\xde\x98\xac\x6c\x35\x62\xc0\x79\x16\xb7\x69\xa0\x75\x84\x0f\x6f\x51\xe5\x90\xfe\xef\xd0\x94\x48\xda\x86\x12\x05\x96\xfa\x1f\x8b\x89\x67\x0b\x22\x17\xc1\x12\x5b\xbc\x86\xce\x59\xfe\x8a\x14\x5d\xf0\x5f\x8d\xaf\x11\x4a\xb4\xae\xc4\x08\x43\x83\xdc\x20\x01\x07\x30\xd6\x62\xc7\x60\xe0\x30\x1a\x16\x63\xaa\x0b\x58\x5a\x81\x49\x3f\x75\x21\x87\xc7\x10\xda\xe4\xe1\xa6\x75\xe8\x79\x02\xdc\x34\x68\x77\x2f\x1e\x0b\x2b\xf7\x98\xb6\xda\x26\xbb\x67\x3f\xa9\x16\xef\x81\x5f\x6d\xfa\x1a\xc8\xd7\x8b\xc0\xfe\xc7\xdd\x9f\xb3\x95\x45\x2f\xc6\xb2\x16\xbf\x20\x6b\xf5\xfb\xd5\x18\xf2\xd5\x14\x88\x92\x31\xfb\xaf\x81\xff\xa9\x2c\xe7\x9b\x1a\xcb\x70\x35\x39\x50\xeb\xb9\x20\xc2\xe3\x2a\xf8\x8e\x7c\x3d\xd7\x7a\x5a\xfd\x61\xd1\xdc\x1f\x3e\x7e\xcc\x5f\x48\xf5\x1f\x7f\x3e\x1e\x19\xb5\xec\xee\x7d\x47\xce\x73\xa5\xbf\x3d\xa9\x3d\xc6\x68\x6a\xd9\x1a\xb5\xac\x3b\x5c\x5e\xc4\x4b\xf0\x81\xa7\xce\x63\xb9\x86\xae\x45\x79\x2e\xfa\xae\x26\x53\xa2\x3a\x7d\x7b\x0e\x2a\x4f\xbe\xce\xaf\xd8\xe9\x9d\x9d\xfa\x27\x00\x00\xff\xff\x64\xba\x96\x25\x93\x06\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x17, 0x73, 0xff, 0x36, 0x5d, 0xfe, 0x40, 0x7c, 0x33, 0x99, 0xc8, 0x19, 0x37, 0x47, 0x59, 0x51, 0x96, 0xb, 0x22, 0x81, 0x1c, 0x7d, 0xe4, 0x3, 0x2f, 0x2d, 0x3c, 0x2d, 0xc5, 0x23, 0x70, 0xb8}}
	return a, nil
}

var __hardcodedTracingGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x39\x6d\x6f\xdb\x38\xd2\x9f\xad\x5f\x31\x6b\xa0\x5b\xa9\x70\xa4\x66\xdb\xe6\x59\x64\x91\x0f\x69\x92\xb6\x06\xda\x6c\xd0\xe4\xb9\xc5\xe1\x70\x28\x68\x69\x2c\x11\x91\x49\x1d\x49\xf9\xe5\x8a\xfc\xf7\xc3\x0c\x29\xd9\x8a\x9d\xa0\x3d\x5c\x3f\xb8\x36\x39\xef\xef\xc3\x34\x22\xbf\x17\x25\x82\x45\xb3\x44\xe3\x8c\xc8\xa5\x2a\xa3\x48\x2e\x1a\x6d\x1c\xc4\xd1\x68\x9c\x6b\xe5\x70\xed\xc6\xd1\x68\x8c\x2a\xd7\x85\x54\x65\x36\x93\x4a\x98\xcd\xe0\xa8\xc2\x35\xfd\x9e\x2f\x18\x52\xa1\xcb\x2a\xe7\x1a\xfa\xae\x2d\x7d\x5a\x67\x72\xad\x96\xf4\xd5\xc9\x05\x8e\xa3\x68\x34\x2e\xa5\xab\xda\x59\x9a\xeb\x45\x56\x88\x25\xe6\x65\x95\x95\xfa\xc8\x36\xb8\xca\xe8\x83\x60\x4b\x9d\xea\x06\x95\xc3\x1a\x17\xe8\xcc\x26\x95\x3a\x23\x89\x8c\x9c\x65\x52\x59\x67\xda\x05\x2a\x27\x9c\xd4\x2a\xdb\x21\x57\x6a\x23\xeb\x5a\x64\x8b\x76\x9d\x69\x87\xf5\xa2\x5d\xff\x2c\xb5\x4e\x05\xc6\xef\x74\x39\x44\x80\xee\x9f\xbb\xcb\x70\x4d\xc6\x44\x63\x33\xed\xea\x67\xc9\x3c\x02\xe5\x8f\xd2\x34\xf9\xb3\x38\x8d\xd1\x8d\x28\x59\xe8\x71\x34\xb2\xc5\x3d\x79\x11\xe1\x69\x04\x5b\xdc\x67\x0c\xf3\x2c\xd9\x1e\xca\x7f\x3a\xb4\xee\x59\xf8\x1d\x8a\xcd\x7d\x99\x4a\x95\x5d\xd4\xb8\x44\x93\xdd\x8b\xcd\x12\xf1\xa8\xd4\xe9\xf2\x24\xab\x75\x59\xa2\x19\x47\x49\x14\x65\x19\x74\xa2\x6b\x03\x4e\x43\x6b\x31\x8d\x96\xc2\xec\x1e\xef\x28\x97\xde\xe1\xda\x7d\x11\xcd\xcd\xf6\xf6\x6c\x78\x4f\x12\x5c\xf8\x70\xfd\xfe\x00\x59\x06\x2c\x53\x23\x0c\x2a\x07\x15\x8a\x02\x0d\xb3\x2d\x70\x2e\xda\xda\x5d\xe8\xba\xc6\xdc\x69\x73\x43\xb1\xbe\x12\x16\xf2\x4a\xa8\x12\x0b\x98\x1b\xbd\x80\x77\xef\x7e\x3f\x79\x0d\x52\xc1\xb5\x5e\xe2\x62\x86\x06\x84\x2a\xc0\x55\x08\x1f\x35\xd4\x72\x66\x84\xd9\x10\xb5\x4a\x58\xf5\xd2\xc1\x0c\x51\x41\xdb\x14\xc2\x61\x01\xab\x0a\x15\x48\x07\xd2\x6e\x8f\x10\x72\xa1\x48\x49\x20\xb7\xa6\x97\x07\x84\x60\xed\x0f\x4a\xd7\x4a\xe5\x8e\x4f\xe0\x0c\xde\xbe\x39\xfe\x3f\x56\xe2\x16\x5d\xdb\x7c\xac\xf5\x4c\xd4\xac\xf8\x8d\xd1\x4b\x59\xa0\x39\x57\xc5\x55\x88\x22\xb0\xe8\x48\x00\x10\x0a\xba\xc8\x22\x43\xfb\xef\x13\xa2\x22\x2c\xac\xb0\xae\xe9\x7f\xd2\x6c\xe0\x5a\x28\x99\xba\xb7\x22\x59\x9a\xe9\xc3\x9c\xbc\xc5\x47\x25\x2a\x34\xe4\x08\xdb\xbb\x2f\xcb\xe0\xae\xc2\x2d\x37\x32\x59\x8f\x29\x0c\x82\x41\xd7\x1a\x85\x05\x19\x56\x9b\x9e\x5e\x45\xd6\xa9\x6b\x2f\x5f\x81\x73\x12\xbe\x6a\x5d\xa1\x57\x2a\x8d\xe6\xad\xca\x7f\x44\xdf\x38\x77\x6b\x08\xf5\x2a\x0d\x81\x90\x40\xdc\xa5\x44\x7a\xdb\x08\xd5\xc1\x4e\xe0\x55\x7f\xce\xf4\x4c\x47\x70\x02\x68\x8c\x36\x09\x7c\x8f\x46\x59\x06\xc7\x29\x99\x91\xac\xd8\x29\xc5\xc7\xbf\xf5\xc7\x4c\xc3\xf4\x5a\xf2\xed\x9b\x14\x84\xb5\xb2\x54\x03\x1b\xee\x00\x31\xd4\x74\x0e\x2b\x7c\x49\x46\x69\x95\x92\xaa\x84\x5a\x93\x11\x36\x13\xb2\x87\x02\x32\x14\xe8\xf9\x1c\xac\x58\x34\xb5\x54\x65\x0a\x7f\xba\x0a\xcd\x4a\x5a\xf4\x67\xe8\x25\x7c\x01\xda\xc0\xaa\x12\x8e\xf2\x0d\xee\xbe\x9e\x5f\x4c\xaf\x3f\x7e\xbb\x3d\xff\x72\xf3\x99\xbe\xdc\x7c\xfd\xf3\xfd\xf9\xfb\xe9\xe7\xe9\xdd\xdf\xc1\x36\x98\xcb\xb9\x44\x9b\x46\xa3\x8e\xec\x8d\xd1\x33\x31\x93\xb5\x74\x1b\x38\x3d\x83\xd7\xe9\xeb\xe3\x68\x24\xed\x67\x12\x86\x0e\xb4\x4d\x3f\xa2\x43\xb5\x8c\xc7\xdf\xa6\xb7\xdf\x3e\xff\x79\x71\xfe\x79\x9c\xc0\xd9\x19\x8c\x9d\x69\x29\xe7\xe5\x1c\x3a\xf8\xef\xd1\x68\x34\x5f\xb8\xf4\xc6\x48\xe5\x6a\x15\x8f\x6f\xd1\x91\x4f\xf9\x76\x9c\x44\xa3\x83\x6c\xcf\xe0\x38\x7d\x1d\x8d\x1e\x00\x6b\x8b\x20\xe7\xb0\x7c\xc4\xf8\x39\xa5\xc6\xc9\x1f\xb0\x84\x5f\xce\x60\x3c\x66\xf6\x07\x18\x7c\x30\x7a\x71\xa5\x96\xec\x5a\xa2\x1c\xba\x51\x7a\x23\x8c\xc5\x0f\xb5\x16\x2e\x5e\x4e\xe0\xe4\x2d\xc9\x27\xe7\x0c\xf5\xcb\x19\x28\xe9\xf5\x19\xf9\xa0\xa5\xdf\x13\xff\x41\x1a\x5e\x51\x94\xcc\xe3\x71\xae\xdb\xba\x00\xa5\x1d\x34\x44\x0d\x5e\xbe\xb0\x2f\x49\xe3\x39\x91\x1d\x4f\x60\x49\x44\x1f\x9e\x54\xfc\x69\x69\xa3\xd1\x43\x14\x8d\x44\x51\xb0\xc8\xc4\xf2\xb6\x21\xab\xce\xe3\xf1\x0b\x7b\xfa\xa2\x18\x4f\x0e\x57\x92\x4f\xda\xba\xc9\xc1\x32\x92\xf8\xb8\xbb\x5a\xa2\xd9\xc0\xf1\x3b\xb0\x98\x6b\x55\x50\x09\x78\x59\x53\x8c\x6e\x48\xee\x5c\x2b\x85\x39\x3b\x6d\x58\x0e\xf2\x8e\x14\x08\xc7\x74\x28\x6b\x03\x1b\x0e\x5d\xaa\xc2\xa0\xe7\x3e\x8c\x2b\x6d\xdd\x29\x57\x2b\x02\xfd\x8b\x02\xba\x0b\x73\xa9\x28\x11\x8a\x36\x67\x04\x57\x49\x4b\x85\x52\x80\x95\x05\xe6\xc2\x4c\xb8\x6a\xac\x76\x30\x98\x44\x48\x8e\x1d\xf8\xee\xa4\xa3\xab\x1b\xec\xa5\x3d\xea\xa5\x4d\xa3\x51\x61\x24\x65\x06\x85\x54\xe8\xa8\xe9\x35\xae\x2e\xf9\x34\x8e\x46\xa3\xfe\xf4\x2f\xe9\xaa\xaf\x18\x2c\x20\xb5\xba\x41\x23\x75\x11\x1f\xbf\x7b\x45\x43\x4b\x7a\xcb\xf6\x4a\x26\x8f\x51\xae\x54\xd1\x68\xa9\x5c\x4c\xce\xda\xbf\x9e\x2a\x8b\x79\x6b\x30\xa6\xab\x24\x1a\xe6\xc7\xd1\xd1\x91\x17\xef\xe8\xe8\x88\xf2\x83\x26\x9f\xf4\xb2\x5d\x34\xb1\x3f\x4e\xa2\x11\xf6\x35\x2b\x84\x2f\xbb\xfd\x1a\x57\x7d\xe1\x8b\x46\xa3\xdc\xad\x89\xb1\x47\xf2\x7c\xf6\x23\xf9\xd9\x40\xe6\xb2\x07\xb9\x41\xe1\xc8\x9a\x1d\xd7\x53\x78\xb1\x1c\x33\xeb\x84\x22\x72\x4f\xfa\x0e\x6e\x4f\xfe\xee\x82\xa2\xce\x35\x24\xb7\xc2\xd5\xb0\xd6\xc6\x5b\xd5\x0e\xe4\x01\x21\x3e\xe6\x36\x6c\x46\x7b\x3c\x5d\x43\x48\x34\x90\xa4\xb7\xe8\x1e\x31\xa3\xcb\x91\x9f\x40\x52\xca\xb1\xd0\x22\xa8\x6d\x24\xe9\x54\xcd\xf5\x65\x3c\xb6\x4e\x18\x52\xdf\x33\x32\xe3\x09\x04\x84\x2f\x64\xbf\x31\x39\x18\xad\x1d\x9f\x82\xff\x47\xbf\xc9\xec\xe3\x4e\xfc\x23\x23\x1c\x8e\x4f\x0f\xa9\x33\x89\x46\x0f\x49\xd4\xf9\x60\xab\xb9\x6b\xd8\x19\xd1\xc3\x8f\xf6\xf5\x0f\xda\xdc\xa1\xe5\xe9\x62\x81\x42\x71\xaa\xce\x90\x1a\x30\xb7\xd5\x56\x49\x07\x34\xaf\x49\x55\xfa\x26\xaf\x0a\x58\x48\xf2\xae\xef\xf2\x96\x58\x80\x98\xe9\x25\x72\xf7\xd5\xad\xa3\xd4\xa3\xd4\xdd\xc5\x4d\x61\xea\x42\xbf\xb6\x34\x42\x48\x75\xb4\xc0\x85\xf6\x53\x4f\xdf\xe1\x89\x00\xae\xc5\x42\x72\x0a\x86\x91\x00\x0b\xb0\x8d\x50\xf6\xc7\x7b\x77\xd0\x29\x4e\x20\x7e\xd5\x0f\x9c\xe9\x54\x7d\x61\x8e\x3f\xd7\xb5\x7b\xd9\x4e\xcf\x60\x4b\xeb\x1a\x57\x8f\xc9\xc5\xc9\x8f\x44\xe6\x71\xfa\x3a\x79\x3e\xa6\x9e\x73\x29\x1b\xe0\x69\xf2\xf0\xc4\x74\x72\xa8\x59\x70\x47\x39\x79\x9b\x3c\x69\x04\x52\x3e\xc8\xd2\x43\x5c\xef\xb1\x8e\x46\x5c\x91\x39\x60\x06\x35\x7c\x7a\xb9\x1d\xe9\x52\x98\xee\x0e\x68\x9d\x38\x14\x6a\x2b\x6d\xee\x21\x16\x0e\x6a\x14\xd6\xc1\x4a\xba\xca\x17\x65\x3f\x93\x50\x99\xe8\xba\xc3\x2e\x45\x58\xb4\xd6\xf5\x01\x12\x46\xc8\xe9\xa5\xa5\x98\x9b\x6b\xb3\xa8\x37\x20\x1c\x18\xa1\x0a\xbd\xf0\x13\x38\x51\x40\xe5\xa4\x41\x8a\xa6\xdc\x07\xa8\x54\xee\xe4\x6d\xea\x59\x7c\x08\xb1\xd7\xd4\x38\x19\x68\xb2\x3e\x32\x62\x33\xe4\x5e\x68\xb4\xdc\xa5\x0b\xcd\xd2\x12\x89\xde\x48\x54\xa5\x6f\xbd\xf4\xdb\x71\xf1\x86\x17\x87\xf7\xc2\x62\x11\x0f\x0d\x3e\xbd\xfc\x4a\xcd\x2e\x5c\x1d\x28\x5c\x09\x37\x01\x3f\x0c\x5b\x84\x85\x58\xcb\x45\xbb\xb0\x3c\x01\x3b\x0d\x65\x2b\x4c\x01\xa2\x14\xb4\x74\x82\xd5\x0b\x74\x15\x67\x8f\xa6\xcf\x95\xd1\xaa\xe4\xa4\xb5\xa8\x68\xb9\x06\x01\xce\x77\xd6\x42\x38\x01\xad\xc2\x75\x83\xb9\xc3\xa2\xde\xec\x29\xd1\x08\xf5\x59\x2e\xa4\xb3\xc3\xb1\xd7\x9f\xf1\x30\x73\xee\x68\xdf\x6d\x1d\x5e\xe8\x56\x39\xbe\x38\x85\xe3\xd7\xaf\x49\xe2\xd1\xd5\x12\x95\xdb\xbd\xa0\x7f\xdd\xe5\x67\xa9\xee\x1f\xdf\x75\x97\x0f\xac\xf1\x50\x96\x8d\xca\x77\x22\xdd\x37\x3f\x5f\xe4\xbe\xb4\xeb\x5b\x7e\x71\xf8\x22\x8b\xa2\xc6\xd5\x76\x33\xb0\xb0\xd8\x1e\xb9\x4a\x38\xb0\x15\x0f\x58\x33\x04\xe1\x9c\xc8\x2b\x2c\xc8\x84\x02\x76\x76\xfc\xf0\x7c\xc1\xeb\xc7\xd4\x79\x57\xbb\x15\xfb\x59\x95\xf6\x14\xb8\xa8\x5b\x5f\x92\xfc\x74\x21\x8a\xc2\xff\xf6\x7b\x2b\x48\x35\xd7\x44\x96\x82\xc8\xe0\xbf\x5a\xb4\xee\x28\xcc\xca\x79\xd7\x01\x88\xfa\x57\x59\x56\x0e\x94\x5e\xd1\x3e\xa7\x55\xbd\x01\xdb\x36\xfc\x56\x42\x40\xe4\x2b\x8a\x68\xa9\x98\x10\x45\xb5\x70\x5e\x8b\x4b\xe1\x44\xa1\xb9\xab\x62\xee\xba\xc2\x78\xc0\x0e\x31\xe9\x22\x73\xbc\x16\x0b\xa4\x39\x55\xaa\x32\x01\x02\x8e\x2b\xe7\x9a\xf4\x93\x50\x05\xe5\x18\xec\xfe\xa2\xbc\xd7\xae\x26\x4b\xf0\x74\xc0\x0f\x1e\xe9\x61\x9a\x93\xfe\x9e\x7c\xb4\x5d\xa6\x6d\xbc\x5d\xbb\x93\x6d\x4d\xf3\x9c\xe1\x79\xde\x1d\xb0\x97\x61\x20\xe8\x07\xc2\x67\x22\x66\xe5\xf1\xbe\xa2\x6d\xb4\xb2\xf8\x97\x91\x5c\xe9\x0c\xbc\x0a\xe7\x6c\xf6\xc4\x8f\xdc\x59\xd6\xc9\x49\x0b\x36\xe0\x9a\xfc\x44\x7d\x85\x1b\x58\x23\x54\xca\x5e\x68\x5a\xc7\x7b\xb6\x0a\xbe\xcb\xdd\xfa\xb1\xdf\x88\x9a\xed\x5b\x01\x67\xc3\x6e\xf3\x37\xdd\xa6\x18\x93\xd6\xc3\xe5\xc5\xf2\x89\x9c\x83\xcd\x79\x6b\x60\xe4\x1e\xfc\x0f\xb0\x79\xfa\x49\xd8\x50\x12\xe2\x20\xf9\x88\xa4\x9b\x5e\x4e\x3c\xbf\xe9\x25\x63\xe6\x8c\x4a\x40\xe9\x2d\xfb\x34\x4e\x26\x74\xda\xe3\xf6\xc7\x4c\x22\xcb\x38\xd5\x29\x60\x68\x4f\x41\x8a\xde\x0a\xd7\x21\x1e\xfc\xba\xcd\x85\x90\xc2\x6d\x02\x56\x93\xae\xe0\x2a\x4d\x9b\xa1\xf6\x11\x97\x6b\x63\xb0\xf6\xb3\x5c\xad\x4b\xeb\x9f\x2f\x88\x9f\xe5\x12\x6e\x99\x93\x9c\x43\x8d\x2a\x0e\xb2\xf2\x56\xf7\xe6\x37\xf8\xf5\x57\x3e\xf5\x9a\xf0\xe1\xf1\x09\x7c\x07\x72\xca\x60\x25\xf0\x11\x7e\x4a\xb7\xb3\x8d\x43\x88\xdf\xfc\x76\x94\x57\xc2\x90\xb0\xc9\x04\x7e\x0f\xa7\xc7\x27\xdb\xd3\x50\xf8\xb9\xbc\x35\x34\x55\x14\x5e\x90\x51\x10\xe1\xbd\x9d\xc0\x37\xb2\x59\x85\xeb\xf4\x12\x73\x5d\x60\x30\x4d\x27\xa3\x07\x3f\x30\xc7\xed\xba\x32\x3d\x2f\x8a\xee\xd7\x98\x31\xbf\xc9\x62\x3c\xf1\xb8\xa3\xe1\x32\x45\x9b\x94\x7f\xd2\x4c\xdf\xcb\xf2\x4a\x15\x52\xa8\xf4\xff\xd9\xbe\x71\x2f\xd5\x3f\x7e\x3f\xfd\x67\x92\x04\xe6\xde\x30\x4f\x8b\x1a\x0c\xf7\xf3\x92\x12\xe2\x7f\x27\x68\x27\x52\x27\x23\x6d\x9c\xfe\xa3\x4a\xb9\xc6\x7c\xba\xbb\xbb\x89\xcd\x6a\x02\xdc\xa0\x1f\x12\x9e\xe6\x7d\x39\xf6\x33\x82\xb2\x5c\xc8\xba\x3a\x4c\xe9\xe4\xfa\x53\xff\xc0\x43\x95\x2d\xaf\x25\x2a\xd7\x55\x49\xdb\x15\x5d\x27\xee\x91\x2b\x9f\x18\x62\xad\x8c\x68\x26\x80\x69\x99\xfa\xf4\x0f\xdb\x6a\xcf\x6f\xd2\x3f\xab\x05\x8a\x44\x2f\x3c\xe2\xc0\x52\xd4\x2d\x77\xcc\xa6\xa5\x05\x35\xa4\x3e\x28\xaa\x8d\xba\x75\x3c\x26\xec\x3f\x3a\x49\xcb\xf0\xdd\x43\x1e\xa1\x85\x67\x18\xad\xa8\x6a\x03\x8a\xbc\xea\xd8\x71\xf6\x84\x87\x6f\x58\x69\xf5\x92\x05\x98\x61\xc9\x93\xb4\x93\xb5\xcf\x26\xda\x39\x95\x74\x52\xd4\xf2\xdf\x58\x40\x8c\x69\x39\x21\x6d\xa5\x0a\xcd\xc7\x37\x99\xb6\x49\x42\x69\xdf\xb5\x69\x3c\x13\x16\xb7\x16\xf6\xe5\x4e\xb7\xaa\xb8\x33\xb2\x69\x78\xdc\x6b\x84\xa2\xea\x7c\xe1\xd6\x7f\x63\x9d\xa5\x72\x68\xe6\x22\xc7\xef\x0f\xc9\x3e\x02\x55\x9a\xc1\x66\x74\xd1\x6d\x6d\x86\xa0\x9c\x87\x1a\xae\x46\x66\x07\xff\xfb\x40\xa0\x53\x18\xfc\xdc\x97\xe6\x74\xef\x64\x67\x91\xf9\xdf\xd2\xa5\x88\x74\x9b\x06\x07\x64\xa9\xee\xb5\xb9\x23\xad\x87\x96\x04\xd8\xb7\x4d\x34\x7a\xce\x98\xfd\x48\x1e\x53\xa8\xef\xa0\x25\xd0\x13\x89\xf7\x5a\x52\xfc\x6a\xd0\xba\x76\x97\x8d\xbe\xf5\xf9\xbf\x12\xa4\x03\xbf\x53\x6b\x74\xe9\xd0\x0a\xfc\x3a\x10\x80\xa9\x05\x3f\x9a\xce\x79\xcf\xf8\xb8\xb7\x67\x24\xc9\x1e\xe2\x13\xbd\x7b\x0f\xee\x36\x98\xe3\x03\x17\x6b\x5a\x7b\xb8\x1f\xd3\xd0\xa9\x8b\xd0\x52\x0e\xb4\x61\x7f\xe1\x7b\xda\x72\x02\xfa\x9e\x2a\xdd\x4e\xe1\x4a\xd9\xb6\xb1\x71\xe9\x63\x7b\x27\x69\x1c\x06\x97\xd0\x3e\xf5\x7d\x68\x8d\x8f\xb7\xf9\x65\x58\xe0\xb9\xa4\x76\x81\xba\xf4\x07\xc1\xb0\xcb\xbe\x8e\x75\xf1\x96\x7e\xf1\x92\x67\x19\x6d\x1c\x08\xc2\xf6\xd6\x7f\x69\xfb\x15\xa0\x2b\x14\xfc\x7c\xe4\xc7\xd3\x24\xdd\x71\x71\x37\x8a\x5e\xf9\xc9\x82\x8c\x74\x4e\x77\xbe\x67\x87\x79\xc3\xcf\x88\xdb\xce\xc9\xd3\x1d\x17\x15\xa1\x38\xf2\xba\x2a\x12\xfe\xbc\x10\x72\xff\x20\xcd\x03\x51\x15\x9a\x8b\x4f\x8c\xe9\x65\x3f\xee\x3d\xce\xee\x40\x8f\xdc\x11\x88\x71\x6a\xff\xe8\x50\x43\x03\xcc\x70\x74\xd9\x1f\x5b\xba\xb5\xf1\x11\xdc\xfe\x7c\x32\xd9\x83\xd9\x9b\x6c\xf8\x99\xc8\x8f\x4c\x5b\xf1\x02\xfc\xae\x94\xdb\xa0\x4d\x83\x86\xbb\x82\x4f\x06\x7f\xdc\xf9\xc4\xf6\xbd\x10\xc6\x48\x34\xb1\x09\xbf\x93\x9d\x49\xf5\x67\x66\xac\x87\xe8\x3f\x01\x00\x00\xff\xff\x81\xa1\xcd\x53\xf9\x1c\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x4f, 0xd1, 0x5e, 0x3d, 0x42, 0xa0, 0xef, 0x25, 0xa0, 0x8d, 0x7c, 0xeb, 0xf5, 0x5f, 0xe6, 0x86, 0xd6, 0xf8, 0x80, 0xde, 0xce, 0x46, 0x7e, 0x29, 0xc, 0xff, 0x42, 0x2a, 0x2c, 0xd3, 0x23, 0x4f}}
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
