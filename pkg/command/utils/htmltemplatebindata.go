// Code generated by go-bindata. DO NOT EDIT.
// sources:
// templates/single_chart.html (18.363kB)

package utils

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

var _templatesSingle_chartHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x7c\x7b\x93\xdb\x36\x92\xf8\xff\xfe\x14\xbd\xca\xee\x4a\xf3\xb3\x00\x91\x20\x08\x92\xca\x68\xea\xe7\x8c\xbd\xb1\x6b\xc7\x49\x2a\x76\x7c\x77\x99\x73\xa5\x20\x12\x23\xd1\xa6\x48\x1d\x09\xcd\x23\x29\x7f\xf7\xab\x06\x48\x3d\x49\x59\x7e\xec\x56\xed\xdd\x31\x65\x11\x04\x1a\xfd\xee\x46\x03\xe4\xe4\xfc\x4f\x4f\x7f\xbc\x7c\xfd\x1f\x3f\x3d\x83\xb9\x5e\x64\x17\x8f\xce\xed\xed\xd1\xf9\x5c\xc9\xe4\xe2\x11\x00\xc0\xf9\x42\x69\x09\xf1\x5c\x96\x95\xd2\x93\xde\x4a\xdf\x90\xb0\x57\x0f\xe9\x54\x67\xea\xe2\x27\x55\xde\x40\x22\xab\xf9\xb4\x90\x65\x72\x3e\xb2\xbd\x16\xa2\x8a\xcb\x74\xa9\xa1\x2a\xe3\x49\x6f\xae\xf5\xb2\x1a\x8f\x46\x71\x92\xd3\x77\x55\xa2\xb2\xf4\xb6\xa4\xb9\xd2\xa3\x7c\xb9\x18\xbd\xfb\xaf\x95\x2a\x1f\xfe\xbf\x47\x7d\xea\x8e\x92\xb4\xd2\x75\x0f\x5d\xa4\x08\xdd\xbb\x38\x1f\x59\x5c\x9f\x8a\x58\x21\xeb\xba\xb2\x38\xeb\x07\xa2\xf2\xe3\x78\xed\x03\x5e\x71\x91\x57\x1a\x32\x35\x53\x79\xf2\x02\x1f\x60\x02\xd7\xeb\x51\xbc\xfa\x4b\xa9\xe7\xe3\xd1\xe8\xa5\xef\x32\xf0\x5d\xb6\x20\x9c\x87\xe0\x48\xfc\x35\x2d\x70\xc1\x81\x30\x12\xe0\xc0\x4e\x1f\x31\x7d\xbf\xf6\x87\xed\xe8\x38\x67\x34\x00\x97\xfb\x34\xbc\x12\x82\xba\x10\x44\x21\x65\x31\xf1\x1c\x1a\x82\xef\x51\x0f\x02\x1c\x67\x0e\x88\x88\x7a\xd8\x98\x07\xd8\x1d\x0b\x97\x22\x2d\xd7\x71\xa8\x4b\x84\xa0\x81\x01\x20\x2e\x73\xae\xfc\xd0\x45\x50\xc4\x69\x11\x11\x9c\x41\x5c\x27\x58\x37\xbd\x10\x67\xff\xde\xc5\x55\x20\x1c\x88\x84\xf3\x9c\x09\x1e\x13\xd7\x75\xa8\x0f\x0e\x61\x0e\x0a\x43\x7d\xd3\x60\x8e\xf3\x06\x47\x9d\x7a\xb8\x19\x80\x7a\x70\xce\x23\x11\xd7\x33\xb1\xcf\x00\x40\x0d\x70\x8b\x83\x0e\x98\x61\xd2\x0c\x34\xb3\x3b\x99\xe2\xc2\x87\x48\xd0\xe8\x2a\x12\xd4\x07\x4f\x70\x2a\x62\xc2\x42\x60\x0e\xe5\xc4\x8b\x50\x5f\x02\x99\x88\xa8\x8b\xd4\x78\xe6\x72\x87\x06\xc0\x3d\x8f\xba\xb1\x8b\x4d\xcf\x03\xee\x52\x0e\xbe\x8f\x7a\x45\x6d\x63\x6b\xce\x7d\x9f\xf2\xd8\xe3\x34\x00\x07\x84\x4f\x39\x61\xac\x06\x20\x08\x70\x15\x21\x62\xe0\x3e\x37\x68\x88\xe7\x11\x97\x88\x88\xba\x86\x16\x0a\xc0\xaf\x7c\x3f\x32\xcc\x21\x47\xc4\x70\x24\x84\xbd\x47\xfc\x88\xa2\x45\xe0\xd3\x08\xd0\x32\xec\xb9\xc7\x43\xea\xc6\x84\x33\x1a\xa1\xdb\x30\x54\x0b\xa3\x11\x71\x1d\x0e\xc2\xa1\xee\x55\xe8\x00\xf7\x19\xc2\x30\x94\xc2\x0b\x90\x03\x6c\x85\xcc\xb0\xee\xba\x11\x0d\x33\x57\x70\xea\x02\x0b\x39\x0d\xe3\x06\x8e\x81\x70\xa9\x6b\xb0\x40\x83\x6e\xee\xb1\x80\x86\xb1\x25\x87\x28\x7c\x14\x1b\x99\xe1\xc4\x90\x8b\x38\x07\x3f\x70\x69\x64\xd0\x10\x24\x07\xa6\x65\xc9\x11\x43\xee\x2a\x08\x22\x9c\x24\x02\xea\x59\xbe\x0c\x20\x41\x7a\x06\x0d\x69\xf0\x75\xea\xc0\x75\x1c\xe4\xd8\x77\xd9\x95\x88\x18\x78\x9e\x09\x30\x70\x11\x2f\x3e\xe0\x3f\xf3\x80\xbd\xf8\x20\x22\x96\xb9\xa1\x03\x9e\xcb\xa8\x6b\xe6\x88\x88\x75\xa2\xf7\xb9\x6b\xa4\x66\x94\x5f\x09\xc1\xc0\xe3\x01\xf5\x32\x16\x38\x94\x81\x17\x51\x2f\x66\x18\x3c\x1e\x12\x0b\xa8\x07\x9e\xa0\x02\xdc\x10\x3d\x83\x5f\x05\x3e\x3a\x85\xf0\x18\xf5\x33\x2e\x28\x03\x86\x66\x8f\x39\x15\xc0\x8c\xaf\x61\x80\x72\x41\x43\xc2\x51\x3b\x1e\xa7\xde\x15\x72\x19\x3a\x21\x0d\x33\xc2\xb8\x6b\x82\x37\x88\x11\x32\x04\x17\x95\xec\xbb\x34\x20\x01\x0d\xcc\x14\x82\x53\x0c\x6a\x62\x50\x5f\x05\x08\xc7\x39\x8b\x89\x6b\xc2\x39\xa4\x21\x09\xa9\xc0\x59\xc8\x16\x31\x6c\x79\x6b\x31\x5c\x86\xe1\xcd\x38\x1a\xdb\x65\x84\x71\xca\x90\x1f\x6e\x5b\x7e\x48\x3d\x72\x4c\xf5\x41\x24\x50\xab\x19\xda\xc8\x0d\x9d\xd8\x15\x34\x42\xa7\xf6\x08\xa7\x8c\x08\x9f\x46\xc4\x0b\xcd\xfd\xb9\xf0\xd1\xc9\x1d\x08\x45\x4c\x36\x60\x7e\x84\x9c\x63\xcb\x84\x93\x73\xe5\x05\x0e\x30\x4c\x64\xcf\x5d\x61\x5c\xd5\x43\x81\x1c\xe2\x73\x54\xb0\xa0\x82\x78\x21\xc6\x58\x84\x34\xc1\x0d\x9d\x2b\x97\x71\xb4\x9f\xc5\x0a\x88\x0b\x90\x77\x84\x81\x1a\x76\xce\x9c\xf0\x8a\x07\x1c\x22\x2f\x8c\x37\x60\x48\xdc\xb6\x2c\xf1\x46\x0c\xea\x3e\x0f\x05\x8b\x2d\x65\x40\xca\xc4\x58\xb5\x16\xe5\xaa\x96\xba\x53\x2b\xa1\x1b\x62\xde\x08\x02\xea\xcf\x89\x11\xc2\x45\x13\xa3\xae\x03\xf3\xcb\x09\x73\x3d\x1a\x10\x86\x29\xb6\x6e\x87\x91\x09\x5b\x57\x60\x66\xe5\x81\xe9\x67\x94\x99\xa0\xe4\x84\x85\x18\x09\x3e\x06\x02\x43\x5f\x10\x24\xe2\xd8\xed\x51\x71\xe9\xfa\x1e\x3a\xa0\x1f\x51\x01\x82\x03\xe7\x28\x91\xc0\x2c\x15\xd1\x08\xd3\xa4\x13\x40\xc8\xc1\x8d\x38\xf5\xc1\x0d\x23\x1a\x60\x9a\xcc\x7c\xe1\x83\x43\xbd\x38\x08\x51\x76\x70\xb9\x4b\x3d\x22\x3c\x93\xf9\xb1\x69\x7e\xc1\x21\x38\x6e\xfa\x4d\x0f\x27\x9b\x51\xde\xed\x17\x3e\xea\x0e\x3d\x23\xf0\x49\xe0\xc7\x02\xe5\xc7\x1f\x30\x2d\x37\xf0\x91\xa6\xf1\x6f\xdb\x4f\x36\xfd\xb6\x69\x86\xc0\xc9\x48\xe0\x83\xc1\x81\x68\x8e\x82\x6e\xe1\x37\xad\x9a\x08\x98\xe1\xac\x41\x03\x0d\x9a\x2e\x48\x58\x0f\xd6\xfd\xa6\xd9\xb0\xd3\xe0\x80\x5a\xa8\x2e\x40\x38\x26\x71\x66\xc5\xf9\xbd\xbf\x56\xde\xdb\xbd\x4a\xa2\x52\x99\x8a\xf5\x93\x2c\xc3\x62\x02\x26\xbb\xd5\x83\x70\x42\x19\x09\x30\xa5\x82\x6b\x4a\x04\x37\x62\xd0\xf4\xd8\x42\xc2\x8d\xd8\xc2\x21\xcc\x17\x31\x09\x43\xca\x18\x37\xae\xe5\x40\xe0\xd2\x20\x10\xa6\xe9\x0a\xa7\xb2\x8f\x50\x3f\x36\xff\xc8\xa6\x9b\x6c\x1e\x49\xf3\x88\x50\x2f\x6d\x96\x72\x62\xe3\xa5\x8e\x40\xfc\x5e\xc8\x91\xa2\xbd\x87\x61\xe5\x62\xee\xf7\x04\xb6\xa1\xee\xc3\x3b\x30\x5f\xd8\x7b\x88\x7a\x31\x30\xa6\x5d\xf7\x2d\x1c\x22\xb8\x73\xc9\x84\x4f\x19\x0f\x0d\x4f\xc6\xab\x3d\xea\x38\xa1\xf1\x6b\x97\xc5\x0e\x88\x90\x46\x11\x03\xe6\xb8\x06\xcc\xf3\x99\xa9\x99\x3c\x9f\x55\x9c\x87\x84\x85\x16\x1e\xdb\x9e\xcf\x62\x87\xd8\x09\xa4\x9e\x80\x9d\xa4\x1e\xec\xef\x69\x7f\x95\xbf\x3a\xa6\x7f\x23\x36\xa6\x28\x97\xa1\x56\x59\x48\x3d\x93\x39\xc3\x48\x10\x37\x64\x54\xf8\x02\xe3\xd4\x89\x44\x16\x71\x1a\x05\x26\x56\xa3\x40\x3c\x71\x7d\x41\x51\xf8\xe6\xee\x98\xff\x8c\x49\x03\x16\x37\x76\x6a\xb1\x80\x6c\x9d\x49\x58\x40\x9d\x80\x91\x30\xa0\x22\xcc\x5c\xc7\xa5\xbe\x27\x48\x7d\xbf\x0c\x71\xd9\x0e\xc1\x8b\x42\x2a\x18\x37\x95\x25\x8f\x3c\xea\x71\xdb\xb6\x4a\xf4\x58\xab\x05\x22\x63\x34\xb9\xe5\x52\xc4\xb4\x63\xe2\x72\x1a\x84\x68\x6c\x16\x52\xc1\x89\x47\x85\x2b\x08\x77\xa9\x13\x86\x24\xa2\x42\xf0\xcc\x65\x01\xe5\x8c\x93\xfa\x1e\x0b\xea\x70\xb3\x72\xa1\x75\x0c\x08\x2e\x7d\x9e\xc3\xeb\x07\x3b\xf9\xa5\xcb\x42\xcb\x13\xc1\xd5\xba\xc5\x73\x1a\x85\x83\x55\x38\x58\x85\x43\xad\x70\xb0\x0a\xbf\x62\x98\x74\x3d\x01\x01\x5a\xd9\x15\x97\x6e\x28\x68\xc0\x40\x30\x9f\x7a\xa8\x4f\xa4\xe2\x39\x66\x46\x4d\x71\x21\x04\x47\x07\x26\xcc\xe3\x34\x74\x45\x16\xb9\xd4\x63\x21\xa9\x6f\xdc\x78\x61\x73\x8b\x02\x2a\x3c\x06\xf6\x76\x29\x02\x8f\xf2\x00\x8b\x0b\x46\x03\x87\x83\x1f\xf9\x94\xfb\x36\x9c\x4c\x05\x22\x70\x2d\x3b\xee\xc6\x58\x41\x44\x66\xb5\x0a\x58\x08\x2e\x0f\x29\x77\x85\x29\x87\x8c\x51\x6a\x9e\x6a\x6e\xa0\xbe\x59\x6e\x9a\x9b\xe5\xa6\xe6\xed\xd2\xf3\x1d\xea\xb3\x10\x42\xcf\xa5\x0c\x8d\xce\x42\xea\xa3\xd1\x2d\x49\xbc\xc7\x8c\x0b\x1a\xf8\xcc\x6e\x35\xf6\x43\x05\x23\xd9\x30\x45\x2c\x53\xa4\x66\x8a\xd4\x4c\x35\x8a\xb2\x39\x00\x63\xeb\x58\x8a\x41\x83\xf9\x86\x69\x46\x43\x1e\xe2\x72\x1b\xa2\xb9\x68\x18\x02\xf7\x68\x18\xb0\xcc\x0f\xa9\x1f\x31\x62\x6f\x32\xf2\xd1\xa7\xa0\xbe\xd9\x7c\x16\x44\x34\x74\x42\x62\x6f\xbb\xf0\x4f\x5c\x3f\xa0\x68\x59\x7b\xdb\x84\x94\x89\xeb\xbd\xc0\xbe\x7a\xf6\xfd\xb3\x1f\x9e\xfe\x76\xf5\xe2\x87\x67\xbf\x3d\x7f\xf1\xfd\xf3\xd7\x30\x01\xc6\xdb\x81\x5e\xbc\x7e\xf6\xf2\xb7\xbf\xbd\xf8\xf7\x67\x4f\x61\x02\x61\x3b\xcc\xe5\xf3\x27\x3f\xaf\x61\x02\xea\x6f\xe8\xdd\xac\xf2\x58\xa7\x45\x0e\x32\x79\xb7\xaa\xb4\x4a\xbe\x2f\xd3\xe4\xbb\x42\xeb\x62\x31\xb0\xdb\xc4\xa7\x52\xcb\x33\xf8\x63\x67\xfd\xbc\x95\x25\xe8\x42\xcb\xec\x12\x37\xd2\x30\x81\x0d\x28\x2d\x55\xb2\x8a\xd5\x60\x60\xc6\x87\x90\x6a\xb5\x18\x82\x3c\x83\xc9\xc5\x1e\x12\xbc\x4a\xa5\x57\x65\x6e\x71\xc1\x63\x03\x4c\x33\x95\xcf\xf4\x7c\x07\xf4\xc3\x10\x9c\xb3\x03\x0e\x2c\xd1\x7f\x4b\x13\x3d\x87\x09\xfc\x79\xd0\xfb\xa6\x07\x8f\xcd\xde\x5e\x3f\x2d\x16\x2f\x92\x33\x7a\x87\x63\x83\x96\x99\x69\xae\x2e\x8b\x55\xae\x61\x02\x83\x2d\x41\xfe\x5f\x8b\xc2\x1e\x6f\x0b\x67\x79\xdb\xc0\x6d\x94\x7f\x06\xa3\x6d\x86\x1e\xb5\x48\xf9\x52\xea\x39\x8d\x55\x9a\x0d\xd6\xf4\xcf\x36\xa8\x36\xc6\x5e\xcf\xfd\xd0\x62\xa7\x99\xd2\xcf\x8c\x88\x3f\x2e\xf1\xf9\xb5\x5a\x2c\x33\xa9\xd5\x20\x4d\x86\x60\x8e\x2a\x86\x5b\xfc\x0e\x21\x93\x53\x95\x55\x43\xa8\x54\x99\xaa\x6a\x08\xf2\x3e\xad\xae\xb0\xef\xe7\x42\x4b\xad\xda\x0c\x6b\xa7\x7f\x5f\x16\xab\xe5\x4b\xb9\x84\x09\xfc\xf1\xa1\x03\xa6\x75\x78\x4b\x5b\x37\x45\xf9\x4c\xc6\xf3\xc1\x20\x97\x0b\xd5\xe1\x01\x88\x6d\x86\xb4\x7e\x90\x0b\x85\x2b\x58\xbf\x15\x64\x29\x4b\x8d\x8e\x86\x98\x68\xb5\xcc\x52\x3d\xe8\xff\xd6\x3f\x3b\x80\x4d\x6f\x60\x60\x60\x1b\x53\x5d\x80\xb3\x2f\x64\x73\x6d\x93\x35\x73\xae\x9d\xb7\x07\x80\x1f\x40\x65\x95\x3a\x01\x03\x72\x76\x38\xbb\x95\xc1\x3f\xed\xaa\x98\xce\x65\xf5\xe3\x5d\xfe\x53\x59\x2c\x55\xa9\x1f\x06\x6b\xac\x67\x5d\x9c\xef\xce\xbf\x5e\x4f\x78\x0b\x13\xf8\x71\xfa\x4e\xc5\x9a\xbe\x57\x0f\xd5\x60\x17\xee\xac\xd1\xc9\x5f\xb6\x4f\x80\xda\xe2\xad\x9d\xf5\xb5\xd1\xaf\xf3\x9a\xd6\x9a\xf0\x6e\xb4\x1e\x06\x5c\xae\xee\xae\xd6\x7e\xb1\x9b\x2f\x16\x72\x79\xd4\x41\xea\xe0\x69\x57\x04\xce\x1b\xb6\x8e\xa4\x71\x91\x8f\xb7\xe5\xbc\xde\x53\xda\x9e\x34\x6f\x5b\x4c\xff\x31\xa9\x16\x0f\xb6\xfe\x7a\x5d\x14\x19\x3a\xe7\x21\x8f\x0d\xc4\x93\x2c\x1b\x77\x88\x50\xcd\x8b\xbb\x31\xe8\x72\xd5\x21\x88\x09\xea\x31\xf4\x6d\xa9\x0d\x32\xcb\xfa\xc7\x24\xde\xa9\xc8\xdb\x01\x8b\x3c\xce\xd2\xf8\xfd\x78\x93\x54\x06\x5d\x8e\xb6\x91\x14\x13\x24\xa6\xcb\xfa\xcc\x91\xa6\x79\xaa\x07\x49\x11\xaf\x16\x2a\xd7\x14\xb3\x52\xa6\xb0\xf9\xdd\xc3\x8b\x64\x90\x26\x67\x87\xa1\xb9\x8d\xaf\x30\xc9\x0b\x26\x0d\x62\x44\x60\x13\xda\xe0\xf8\x44\x2b\x9e\x4a\x0e\xf3\xce\x8e\x84\x06\x15\xb5\x26\xbe\x76\xde\xd2\x64\x27\x1d\xe1\x32\xd3\xe1\x6d\xdb\x57\x43\xeb\xda\x2c\x4b\x8d\xcf\xa3\xa5\x3a\xa7\x7d\xe8\xe6\xfe\x80\xa7\x2d\x51\x9a\x66\xe7\xe4\x46\x4d\xd5\x5a\x4d\x16\x5d\x3b\xb9\x43\xc5\x7c\x38\xf4\x84\xc5\xc3\x2f\xf9\x57\x73\xce\x55\x7e\xa2\x7b\xee\x6d\x59\xfe\xcf\x41\xbf\xbe\x83\xde\xc8\xac\xfa\x17\xf4\xd0\x47\xdd\x4f\x37\x45\x09\x83\x75\xa1\x00\x69\xbe\xb7\x00\xb6\x39\xc7\x4e\x6e\xbe\xee\x37\x8f\x66\x4a\x1f\x1e\x5b\x54\x6f\x5b\x93\x36\x9c\xee\xf8\x06\xcb\x10\x46\xa3\x26\x3d\xaf\x51\x7f\xde\xa2\x64\xb9\x7a\x7b\x72\x58\x2c\x65\xb9\x90\x43\x50\x43\xb0\x2b\xe8\xf1\x20\xb1\xea\xab\x8b\xa8\x52\x2d\x33\x19\xab\xc1\x9e\x6a\x86\xd0\xef\x0f\xc1\x3d\xee\xe4\xff\x2b\xa3\x0d\x2b\xb7\x4d\xd1\xb0\x1d\x72\x75\x21\x74\x4c\xfd\xcd\xf5\x19\x6b\x0a\x7c\xac\x0a\xfd\x38\xfe\xe3\x29\x01\x5a\x03\x72\x3d\xf2\x2f\x98\x30\xd0\x35\xe2\x62\xb1\x28\xf2\x8f\x16\x68\x72\x96\xc6\xaf\x1f\x96\xaa\x6b\x05\xd4\x66\xec\xba\x8f\x5b\x36\x0c\x8e\xa9\x2c\xf1\x56\x69\x19\xbf\xc7\x86\x4e\x33\x95\xf4\x5b\x8a\xc7\xc3\x18\x46\xaf\xfb\xb5\x28\x16\x63\xf8\xa3\x65\xb4\x54\x95\x2e\x4a\xd5\x3e\x88\x53\xdf\xa4\xea\xae\x7d\xb4\x92\xb7\xea\x49\xf5\x62\x21\x67\x66\xfa\x11\xc5\xd4\x9b\x03\x59\x55\xe9\x2c\x1f\xec\xe4\xc8\xe1\xa1\xc6\xce\xbe\x6d\xdb\xc6\x1e\xea\xa9\x4e\x86\x1d\x0a\x54\xf7\x7a\x5c\x6f\x4d\x3b\x01\x5e\xe9\x87\x6e\x0c\x60\x96\x80\x5c\xbf\x4a\x7f\x57\x63\x70\xc5\xe7\xd7\x3c\xba\x28\x32\x9d\x2e\x3b\x59\x2d\xd3\xd9\x4c\x95\x63\xe8\xe3\x36\xb9\xa3\x8e\xc1\xa1\x9f\x8a\x34\xd7\x08\xd8\xcd\xb0\xf5\x9b\x7e\x5c\x16\x55\x17\x26\xbc\xcc\xfe\xfc\x18\x1e\xbc\xa6\x32\x7e\x8f\x19\x26\x4f\x2e\x8b\xac\x40\xfe\xbe\x11\x32\x88\x42\xff\x70\xaf\xdc\x5c\xed\xf1\xdc\xa2\x12\xf8\x0c\x0b\xb0\x4f\x41\xbe\x2c\xaa\x14\xe3\x7a\x67\xd9\x42\xfd\x1d\x4b\x97\xb5\xab\x19\xb8\x4f\x21\x76\x53\x94\x0b\xa9\x8d\x69\x76\x16\x49\xb9\xa8\x3e\xb6\x38\xc6\x45\xae\x95\x39\x11\xb2\xf0\x98\xcf\xd0\xd6\x6f\x64\xb6\x52\xf0\x18\xfa\xe7\xd3\x72\x74\xd1\xad\x71\x3b\x6b\xbd\xaa\x60\xfe\xfd\xf8\x9a\x82\x84\xd3\xfa\xe8\xfe\xbc\xba\x9d\xc1\x6d\xaa\xee\xbe\x2b\xee\x27\x3d\x73\xa0\xe8\x30\x6e\x7e\x7a\x70\xab\xca\x2a\x2d\xf2\x49\xcf\xa5\x6e\x0f\xee\x17\x59\x5e\xd9\x2f\x45\xc6\xa3\xd1\xdd\xdd\x1d\xbd\xf3\x68\x51\xce\x46\xcc\x71\x9c\x51\x75\x3b\xeb\x81\x39\xfa\x9a\xf4\x5c\xd6\x83\xb9\x4a\x67\x73\x6d\xda\x17\xe7\x4b\xa9\xe7\x50\xe9\xb2\x78\xaf\x26\xbd\x7e\x73\xe6\x16\xa3\x5f\xa1\x90\x3d\xb8\x49\xb3\xac\x7d\x24\xb1\xdd\x27\xed\xa9\xcd\x5c\x7b\xe8\xf4\x83\xdd\x5e\x6f\x0a\x8e\xfa\x1d\x85\x29\x35\xce\x0c\xea\xd1\xc5\xf9\x68\x76\x71\x8e\x9c\x1f\xd1\x30\xd8\x63\x4d\x63\xa5\xc7\x13\xab\xb7\xc7\xd0\x87\x35\xaf\x1b\x7a\xd8\x3f\xde\x0c\xdc\x9e\x68\xc3\x0f\x7b\x19\x6f\xfb\xaa\x5d\xb2\x66\xa0\x1d\xee\xa4\x24\x64\x95\xd4\x15\x6a\x53\x73\xf6\x3a\x06\xa7\xdd\xc1\x71\x21\x18\xef\x1e\xa7\xb4\x03\x36\xeb\xef\xcb\x22\xc1\x3c\xb4\x58\x61\xe2\xcb\xd4\xa1\xec\x1d\x69\x72\x5a\xdc\x77\xb1\x78\xa3\xa4\x5e\xe1\x4a\xb5\xb3\x7e\x9c\x82\x78\x56\xa6\x9d\x82\x67\xea\x46\x8f\xa1\xef\xfd\xa5\x23\x5b\x96\xe8\xc4\x63\xe8\xf3\x2e\x80\x46\x73\xc7\x4f\xb2\xdb\xe7\xa2\x55\x65\x9a\x5f\xd9\x6c\xdc\x5d\xf6\x4f\x8b\x32\x51\x65\x93\x85\x4b\x95\x9c\xa4\xcf\xfb\x27\xf7\x69\x35\xde\xfb\xd0\xaa\xb9\x3e\xbe\x84\x48\xad\x66\x45\xf9\x70\x64\x15\x99\xe2\xda\x20\xcb\x87\xef\xe5\x72\x6c\x8b\xbd\x6e\x58\xeb\x41\xf5\xb9\x70\x27\x94\x39\x62\xbd\x4a\xf3\xa3\x4b\x02\x9c\xb0\x51\x6a\x2e\xac\x9f\x3e\xba\xc6\x34\x57\x6c\x15\xdc\x2b\x67\x53\x39\x60\xdc\x1d\x02\xf3\xc2\xfa\xc7\x3d\xeb\x7d\x6e\x25\xdb\xcd\xe1\xfa\x58\xfc\x63\xdc\x95\xe6\xd8\x7c\xbc\x7f\x8e\xfe\xa9\xeb\xf0\x41\x6f\xcb\x96\xef\xe1\x8b\xbc\xc6\xa4\xbc\x23\x2e\xf3\x3f\xd2\xc0\x9f\xaf\x6c\xbb\x74\x74\x54\xce\x9b\xd6\xf6\x6b\x18\xb3\x7d\xf9\x5b\x59\x2c\x2e\x5f\xbd\xd9\x7e\xfd\x92\x74\xbc\x30\xb3\x21\x07\x13\xb8\x7e\x7b\x30\x66\xc9\xb7\x8f\x65\xdb\x07\xe7\x2d\xe3\x5a\x4e\x33\x05\x13\xe8\xf5\x0e\x86\x4a\x95\x3d\x29\x4b\x98\x18\x9e\xea\xd7\x26\xbd\xff\xcc\x7b\xbb\x7b\x2c\xf3\x4a\xe2\xcf\x34\xad\x9e\x2d\x96\xfa\xc1\x6e\x11\x06\x76\xea\x19\xfc\xf5\xaf\x35\x96\xcd\x0b\x15\xb7\xad\x94\x5a\x9f\xd1\xa4\x30\x01\xe7\x5b\x48\xe1\x7c\x77\xe2\xb7\x90\x3e\x7e\xdc\x55\x84\xe1\x44\xe3\xb1\xa8\x03\x3b\xed\x3a\x7d\xdb\xbe\xc8\x5a\x79\x1f\x4f\xa0\x77\xae\xcb\x8b\x76\x5f\x69\x15\xc9\x12\xa0\xba\x4c\x17\x83\xb3\xce\x77\x2b\x0d\x3b\xc5\xf4\x9d\xd5\xdd\xce\xb4\x46\x89\xc3\xde\x91\x52\x61\xad\x8b\x77\x56\x17\xef\xe0\xbc\x46\xb7\xd6\xc5\xbb\x6e\x5d\x6c\xcb\x90\xc2\x64\xd2\xfd\x02\x6b\x1f\xfa\xdd\xc9\xd0\xb0\xf3\xa2\x0e\x26\x35\x7b\x1f\x9d\x77\xf2\x39\x04\xac\x63\xea\xfa\x1d\x10\x70\xbb\x4f\xda\xda\xae\x5c\x2e\xd4\xb8\xe6\xe9\xfa\x5d\xc7\x71\x58\xdb\x65\x17\xb7\xeb\x4f\x98\x51\xa7\x4b\xbb\xbd\x3f\x79\x56\xf5\xb0\x98\x16\x99\xdd\x0e\xb5\xef\x86\xf6\xaf\xee\xac\x75\x1a\xc4\xc9\x9a\xff\x0c\x47\x30\x89\xe9\x3a\x6d\xec\xb4\xd6\xfb\x3f\xd0\x1d\xcc\x11\xdc\x67\x91\xfc\xcc\xb5\x01\x3e\x59\x35\x9f\x16\x7f\xb0\x97\x99\xe6\x17\x69\x9e\xa8\xfb\xf3\x91\x9e\x77\xe4\xa8\x1d\xbe\x3f\x45\x91\xbb\x64\x7a\xb8\xc9\x81\xc7\xd0\x3b\x95\xd4\x17\x2a\xf0\x74\x95\x1c\xf2\xb9\x36\xf3\x89\xfc\x9e\xac\x96\x6d\x52\x49\x2b\xa9\xe4\x63\xa4\xbe\xb8\xaa\xd8\xe3\x63\xd4\xba\x38\x1d\x3b\x9c\xdb\x7a\x29\x5e\xcd\xd3\x1b\xdd\xf2\xe1\x4a\x5c\xe4\x37\xe9\x0c\x26\x5f\xfa\x19\x08\xf7\xcf\x4e\x3b\xda\xb3\x04\x5b\xb6\x88\x28\xe8\xe9\xe5\xd2\xbb\xea\x67\x25\x93\xbf\xa5\x99\xaa\x06\x37\xf8\xbb\xef\x3f\xe8\x59\x66\xa0\x5e\x21\xdb\xfc\x0b\x15\x80\x30\x30\x31\xb7\xea\xda\x69\xa9\x10\x6c\xd9\x23\x13\x85\x4b\x77\xae\xee\x00\x89\xfe\x6c\x3a\x06\x2d\x4b\x36\xd2\x1d\x69\x75\xaf\x1f\x8f\xa8\x56\x95\x36\x4c\x50\x5c\x14\x3a\x0b\x04\x8b\x9d\x16\x79\x56\xc8\x04\x79\x39\xe9\x25\xe5\x96\xb1\xac\xf9\x76\xca\xc7\xcd\x47\x4b\x43\x23\x9b\x39\xb9\x1f\x82\x9e\xa7\x15\x2d\x55\xb5\xca\x74\xf7\x29\xbc\x99\x4b\xe3\x4c\xc9\xf2\xc8\xdb\x92\x78\xef\xb0\x7d\x8b\x1f\x6a\x8d\xdc\x3d\xf7\xcf\x83\xde\x37\x4b\x55\xde\x90\x44\x69\x99\x66\xc4\x18\xbf\x77\x46\xe7\x7a\x91\x0d\x7a\xe7\x7a\x5a\x24\x0f\x17\xeb\x6f\xaf\x6a\xa4\x75\x28\x98\x48\xb0\x00\xa7\x9e\xec\x6f\x29\x19\x6f\x4f\xaa\xd7\xea\xde\x1a\xa6\xc5\x80\x47\x53\x84\xcc\x54\xa9\x07\xfd\x5f\xf2\x6a\xb5\x5c\x16\xa5\x56\x89\xd1\x6e\xcb\x77\x43\x5d\x71\x69\x5b\xfb\x7f\x55\x85\xdb\x1c\x53\x39\x4c\x7a\xe8\x3b\xa3\xb8\xaa\x7a\x9b\x3f\xb2\x42\x69\xf7\x18\xba\x29\x72\x4d\x2a\x7b\x7a\xca\x97\xf7\xdf\x1e\x0e\xde\xc8\x45\x9a\x3d\x8c\xa1\xff\x4a\xcd\x0a\x05\xbf\xbc\xe8\x0f\xe1\xb5\x9c\x17\x0b\x39\x84\xef\x55\xae\x6e\xe5\x10\xde\xa8\x32\x91\xb9\x1c\x42\x25\xf3\x8a\x60\x2c\xdf\xec\x62\x5a\xca\x24\x49\xf3\xd9\x18\x3c\x67\x9b\xc8\xd6\x77\x64\xd6\x2a\xbb\xcc\x2d\x64\x39\x4b\x73\xa2\x8b\xe5\x18\x98\xb3\xcf\x9d\x3d\xf2\x20\x71\x91\x65\x72\x59\xa9\x31\x34\xad\x23\xf8\xf5\x7c\xb8\xdf\x93\xec\x11\xb5\x68\xc7\xe0\x2e\xef\xa1\x2a\xb2\x34\x81\x6f\x12\xa5\x98\x12\xbb\xd4\x51\xbd\x44\x66\xe9\x2c\x1f\x43\xac\x72\xad\xca\x0e\x89\x29\xf3\x4b\xb5\x68\xe5\x89\x1a\xdf\x35\x59\xf1\x88\x59\x44\xab\x59\xee\x94\x3d\x76\x9a\x16\x59\xb2\x3b\x9c\xa4\xd5\x32\x93\x0f\x63\x48\x73\x2c\x1e\xc9\x34\x2b\xe2\xf7\xad\xf4\x77\x62\x27\x96\xf9\xad\xac\xf6\xf8\x98\xd7\x54\x7c\xa7\xcb\x6e\x87\xf1\x77\x44\x14\xd6\x81\xa4\x56\x04\xce\x26\x89\xb2\x2e\x8d\xf9\x68\xdb\x5c\x9d\x30\x5d\x06\x74\xba\x0d\x96\xa9\x1b\xdd\x65\xae\x7d\x6b\x81\x0d\x33\x8c\xac\x8b\x47\xe7\x23\xfb\x17\x9b\x8f\xce\x4d\x2c\xc5\x99\xac\xaa\x49\x6f\x5b\x05\x4b\x39\x53\xcd\xdf\x6d\x26\xe9\xed\x0e\x08\x86\xf8\x56\x38\x1e\x8c\x1b\x4f\xe8\x5d\x98\x9d\xcf\xab\x62\x55\xc6\x6a\x7c\x3e\x4a\xd2\xdb\xad\x29\x69\xbe\x5c\xe9\x3a\xc2\x0d\x3a\x28\xf2\x78\x2e\xf3\x99\x9a\xf4\xb6\xd7\x32\x93\xa0\xed\x82\xd6\x83\x51\xcd\xcf\x06\xd7\x01\xe9\xfa\x9c\x51\x95\xfb\xfc\xa5\xc9\xae\x7c\xd6\x4d\x7a\x17\xfb\x8c\x59\xd3\xef\x43\xdb\x84\x7c\x71\x3e\x32\x8d\x43\x36\xa6\xe5\x9a\xb9\xf9\x56\xd3\xbb\xf8\x4e\xe5\xf1\x7c\x21\xcb\xf7\xf0\xd4\x60\xaa\xce\x47\x73\xaf\xf9\x73\x58\x43\x6a\x47\x71\xfb\x5e\xb1\x2d\x85\x4d\xf2\x3b\xd6\xc6\xfd\xf9\x41\xa6\xc5\x5a\xf0\xa7\xb2\xc0\x1d\x39\xfc\x20\x17\xca\x94\x81\x2d\x50\xc9\xc5\xdf\x73\xa9\xd3\x5b\x65\x8a\xb7\x5d\xbc\xa3\x7d\xc4\x9d\x94\x36\x02\x9e\x44\x0b\x0b\xa4\xdb\x34\x56\x10\x97\x4a\xa2\x84\x5f\x44\xfc\xe9\x46\x53\xdd\x94\x5f\x2a\x59\xad\x4a\x05\xef\x3b\x38\x00\xac\xeb\xf2\xf8\xe1\x24\x4e\xfe\x44\xc8\x29\xba\xf8\xbb\x7a\xe8\x66\xa8\x83\x10\x10\xb2\x65\xed\xd1\x96\xb9\x77\xdd\xae\xd3\xd7\x4c\xc1\x93\xe6\x33\x78\x9d\x2e\xff\x49\x7e\x66\x5f\x4f\xc0\x56\xa4\x77\x8b\x7d\x39\x2f\x8a\x4a\x81\xcc\x0b\x3d\x57\x25\x5c\xbe\x7a\x63\x8b\xcd\x9b\xb2\x58\x40\x56\xc4\x32\x03\x5d\xc0\x54\x41\xa9\xf2\x44\x95\x2a\x81\x14\x33\xa7\x02\x4c\x45\xf4\x8b\xfc\xc4\xbc\x62\xec\xe6\xec\x79\x71\xab\x4a\x28\x2c\x35\x53\x62\x81\x2c\x95\x44\x76\x2a\xa5\xc0\xe6\x00\x95\xc0\x42\xe9\x32\x8d\x9b\x73\xb4\x15\x72\x69\xa6\x54\x72\xa1\xe0\xde\x1c\x59\x7f\x19\x9f\xf6\x9d\xd3\x11\x15\x66\x69\xfc\xbe\x61\xd4\x6e\x3f\x0c\x93\xf3\xe2\x6e\x34\x4f\x13\x65\xfa\x6b\x2e\x71\xcd\x1c\x4d\x65\xd9\x68\xd1\xd6\xa7\x5f\xc4\xde\xeb\xe6\xb5\x55\x6d\xf6\x27\x59\x76\x02\xaf\xdb\x1f\xd6\xda\x77\x8a\xc5\x8d\xe1\xa8\x7e\x0b\xd6\x48\x80\x00\xcd\xd0\x96\x0c\x15\x0a\x51\xfd\x23\xa4\xf8\xa5\xf9\xa6\xf2\x44\x39\x76\xbe\xc1\xec\x92\xc4\x98\xe1\x9f\x2d\x49\x6d\x0f\xf3\x9a\xf8\x04\x49\x64\xfe\x00\xfd\x99\xfd\x50\xef\x98\x41\xda\xf9\xdf\xf8\xbd\x96\xe5\x4c\xe9\xcd\x47\x83\x7a\xde\xfe\x9e\xa6\x0e\xa9\x3c\xb1\xda\x31\xf1\xff\x85\x91\xb2\x11\xfd\x2e\xd5\xf1\x1c\xec\xe7\x72\xaf\x1f\x96\x47\xc2\xbc\x06\xb5\xdc\x60\xe9\x51\xc1\x54\xe9\x3b\xa5\x72\x23\x1e\x9c\x93\x0b\x98\xca\x72\x08\x95\x96\xef\x55\x82\x6d\xd3\xb7\xca\xcd\xc7\x48\xb6\x07\x77\xff\x98\x0b\xd6\x33\x56\xf9\xa6\xa3\x73\x83\xd7\x5c\xa8\x03\xa5\x63\x7a\xc8\xdf\x57\x51\x86\xc9\xc3\xbf\x16\xc5\xa2\x5b\x09\xcf\x72\xbb\x5d\x28\xa0\x71\x66\xa8\x96\x2a\x4e\x6f\xd2\x78\x9d\xf4\xa6\x0a\x7e\x2f\x8a\x85\xc9\xc0\xa3\x62\xf5\xb5\xbc\xf4\x67\xfb\x01\x56\x37\x6f\x35\x00\xb2\x90\xe6\xa9\x4e\x65\x56\x1b\xcb\xee\x9f\x57\xa5\x5d\xaf\xe5\x8d\x56\x25\x54\xc6\x9a\xb8\xdc\x6d\x0c\x6a\xdf\x0d\x19\xe6\x71\xa0\xd1\xf6\x57\x54\xee\x9b\x54\xdd\x75\x0b\xf0\xd4\x6e\x59\xa0\x94\x77\x96\x93\x34\x87\x78\x55\x96\x2a\xd7\x5b\x41\xb0\x5a\x26\x52\xaf\x17\x1a\x23\xcc\x54\x21\xbf\x2a\x49\xb5\x4a\xbe\x56\x60\xc8\x5b\x05\xb2\x02\xf3\xd5\xda\x91\xa0\x40\xb0\x5d\x1e\x6f\x53\x75\x87\x33\x65\x0e\xe9\xe2\x94\xb5\xf7\x48\xb1\xb2\xff\xff\x47\x31\x67\x6c\xd5\xed\xcf\xe6\xc8\x05\x26\xd0\xfb\xe3\x0f\x8a\x8a\xfd\xf0\xa1\xb7\x0b\xb3\x3e\xaf\x41\xa0\x96\xc2\xfd\x10\xfa\xa4\xb3\x9f\x5e\x6f\xb8\xa1\x7f\x76\x88\xe4\xd4\x2f\x6e\xb7\xfe\x08\x6e\x08\xbd\x0c\x77\x97\x5b\xc7\x2f\x9f\x72\x16\xf4\x55\xcf\x7e\x36\xe7\x28\xe7\x23\x6b\x12\xdc\xea\xe9\x45\x76\xf1\xdf\x01\x00\x00\xff\xff\x84\x2d\x5e\x81\xbb\x47\x00\x00")

func templatesSingle_chartHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesSingle_chartHtml,
		"templates/single_chart.html",
	)
}

func templatesSingle_chartHtml() (*asset, error) {
	bytes, err := templatesSingle_chartHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/single_chart.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x88, 0xdd, 0xa4, 0x9a, 0x6d, 0xb9, 0xe5, 0xe9, 0xcd, 0xd9, 0xa3, 0x2e, 0x1d, 0x56, 0xb6, 0xd1, 0xc, 0x75, 0xd2, 0x38, 0xe8, 0xab, 0x22, 0xf3, 0xe5, 0x40, 0xde, 0x22, 0x34, 0x97, 0xc1, 0x42}}
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
	"templates/single_chart.html": templatesSingle_chartHtml,
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
	"templates": {nil, map[string]*bintree{
		"single_chart.html": {templatesSingle_chartHtml, map[string]*bintree{}},
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
