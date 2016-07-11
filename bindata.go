// Code generated by go-bindata.
// sources:
// config.yaml
// DO NOT EDIT!

package main

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

var _configYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x5a\x7b\x6f\x1b\x37\xb6\xff\xdf\x9f\x82\x55\x6e\x70\x13\xc0\x55\xe2\xdc\xbe\x60\xe4\xa6\x48\x9d\xec\xd6\x8b\xb8\x29\x12\xb7\x40\xff\x1a\x50\x1a\x8e\xc4\x7a\x86\x9c\x1d\xce\x48\x56\x3f\xfd\x9e\x07\xc9\x79\xca\x92\x9c\x02\xeb\x7b\xd1\x8d\x38\x87\xe7\x90\xe7\xc5\xdf\x39\xe4\x13\x71\xd3\x14\x8b\x5c\xbd\xfb\xd7\xd9\x13\xf1\xd3\x4e\xdc\xc8\xba\x5e\x6b\xd5\x88\x7f\x56\x5a\xad\x54\x05\xa3\x57\xb6\xdc\x55\x7a\xb5\xae\xc5\xb3\xe5\x73\xf1\xea\xe5\xc5\x77\x23\x2a\xf1\xec\xe6\xfa\x56\x7c\xd0\x4b\x65\x9c\x7a\x0e\x73\x96\xd6\x64\x7a\x35\xdf\xc9\x22\x3f\x3b\x93\xa5\x4e\xee\xd4\xce\x5d\x9e\x9d\x09\xf8\x7b\x22\xfe\xb0\xcd\x6d\xb3\x50\xe2\xed\xaf\xd7\x02\x3e\xcc\x69\x78\x67\x9b\x1a\x06\x2f\xc5\x6c\x16\xe8\x3e\xdb\xc6\xa4\x57\xb9\x6d\xd2\x3e\xe9\x13\xf1\xcb\xc7\xdb\xf7\x97\xe2\x76\x1d\x79\x08\xed\x90\x43\x25\x96\xb9\x56\xa6\x16\xd7\xef\x98\xd4\x21\x8b\x25\xb2\x60\xc6\x67\xa9\xca\x64\x93\xd7\xed\x62\xde\xf1\x00\x2c\xb9\x28\x70\x66\x6d\x05\x2c\x4d\x96\x25\x30\x4a\xe9\x97\xad\xfb\x62\xaf\x33\x14\x25\x52\x2b\x8c\xad\xc5\x56\xc2\x24\x19\xa7\x2f\x76\xc2\x8b\x38\x17\x4e\x11\x3b\x55\x94\xf5\x4e\xb8\xba\xd2\x66\x25\x9e\xcd\x66\xcf\x99\x9d\x9f\x01\xeb\xfa\x59\xe5\xb9\xfd\x4a\x5c\x0b\x59\x00\x27\x94\x27\x6e\x77\xa5\x12\x5f\xad\x55\x5e\x8a\xcc\x56\x30\x9a\x6b\x57\x0b\x9b\xd1\x2c\x69\x52\x37\x9f\x8d\x36\xb0\x96\xc6\xa8\x9c\xe8\x6b\xd0\x0c\xf0\x21\xe9\xa6\x06\x03\x35\xa5\x35\x68\x15\xa3\x96\xb5\xb6\x66\x72\x43\x5b\xed\xd6\xc3\xd9\x7e\x0a\xfe\x13\x47\x2b\x6b\xa3\xa0\x83\xfb\x63\xb2\xae\x41\xaf\x78\xf1\x38\xa9\x71\x0a\xff\xa7\xcc\xe5\x4e\xc8\x26\xd5\x56\x64\x3a\x57\x6e\x4e\x46\xad\xb7\x56\xb8\xa6\x2c\x6d\x55\x83\x0d\x96\x6b\x0b\x9e\xe5\x84\xac\x94\x98\x65\x59\x51\xaa\xd5\x4c\x20\x9b\x99\xdc\xc0\xfa\x36\x33\x96\x87\xac\x54\x95\x78\x05\x5d\x46\x52\x30\xfa\xbf\x1b\xd5\xa8\x68\xf1\x4f\x12\x54\x00\xdb\x91\xb5\x28\x1a\xd0\x2a\x98\xbb\x80\x9d\xc0\xc6\xd5\xfd\x52\xa9\x94\xcd\x0e\xdb\x59\xa1\x6b\x4b\xf8\x97\x5c\xde\x09\x77\xa7\x4b\x16\x44\xbf\x13\xfc\x9d\x54\xc8\xea\x52\xbc\x9c\x7f\xfb\x58\xe6\xb8\x6a\xb2\x6d\xcb\x3f\x0c\xed\x13\x71\x23\xef\x75\xd1\x14\x7e\x5d\x69\x43\x14\x46\x68\x03\x06\x01\x7d\x80\x6f\x88\xcf\x6c\x99\x97\x64\xce\xc6\x54\x0a\xad\xb3\x44\x65\x06\x72\x16\x55\xc8\xfb\x84\xb7\x13\xc6\x41\xd2\xa4\x1c\x27\x4a\x58\x6f\x58\xda\x43\x12\x02\x8d\x1b\x88\x70\x09\x70\x48\xc2\xd7\x4b\xf1\x6d\x14\x74\xed\x84\x5b\x37\x59\x96\xa3\x03\x29\x23\x21\x1f\xa5\x62\xbb\x56\x26\x7a\xa2\xab\x65\x55\xbb\x1f\x89\x5e\x36\xb5\x2d\x60\xad\xcb\x84\x27\xa9\x04\x57\x9d\xc9\xdc\xa9\xc0\xf0\xad\x31\x10\xf7\x4b\xe5\x55\xa4\x0d\x2c\xb2\x60\x2d\x81\x5d\x88\xa9\x5a\x69\x63\x50\x1e\xc4\x14\xfb\x1f\xae\x6c\x01\xe4\x5e\x8a\x67\x91\x18\xb5\xf5\xeb\xbf\x04\x76\x0d\xc8\x38\x6b\xe3\x28\xfa\xd4\xdb\x34\x05\x15\x38\x5e\xec\xda\x36\x79\x0a\x92\x6a\x8c\x8c\x7e\x14\xb1\x4e\x24\x53\x83\x8f\x5e\xbc\xfa\x7e\xfe\x12\xfe\xef\x22\xc6\xc8\xaf\xe0\xf4\x47\xb2\xc1\xf8\x00\x1e\xdf\x7d\xf3\xfd\xff\xfd\xd0\xce\x97\xce\x6d\x6d\x95\x92\x61\xc2\x4a\x61\x9f\x30\xdf\xa9\x6a\xa3\xaa\x51\xec\x1b\xd8\xbb\x9f\x74\x28\xa6\x03\x5d\x37\xa8\x7f\x03\xb6\x46\x16\x8a\x04\x86\xd3\x84\xc9\x1b\xff\x09\xc8\xc3\x87\x38\xed\x1f\x10\xee\xa5\xac\xd7\x3e\x19\x80\x6f\x5d\xbc\xa2\x1c\xc0\x09\xaf\x01\x2b\x19\x30\xb2\xa4\xc5\x4b\x88\x7e\x51\x81\xcd\x1c\x24\x33\x70\x0e\x9c\x30\xb9\x8f\xc0\x03\x8e\x03\x43\xd1\x76\x68\x47\xc8\x29\x81\x69\xbd\x73\x87\x35\x1f\x9c\x2f\x58\x40\x62\x8c\x41\x88\x35\x90\x87\x5a\x17\xf8\x31\xfa\xdc\xd4\x57\x38\x25\x20\x73\xe1\x39\x01\x9a\xd7\xd9\x8e\x98\x2e\x55\x55\xeb\x0c\xf7\xa6\xd0\xfd\x70\x88\x4d\x83\x5b\xf7\xec\x80\x85\xc3\xdd\x9a\xe5\x6e\x2e\xae\x6b\xdc\xd0\x02\x62\x0c\x77\x92\x2b\xb9\x01\xd7\x5e\xc3\x90\x35\xe7\x62\xd1\xd4\x22\xd5\x0e\xa3\x46\x68\x20\xe4\x64\x8e\xc9\x72\x2d\x37\xb0\x59\xcf\x50\x3b\xd7\xc0\x52\xfa\x1e\x21\x83\x60\x54\x39\xcc\xa8\x1a\x0e\x8a\x02\x0e\x14\x5d\x22\x43\x03\xa1\x67\x30\xfb\xc2\x42\xfb\xc6\x0d\xbb\x1d\xc4\x5e\xd7\xae\xdd\x8d\xa2\x59\xa6\x4c\x36\xa4\x39\xde\x74\x38\xb3\x6b\xb6\x7d\x92\x11\x1e\xec\x93\xee\xa1\xc3\x71\x02\x81\xb8\x2b\xef\xed\x72\x89\x21\x5f\xdb\x3b\x80\x3d\x38\x4d\x1b\x5d\x6b\x99\xeb\xbf\x54\xf4\x9d\xad\xae\xd7\xc8\xb6\x94\x90\x5e\xc1\x71\x01\x1e\xd0\x01\xe5\xa6\x16\x23\x7b\x0c\xd1\x1e\xc7\xad\x8b\xe7\x25\x3c\xef\x21\x47\x0e\x99\x53\xe6\xf9\xae\x9b\x58\x2a\x55\x57\xbb\xae\xd7\x76\x5d\x43\x66\x08\x20\xc0\xc3\x5a\xd7\x61\x9f\xa7\x59\x89\xcf\xd7\x21\x39\xb2\xe0\x9f\xed\x16\x12\xbf\x01\x77\xd7\x85\x72\x21\x95\x0d\x03\x8a\x24\x0f\x10\x06\x0b\xed\x0a\xf0\xd4\xb0\xb1\x8b\x97\x23\xfe\xfe\xc4\x1b\x4a\xd8\x4a\x8c\x04\xf3\xf5\x42\xd5\x5b\xa5\xba\xc8\xc7\xef\x35\x30\xed\x0a\xd2\x88\x94\x36\x12\x10\xcb\xb7\x98\xe4\xe5\x72\xdd\x62\x86\x2b\xfc\x05\x68\xd2\xac\x1c\x26\x23\x90\xb3\x23\x03\xa5\x76\x6b\x72\x2b\xc1\x48\xcc\x29\x6a\xa3\x17\x13\xf1\x24\xb5\xb5\xcc\xd9\xcb\x1d\x7a\x09\xe2\x39\x62\x9c\x6a\x50\x44\x6d\x61\x61\x70\x8a\xdf\xe8\x9f\xe2\xd1\x89\xd3\x12\xa4\x85\x45\x5d\xbc\x8a\x39\x1e\x72\x89\x4d\x29\x77\x80\x7e\x19\x6d\x78\x0d\xa8\x5c\x96\x0e\xcf\x37\x48\x25\x0a\x8d\x08\x4b\x46\x0f\x5f\x42\xd6\xc0\xcc\x99\x55\xb6\xe0\x24\x84\x82\xcf\x51\x1e\x4c\xac\xbc\x3f\xaa\xfb\x12\x56\x92\x20\xd7\x4b\xf1\xea\x9b\x3d\xf2\x82\x56\x15\xb0\x00\x74\xa6\xe0\x88\xf5\x69\x8c\x77\x93\x11\xda\x41\x4e\x29\x64\x24\x55\x38\x12\x53\x68\xd3\xd4\xca\x05\x70\x08\xb3\xfa\x1a\xf7\x68\x36\x6a\x02\x0f\xac\x1a\x37\x41\x4c\x3d\xa7\xb9\x78\x6f\x36\xba\xb2\x86\xc0\xf6\x46\x56\x1a\xf5\xcd\xc1\x42\x19\x90\xe1\x3b\x64\xf5\x54\xac\xe1\xa8\x60\x69\x51\xbd\x10\x1c\xff\xf3\xf3\xc7\x9b\xf7\x2f\xe6\xc4\xf4\x45\x41\x19\x2d\xfd\x13\x41\xe2\xc6\xe6\x4d\xa1\x46\x75\x01\x0f\x7b\x3e\x3c\x86\x68\x2c\xda\xe2\x83\xdd\x62\x5e\x66\x32\x01\x91\x05\xbf\x53\x26\xcf\xe9\x13\x52\xbf\xbc\x88\x9e\x0b\x95\xd4\x3e\xfa\x35\x7f\xc3\x09\x3f\xc0\x82\x64\x0a\x2a\x6b\x0b\x95\xf7\xe4\x5a\x82\x47\x7f\x1c\xa6\x0f\x3a\x0e\xe0\xff\x7d\xa6\x20\xf7\x3b\x17\x18\x22\x1e\x10\x83\x1a\x0d\xaa\x46\xdd\x43\xd6\xf6\xa9\x08\x3f\xb7\x47\xe9\x64\x24\x7f\xf0\x75\x07\x89\x15\x78\x98\x0f\x53\x17\x9d\x4d\x18\xc7\x58\xce\x10\xbe\x5d\x7b\x90\x45\xd4\x68\x7a\x5a\x1c\xa1\x5c\x3a\x64\xda\x73\xdc\x12\xc4\xf3\xfc\x7c\xbe\x71\x1e\x3e\xeb\xa2\xb4\x48\xe6\x70\xe5\x78\x82\xfa\x95\xfb\xa5\xc4\x42\x88\x66\x93\xa8\x4b\xfa\x27\xfe\x7d\x2d\x66\x9f\x1b\x80\x9b\x88\x4d\x66\x84\xd8\x98\xb8\x8d\xe7\x35\x24\xe4\x25\x55\x46\x8e\x21\x79\xaa\x9c\x5e\x19\x3c\x2f\x02\x31\xc7\x8a\x41\xe8\x98\x8b\x5a\xdd\x03\xa2\x87\x34\x2b\x57\x43\x0d\x7c\x34\x90\x4d\xad\x51\x58\xf1\x78\xa6\xcf\x70\xfb\x99\xae\x5c\xfd\x1c\xb5\x83\x32\x3c\x80\xaa\x54\xa6\xef\xc1\x0d\xbf\xf2\x49\x1a\x85\x59\x93\x04\xce\xed\x16\x8c\x0d\xb0\x59\x55\x95\xad\x60\xca\x2d\x3a\x34\x1f\x0b\x36\x80\x72\xcd\x18\x99\x0a\x1c\x28\x09\xc3\x64\xf4\x6e\xc8\xd3\x89\xc7\x05\x69\xe4\x71\xc5\x1f\x28\x23\x34\x55\x05\x41\x04\x6b\x0f\x54\x6d\x4d\xf9\x93\x02\xb7\x44\xa2\xb6\xf0\xa4\xf0\x0e\x9a\x69\x8b\x33\xf0\xa2\x08\x0b\xc4\x7b\x4a\x08\xde\xdf\xd6\xd2\x79\x6e\xf5\xba\x52\xca\xf7\x04\xa0\xcc\x40\x2f\xb6\x25\x26\x63\xbf\xdd\x27\x10\x06\x5a\x3a\xd8\xbd\x78\x1b\xe5\xb1\xf3\x90\x27\x78\xcf\x0d\x96\x0a\x7e\xd0\x59\xd1\x3c\x82\x9c\x84\xbc\x83\x7d\x58\xfc\x3f\x98\x05\x4f\x3a\x0a\x19\x62\x33\x31\xf7\x9c\x83\x05\x88\x21\x1c\xc8\x8c\xd3\x74\x41\x06\x38\xca\xb2\xd2\x25\x17\x4b\xef\xda\x1f\x98\x81\xb7\x26\x16\xd0\x41\x0d\xb1\x8e\xa1\x62\x3e\x8c\x82\x6a\x43\x20\x06\xbe\xd1\x05\xc4\xef\x90\xd6\x6c\xe3\xe2\x88\x2f\x27\xa1\x52\x5e\xe0\x01\x89\xfd\x06\xb4\x4c\xd7\x25\x3b\x79\xdd\xaf\x16\x71\x79\xd6\xf8\x76\x40\x25\x8d\xcb\x09\x4a\x7b\x61\xed\x1f\xa3\x09\xc2\x2f\x16\xe6\x57\x22\x97\x66\xd5\x90\x97\x8b\x77\x16\x5d\x1c\x02\xb6\xb0\x80\x38\x23\x25\xae\x86\x0a\x28\x82\x37\x62\xf6\x74\x26\x9e\xb9\x06\x4c\x0f\xcb\x9a\x3d\x75\xb3\x73\xf8\x6f\x0a\xff\x55\xf5\x72\xfe\x7c\x24\x30\x1c\x9f\xae\x59\xb8\x5a\xd7\x94\x8b\x88\x0f\x80\x4f\x3a\x5e\x52\x59\xcb\xb9\xf8\x84\x42\x09\xa6\x42\x4e\x6c\x85\x6f\x75\x9e\x83\x85\xa8\x7d\xd0\xb6\x29\x0a\xed\x16\x0a\xe0\xae\x8a\xf5\x55\x1b\x48\xc1\xb7\xce\x3a\x6b\xc0\x04\x01\x44\xb3\xd1\x58\x3b\xd2\xba\x12\x1f\xe5\x61\xbc\x67\xfe\x19\xd4\x7d\x2e\xb6\x08\x6c\x5b\x20\xb3\x3d\x24\x98\x27\xd5\x70\x08\xeb\x5a\x05\x80\x33\x0c\xd5\x71\xe4\xfb\xe8\x6f\xaa\x3c\x86\xed\x5b\xf1\xdb\xa7\x0f\xb1\xa1\x80\xd1\x47\xdd\x29\x52\x1b\x32\x85\xbd\x44\xc3\xcf\x86\x8c\xe0\x80\xd5\xe9\x30\x99\xfc\x62\x05\x8d\x87\x44\xb2\xc5\xdc\x92\x61\xb7\xac\xe5\x5a\x56\x60\x01\xcc\xe8\x20\xfc\x99\x7b\x3e\xe0\xec\x19\xd6\xd6\x26\x39\x80\x8c\xc8\xf9\x0f\x6c\xc3\xd1\x47\x98\xc3\x7c\x95\x26\xcf\x02\x52\x81\xa4\x54\x6d\x40\x8c\xe1\x04\x61\x97\x94\x88\x30\x50\x10\x17\x81\x4c\x04\xb3\xde\xf0\xc5\x5c\xfc\x62\x5b\x66\x68\x61\x50\x00\x04\x14\x6c\x58\x0d\xb7\x0a\xb1\xeb\x9b\x19\xf4\x15\x96\xf2\x7a\xf1\xe6\xa9\x7b\xfd\x62\xf1\x86\xe9\x05\xfc\xbe\xa0\x9f\x6c\xaf\xae\x45\x2e\x5f\x2f\xaa\x37\xaf\x35\xd1\xeb\x37\x6c\x3e\x70\xe5\x9e\x00\x44\x9c\x41\x8f\x0f\x88\x78\x9a\xb6\x32\xdc\x3e\xb3\x93\x6d\x00\xdb\x0d\xb4\x48\x1c\x61\x21\x43\x2e\x4b\x02\x77\x78\x0a\x2e\x94\x97\x94\x36\xe4\x53\x5e\x8b\x15\x7c\x88\x61\xc1\xd8\x34\xa8\x3b\xa4\x75\x98\x66\x20\x65\x1c\x15\x19\x48\x38\x8e\x0e\x33\x15\x1e\x84\x16\xbe\x38\x3a\x38\x2b\x08\x94\x4b\x48\x6f\xdf\xc9\xf6\x24\x6c\x03\x8f\x03\x9e\x13\xd3\x24\xc0\x33\x6d\x14\x77\x3e\x80\x6a\xee\x4f\x58\x44\x7a\x04\xa1\x0f\x6e\x3c\x92\x8e\xb6\xbe\x74\x27\x6e\xfd\x63\x53\x97\x4d\xcd\x0b\xec\x01\xfe\x16\x26\x33\xd4\xc7\x82\x7d\xd9\x9e\xca\x1e\x84\x1d\x4c\x10\xfe\xf4\xf6\xb5\x01\x62\x83\x30\x34\x25\xc9\x91\x5f\xce\x5f\x6d\x50\x22\xfa\x55\xf0\x09\x3f\x87\x2c\x74\x84\x7e\x3a\xd4\x63\x15\xf1\x47\x2c\x38\xf6\x7d\x3b\x35\xbb\x06\x25\xf6\x1a\x78\x0b\xdb\x30\xbc\x0c\xfb\x0d\x4d\xbe\xd6\x5f\x50\xa7\x78\x92\xab\x7b\xea\x43\x1e\xab\x4b\xd6\x42\x5f\x99\x9e\x39\x28\x30\xe4\x86\x73\x1f\x7f\x80\x7c\x62\xf0\x07\x75\xc2\x32\x21\x7b\xdd\xe9\xf2\xb0\x2e\x23\xe9\x48\x59\xd9\xa9\xbe\x76\x5d\x50\x20\xd5\x0a\xb0\x0e\x72\x74\x63\xf5\x1c\xd4\x41\xdb\xd4\x2e\x29\xaf\x8d\x75\x00\x78\x8e\x73\x2f\xae\x5c\x2f\xbc\xac\xf2\x90\x26\x62\xc3\xf7\x78\x8d\x84\x29\x13\x9a\x29\xff\x56\xd5\xc4\x76\xf6\x11\xc7\x71\xec\xca\x77\xe0\xf8\xd8\x4b\x30\x43\x97\xb2\xe2\x92\x69\x8a\x3f\xfe\xf5\x1a\xfc\x63\x75\xc7\x2c\x79\x9a\xc6\x11\x5f\x1e\x56\x32\x52\x8d\xf4\xba\x7e\x6c\x60\xb6\x85\x5d\xff\x6a\xea\x61\x6d\x06\xc2\x64\xad\x64\xaa\xaa\xf6\xcc\xbb\x0a\x65\x1a\xee\x0b\xc7\xfa\x2b\xa5\x85\x25\x7b\x67\xbf\xa5\xba\x70\x82\x07\x31\xf9\xd3\x6a\x53\x1c\x71\x06\x30\xdd\x48\x45\x38\x7c\xa2\xef\xdd\x00\x68\x76\xb1\x3a\x82\x04\x05\xa7\x35\xdf\x51\x7a\x43\x87\x1b\x3b\x9d\xb1\xdf\x80\xdd\xa9\xd9\x4d\xd7\x11\xd8\xfe\xb0\x85\xa2\x34\x06\x86\x38\xa8\x54\x02\xef\xb0\xac\x4a\x25\x39\x35\x8e\x75\x07\x93\x61\x1d\x4c\xc5\xa2\x34\x0c\xf2\x83\x68\xc4\x09\x91\x9c\x90\xf4\x10\xa4\x80\xc6\x71\xd1\x49\x7b\x9d\x47\xf7\x94\x06\xeb\x43\xe3\xf7\xc3\x9f\x42\x49\x7b\x07\x10\xfd\xb0\x9e\x91\x6a\xa4\xe5\xbb\x13\x55\xfc\xb9\xb6\x3e\xa4\xa9\xaf\x88\x5d\x8e\x5c\x41\xad\x03\x10\xc2\x0d\x5b\x6b\x21\x4e\x70\xbb\xfe\x42\xe7\xe0\x22\x5b\xda\xd1\x52\xe9\x8a\x0b\xfb\x81\x93\x5f\xc6\x83\x8f\x0d\xb1\x7e\x05\x1e\xe0\x60\xac\xdd\xf7\xc0\xa4\x69\x1f\xd1\x86\x6b\x01\xec\xbb\xad\x54\xd5\x96\x17\x26\x7c\x12\xfe\x93\xd8\x4a\x17\xeb\x8c\x29\xe0\x4f\x4e\xa6\x3d\x60\xf5\x60\xf5\xf2\xc0\x21\xd9\x89\x46\x2c\x28\x0f\xab\x1f\xa9\x46\x9a\x2c\x1e\x15\x86\xc1\x47\x28\x0a\xf1\x07\xc7\x65\x0c\x84\x58\xeb\x6c\x00\x91\xca\x6a\xd5\x60\x83\xf1\x98\x73\xc1\x33\x48\x02\x83\x4e\xcd\x06\xeb\x00\x15\x31\x6c\x09\x72\x46\x35\x1c\xc6\x9c\xdd\xc4\x6a\x76\xa0\xeb\xc0\x1d\xaf\x91\x00\xa1\x10\xa0\xe9\x9d\x40\x71\xdd\x41\x40\xbc\x70\x22\xda\x01\x3b\x94\x94\x40\xa9\x8e\x25\x56\xd6\xe4\x5c\xad\x71\x59\xd5\x8e\x82\x57\x21\x5d\xda\x2d\xb0\x3b\xa7\x8d\x60\x0b\x22\x06\x3f\x12\x36\x46\xd2\x91\x2d\xf1\xcb\x24\x60\xec\xd7\x1f\x7f\x07\x5a\xa4\x9a\xe1\xef\x85\x8a\x09\x76\x97\x1e\xc6\x03\x28\x87\x7a\x50\x63\xc9\xc3\x62\x10\xd6\xd7\x43\xa0\xdd\x05\x1f\x09\x3f\xa1\xa0\xe4\x66\xf9\x11\x36\x09\xa4\x63\xd5\x2f\xbf\xa0\xd4\x01\xb6\x0b\xbc\x4e\xcc\x62\xa6\xe2\xe6\x3d\xde\x84\x6a\x77\xf7\xc8\x62\x07\x0b\x65\xbf\xb1\x6e\x1b\xb4\xcd\x82\x6d\xbd\x4c\xb7\x04\x7c\x71\x90\x06\x75\xd3\xd4\x8e\x8e\x8e\xcd\xfe\x91\x74\xac\xa3\xa6\x98\xce\xfd\x8f\xaf\x71\xa6\xb5\xf7\xb8\x3c\x1f\x3b\x21\x51\x5d\xbd\x7e\xef\xa0\x0d\xf2\x80\x53\x96\x79\x53\xc9\x3c\xbe\x7f\x38\xa0\xfb\xe9\x9e\x34\x31\x2c\xb1\x35\x71\x58\xe3\x44\x76\xaa\x06\x7f\x95\xd4\x0a\xe8\xbf\xe2\x38\x26\x75\xd3\x8c\x18\xbf\xef\x7d\x93\x0a\xdb\x69\xc4\x0a\xdb\xdf\x79\x05\x20\x73\xc7\xcb\x4f\xcf\x05\xf7\x76\x8f\xed\xc2\xc7\x8d\xf7\x1b\x45\x08\xeb\x79\x78\xbc\x66\x9a\x1b\x2e\x48\x0e\xeb\x2b\x50\x8e\xfc\xb0\x52\xab\x13\xa3\xf8\x93\x67\xd5\x1e\x95\x7c\x39\x13\x5e\x8d\x1c\xd2\xa7\x57\x55\xd2\xde\xee\x44\xcd\xf2\xc3\x36\xaf\xca\xd1\xed\xcf\x58\x40\x57\x07\xa4\xbb\x88\x38\x1f\x98\xec\x35\x87\x97\xb1\xc7\xe8\x0d\xe9\xc6\x5a\x3b\x59\x67\xc8\xc6\xd7\x94\xe1\x7a\x83\xce\x1d\x7a\x6f\x70\x48\x65\xbc\x8a\xb6\xfe\x1b\x71\x68\x2b\xc0\xde\xe9\x1c\xe6\xb5\xbb\x76\xea\x88\xfa\x9a\xc8\x26\x3c\xe5\xe4\x4d\x03\x1b\xd7\x39\x41\x17\x3b\xbe\x21\xa0\xda\x25\xcf\xc3\xb9\x4a\xd7\xb5\x87\x54\x40\xb4\x09\x6f\x60\x18\x23\x34\x3a\x4e\x25\x30\xdc\x1c\x53\xc8\x31\xdd\xa9\xc9\xe4\x13\xcd\x3a\x39\x9b\x9c\x90\x4a\xb8\xca\x7b\x4c\x2e\xe1\x1d\x8d\x93\x89\x1f\xdf\x93\x4d\x40\x89\xe1\xa9\xe9\x41\x9d\xb5\xb4\xe3\x16\xde\x9e\x71\x77\x2a\x5c\xf8\x1c\xbc\x27\x3c\x99\x05\x60\x40\x6f\x37\x53\x0f\x79\x6c\xac\x99\xff\xd7\xc5\x27\x64\xd4\x2d\xa5\xe1\xa3\xda\x0b\x88\xd1\xf8\xda\xaa\x8d\x2e\x96\xd6\x7d\xe0\xba\x2f\xbc\x68\xde\x10\x89\x7b\xae\x88\xb3\x57\x8f\xe0\xea\xe7\x85\xfb\x80\xcc\xe2\xcd\x3f\x15\x50\x78\xcd\xc0\x96\xe2\xf7\x8c\x47\x98\x89\x09\xc7\xb6\x80\xf1\x89\xc1\x53\x03\x1c\xea\x68\x5b\xe8\xbf\x7c\xd5\xf4\x65\x50\x04\x2a\x91\x44\x19\xdb\xac\xd6\x0f\x5d\x67\x43\xb5\x42\x34\x53\x41\xd0\xbd\xf2\x95\x41\x47\x03\xe3\xf8\xd1\x60\x15\x0e\x04\x9e\xdd\x5a\xc3\xd3\xc4\xb8\x38\xaa\x4d\x3b\xd9\xa1\x9d\x6c\xd0\x3e\x08\x51\x72\x49\x2f\x98\xc5\xc6\xf2\x4d\x20\xb2\x7d\x44\x97\x36\x1c\xb2\xc8\x26\xed\xde\xb8\x71\x0d\x17\x72\x0c\x7d\xee\x88\xc1\x42\x64\xc0\x1f\xff\x88\x6c\x94\x4d\x86\x93\xc7\x6b\x8c\xea\x3b\xbe\xb7\xfb\x60\x5b\x77\xba\xab\xfb\x65\xfa\xfc\x2f\xb5\x76\x1f\x6f\xa0\x3d\x0c\x4f\xb5\x51\x87\x0d\xf1\xa9\xed\x6a\x95\xab\xa3\x33\x4b\x8f\x7c\x64\xa8\xf6\xeb\xd4\xa7\xe9\xf1\x93\xd3\xcf\x2d\x0b\x69\x5f\x2b\x86\x07\xf8\xf1\x89\xb8\x35\x2f\x6c\x96\x1d\xbe\xd2\x20\x46\x69\x02\xb4\xd8\x8e\x89\xec\x5a\x46\x31\x37\x78\x52\xd1\x67\xdb\x63\x62\x8e\xe6\x61\x50\xf7\xc4\x04\x90\xa9\xa3\x57\xe2\x87\xd4\xee\x09\x47\xda\xdb\x7c\x49\x35\x19\x5c\xc2\x33\xef\xbd\xe0\x3d\xa4\xbb\xb0\xf2\xf6\xe1\x76\x3b\x14\x5d\xd1\xbb\x58\x78\x38\x77\x70\x93\x44\x37\xde\xa3\x3d\xb9\x37\x7a\x45\x67\x29\xef\xd2\x3f\xa4\xd3\x19\x3e\x29\x08\x8d\x3b\x8c\xd7\xd0\x19\x03\xa4\x35\xa5\x14\x9e\x46\xdd\xf0\xad\x3e\xa2\xbf\x0e\xa1\xef\xba\x2d\x75\xcc\x11\x95\x82\x33\xc6\x61\x28\x7a\x76\xbd\x6b\x7a\x9c\x31\x7e\xa4\xd0\xd4\xe0\x90\x49\x85\x1b\x88\xbc\x7e\xa7\xd9\x2e\x36\x07\xc3\x03\x4b\xda\x9f\xcc\xf1\xdd\x36\xdf\xdf\x66\xfc\xd2\xc0\xa4\xdd\xdf\x43\xa0\xe2\x5b\x54\xde\x2c\xfd\x6c\x16\xb4\xe5\x1e\x60\xc0\x34\x1d\xa0\xd3\xcf\x3d\x11\xc8\xb4\xca\xf7\x0d\xc2\xc8\xee\x3f\x01\x00\x00\xff\xff\x03\xe4\x68\x65\x56\x36\x00\x00")

func configYamlBytes() ([]byte, error) {
	return bindataRead(
		_configYaml,
		"config.yaml",
	)
}

func configYaml() (*asset, error) {
	bytes, err := configYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config.yaml", size: 13910, mode: os.FileMode(420), modTime: time.Unix(1468209663, 0)}
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
	"config.yaml": configYaml,
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
	"config.yaml": &bintree{configYaml, map[string]*bintree{}},
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

