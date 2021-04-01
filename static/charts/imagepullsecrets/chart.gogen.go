// Code generated by vfsgen; DO NOT EDIT.

package imagepullsecrets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Chart statically implements the virtual filesystem provided to vfsgen.
var Chart = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Time{},
		},
		"/.helmignore": &vfsgen۰CompressedFileInfo{
			name:             ".helmignore",
			modTime:          time.Time{},
			uncompressedSize: 349,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x4c\x8e\x41\x6e\xe3\x30\x0c\x45\xf7\x3c\xc5\x1f\x78\x33\x63\x0c\xe4\x43\x24\xb3\x98\x55\x0b\xa4\xc8\xb6\x90\x6d\x46\x62\x22\x8b\x82\x44\x27\x6d\x17\x3d\x7b\x91\x04\x41\xbb\x79\x20\x3f\xc8\x8f\xd7\xe1\xd9\x9b\x71\xcd\x0d\xa6\x90\x90\xb5\x32\x2e\x91\x33\xc6\x55\xd2\x2c\x39\xa0\xf8\xe9\xe4\x03\x37\x47\x1d\x5e\xa2\x34\xb4\xb5\x14\xad\xd6\xd0\x22\xa7\x84\x90\x74\xc4\xe2\x6d\x8a\x92\xc3\x5f\x54\x4e\xde\xe4\xcc\x28\xde\xe2\x8f\xdc\xe7\x99\x3a\x64\x0e\xde\x44\x33\x7e\x97\xca\x07\x79\xe3\x19\x17\xb1\x88\x5f\x7f\x1c\x9e\x72\x7a\x87\xe6\xdb\xe7\x55\x09\x85\x2b\x92\x64\x76\xe4\xb6\xbb\xd7\x9d\x69\x65\xea\xb0\xd1\x65\xd1\x8c\xfd\x66\x87\x59\x6a\x23\x17\xc4\x86\x1b\xef\xfa\xe4\xc6\x8f\x3a\xdc\xf8\x08\x62\x18\xae\x78\xac\xed\x9c\x87\xef\xa2\xd1\x4f\xa7\xb5\xe0\x20\x89\x1b\xf5\xae\x5d\x0a\xf5\x6e\xf4\x27\xea\x9d\x2d\xd7\x59\xab\x04\xea\x3f\xa9\xc3\xde\x57\xd1\xb5\xe1\xff\xf6\x5f\x23\x57\xaa\x1e\x79\x32\x72\x32\xb3\x1f\xee\xe7\x55\x8f\xe4\xce\x6d\xd2\x99\x07\xfa\x0a\x00\x00\xff\xff\x16\xec\x32\x27\x5d\x01\x00\x00"),
		},
		"/Chart.yaml": &vfsgen۰CompressedFileInfo{
			name:             "Chart.yaml",
			modTime:          time.Time{},
			uncompressedSize: 276,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x5c\x8f\x41\x4e\xc5\x30\x0c\x44\xf7\x3e\x85\x2f\x40\xcb\x67\x99\x15\x02\xce\xf0\xf7\x6e\xea\xdf\x5a\x24\x4e\x14\x27\x95\xca\xe9\x51\x5a\x2a\x01\xdb\xf1\x9b\x19\x0f\x65\xb9\x73\x31\x49\xea\x70\xbb\x81\x52\x64\x87\x12\x69\xe1\xdc\x42\x30\xf6\x85\xab\xc1\xcc\xe6\x8b\xe4\x7a\x60\x1f\xbb\x52\x14\x4f\x21\xec\x18\x49\x69\x61\xec\x2c\xfe\xc0\xf8\x48\xe5\x4c\x30\x58\x53\x8f\x5b\x6b\xcd\xe6\xc6\x71\x22\xfd\x22\xf1\x21\xb5\x79\xf0\x29\x02\x44\x12\xad\x24\xca\xc5\x1c\x3c\x21\x47\x92\xe0\x50\xf4\x91\x5e\xff\xb3\x88\xe7\x6b\x6f\x87\x8e\xef\xfd\x00\x96\x5a\xf1\x7c\x78\xaf\x8e\x45\xea\xda\xa6\x6e\xf9\x5d\x37\x4e\xe4\x3f\x77\x2a\xb3\x01\x6c\xd7\xdc\xe7\xe1\x36\xbc\x00\xe5\x7c\xff\xab\xc0\x77\x00\x00\x00\xff\xff\x01\x13\x31\xf0\x14\x01\x00\x00"),
		},
		"/crds": &vfsgen۰DirInfo{
			name:    "crds",
			modTime: time.Time{},
		},
		"/crds/crds.yaml": &vfsgen۰CompressedFileInfo{
			name:             "crds.yaml",
			modTime:          time.Time{},
			uncompressedSize: 17682,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x5c\x4b\x6f\xe4\xb8\xf1\xbf\xf7\xa7\x28\xf8\x7f\x98\x7f\x00\xb7\xbc\x93\xb9\x04\x7d\x33\x3c\x03\xc4\xc9\xc6\x6b\xb8\x8d\xd9\x43\x90\x03\x5b\xaa\x56\x33\xa6\x48\x2d\x49\xf5\x4c\xef\x62\xbf\x7b\x50\x24\xf5\x7e\xda\x3b\x8b\x4d\x00\xf5\xc5\x10\x1f\xc5\x7a\xb1\xea\x57\xa2\x29\x96\xf3\xcf\xa8\x0d\x57\x72\x07\x2c\xe7\xf8\xd5\xa2\xa4\x27\x13\xbd\xfc\xc5\x44\x5c\xdd\x9c\xdf\x1f\xd0\xb2\xf7\x9b\x17\x2e\x93\x1d\xdc\x15\xc6\xaa\xec\x09\x8d\x2a\x74\x8c\x1f\xf1\xc8\x25\xb7\x5c\xc9\x4d\x86\x96\x25\xcc\xb2\xdd\x06\x80\x49\xa9\x2c\xa3\x66\x43\x8f\x00\xb1\x92\x56\x2b\x21\x50\x6f\x53\x94\xd1\x4b\x71\xc0\x43\xc1\x45\x82\xda\xad\x50\xae\x7f\xfe\x2e\xfa\x10\x7d\xb7\x01\x88\x35\xba\xe9\xcf\x3c\x43\x63\x59\x96\xef\x40\x16\x42\x6c\x00\x24\xcb\x70\x07\x3c\x63\x29\xe6\x85\x10\x06\x63\x8d\xd6\x44\xae\xc1\x44\x07\x26\x7f\x66\x3c\x16\xaa\x48\x22\xae\x36\x26\xc7\xd8\xf1\x93\x24\x8e\x49\x26\x1e\x35\x97\x16\xf5\x9d\x12\x45\xe6\x99\xdb\xc2\xdf\xf6\x3f\x3c\x3c\x32\x7b\xda\x41\x64\x2c\xb3\x85\x09\x7f\x1c\xe7\x09\x9a\x58\xf3\xdc\x3a\xfe\x9e\x30\xd7\x68\x50\x5a\x03\xfc\x08\xf6\x84\xa0\x0e\xff\xc6\xd8\xc2\x89\x19\x38\x20\x4a\x30\x45\x1c\xa3\x31\xc7\x42\x88\x0b\x68\x8c\x95\x8c\xb9\xc0\xc4\x91\x3a\x2a\x9d\x31\xbb\x83\xc3\xc5\xa2\x6b\xf0\xb2\xec\x2d\x0b\xcf\xf6\x92\xe3\x0e\x8c\xd5\x5c\xa6\x23\x9c\x09\x66\xec\xbe\x5a\xe3\x29\x2c\xc0\x9d\xb2\xfa\xfc\xfe\x78\x42\xf9\x66\x36\x13\xd6\x62\xf3\xa9\x3d\xca\xf3\x1a\xc6\x0c\x71\x7a\x66\x82\x27\xdc\x5e\xf6\x34\x2f\x19\x50\xe6\x5f\xd5\x17\x10\x4a\xa6\x8e\xc1\x14\x25\x6a\x66\x31\x21\xd3\x27\x28\x2d\x67\x02\xb8\x01\x47\x85\x58\x02\x76\xb4\xa8\xdd\x58\xd2\x41\xc5\x74\x43\xf4\x92\x71\x2e\xed\x87\x3f\x37\x38\xff\x1c\x38\x01\xd3\x60\xc5\xf3\x4f\xce\x90\xa2\xee\x89\x90\x63\x1c\x59\xa6\x53\xb4\x91\xf7\xb0\x88\x48\xf5\x65\x78\x60\x19\x82\xf2\xae\xe0\x07\xd6\x92\x4c\xda\xdc\x8f\x7d\x28\x89\x2e\xb0\x7c\xc6\x24\x4b\x31\xa1\x29\x26\x67\x31\x9a\x69\x6e\x64\x35\xae\xc9\x1c\x37\x0d\x4d\x73\x39\xc5\x62\x67\xa1\x0e\x87\xa9\x56\x45\x1e\xf6\x61\x6f\xdb\x79\x12\x61\xeb\xfb\xb0\x71\x4f\x03\x1f\x0b\x21\xbc\xe4\xae\x47\x70\x63\xff\x3e\xd4\xfb\x3d\x37\x7e\x44\x2e\x0a\xcd\x44\x7f\xbb\xbb\x4e\x73\x52\xda\x3e\xd4\x0b\x6d\x81\x67\x79\xe8\xe2\x32\x2d\x04\xd3\xbd\x99\x1b\x00\x13\x2b\x92\xe4\x4e\x14\xc6\x3a\xcb\x9b\xe2\xa0\x43\x38\x0b\x94\xbc\xc6\x77\xf0\xcb\xaf\x1b\xf0\x1e\xe8\x9c\xcc\x77\xaa\x1c\xe5\xed\xe3\xfd\xe7\x0f\xfb\xf8\x84\x19\xf3\x8d\x1d\x43\x74\xe4\x21\xb5\x93\x11\xfc\x0c\xe7\xcd\xf4\xd8\x95\x0a\x6e\x1f\xef\x03\xb5\x5c\xab\x1c\xb5\xe5\x25\x47\xf4\x6b\x04\xea\xaa\xad\xb3\xee\x3b\x62\xcc\x8f\x81\x84\x42\x73\xb0\x7e\x08\xb0\x98\x80\xf1\x2c\x38\x27\xe1\x06\x74\x19\xd0\xea\x5d\x54\xfe\xd4\x11\x98\x0c\x81\x23\x82\x3d\x6a\x22\x42\x4a\x2f\x44\x42\xf1\xfc\x8c\xda\xef\xc2\x54\xf2\x9f\x2b\xca\x06\xac\x0a\x9b\xd4\x62\xb0\x62\xf9\x73\x91\x57\x32\x41\x2a\x2d\xf0\x1a\x98\x4c\x20\x63\x14\x7f\x5c\x70\x2a\x64\x83\x9a\x1b\x62\x22\xf8\x87\xd2\x08\x5c\x1e\xd5\x0e\x4e\xd6\xe6\x66\x77\x73\x93\x72\x5b\xa6\xa6\x58\x65\x59\x21\xb9\xbd\xdc\xb8\x04\xc3\x0f\x85\x55\xda\xdc\x24\x78\x46\x71\x63\x78\xba\x65\x3a\x3e\x71\x8b\xb1\x2d\x34\xde\xb0\x9c\x6f\x1d\xe3\xd2\x65\xa6\x28\x4b\xfe\xaf\x32\xfc\xbb\x06\xa7\x1d\x57\xf7\x3f\xe7\xc6\xa3\x7a\x27\x37\x26\x2b\xb3\x30\xcd\xf3\x5f\xab\x97\x87\x30\xf7\xf4\x69\xff\x0c\xe5\xa2\xce\x04\x6d\x9d\xfb\x30\x5d\x4d\x33\xb5\xe2\x49\x51\x5c\x1e\x5d\x08\xe4\x06\x8e\x5a\x65\x8e\x22\xca\x24\x57\x5c\x5a\xf7\x10\x0b\x8e\xb2\xad\x74\x53\x1c\x32\x6e\xc9\xd2\x3f\x15\x68\x2c\xd9\x27\x82\x3b\x97\xa0\xe1\x80\x50\xe4\x14\xc2\x93\x08\xee\x25\xdc\xb1\x0c\xc5\x1d\x33\xf8\xbb\xab\x9d\x34\x6c\xb6\xa4\xd2\x79\xc5\x37\x71\x45\x7b\xa0\xd7\x56\xd5\x5c\x26\xfb\x41\x0b\x75\x76\xe4\x3e\xc7\xb8\xb5\x43\x12\x34\x5c\x93\x17\x53\x32\x26\xdf\x1f\x0a\x58\xe3\x7b\x93\x7e\x1a\x53\x6e\xac\xbe\xb4\x5b\x7b\x08\xc2\x0f\x72\x88\x88\x71\x59\xae\x6e\x19\x17\xa6\x93\x49\xac\x22\x0b\x39\x28\x14\xb2\x49\x7b\x37\x01\xb2\xf8\x54\x07\xfa\xce\x88\x31\x36\x1d\x1c\xab\x72\xec\x40\x67\x87\xe3\xbb\x7a\xac\x53\x31\x3f\x72\x34\xf0\xe5\xc4\xe3\x53\x9b\xd1\xc2\x60\x02\xcc\x09\x34\x40\x13\x20\x38\x3d\x45\xbf\x44\xc5\x2f\xa8\x41\xa8\x94\xcb\x26\x33\x03\xf3\xb8\xc5\x6c\x90\xc9\x69\x09\xfd\xcf\x65\xb2\x91\xbe\x11\x77\xeb\x13\x70\xca\xfd\x0d\x54\x68\xe7\x91\x6f\x0d\x93\xd8\x42\x85\x2c\x86\xbb\x86\x6c\x3b\xb1\x0b\xba\xdd\x4c\x6b\x76\xd9\x2c\xe3\x68\x3b\x61\x8b\xd1\xc5\x3c\x4c\x9a\x74\xfa\x67\x37\xa4\xe5\x3d\xcc\x96\xa9\xe4\x80\x15\x5e\x69\xfb\xff\x90\xd7\x06\xdc\x02\x0c\x52\x7e\x46\xf9\x26\xe7\xaf\x91\xd1\xac\xef\xd7\x18\x28\x30\xef\xf6\xad\xaf\x25\x0c\x28\xd9\x47\x5a\xcc\x0e\xfb\xbe\x97\xf5\xc4\xce\x5e\x5a\xaf\x91\xfd\x10\x62\x7c\x9d\x8f\xf7\x2a\xad\x19\x89\xf6\x28\x30\xa6\x68\x5d\x09\xe4\xf2\x46\x25\xd4\x75\xd8\xd9\x4c\x23\x64\xcc\xc6\xa7\x41\xae\xc2\xda\x29\x05\x30\xdb\x55\x82\x60\x07\x14\x0e\x07\x24\x18\xf3\x04\x7d\xa9\xc4\xcd\x44\x48\x1d\x54\xd6\x01\x81\xe5\xb9\xe0\x98\x94\x98\xa2\x63\xf3\x6b\xa2\x9c\x15\xc2\xf2\x5c\x90\xd3\x04\xc9\xc6\xf9\xd5\x58\xf9\x60\x42\x73\x95\x44\x4a\xda\x5e\xce\x16\x4e\xfe\xc2\x85\x20\x06\x02\xe8\x1e\x25\xf9\xff\x3f\x3c\xfd\x69\xa4\x73\x22\x70\x2d\x0b\x5e\x2e\xf9\x11\x6b\xb7\xf3\x36\x0e\x02\x36\xca\xdc\x25\xd4\x17\x06\xc1\x45\xd1\xa6\xc5\xf1\xa7\xaf\x04\x60\xcc\x3c\xc7\x33\x4a\x1a\x70\xe0\x5b\xef\x5f\x95\xb9\xcb\x60\x96\xa1\xb4\x01\x81\x85\x9e\x19\xaa\xe0\x76\x6b\x9d\x87\x3d\xe2\xbc\x06\x06\x2f\x78\xf1\xe0\x94\xf0\x6f\x4e\x5b\xd3\x01\x76\x66\x67\x29\x6a\x74\xb0\xd7\x79\xd2\x0b\x5e\x1c\x91\x80\x64\x67\xe6\xe6\x8b\x0d\x06\x44\x79\x7e\x50\x47\x6d\xc4\x4d\xa8\x42\xbc\xfe\xa8\xc1\x29\xc0\x7b\xfd\x42\x95\xf9\x3a\x84\x36\xa5\x03\x92\x0b\xc6\x2f\xf6\xaf\x50\x5b\x39\x6d\xbf\x5a\xbc\xca\x4c\x35\x74\xf6\x86\x7c\x67\xbc\x51\xc8\x15\x4f\x3c\x5f\x24\xa0\x55\xce\x8b\x2c\xa5\xa3\xb2\x0e\x71\xef\x10\xaa\x65\x8c\x8b\x25\xf7\xf2\x1a\x1e\x94\xbd\x97\xd7\x8b\xc8\x7e\xfa\xca\x09\x7f\x93\x4f\x7c\x54\x68\x1e\x94\x75\x2d\xdf\x5c\x89\x9e\xe5\x57\xab\xd0\x4f\x73\x5b\x48\x7a\xd4\x40\xf2\x37\xcb\x19\x13\x2d\x92\xf3\x3e\xbc\x19\x2b\x4d\x42\x91\x5f\x82\xd2\x41\x57\xbe\x20\xf5\x8b\x0d\x81\x93\x91\xb0\x52\x18\x57\xaf\x48\x25\xb7\x98\xe5\xf6\x12\x0d\xad\x13\x54\xbc\xd0\x91\x9b\x56\xe8\xb3\x55\x2d\xe9\x97\x5b\x44\xf1\x99\xd2\x9c\x9f\xed\x8b\x6b\xc1\x62\x4c\x20\x29\x9c\x12\x5d\x71\xc8\x2c\xa6\x3c\x86\x0c\x75\x8a\x8b\x68\xe6\x14\x50\x97\x2c\xbf\x28\x96\xbe\xc1\x9f\xa6\x90\xe4\xeb\xb0\x6e\x13\x69\xbe\xe0\x65\x76\x4c\x69\xda\x45\x29\x6c\x36\x33\x2d\x93\x63\x01\xb1\x39\x32\x1e\x00\x2d\x82\x63\xdf\x7b\xac\xb4\x62\xb1\x3f\x14\x8b\x4d\x03\x0c\x07\x2a\x7c\xdb\x4f\x05\xea\x0b\xa8\x33\x6a\x9f\x21\x26\xbc\x4d\x1d\xab\xf7\x3c\x26\x82\xe7\x13\xd2\x63\x21\x5c\x52\x71\x42\x06\xd3\xfb\x97\x60\x6d\xcc\x34\x41\x96\xf4\x76\xfb\xf0\x11\x93\x08\x6e\xa5\x0f\x4d\x5d\x7e\x4b\x15\x32\x21\x82\x1f\x4f\x06\xee\x5b\x77\xb2\x32\x46\x44\xaa\x79\x1a\xaf\x80\xb2\x8b\x81\x61\xcb\x24\xdd\xa9\xc1\x24\xdc\x38\x6d\xb6\x39\x5f\x16\xa0\x32\xff\x72\xcd\x9b\xa5\x6e\x69\xa8\x77\x85\xad\x2b\x6c\x5d\x61\xeb\x0a\x5b\x57\xd8\xba\xc2\xd6\xff\x01\xd8\x0a\x4d\x50\xf1\x07\xbe\x25\xea\xe7\xed\x80\x73\x5c\x8e\xcb\x58\x4e\xbb\xf3\x17\x4a\x55\xce\x69\x7f\x85\x9c\x71\x3d\xbb\x43\x6f\xdd\x89\xae\xc0\xd6\x4c\xee\x5f\xbf\x36\x17\x21\xfa\xdc\x00\x59\xf3\xcc\x44\xf7\x28\x6c\x38\x64\x4a\x40\xe1\xd3\x70\x89\xcc\x1a\x48\x83\x20\xb8\x32\x3e\x2b\x1e\x39\x8a\x04\xb8\x99\xa1\x79\xf5\x82\x97\xab\xeb\xde\x1e\xbf\xba\x97\x57\x3e\x3d\xf7\x76\x6c\x99\xcb\x67\x08\x2b\x29\x2e\x70\xe5\x66\x5e\xbd\x1d\xba\x2c\xf2\xba\x6f\x50\x04\x35\xfe\x03\xe0\xb5\x2f\xd9\x29\x4b\xd7\x4e\xda\x2c\x67\x0e\x97\xf1\xd3\x92\x60\xcf\xf2\x65\x7a\xeb\xbf\x1e\xa4\x55\x6f\xab\x0c\x16\xf8\xfd\xb4\x22\x26\x35\x59\x4b\xf6\x23\xb7\xa7\x47\x95\xcc\x9f\x48\xd0\xa0\xe5\xe5\x62\x59\x16\x0e\x1f\x1f\x9d\x10\x72\x22\xd7\x3b\x49\x9c\xaf\x17\x43\x5d\x38\x48\x77\xac\x56\xcc\x55\xf2\xce\xbc\xad\x62\x7c\x7b\xb5\x38\x56\x15\xae\xe7\x2a\xeb\xb9\xca\x7a\xae\xb2\x16\xa8\x6b\x81\xba\x16\xa8\x6b\x81\xba\x16\xa8\x6b\x81\xba\x9e\xab\xac\x58\x6c\x3d\x57\x59\xcf\x55\xd6\x73\x95\x15\xb6\xae\xb0\x75\x85\xad\x2b\x6c\x5d\x61\xeb\x0a\x5b\xd7\x73\x95\xf5\x5c\xe5\xbf\xe8\x5c\x65\x92\x80\xc7\xf5\xb3\x47\x08\xcd\xbb\x07\x77\x4a\x1e\x79\x1a\xfa\x0f\x01\xaa\xd4\xe0\x23\x5c\xc8\x18\xbe\xd2\x10\x2e\x2d\x36\x6e\x65\xb8\x53\x04\x0f\x1e\x30\x99\xbc\xbc\xf2\x0d\x5e\xbf\xbf\x6e\xbb\x2d\xd8\x68\x6d\x5c\x59\x33\x30\x7c\x2a\xd5\xe0\x10\x46\x4f\x97\xdc\x61\x48\x5e\xd8\xf2\xaa\x48\x38\x3c\x59\xe0\x00\x13\x43\xa6\xcb\xe0\xdf\x57\x2d\xad\xba\xba\xa3\x91\xba\x26\xae\x84\x1e\x5f\xf7\x5b\x29\x63\xea\x86\x57\xff\x6e\x74\xcd\xf7\xf0\x95\xa3\xf2\x02\x26\x97\x86\x27\xe3\xf9\x98\x2a\xc0\x1a\x27\xcf\xb8\xfa\x02\x35\x4f\xa5\xc9\xd1\xfb\x61\x13\xaa\x19\xbf\xdf\x35\x78\xb3\x6a\x84\xd2\x10\x95\x6d\x75\xcb\xb1\xd5\xe8\xef\x80\xcd\x5e\xd0\xf4\xd7\xaa\x97\x5e\xd1\x74\xa3\x5b\x97\x34\xd5\xc1\xa0\x3e\xff\xb6\x5b\x9a\x53\x1f\x50\xe8\xea\xab\xf9\x25\x84\xad\xe5\x3d\x33\x8c\x9a\xb5\x77\x57\xbf\x4b\x79\x04\x9c\x4d\xf8\xc9\x58\x5e\xe8\x2b\x75\x92\x50\xe7\xab\x0c\x63\x12\xd7\x9f\x50\xe8\xd2\xac\xbf\x98\x30\x61\xeb\x4e\x53\xfd\x7d\x8f\xf7\x4c\xe4\x27\xf6\xbe\x6e\x0b\x9f\xe0\xf0\x9f\x1c\x68\x74\x53\x7e\x21\x53\xef\xc0\xea\x02\xc3\xad\x7c\xa5\x59\x8a\xa1\xa5\x96\x9b\xc5\x31\xe6\x36\x68\xbb\xf9\xcd\x81\xab\xab\xd6\x47\x04\xdc\x63\xfd\x16\x70\x07\xff\xfc\xd7\xc6\x53\xc5\xe4\x73\xc9\x0d\x35\xfe\x27\x00\x00\xff\xff\x03\x0f\xf4\xdc\x12\x45\x00\x00"),
		},
		"/examples": &vfsgen۰DirInfo{
			name:    "examples",
			modTime: time.Time{},
		},
		"/examples/test.yaml": &vfsgen۰CompressedFileInfo{
			name:             "test.yaml",
			modTime:          time.Time{},
			uncompressedSize: 728,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x94\x92\x3f\x4f\xc3\x30\x10\xc5\x77\x7f\x8a\x93\x77\x17\x85\x3f\x8b\x67\x3a\xb0\xa0\x4a\x48\x30\x20\x86\x23\x39\x25\x56\x9d\xd8\xf2\x5d\xaa\x96\x4f\x8f\x12\x27\x21\xa8\x50\xc1\x14\xcb\xef\x3d\xff\x9e\x2f\x36\xc6\x28\x8c\xee\x99\x12\xbb\xd0\x59\x38\x14\x6a\xef\xba\xca\xc2\x23\xb6\xc4\x11\x4b\x52\x2d\x09\x56\x28\x68\x15\x40\x87\x2d\x59\x70\x6d\x64\x23\xc4\xa2\xd4\xaf\xf9\x27\x2a\x13\xc9\xa5\xb0\x29\xa6\xbd\x11\xb3\x3e\x95\x25\xb9\xae\xbe\x9f\x62\xc3\x96\x05\x3d\x7c\x8a\xeb\x9b\xdb\x3b\x7d\x06\x75\x2d\xd6\xc4\x9b\x77\xec\x3e\xd0\x95\x3e\xf4\xd5\xc6\x85\xab\x43\x81\x3e\x36\x38\x17\x7a\x18\x4c\xbb\xde\xfb\x3f\x35\xe3\x48\xe5\x08\xc7\x54\x93\x0c\x2b\x00\x1e\x83\x79\x3d\x67\xf4\xc8\x36\xb1\xf7\xde\x64\xdd\xe4\x1a\x5a\xcd\xa6\xf1\x7a\xbc\x8e\xb1\x85\x57\xbd\xd0\xf4\xdb\x24\x31\x79\x2a\x25\xa4\xc5\x0b\x60\xa0\x45\x29\x9b\xed\x31\x26\xe2\xe1\xb2\x2b\x2d\xeb\x7b\x3a\x59\xd0\x99\x69\xb8\x09\xbd\xaf\x4c\x83\x07\x9a\xda\xb0\xfe\xe6\x07\x08\x91\x12\x4a\x48\x16\xf4\xf6\xe8\x78\x36\x7c\x15\x7d\x71\xd2\xec\x42\xb5\x80\x2e\x57\xf8\x6f\x81\x1f\xf1\x89\x6a\xc7\x92\x4e\xf9\xe0\x32\x51\x45\x9d\x38\xf4\x7c\x36\xeb\xe5\x07\xe9\xf5\x38\xf3\xfb\x59\x4d\x54\xa9\xcf\x00\x00\x00\xff\xff\x34\x19\xfb\x89\xd8\x02\x00\x00"),
		},
		"/templates": &vfsgen۰DirInfo{
			name:    "templates",
			modTime: time.Time{},
		},
		"/templates/_helpers.tpl": &vfsgen۰CompressedFileInfo{
			name:             "_helpers.tpl",
			modTime:          time.Time{},
			uncompressedSize: 2469,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xac\x55\x4d\x6f\xdb\x38\x10\xbd\xeb\x57\x0c\x84\x0d\xbc\x9b\x85\xe8\xc3\x02\x7b\x30\x90\x43\x91\xf6\x50\xb4\x48\x8b\x06\x48\x0f\x45\x61\x50\xd4\xc8\x22\x42\x51\x0c\x87\x74\xe3\x26\xfe\xef\x05\x29\x5a\x91\x1d\xdb\x75\x8a\xdc\x68\xcd\x9b\x99\x37\x6f\x3e\xfc\xf0\x30\x3d\x87\xa5\x6c\x67\x40\xe8\xa0\x96\x0a\xdd\xca\xe0\x45\xeb\xc9\x71\xd1\xe0\x0c\xce\xa7\xeb\x75\x16\x50\xd9\xbb\x7b\xc3\x75\x05\xae\x41\xd0\xbc\x45\xe8\xea\xf8\x16\x0d\xb7\x8e\x65\x09\x57\x40\x85\xb5\xd4\x08\xb9\x6c\xf9\x02\x8d\x57\x8a\x50\x58\x74\x2c\xf8\xe4\x50\x3c\xa1\xb8\x57\x0e\xd8\x65\x74\xbf\x0a\x01\xd9\x0d\x57\x1e\x29\x22\x3f\x2d\xd1\x5a\x59\x21\x3c\x82\xb3\x5e\x0b\xf8\xff\xbf\xf8\x94\xed\xb5\xaf\x6b\x79\x0f\x79\x91\x43\x8a\x85\xba\x0a\xcf\x9e\xe5\xa5\x45\xee\x10\xf8\x90\xa1\xf6\x4a\xad\xe0\xce\x73\x25\x6b\x89\x15\x70\x63\x22\x7f\x96\x7d\xc5\x3e\x76\xc4\xbb\x90\x21\xd4\x42\x50\xa2\xe0\x9e\x10\xa8\x6b\x11\x3e\xf8\x12\xad\x46\x87\xd4\x57\x5d\x4b\x54\x15\x01\xb7\x08\x4a\xb6\xd2\x61\x05\xae\x03\xd7\x48\x82\xbf\xcb\x55\x54\xe4\xed\xd5\x75\xc0\x4a\xbd\x00\x32\x28\xfe\x61\xd9\xfb\x1a\x2c\x2a\xe4\x94\xa4\x13\x9d\x76\x5c\x6a\xea\xc5\xeb\xbf\x49\x07\x3f\xa4\x52\x50\x22\x78\x0a\x3c\x09\x78\x24\x9f\xd8\xfe\x56\xe0\x80\xdd\x16\x59\xd6\x83\xa6\x1b\xe3\xa0\x6b\x82\x1c\xb4\x9f\xa2\xbb\xa2\x21\xce\x5f\xb1\x86\xd9\xc5\xe9\x8d\x7d\xe2\x38\xa8\xd1\x07\x61\x5f\x7a\xa9\x7a\xdf\x0d\xcf\xad\x8f\x2f\x24\x67\xac\xd4\xae\x86\xfc\x8c\x8a\x33\xca\x77\x62\xf5\x49\x4f\x1f\xb3\xfd\xcf\xad\xe1\x1b\x75\x35\x6c\xcc\x12\x2d\xc9\x4e\x87\x8e\xc6\xce\xa6\x31\xe9\x51\x8a\x97\xa8\x4e\xe8\x6e\x44\x3f\xb5\x76\xb7\xa4\xb1\xda\xfd\xfb\x26\x65\x7d\x04\x8b\x46\x71\x81\x90\xff\x9b\x43\x3e\xcf\x5f\xb2\x52\xc7\x28\x15\xa1\x6f\xb6\x53\x0a\xed\xb3\xd9\x03\xa9\x85\xf2\xd5\xd1\x29\x65\xb0\x5e\x8f\x62\xec\xa8\x79\x52\xda\xd3\x52\xbe\x5a\xba\xd8\x2b\x8a\x22\x71\x63\x66\x70\x24\xe9\x7e\x71\x58\x72\x65\xb7\xc3\x55\x61\xb2\x9b\x06\xe3\xc9\xd1\x46\x91\x4a\x2e\x6e\x57\xdc\x56\xc4\x4a\xae\x7f\x72\x29\x54\xe7\xab\x10\x50\x28\x4f\x0e\x6d\x31\x04\xb6\x78\xe7\xa5\xc5\x0a\x72\x96\x6c\x71\x54\xc2\x99\x0f\x07\x27\x9c\xa9\x78\x1d\xf3\x61\x57\xc7\xb0\xf5\x3a\x6b\x50\xb5\x8c\x9a\x69\x1c\xc3\xa3\x5c\x37\x93\x7a\xa0\xd6\x96\x6b\xbe\xc0\xaa\x28\x57\x31\xca\xb0\x8b\xd7\x68\x97\x52\xe0\x7e\x27\xa9\xc9\x71\x2d\x70\xdb\x65\xc3\xed\x39\x3e\x6d\x5c\x0f\xef\xd7\xe1\x8d\x31\x87\x37\x62\x6f\x10\xd1\xb5\xa6\xd3\xa8\xdd\x0c\x0e\xb7\x63\x8f\x9f\xe1\xd6\x15\x5d\x7d\x5c\xa4\x51\x17\xff\x60\x10\x09\x15\x0a\xd7\xd9\x8f\x69\x20\x8b\x57\x1d\xab\x97\xea\xbf\x73\x30\xa6\xe7\x59\x06\xf0\xd9\x76\x2d\xba\x06\x3d\x41\xe9\xa5\x72\x85\xd4\x94\x3d\x3f\x73\x66\x80\xb1\x16\x9d\x95\x82\x58\x8d\x15\x5a\xee\xd2\x62\x17\x9b\x66\xb5\xb1\x13\x93\x15\xd2\x24\x03\x70\xdc\x2e\xd0\xc5\xfa\x67\x30\x19\x16\x61\xbe\x71\x9e\xec\xde\xe9\x23\x09\xd3\x5f\x56\x91\x7e\x6f\xf2\x52\xe7\xad\xc0\x5e\xe2\x19\x7c\x9b\xcf\x93\x7d\x6e\xb8\x6b\xe6\xf3\xef\xbb\x2c\xc6\xe6\xac\xd8\xb6\xa5\x6d\xca\x60\xbb\x9a\x3c\xc8\xb9\x77\xe1\xf2\x31\xff\x5f\x01\x00\x00\xff\xff\xf9\xf5\xdb\x31\xa5\x09\x00\x00"),
		},
		"/templates/default_imps_cr.yaml": &vfsgen۰CompressedFileInfo{
			name:             "default_imps_cr.yaml",
			modTime:          time.Time{},
			uncompressedSize: 762,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x94\x90\xcd\x4e\xeb\x40\x0c\x85\xf7\x79\x0a\xab\xfb\xcc\x55\xa5\xbb\xb8\xca\xf6\xae\xd8\xa0\x0a\xa4\x22\x96\xee\x8c\x93\x5a\x38\x33\xd1\xfc\x80\x20\xe4\xdd\xd1\x74\x12\x08\x52\xf9\xe9\xce\xb2\x7d\xce\x77\xec\x71\xac\x81\x5b\x50\x7b\x94\x44\x41\x19\x6a\x31\x49\xfc\xef\x6c\xcb\x9d\x22\x8b\x07\x21\x03\xd3\x54\xd5\x75\x5d\xe1\xc0\x7b\xf2\x81\x9d\x6d\x80\x7b\xec\x28\xa8\x03\xda\x17\x64\x2d\x2e\x19\xc5\xee\xcf\xe3\x16\x65\x38\xe2\xb6\x7a\x60\x6b\x1a\xb8\xca\x4b\xbb\x24\x72\x4b\xda\x53\xac\x7a\x8a\x68\x30\x62\x53\x01\x58\xec\xa9\x81\x71\x04\xb6\x5a\x92\x21\xd8\x9c\x2c\x87\x24\x12\x4e\xdb\xb5\x76\x36\x7a\x27\x42\x5e\xb5\x49\x24\x0b\x36\xa0\x60\x9a\xea\x39\xe5\xec\x12\x06\xd4\xc5\x4a\xdd\x90\x10\x06\x52\xd7\x4b\x3b\x47\x07\x10\x3c\x90\x84\x4c\x3d\x9d\xfb\x0b\x60\x51\x64\xdc\x2b\x58\xb6\x86\x6c\x84\xbf\xd9\x2c\x0c\xa4\xb3\x51\x44\xdf\x51\xcc\x15\x40\x91\x97\x7a\x75\xd8\xf9\x9f\x16\x61\x79\x48\x8e\x59\x12\x96\x64\x4f\x1c\x8f\x5f\xc8\xde\x0f\x0d\x8b\xe0\xa3\xd3\xcc\x06\x10\xdd\x3d\xf6\xf2\x29\xf4\xbf\xb5\x3f\x59\x73\x31\xee\x8e\xe3\x71\xe7\xcc\x19\xec\x32\xb9\x1c\xff\x03\x5c\x7b\xca\x62\x46\x99\xa9\x9e\x3a\x0e\xd1\x3f\x17\xd2\x6a\xbc\xfc\xfc\x5b\xf8\x0a\x5d\xad\xea\xb7\x00\x00\x00\xff\xff\x0f\x9c\x61\x79\xfa\x02\x00\x00"),
		},
		"/templates/deployment.yaml": &vfsgen۰CompressedFileInfo{
			name:             "deployment.yaml",
			modTime:          time.Time{},
			uncompressedSize: 2774,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xbc\x55\x4d\x8f\xdb\x36\x10\xbd\xef\xaf\x20\x7c\x97\x36\x29\x8a\x62\x21\x20\x87\xc5\x36\x08\x0a\x38\xa9\xd1\x2d\x16\xe8\xa9\x18\x93\x63\x9b\xc8\x88\x43\x90\x23\xb5\xee\xc6\xff\xbd\xa0\x24\xcb\xa2\x6c\x27\xdb\x6e\x51\x9d\xec\xf9\x7a\x8f\x9c\x99\x47\xf0\xf6\x09\x43\xb4\xec\x2a\x05\xde\xc7\xdb\xf6\xed\xcd\x67\xeb\x4c\xa5\x7e\x44\x4f\xbc\xaf\xd1\xc9\x4d\x8d\x02\x06\x04\xaa\x1b\xa5\x1c\xd4\x58\xa9\xe7\x67\x65\x9d\xa6\xc6\xa0\x5a\xd8\x1a\xb6\xe8\x1b\xa2\x88\x3a\xa0\x14\x9a\x9d\x04\x26\xc2\x50\x6e\x1a\xa2\x94\xb0\x50\xa5\x3a\x1c\x86\xec\xe8\x41\xf7\x25\xca\x5f\x90\x10\x22\x96\x9f\x8e\xe6\x3e\x8a\x60\x8d\x14\x13\x9a\x52\xcf\xcf\xc5\x8b\xa0\xfa\x9c\x04\xf4\x45\x39\xeb\x0c\x3a\x51\xdf\xa7\x72\xd1\xa3\x4e\xa5\x02\x7a\xb2\x1a\x62\x8f\xfc\x04\xd4\x60\x2c\x8f\xc6\x1e\x37\x22\xa1\x16\x0e\x3d\x72\x0d\xa2\x77\xcb\x09\x95\x97\x93\x39\x16\x5a\x5e\x20\xf5\x43\x8f\x25\x58\x7b\x02\xc1\x01\x6b\x72\xc3\xe9\x03\xe7\x58\x40\x2c\xbb\x11\x5b\xa9\x35\xe8\xcf\x7b\x08\x26\x96\x6b\x70\x7f\x81\xd5\xc4\x8d\x29\x2d\xdf\x46\x1d\xc0\x5b\xb7\x2d\xac\x13\x0c\x2d\x50\xa5\x16\x93\x43\x46\x0c\xad\xd5\xf8\x91\x9d\x15\x0e\x65\x17\x8d\x3f\x0d\xa1\xea\x70\x58\x8c\x08\xdd\xf9\x36\x63\xa2\xc1\x16\x89\x7d\x9a\x81\x8f\x6c\xb0\x44\x07\x6b\x42\xd3\x1f\xe0\x1b\x9c\xd8\x63\x00\xe1\x50\x84\xbe\xc7\x39\xa5\xa3\x77\x98\xbd\x33\x12\x48\x11\xff\x25\x4c\x94\x44\x32\xaf\xe6\x32\xce\xb3\x53\xda\x28\x96\xcb\x80\xad\x1d\xa8\x8c\x81\xbd\x67\x86\x18\xb0\xcd\x46\xe8\x7a\xfa\x0c\x3a\xfd\xfd\xc3\xca\x6e\xcc\xf4\x6c\xee\x4f\x7d\x9e\x67\x0a\xff\x06\x35\x65\xa3\x73\x37\x2d\x95\x1d\x8a\xb2\x31\x7d\xdd\xd6\x8c\x28\xc7\xcd\xe9\x7e\xa3\x6e\x82\x95\xfd\x03\x3b\xc1\x3f\xa5\xba\x48\xf5\x74\xae\xc7\x3c\xfc\x52\xf5\x54\xb3\x9b\xcb\x7b\xad\xb9\x71\xf2\xe9\x55\xc2\x92\xbe\xe4\x07\xeb\x30\x4c\x2e\xa2\x18\x04\xeb\x94\x3b\xba\x52\x42\x5d\x83\x33\xd5\xc4\x94\x32\x6e\x6b\x70\xb0\xcd\x22\x93\xb9\x28\x6a\x94\x60\x75\x2c\xc0\x98\xf0\xae\xba\x7b\x73\xf7\xe6\x2c\xa4\x5f\x91\x82\x10\x0c\x86\xa2\x93\x01\xcb\xee\x2c\x4c\xb3\xdb\xd8\x6d\x31\xca\xe1\xbb\xaf\xa8\xe1\xf1\xbb\xda\x82\x2b\x6d\x88\x57\x7b\xf0\xf6\xbb\xbc\x72\x77\xd1\xf9\x82\x76\xa6\x24\x8f\x1c\x93\x68\xec\xd5\xe1\x50\x9d\xb9\x05\xb6\xea\x8b\x32\xb8\x81\x86\x44\x95\x0f\x3b\x08\x52\xde\x7b\x7f\x69\xab\x07\x94\x55\x43\xb4\x62\xb2\x7a\x9f\x2f\x51\x57\xcf\x8f\xce\x9c\x9f\xe7\x20\x71\xde\xa5\xbe\xaf\x43\x4b\x32\xdf\x64\x12\x56\x1c\xa4\x52\x67\x9d\x52\xca\x07\x16\xd6\x4c\x95\xfa\xf5\x61\x35\xf1\x91\x6d\xd1\x61\x8c\xab\xc0\x6b\xcc\x21\x77\x22\xfe\x03\xce\xee\x5d\x29\x0f\xb2\xab\xd4\xed\x65\x22\xbe\xc3\x3f\xf7\x05\x04\x63\xff\x17\x9c\xc8\x4d\xd0\x18\xbf\x39\x2d\x63\xe4\x57\xe6\x04\x5d\x3b\xad\x33\xd3\x51\xe2\x6d\x49\xe9\xc5\xc8\x93\x4e\xbd\x5a\xfe\xfc\xe1\xf7\xe5\xfb\xa7\xf7\xcb\x19\xf9\x36\xe5\xe7\xe3\x37\xad\xb5\x98\x41\xce\xd4\xfc\x82\xb0\xa2\x6b\xe7\x1c\xae\x28\xea\xfc\x84\x67\xf5\xcf\x6a\x3b\x36\xf8\x38\x3c\xef\xa7\xb0\xa9\xb5\x7a\xb1\x8c\xbf\x04\x0f\x36\x1b\xeb\xac\x4c\x36\xe2\x68\xf9\x6f\x71\x84\x29\xbd\xa4\xf9\x4b\x34\x31\xbe\x12\x2d\x07\x1b\xa5\xe0\xb1\x53\xf6\x09\xe2\xdc\xf3\xcf\x60\x8f\xa8\x7f\x07\x00\x00\xff\xff\xcf\x82\xe1\x1e\xd6\x0a\x00\x00"),
		},
		"/templates/poddistruptionbudget.yaml": &vfsgen۰CompressedFileInfo{
			name:             "poddistruptionbudget.yaml",
			modTime:          time.Time{},
			uncompressedSize: 431,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x8c\x91\xbd\x6a\x2b\x31\x10\x46\xfb\x7d\x8a\xc1\xbd\x75\x31\x5c\x52\x6c\x97\x90\xd2\x84\x90\xc2\xfd\xac\xf4\xd9\x19\x32\x3b\x12\xfa\x31\x84\x8d\xdf\x3d\xc8\x8b\x21\x90\x14\x6e\x85\xce\x39\x23\xcd\xb2\x6c\x49\x8e\xe4\x0e\xac\x0d\xc5\xa5\x18\x9e\xa5\xe4\x96\xaa\x44\x7b\x6a\xe1\x84\xea\x60\x3c\x29\x02\x5d\x2e\x03\x27\x39\x20\x17\x89\x36\x52\x8a\x2a\xfe\xf3\xdf\x79\x37\xa1\xf2\x6e\xf8\x10\x0b\x23\xbd\xfe\xe6\x87\x19\x95\x03\x57\x1e\x07\x22\xe3\x19\x23\x2d\x0b\x89\x79\x6d\x01\xb4\x91\x99\x4f\x48\x4d\xb5\xc0\x67\xd4\xad\x8f\x56\x73\x54\x45\x76\xc7\xa6\xda\x81\x0d\xb9\x1e\x5f\xe9\x92\xd8\xaf\x0a\xf7\x06\x05\x17\xb8\x97\xdb\xf1\x7a\x4b\x79\x82\x96\x5e\x23\xba\x3e\xef\x8e\xd4\xca\xf4\xd0\x17\x99\x58\x80\x55\xfa\xdf\x75\x25\xc1\x77\xd5\x2c\xf6\x78\x66\xd1\xfe\x15\x23\xed\x06\xa2\x02\x85\xaf\x31\xaf\xa1\x99\xab\x7f\xdf\xff\x28\xdf\xdf\xbe\x89\xf6\x7f\xcc\xf0\xd0\x67\xe8\x22\xd8\x75\x01\xdf\x01\x00\x00\xff\xff\x4c\xf7\xc1\xf2\xaf\x01\x00\x00"),
		},
		"/templates/rbac.yaml": &vfsgen۰CompressedFileInfo{
			name:             "rbac.yaml",
			modTime:          time.Time{},
			uncompressedSize: 1383,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xcc\x54\x4d\x6b\x1c\x31\x0c\xbd\xcf\xaf\x10\x7b\x2c\x8c\x4b\xa1\x87\x32\xb7\xb6\x87\xde\x7a\xd8\x42\xa0\x94\x1e\x34\xb6\x76\x57\x8d\x46\x1e\x2c\x7b\x03\xd9\xe6\xbf\x97\xb1\x93\xec\x26\xfd\x60\x21\xa1\xe4\x34\xb2\x90\xde\x7b\x7e\xf3\x30\xce\x7c\x41\xc9\x38\xea\x00\xfb\x37\xdd\x25\x6b\x18\xe0\x0b\xa5\x3d\x7b\x7a\xef\x7d\x2c\x9a\xbb\x89\x32\x06\xcc\x38\x74\x00\x8a\x13\x0d\x70\x38\x00\xab\x97\x12\x08\x56\x3c\xe1\x96\xe6\x22\x62\xe4\x13\xe5\xde\x47\xcd\x29\x8a\x50\x72\x9b\x22\xb2\x2c\xac\xc0\xc1\xcd\xcd\xed\xb6\xcd\xe8\x1b\x84\x5b\x93\x10\x1a\xb9\xcf\x77\xed\x36\x25\x38\x92\xd8\xc2\x06\x70\x38\xf4\x67\x51\xb5\x9d\x85\xe8\x27\x28\x6b\x20\xcd\xf0\xb6\xc1\x2d\x10\x57\x9c\x77\xe0\x2e\x50\x0a\x99\xb3\x07\xd7\x73\xa8\x1a\x33\x66\x8e\x6a\x6d\xe1\xa4\x71\x14\x91\xe3\x57\x9c\xe4\x6f\xf8\xa4\x61\x39\x74\x7d\xdf\x77\xa7\x96\xa6\x11\xbd\xc3\x92\x77\x31\xf1\x75\x85\x74\x97\xef\xcc\x71\x7c\x7d\x6f\xf6\x47\x29\x96\x29\xad\xa3\xd0\x33\x3a\xfd\xac\x1e\xa6\x22\x64\x43\xd7\x03\xce\xfc\x29\xc5\x32\xdb\x00\xdf\x1a\x92\xb9\x11\xf5\x1a\xd9\x4b\x2c\xc1\x71\x5c\x7d\xef\x00\x12\x59\x2c\xc9\x53\x1d\x7b\x55\x5b\x7b\x4a\x63\x95\xd3\xc3\x96\x72\xfd\x0a\x5b\x2b\xae\x30\xfb\x5d\xad\xca\x1c\x30\xd3\x63\xa2\x47\x98\x75\xd2\x47\xdd\xf0\x76\xc2\xd9\xea\xb1\x5d\xa6\xd5\xb4\x27\xad\xe5\x79\x9c\x3e\xd1\xc2\x79\x42\xbf\x94\x81\x84\xce\x51\xb2\xcc\xde\xa7\xda\x6e\x1b\x73\x0c\x0f\xf8\x8f\x0a\x4e\x34\x1c\x55\x3c\x2d\x34\x1f\x58\x03\xeb\xf6\xc5\x66\x27\x0a\xad\x69\xb3\xa0\xfd\x9e\xf7\xa7\x2a\xbd\xfb\x39\xff\xf0\xac\xb3\x32\xfe\x20\x9f\x6b\x7e\xff\xf8\xbc\xfd\x8f\x47\xed\x57\x00\x00\x00\xff\xff\x31\xc9\x67\xb9\x67\x05\x00\x00"),
		},
		"/templates/service.yaml": &vfsgen۰CompressedFileInfo{
			name:             "service.yaml",
			modTime:          time.Time{},
			uncompressedSize: 467,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x94\x90\x4f\x4b\xc4\x30\x10\xc5\xef\xfd\x14\xc3\xde\x1b\x10\x3c\xf5\xea\x55\x64\x51\xd9\xfb\x98\x3e\xd7\xe0\x34\x09\x99\xe9\x82\xd4\xfd\xee\x92\x4d\x16\x2f\x0a\xda\x5b\xe7\xfd\xf9\xf1\xc2\x39\x1c\x50\x34\xa4\x38\xd1\xe9\x66\x78\x0f\x71\x9e\xe8\x09\xe5\x14\x3c\x86\x05\xc6\x33\x1b\x4f\x03\x51\xe4\x05\x13\x6d\x1b\x85\xe8\x65\x9d\x41\xbb\xb0\xf0\x11\x79\x15\x51\xf8\x02\x1b\x7d\x8a\x56\x92\x08\x8a\x7b\x5d\x45\x6a\x60\x47\x8e\xce\xe7\x9e\xd6\xcc\xbe\x55\xb8\x47\x08\x58\xe1\x1e\xae\xe7\xe6\x12\x7e\x81\x68\xa5\x11\x6d\xdb\xf8\x27\x54\xcb\x54\xd0\x27\xc5\x10\x67\x44\xa3\xdb\x5a\xa7\x19\xbe\x56\xd9\x47\xee\xd4\x03\xcb\x0a\x75\xda\xd6\xb9\x2a\x34\x6e\x4e\xc5\x3a\x76\xbc\xfc\xfc\xe8\xaf\x42\xf3\xd7\xcf\xb8\x1c\x61\xfb\x8b\x79\x81\x95\xe0\xb5\x2b\xb9\x24\x4b\x3e\xc9\x44\xcf\x77\xfb\x7e\x6b\xaf\xf7\x66\x96\xc7\x6f\xb3\x42\xe0\x2d\x95\xff\x0d\xbe\xa6\xee\x7f\x1b\xfe\x15\x00\x00\xff\xff\xc7\x19\xdc\x99\xd3\x01\x00\x00"),
		},
		"/values.yaml": &vfsgen۰CompressedFileInfo{
			name:             "values.yaml",
			modTime:          time.Time{},
			uncompressedSize: 1115,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x6c\x53\xc1\x6e\x1c\x37\x0c\xbd\xeb\x2b\x08\xef\x75\xbd\x59\x27\x75\x11\xcc\xcd\x75\x8a\xc2\x40\x1d\x2c\xe0\x16\x45\x51\xf4\x40\x6b\x38\x63\x16\x1c\x71\x4a\x51\x93\x6e\x0d\xff\x7b\xa1\x91\x37\x0d\xb2\xb9\xec\x4a\x6f\x1e\x25\xf2\xe9\xbd\x0d\x7c\xa0\x01\x8b\x38\x2c\x28\x85\x32\x0c\x6a\xc0\x13\x8e\x34\x17\x91\x4c\xd1\xc8\x2f\xa3\x26\x37\x15\x21\x0b\x1b\xf8\xe5\x89\x33\x70\x06\x84\xdf\x6f\xee\x7f\xbe\x1c\xd4\x26\x74\xa7\x1e\x06\x16\xda\x85\x7a\x60\x14\x34\x82\x05\x8d\xf1\x51\x28\x83\x2b\x3c\x12\xcc\x98\x33\xf5\xc0\xc9\x15\x8e\x5a\x0c\x9c\xa6\x59\xd0\x29\xef\x42\x88\x52\xb2\x93\x7d\xc4\x89\x3a\x98\xb0\xae\x43\x30\x9a\x85\x23\xe6\x0e\xae\x42\xe0\xec\xac\x5d\x00\x30\x5a\x38\xb3\xa6\x0e\x2e\x2e\x42\x98\xb5\xbf\x49\x49\x1d\x9d\x35\xe5\x0e\xfe\xf8\x73\xc5\x1e\x28\x16\x63\x3f\xde\x6a\x72\xfa\xc7\xd7\xba\x92\x6e\xf2\xaf\x99\xac\x83\xef\xaf\xaf\xdf\x7d\x77\x82\x7e\x32\x2d\xf3\x09\xcb\xe7\x75\x28\xa2\x9f\x0e\xc6\x0b\x0b\x8d\xf4\x63\x8e\x28\xeb\x65\x1d\x0c\x28\x99\xc2\xaa\x56\x6b\x6c\xd6\xcc\xae\x76\xec\x60\x7c\x8a\xb6\x63\x7d\xf3\x88\xe9\x5f\xe4\x28\x5a\xfa\x37\x5f\xc9\x9a\x03\x80\xe3\xd8\xc1\xb2\xdf\x5d\xed\xde\x06\x80\xfa\xed\xa0\xc2\xf1\xd8\xc1\xdd\xf0\x51\xfd\x60\x94\x29\x79\x68\x77\x1c\x8a\xc8\x43\x2b\x6d\x73\x26\xed\xe9\x81\x84\xa2\xab\x75\xf0\xfc\x12\x70\x18\x38\xb1\x1f\xd7\x8d\xab\x90\x7d\x21\x8b\x51\xd6\x62\x91\x72\xeb\xf5\xef\x42\xd9\xd7\x35\xc0\x44\xd3\xda\xf5\xc5\xd5\x7e\x7f\xcf\x17\x2b\x16\xe7\xd2\x80\xa9\xee\x85\x27\x3e\x63\xbf\xfd\x9a\xfd\x6e\x65\x87\x4c\xb6\x70\x5c\x35\xf1\xe3\x4c\x1d\xdc\xb6\xd7\xbd\x3b\xd4\x21\xd5\xbc\x83\xf7\xfb\xf7\xfb\xcf\xc4\x9b\x18\xb5\xa4\x26\xf6\x97\x6f\xf9\xfc\xf2\x99\x72\xaf\xa9\x2a\x5b\x29\x39\x1a\xce\x74\x97\x9c\x6c\x41\xe9\xe0\x7a\x15\x52\xf2\xad\xa6\x81\xc7\x56\xd5\xd3\x42\xa2\xf3\x44\xc9\xef\xb5\x5f\x3b\xa1\x54\xcd\xd8\x9f\x5e\xad\x9a\xe4\x03\x67\x2b\x73\xbd\xec\x87\xd2\x8f\xe4\xdf\xa2\x89\x8e\x15\x7e\x7e\x09\x00\x1b\x90\x7a\x6e\x07\x9c\x06\x85\x0d\x60\x8c\x34\x57\xeb\xb7\xec\x74\x30\x63\xe2\xb8\x85\x01\x1d\x65\x0b\x64\xa6\xb6\x85\x4f\x68\xa9\xfd\x72\x1a\xb7\x6b\xed\x16\x7a\x7a\x2c\xe3\x16\xdc\x30\x52\x08\x1b\xc0\xbe\xe7\xda\x08\x0a\x50\x5a\xd8\x34\xd5\xde\xcf\x32\xc4\xe9\x2f\x8a\x7e\x4a\x91\x3f\x11\xd4\x6c\x22\x27\xb2\x40\x69\x79\xed\xb4\xce\xbf\x86\xfa\x55\x92\xf3\xb1\xaa\xf3\x6c\x24\x6f\x76\x6a\xa9\x7b\xad\xb9\x7c\x0d\x7d\xc2\xa9\xf2\xea\x5f\x9e\xb1\xda\xa6\x89\xf0\x3f\xf0\x1b\xfb\xd3\x41\xfb\xe6\x2e\x80\x68\xd4\x53\x72\x46\x69\xc8\x7f\x01\x00\x00\xff\xff\x69\x14\x1d\x66\x5b\x04\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/.helmignore"].(os.FileInfo),
		fs["/Chart.yaml"].(os.FileInfo),
		fs["/crds"].(os.FileInfo),
		fs["/examples"].(os.FileInfo),
		fs["/templates"].(os.FileInfo),
		fs["/values.yaml"].(os.FileInfo),
	}
	fs["/crds"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/crds/crds.yaml"].(os.FileInfo),
	}
	fs["/examples"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/examples/test.yaml"].(os.FileInfo),
	}
	fs["/templates"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/templates/_helpers.tpl"].(os.FileInfo),
		fs["/templates/default_imps_cr.yaml"].(os.FileInfo),
		fs["/templates/deployment.yaml"].(os.FileInfo),
		fs["/templates/poddistruptionbudget.yaml"].(os.FileInfo),
		fs["/templates/rbac.yaml"].(os.FileInfo),
		fs["/templates/service.yaml"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
