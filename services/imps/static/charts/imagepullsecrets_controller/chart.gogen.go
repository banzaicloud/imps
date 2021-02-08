// Code generated by vfsgen; DO NOT EDIT.

package imagepullsecrets_controller

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

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\x41\x6e\xe3\x30\x0c\x45\xf7\x3c\xc5\x1f\x78\x33\x63\x0c\xe4\x43\x24\xb3\x98\x55\x0b\xa4\xc8\xb6\x90\x6d\x46\x62\x22\x8b\x82\x44\x27\x6d\x17\x3d\x7b\x91\x04\x41\xbb\x79\x20\x3f\xc8\x8f\xd7\xe1\xd9\x9b\x71\xcd\x0d\xa6\x90\x90\xb5\x32\x2e\x91\x33\xc6\x55\xd2\x2c\x39\xa0\xf8\xe9\xe4\x03\x37\x47\x1d\x5e\xa2\x34\xb4\xb5\x14\xad\xd6\xd0\x22\xa7\x84\x90\x74\xc4\xe2\x6d\x8a\x92\xc3\x5f\x54\x4e\xde\xe4\xcc\x28\xde\xe2\x8f\xdc\xe7\x99\x3a\x64\x0e\xde\x44\x33\x7e\x97\xca\x07\x79\xe3\x19\x17\xb1\x88\x5f\x7f\x1c\x9e\x72\x7a\x87\xe6\xdb\xe7\x55\x09\x85\x2b\x92\x64\x76\xe4\xb6\xbb\xd7\x9d\x69\x65\xea\xb0\xd1\x65\xd1\x8c\xfd\x66\x87\x59\x6a\x23\x17\xc4\x86\x1b\xef\xfa\xe4\xc6\x8f\x3a\xdc\xf8\x08\x62\x18\xae\x78\xac\xed\x9c\x87\xef\xa2\xd1\x4f\xa7\xb5\xe0\x20\x89\x1b\xf5\xae\x5d\x0a\xf5\x6e\xf4\x27\xea\x9d\x2d\xd7\x59\xab\x04\xea\x3f\xa9\xc3\xde\x57\xd1\xb5\xe1\xff\xf6\x5f\x23\x57\xaa\x1e\x79\x32\x72\x32\xb3\x1f\xee\xe7\x55\x8f\xe4\xce\x6d\xd2\x99\x07\xfa\x0a\x00\x00\xff\xff\x16\xec\x32\x27\x5d\x01\x00\x00"),
		},
		"/Chart.yaml": &vfsgen۰CompressedFileInfo{
			name:             "Chart.yaml",
			modTime:          time.Time{},
			uncompressedSize: 287,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8f\x41\x6e\xc3\x40\x08\x45\xf7\x9c\x82\x0b\xc4\x4e\xb6\x5e\x55\x6d\xcf\x90\x3d\x1e\x13\x1b\x75\x06\x46\x30\x8e\xe4\x9e\xbe\xb2\xd3\x48\x6d\xb6\x1f\xde\x7f\x40\x55\xae\xec\x21\xa6\x03\xde\x2f\xa0\x54\x78\x40\x29\x34\x73\x5d\x73\x0e\x4e\xce\x2d\x4e\xc9\xb4\xb9\xe5\xcc\x0e\x13\x47\x72\xa9\xed\x20\x3e\x37\xa5\x22\x89\x72\xde\xb0\x90\xd2\xcc\xb8\x63\xf8\xcb\xe1\xcd\xfc\x51\x16\xb0\xd8\xde\xbc\xb4\x56\x63\xe8\xfb\x91\xf4\x9b\x24\x65\x5b\xa7\x2e\x59\x01\x28\x24\xda\x48\x94\x3d\x06\x38\x21\x17\x92\x3c\xa0\xe8\xcd\xde\x5e\x77\x11\x1f\x57\xbe\x1f\x39\x7e\xec\x03\x08\x5b\x3d\xf1\xc1\x3e\x1d\xb3\xb4\x65\x1d\x77\xe4\xaf\xae\x1f\x29\x7d\x6d\xe4\x53\x00\xdc\x9f\x9f\x9f\xbb\x4b\x77\x06\xaa\xf5\xfa\x3f\x81\x9f\x00\x00\x00\xff\xff\xf5\x64\x35\xa7\x1f\x01\x00\x00"),
		},
		"/crds": &vfsgen۰DirInfo{
			name:    "crds",
			modTime: time.Time{},
		},
		"/crds/crds.yaml": &vfsgen۰CompressedFileInfo{
			name:             "crds.yaml",
			modTime:          time.Time{},
			uncompressedSize: 11061,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x5a\x4b\x73\xe3\xb8\x11\xbe\xeb\x57\x74\x39\x87\x49\xaa\x2c\x79\xa7\xf6\x92\xd2\xcd\xe5\x99\x83\x92\x8d\xd7\x35\x76\x39\x87\x54\x0e\x10\xd9\x22\x11\x83\x00\x17\x00\xb5\xa3\xdd\x9a\xff\x9e\x6a\x00\x7c\x8a\x0f\x48\x33\xc6\x45\x25\xb0\xd1\xe8\x37\xfa\x23\xc1\x4a\xfe\x8a\xda\x70\x25\xb7\xc0\x4a\x8e\x5f\x2d\x4a\xfa\x67\x36\x6f\x7f\x37\x1b\xae\xee\x8e\x1f\xf7\x68\xd9\xc7\xd5\x1b\x97\xe9\x16\x1e\x2a\x63\x55\xf1\x05\x8d\xaa\x74\x82\x9f\xf0\xc0\x25\xb7\x5c\xc9\x55\x81\x96\xa5\xcc\xb2\xed\x0a\x80\x49\xa9\x2c\xa3\x69\x43\x7f\x01\x12\x25\xad\x56\x42\xa0\x5e\x67\x28\x37\x6f\xd5\x1e\xf7\x15\x17\x29\x6a\xb7\x43\xbd\xff\xf1\xa7\xcd\xcf\x9b\x9f\x56\x00\x89\x46\xb7\xfc\x85\x17\x68\x2c\x2b\xca\x2d\xc8\x4a\x88\x15\x80\x64\x05\x6e\x81\x17\x2c\xc3\xb2\x12\xc2\x60\xa2\xd1\x9a\x8d\x9b\x30\x9b\x3d\x93\x7f\x30\x9e\x08\x55\xa5\x1b\xae\x56\xa6\xc4\xc4\xc9\x93\xa6\x4e\x48\x26\x9e\x34\x97\x16\xf5\x83\x12\x55\xe1\x85\x5b\xc3\x3f\x9e\x7f\x7d\x7c\x62\x36\xdf\xc2\xc6\x58\x66\x2b\x13\x7e\x9c\xe4\x29\x9a\x44\xf3\xd2\x3a\xf9\xbe\x60\xa9\xd1\xa0\xb4\x06\xf8\x01\x6c\x8e\xa0\xf6\xff\xc3\xc4\x42\xce\x0c\xec\x11\x25\x98\x2a\x49\xd0\x98\x43\x25\xc4\x09\x34\x26\x4a\x26\x5c\x60\xea\x58\x1d\x94\x2e\x98\xdd\xc2\xfe\x64\xd1\x4d\x78\x5d\x9e\x2d\x0b\xff\xed\xa9\xc4\x2d\x18\xab\xb9\xcc\x56\x00\x99\x56\x55\x19\x74\x3d\x53\xcd\xaf\x0e\xe6\xf5\xae\xd9\x11\xe1\x53\x25\xc4\xb3\x33\x8a\x7b\x22\xb8\xb1\xff\x1c\x7b\xfa\x0b\x37\x9e\xa2\x14\x95\x66\xe2\xdc\xa4\xee\xa1\xc9\x95\xb6\x8f\xed\x46\x6b\xe0\x45\x19\x1e\x71\x99\x55\x82\xe9\xb3\x95\x2b\x00\x93\x28\xd2\xe4\x41\x54\xc6\xa2\xa6\x89\x6a\xaf\x43\xc8\x04\x4e\xde\xc4\x5b\xf8\xf3\xdb\x0a\xe0\xc8\x04\x4f\x9d\xc3\xfd\x43\x55\xa2\xbc\x7f\xda\xbd\xfe\xfc\x9c\xe4\x58\x30\x3f\x39\xf0\xc5\x40\x1f\xe0\xc6\x39\xc4\xaf\x20\x5b\xbb\xbf\x43\xad\xe0\xfe\x69\x17\xb8\x95\x5a\x95\xa8\x2d\xaf\x25\xa2\xd1\x49\x86\x66\x6e\xb0\xef\x07\x12\xcc\xd3\x40\x4a\xe1\x8f\x7e\xe3\x10\xc4\x98\x82\xf1\x22\x28\x8a\x10\x6e\x40\xd7\x41\xe3\x14\xec\xb0\x05\x22\x61\x32\xc4\xd0\x06\x9e\x51\x13\x13\x32\x7a\x25\x52\xca\x99\x23\x6a\xeb\xa2\x28\x93\xfc\x8f\x86\xb3\x01\xab\xdc\x96\x82\x59\x0c\x5e\xac\x87\x8b\x6e\xc9\x04\x99\xb4\xc2\x5b\x60\x32\x85\x82\x51\x28\xba\x38\xad\x64\x87\x9b\x23\x31\x1b\xf8\x97\xd2\x08\x5c\x1e\xd4\x16\x72\x6b\x4b\xb3\xbd\xbb\xcb\xb8\xad\xd3\x3f\x51\x45\x51\x49\x6e\x4f\x77\x2e\x89\xf9\xbe\xb2\x4a\x9b\xbb\x14\x8f\x28\xee\x0c\xcf\xd6\x4c\x27\x39\xb7\x98\xd8\x4a\xe3\x1d\x2b\xf9\xda\x09\x2e\x5d\xf6\x6f\x8a\xf4\x2f\x8d\xe3\x3f\x74\x24\x1d\x84\xba\x1f\x2e\x8c\x27\xed\x4e\x61\x4c\x5e\x66\x61\x99\x97\xbf\x35\x2f\x4d\x91\x55\xbe\x7c\x7e\x7e\x81\x7a\x53\xe7\x82\xbe\xcd\x7d\xc6\x36\xcb\x4c\x6b\x78\x32\x14\x97\x07\xd4\xde\x71\x07\xad\x0a\xc7\x11\x65\x5a\x2a\x2e\xad\xfb\x93\x08\x8e\xb2\x6f\x74\x53\xed\x0b\x6e\xc9\xd3\xbf\x55\x68\x2c\xf9\x67\x03\x0f\xae\x08\xc2\x1e\xa1\x2a\x53\x66\x31\xdd\xc0\x4e\xc2\x03\x2b\x50\x3c\x30\x83\xef\x6e\x76\xb2\xb0\x59\x93\x49\x97\x0d\xdf\xad\xdd\x7d\x42\x6f\xad\x66\xba\x2e\xa8\xa3\x1e\x1a\x64\xe4\x73\x89\x49\x2f\x43\x52\x34\x5c\x53\x14\x53\xc1\xa3\xd8\x1f\x2b\x58\xd3\xb9\x49\x43\x63\xc6\x8d\xd5\xa7\xfe\xec\x59\x95\xf6\x44\xee\xd4\x61\x5c\xd6\xbb\x5b\xc6\x85\xf1\x69\x89\xe0\xeb\x01\x65\xd2\x1e\xfd\x71\x13\xaa\x74\x3f\x9b\x00\x59\x92\xfb\x52\x5b\xb2\x04\x07\x14\x53\x62\xd2\x48\x34\xa6\xe4\x0e\x26\x46\x1e\x0e\x24\x7e\x68\x69\x9d\x89\xf9\x81\xa3\x81\xdf\x73\x9e\xe4\x7d\x41\x2b\x83\x29\x30\xa7\xd0\x08\x4f\x80\x10\xf4\x54\xfd\x52\x95\xbc\xa1\x06\xa1\x32\x2e\xbb\xc2\x8c\xac\x9b\x53\x03\xea\x73\x6a\xf4\xc9\x44\x3c\x0d\x17\x3b\xdb\x5d\xc9\x81\x92\x8a\xc2\x66\x6c\xf9\xda\x71\x9f\x7c\x30\xe6\x32\x98\x0a\xed\xf9\xdd\xd6\x33\x26\x9c\x64\x67\x99\xce\xd0\xce\xc6\xea\x8b\x23\xe9\x39\x9d\xd9\xfa\x04\xd8\xa3\x0b\x55\x52\xa5\x1f\xb6\x63\xc1\x46\x01\x4c\x01\xcb\x20\xe3\x47\x94\x57\xc5\x6c\xb3\x66\x39\x64\x1f\x1b\xd2\x20\xbc\x4b\x37\xdf\x66\x19\x50\xb2\x91\x3c\x10\xd9\x9c\x9d\x8b\x0d\xbe\xc5\x20\x5d\x73\x76\xf4\xda\x7a\x8b\x84\x43\x3d\x43\x89\x7a\x34\x35\xe3\xa2\x76\xe2\xd1\x92\x2e\xe4\x88\xb6\x67\xec\xaa\xb1\x3f\x4d\x85\x9c\x1f\x56\x35\x32\x77\xcb\x0c\x97\x56\x4d\xac\xe1\x16\x8b\x49\x31\x23\xf2\xab\x26\x61\x5a\xb3\xd3\xb8\x85\x51\x60\x42\xa7\x47\x94\x31\x9e\x6b\xea\xc6\xaf\xee\xd4\x6b\x7c\x7b\x1b\xea\x12\xd3\x08\x05\xb3\x49\x3e\xea\x1c\x3f\x58\x46\xe5\xd7\x0e\x63\x41\xb0\x3d\x0a\xd7\xc5\xa4\x98\xf0\x14\x7d\x33\xcd\xcd\xcc\x81\x30\xd0\xa8\xc9\x0f\x56\x96\x82\x63\x5a\x77\x44\x83\xd0\xbf\x25\xce\x45\x25\x2c\x2f\x05\xb6\x76\x98\x96\x57\x63\x13\x01\x29\xad\x55\x12\xa9\xe5\xf0\x7a\x9a\xae\x47\x7f\xe7\x42\x90\x00\x05\x93\x2c\x9b\x31\xc1\x5f\x7f\xfd\xf2\xb7\xeb\x3c\xdf\xf3\xca\xbd\x37\x5a\xa3\x83\x6f\x84\xfc\xdc\x6f\x15\xea\x13\xa8\x23\x6a\xea\x8d\x66\xcc\xe6\x1a\xce\xa6\x1d\xdb\xc0\x4b\x4e\xed\x93\xa9\x84\xa5\x07\x4e\xc9\x5f\xbc\x6b\x7c\xdb\x68\x93\xfc\xf3\x57\x6a\x93\x1c\x1e\x9c\x61\x4b\x76\xbb\x7f\xfc\x44\x3d\xce\xbd\x04\x2c\x4a\x7b\x1a\xca\x5b\x9b\x90\x09\x11\x4a\xa5\xd9\xcc\x70\xbc\x77\x78\x6f\x8a\x89\x54\xcb\x3c\x96\xca\x83\x1f\x43\x25\xe7\x68\x07\x2e\x19\x2e\x0d\x2e\xe1\xc6\x59\xb3\x2f\xf9\x2c\xd7\xe6\xd4\x29\x7c\x3b\xea\xdd\xd2\xce\x74\xcc\x3b\xcb\x67\x21\x9e\x46\x54\x38\x8b\xaa\xce\xb6\xa1\xd5\x8e\x53\x00\x5c\x7d\x6f\x1b\x2e\x0f\x2d\x6e\x81\xc1\x1b\x9e\x3c\x0a\x21\xa0\x53\x52\x61\x74\xc8\x6c\xe2\x30\xe8\x1b\xc5\xe1\x1b\x97\x74\x6f\x78\x72\x4c\x02\x64\x59\x58\x1b\xe7\x7a\x3f\xde\xf0\xac\x99\x1c\x1b\x3d\xb3\x91\x34\x01\x6e\x7a\xfb\xd1\x84\x33\x80\x2f\x10\x91\x26\x03\x07\x38\xa9\x7e\x39\xc4\x10\x41\x1f\x71\x18\xb4\xa3\xb6\xf6\xc5\xea\x35\x6e\x6a\x31\x92\x77\xe4\x07\xe3\x9d\x42\xd1\x9e\xf3\x32\x4a\x41\xab\x7c\x51\xa2\x9c\xa8\x01\xe7\x2b\xc1\xfd\x66\x1b\x1f\xdf\x3b\x79\x0b\x8f\xca\xee\xe4\x6d\x14\xdb\xcf\x5f\x39\x01\x2d\x8a\x89\x4f\x0a\xcd\xa3\xb2\x6e\xe6\x87\x1b\xd1\x8b\x7c\xb1\x09\xfd\x32\x97\x42\xd2\x9f\xcc\xa4\x7f\x17\xb7\x2e\x06\xb1\x1f\xbb\xf0\x9a\xa9\x76\x09\x1d\x92\x12\x94\x0e\xb6\xf2\x6f\x1e\xfc\x66\xd3\x0d\xc0\x70\x14\x95\x71\xc0\x54\x2a\xb9\x76\xc5\x7a\x33\xb6\x4f\x30\x71\x64\x20\x77\xbd\x70\x2e\x56\xb3\xa5\xdf\x2e\x8a\xe3\x0b\x75\x04\x7e\xb5\x7f\x8b\x22\x58\x82\x29\xa4\x95\x33\xa2\x7b\x0b\xc0\x2c\x66\x3c\x81\x02\x75\x36\xdd\x92\x75\x47\x49\x35\x3b\x66\xfb\xa8\x5a\xea\xc7\x45\xf1\xb4\xdc\xad\xb5\x63\x0e\xf9\x74\xc7\x9a\x72\x73\x91\xa6\x76\xed\x02\xe1\x0c\x36\x1a\x23\x5c\xd2\xa3\xd3\x54\xcc\xab\xd1\x7d\x45\x1b\x5b\xbd\xa3\x2d\x7f\x7e\x6e\x87\x3e\xc7\x9d\x71\x05\x2b\x29\x3b\xff\xa4\xa3\xca\x05\xed\x37\x28\x19\xd7\x8b\x19\x7a\xef\xde\x81\x0a\xec\xad\xe4\x1e\xf9\x74\x37\x21\xfe\xdc\x00\x79\xf3\xc8\xc4\xf0\xe5\xd1\x88\x5a\x8a\xaa\x06\x0a\x7f\x0c\xd7\x9d\x59\xa7\xd3\xa0\x16\x5c\x19\x7f\x2a\x1e\x38\x8a\x14\xf8\x5c\x77\x46\xe3\xe6\x0d\x4f\x37\xb7\x67\x39\x7e\xb3\x93\x37\xfe\x78\x3e\xcb\xd8\xfa\x2c\x5f\x60\xac\xa4\x38\xc1\x8d\x5b\x79\x73\x7d\xeb\x12\x15\x75\x11\x44\xf3\x41\x39\xcb\xa0\xc5\x27\xff\xe6\x36\x7f\x52\xe9\x32\x06\x26\xa2\x78\x94\x54\xa3\xa1\x51\xd1\x69\x75\x49\xec\xce\x5e\x39\x2d\xc3\xa4\x00\x87\x46\xf9\x4e\x41\xa4\x52\xa5\x1f\xcc\x75\x40\xe9\x7a\x90\x34\x05\x86\x66\x8a\xed\xd5\x20\xa8\x0f\x75\x26\xc2\xe5\x32\x00\xd4\x87\x39\x13\x2c\xa3\xc1\xcf\x38\xc4\x99\xe0\x1a\x0b\x7c\x96\x3b\xdf\x78\xc0\xf3\x1e\x60\xe7\xfb\x81\xce\xe2\xc1\xfc\x1e\x00\xe7\x62\x70\x53\x83\x97\x05\xae\x97\x00\x9b\x58\x50\x13\x01\x68\xde\x03\xcc\xc4\x03\x99\xe8\x53\x3b\x16\xc0\xbc\x0f\x78\x79\x07\xe0\x72\x25\x68\x89\x36\x58\x1c\x58\xb9\x18\xa8\xc0\xee\x10\x01\xff\x2f\x00\x29\x0e\x0d\x2c\xb2\x8c\x07\x28\x7d\xf0\x11\x25\xeb\x0c\x38\x99\x00\x1e\x8b\x6c\xa7\x80\x49\x1c\xe8\x88\x04\x1c\x17\x80\x8d\xb8\x06\x3d\x06\x64\x2c\x01\x8c\x28\x70\x71\x41\x8b\x37\x27\x73\x14\xa0\xb8\x14\x4c\x44\x59\xf5\x4a\x10\x01\xf7\x33\x1b\xff\x78\x00\x71\x15\x78\x08\xf0\x60\x8e\xed\xe5\xc0\xc1\x41\x83\x19\x96\xdf\x05\x1a\x16\xa3\x69\x81\x60\x2e\xd0\x7c\xf3\xba\xd8\xfd\x77\x3f\x54\x3d\x28\x79\xe0\x59\x78\xbe\x0f\x4d\x70\x7b\x62\x87\xaf\x77\x13\x5f\x67\xfc\xc5\x94\xce\x27\x3c\x07\x00\xfc\xc9\x8b\xe9\xec\x97\xcc\xa5\xa6\xe0\xec\x22\xd6\x08\xc9\x45\xb9\x12\x91\x29\xfd\x06\xac\x15\x60\xfc\xdb\x5a\x47\x42\x98\xfc\x46\xe6\x0e\x83\xb2\xb2\xf5\x77\xc5\x80\x7b\x22\xa0\xe0\x0c\x89\x98\x2d\x22\xef\x6b\x96\x90\xd8\xa3\x16\x69\xbf\x8e\x35\x4a\x4f\xef\xfb\xa3\x8c\x31\xf7\x89\xff\xec\x5b\x69\x47\xee\xf1\xef\xd3\xf5\x25\x1b\x2e\x0d\x4f\xa7\xdf\xcc\x11\x1c\x6a\x9b\xcc\x85\x50\x87\x77\xba\x2a\x70\xd5\x65\x80\xd1\xcf\xf0\x13\x9c\xc6\xb8\xac\x9b\x9b\x2c\xbd\x49\x7f\x61\x60\xb5\xc0\x32\x5c\x9d\xeb\x90\xcd\x5e\xc3\x71\xd4\xbd\x8b\x38\x6a\x6f\x50\x1f\xbf\xef\x26\xce\xb9\x10\x30\xed\xa0\x11\x35\x06\x53\xed\x15\xd0\x8f\x4c\x94\x39\xfb\xd8\xce\x85\x5b\x9a\xfe\xb2\x64\xe7\x31\x95\x4e\xd2\x62\x0b\x56\x57\xde\xb3\xc6\x2a\xcd\x32\x0c\x33\xad\x88\x2c\x49\xb0\xb4\x98\x3e\x0e\xaf\x4c\xde\xdc\xb8\x3f\xf5\x1d\x48\xf7\xb7\x7d\x89\xb3\x85\xff\xfc\x77\xe5\xb9\x62\xfa\x5a\x4b\x43\x93\xff\x0f\x00\x00\xff\xff\x5e\xc5\x4b\xcd\x35\x2b\x00\x00"),
		},
		"/examples": &vfsgen۰DirInfo{
			name:    "examples",
			modTime: time.Time{},
		},
		"/examples/test.yaml": &vfsgen۰CompressedFileInfo{
			name:             "test.yaml",
			modTime:          time.Time{},
			uncompressedSize: 728,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\x3f\x4f\xc3\x30\x10\xc5\x77\x7f\x8a\x93\x77\x17\x85\x3f\x8b\x67\x3a\xb0\xa0\x4a\x48\x30\x20\x86\x23\x39\x25\x56\x9d\xd8\xf2\x5d\xaa\x96\x4f\x8f\x12\x27\x21\xa8\x50\xc1\x14\xcb\xef\x3d\xff\x9e\x2f\x36\xc6\x28\x8c\xee\x99\x12\xbb\xd0\x59\x38\x14\x6a\xef\xba\xca\xc2\x23\xb6\xc4\x11\x4b\x52\x2d\x09\x56\x28\x68\x15\x40\x87\x2d\x59\x70\x6d\x64\x23\xc4\xa2\xd4\xaf\xf9\x27\x2a\x13\xc9\xa5\xb0\x29\xa6\xbd\x11\xb3\x3e\x95\x25\xb9\xae\xbe\x9f\x62\xc3\x96\x05\x3d\x7c\x8a\xeb\x9b\xdb\x3b\x7d\x06\x75\x2d\xd6\xc4\x9b\x77\xec\x3e\xd0\x95\x3e\xf4\xd5\xc6\x85\xab\x43\x81\x3e\x36\x38\x17\x7a\x18\x4c\xbb\xde\xfb\x3f\x35\xe3\x48\xe5\x08\xc7\x54\x93\x0c\x2b\x00\x1e\x83\x79\x3d\x67\xf4\xc8\x36\xb1\xf7\xde\x64\xdd\xe4\x1a\x5a\xcd\xa6\xf1\x7a\xbc\x8e\xb1\x85\x57\xbd\xd0\xf4\xdb\x24\x31\x79\x2a\x25\xa4\xc5\x0b\x60\xa0\x45\x29\x9b\xed\x31\x26\xe2\xe1\xb2\x2b\x2d\xeb\x7b\x3a\x59\xd0\x99\x69\xb8\x09\xbd\xaf\x4c\x83\x07\x9a\xda\xb0\xfe\xe6\x07\x08\x91\x12\x4a\x48\x16\xf4\xf6\xe8\x78\x36\x7c\x15\x7d\x71\xd2\xec\x42\xb5\x80\x2e\x57\xf8\x6f\x81\x1f\xf1\x89\x6a\xc7\x92\x4e\xf9\xe0\x32\x51\x45\x9d\x38\xf4\x7c\x36\xeb\xe5\x07\xe9\xf5\x38\xf3\xfb\x59\x4d\x54\xa9\xcf\x00\x00\x00\xff\xff\x34\x19\xfb\x89\xd8\x02\x00\x00"),
		},
		"/templates": &vfsgen۰DirInfo{
			name:    "templates",
			modTime: time.Time{},
		},
		"/templates/_helpers.tpl": &vfsgen۰CompressedFileInfo{
			name:             "_helpers.tpl",
			modTime:          time.Time{},
			uncompressedSize: 2469,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x4d\x6f\xdb\x38\x10\xbd\xeb\x57\x0c\x84\x0d\xbc\x9b\x85\xe8\xc3\x02\x7b\x30\x90\x43\x91\xf6\x50\xb4\x48\x8b\x06\x48\x0f\x45\x61\x50\xd4\xc8\x22\x42\x51\x0c\x87\x74\xe3\x26\xfe\xef\x05\x29\x5a\x91\x1d\xdb\x75\x8a\xdc\x68\xcd\x9b\x99\x37\x6f\x3e\xfc\xf0\x30\x3d\x87\xa5\x6c\x67\x40\xe8\xa0\x96\x0a\xdd\xca\xe0\x45\xeb\xc9\x71\xd1\xe0\x0c\xce\xa7\xeb\x75\x16\x50\xd9\xbb\x7b\xc3\x75\x05\xae\x41\xd0\xbc\x45\xe8\xea\xf8\x16\x0d\xb7\x8e\x65\x09\x57\x40\x85\xb5\xd4\x08\xb9\x6c\xf9\x02\x8d\x57\x8a\x50\x58\x74\x2c\xf8\xe4\x50\x3c\xa1\xb8\x57\x0e\xd8\x65\x74\xbf\x0a\x01\xd9\x0d\x57\x1e\x29\x22\x3f\x2d\xd1\x5a\x59\x21\x3c\x82\xb3\x5e\x0b\xf8\xff\xbf\xf8\x94\xed\xb5\xaf\x6b\x79\x0f\x79\x91\x43\x8a\x85\xba\x0a\xcf\x9e\xe5\xa5\x45\xee\x10\xf8\x90\xa1\xf6\x4a\xad\xe0\xce\x73\x25\x6b\x89\x15\x70\x63\x22\x7f\x96\x7d\xc5\x3e\x76\xc4\xbb\x90\x21\xd4\x42\x50\xa2\xe0\x9e\x10\xa8\x6b\x11\x3e\xf8\x12\xad\x46\x87\xd4\x57\x5d\x4b\x54\x15\x01\xb7\x08\x4a\xb6\xd2\x61\x05\xae\x03\xd7\x48\x82\xbf\xcb\x55\x54\xe4\xed\xd5\x75\xc0\x4a\xbd\x00\x32\x28\xfe\x61\xd9\xfb\x1a\x2c\x2a\xe4\x94\xa4\x13\x9d\x76\x5c\x6a\xea\xc5\xeb\xbf\x49\x07\x3f\xa4\x52\x50\x22\x78\x0a\x3c\x09\x78\x24\x9f\xd8\xfe\x56\xe0\x80\xdd\x16\x59\xd6\x83\xa6\x1b\xe3\xa0\x6b\x82\x1c\xb4\x9f\xa2\xbb\xa2\x21\xce\x5f\xb1\x86\xd9\xc5\xe9\x8d\x7d\xe2\x38\xa8\xd1\x07\x61\x5f\x7a\xa9\x7a\xdf\x0d\xcf\xad\x8f\x2f\x24\x67\xac\xd4\xae\x86\xfc\x8c\x8a\x33\xca\x77\x62\xf5\x49\x4f\x1f\xb3\xfd\xcf\xad\xe1\x1b\x75\x35\x6c\xcc\x12\x2d\xc9\x4e\x87\x8e\xc6\xce\xa6\x31\xe9\x51\x8a\x97\xa8\x4e\xe8\x6e\x44\x3f\xb5\x76\xb7\xa4\xb1\xda\xfd\xfb\x26\x65\x7d\x04\x8b\x46\x71\x81\x90\xff\x9b\x43\x3e\xcf\x5f\xb2\x52\xc7\x28\x15\xa1\x6f\xb6\x53\x0a\xed\xb3\xd9\x03\xa9\x85\xf2\xd5\xd1\x29\x65\xb0\x5e\x8f\x62\xec\xa8\x79\x52\xda\xd3\x52\xbe\x5a\xba\xd8\x2b\x8a\x22\x71\x63\x66\x70\x24\xe9\x7e\x71\x58\x72\x65\xb7\xc3\x55\x61\xb2\x9b\x06\xe3\xc9\xd1\x46\x91\x4a\x2e\x6e\x57\xdc\x56\xc4\x4a\xae\x7f\x72\x29\x54\xe7\xab\x10\x50\x28\x4f\x0e\x6d\x31\x04\xb6\x78\xe7\xa5\xc5\x0a\x72\x96\x6c\x71\x54\xc2\x99\x0f\x07\x27\x9c\xa9\x78\x1d\xf3\x61\x57\xc7\xb0\xf5\x3a\x6b\x50\xb5\x8c\x9a\x69\x1c\xc3\xa3\x5c\x37\x93\x7a\xa0\xd6\x96\x6b\xbe\xc0\xaa\x28\x57\x31\xca\xb0\x8b\xd7\x68\x97\x52\xe0\x7e\x27\xa9\xc9\x71\x2d\x70\xdb\x65\xc3\xed\x39\x3e\x6d\x5c\x0f\xef\xd7\xe1\x8d\x31\x87\x37\x62\x6f\x10\xd1\xb5\xa6\xd3\xa8\xdd\x0c\x0e\xb7\x63\x8f\x9f\xe1\xd6\x15\x5d\x7d\x5c\xa4\x51\x17\xff\x60\x10\x09\x15\x0a\xd7\xd9\x8f\x69\x20\x8b\x57\x1d\xab\x97\xea\xbf\x73\x30\xa6\xe7\x59\x06\xf0\xd9\x76\x2d\xba\x06\x3d\x41\xe9\xa5\x72\x85\xd4\x94\x3d\x3f\x73\x66\x80\xb1\x16\x9d\x95\x82\x58\x8d\x15\x5a\xee\xd2\x62\x17\x9b\x66\xb5\xb1\x13\x93\x15\xd2\x24\x03\x70\xdc\x2e\xd0\xc5\xfa\x67\x30\x19\x16\x61\xbe\x71\x9e\xec\xde\xe9\x23\x09\xd3\x5f\x56\x91\x7e\x6f\xf2\x52\xe7\xad\xc0\x5e\xe2\x19\x7c\x9b\xcf\x93\x7d\x6e\xb8\x6b\xe6\xf3\xef\xbb\x2c\xc6\xe6\xac\xd8\xb6\xa5\x6d\xca\x60\xbb\x9a\x3c\xc8\xb9\x77\xe1\xf2\x31\xff\x5f\x01\x00\x00\xff\xff\xf9\xf5\xdb\x31\xa5\x09\x00\x00"),
		},
		"/templates/deployment.yaml": &vfsgen۰CompressedFileInfo{
			name:             "deployment.yaml",
			modTime:          time.Time{},
			uncompressedSize: 2640,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x55\x4d\x8f\xdb\x36\x10\xbd\xef\xaf\x20\x7c\x97\x36\x29\x8a\x62\x21\x20\x87\xc5\x36\x08\x0a\x38\xa9\xd1\x2d\x16\xe8\xa9\x18\x93\x63\x9b\xc8\x98\x43\x90\x23\xb5\xee\xc6\xff\xbd\xa0\x24\xcb\xa2\x25\x67\x83\x6c\x51\x9d\x6c\xce\xc7\x7b\x9c\x8f\x47\xf0\xf6\x09\x43\xb4\xec\x2a\x05\xde\xc7\xdb\xe6\xed\xcd\x67\xeb\x4c\xa5\x7e\x46\x4f\x7c\xd8\xa3\x93\x9b\x3d\x0a\x18\x10\xa8\x6e\x94\x72\xb0\xc7\x4a\x3d\x3f\x2b\xeb\x34\xd5\x06\xd5\xc2\xee\x61\x8b\xbe\x26\x8a\xa8\x03\x4a\xa1\xd9\x49\x60\x22\x0c\xe5\xa6\x26\x4a\x01\x0b\x55\xaa\xe3\xb1\x8f\x8e\x1e\x74\x97\xa2\xfc\x0d\x09\x21\x62\xf9\xe9\x74\xdc\x79\x11\xac\x91\x62\x42\x53\xea\xf9\xb9\xf8\x26\xa8\x2e\x26\x01\x7d\x51\xce\x3a\x83\x4e\xd4\x8f\x29\x5d\xf4\xa8\x53\xaa\x80\x9e\xac\x86\xd8\x21\x3f\x01\xd5\x18\xcb\xd3\x61\x87\x1b\x91\x50\x0b\x87\x0e\x79\x0f\xa2\x77\xcb\x11\x95\x6f\x27\x73\x4a\xb4\x9c\x21\xf5\x53\x87\x25\xb8\xf7\x04\x82\x3d\xd6\xa8\xc2\xe9\x03\xe7\x58\x40\x2c\xbb\x01\x5b\xa9\x35\xe8\xcf\x07\x08\x26\x96\x6b\x70\xff\x80\xd5\xc4\xb5\x29\x2d\xdf\x46\x1d\xc0\x5b\xb7\x2d\xac\x13\x0c\x0d\x50\xa5\x16\xa3\x4b\x46\x0c\x8d\xd5\xf8\x91\x9d\x15\x0e\x65\xeb\x8d\xbf\xf4\xae\xea\x78\x5c\x0c\x08\xed\xfd\x36\x43\xa0\xc1\x06\x89\x7d\x9a\x81\x8f\x6c\xb0\x44\x07\x6b\x42\xd3\x5d\xe0\x05\x4e\xec\x31\x80\x70\x28\x42\xd7\xe3\x9c\xd2\xc9\xda\xcf\xde\x84\x04\x52\xc4\xef\x84\x89\x92\x48\xe6\xd9\x5c\xc6\xf9\xe2\x96\x36\x8a\xe5\x32\x60\x63\x7b\x2a\x83\x63\x67\xb9\x40\x0c\xd8\x64\x23\x74\x3d\xfc\x02\x3a\xfd\xfd\xcb\xca\x6e\x88\xf4\x6c\xee\xcf\x7d\xbe\x8c\x14\xfe\x03\xf6\x94\x8d\xce\xdd\x38\x55\x76\x29\xca\xc6\xf4\x75\x5b\x33\xa0\x9c\x36\xa7\xfd\x8d\xba\x0e\x56\x0e\x0f\xec\x04\xff\x96\x6a\x96\xea\xf9\x5e\x8f\xb9\xfb\x5c\xf6\x94\xb3\x9d\xcb\x7b\xad\xb9\x76\xf2\xe9\x55\xc2\x92\xbe\x64\x07\xeb\x30\x8c\x0a\x51\xf4\x82\x75\x8e\x1d\x4c\x4a\x41\xd8\x8e\x5c\x3b\xf7\xa2\xd8\xa3\x04\xab\x63\x01\xc6\x84\x77\xd5\xdd\x9b\xbb\x37\x13\x97\x6e\x11\x0a\x42\x30\x18\x8a\x76\xd9\x2d\xbb\x89\x9b\x66\xb7\xb1\xdb\x62\x10\xbd\x77\x5f\xd1\xbc\xd3\x77\xb5\xd0\x57\x8a\x1d\xaf\x56\xfa\xed\x0f\x79\xe6\xb6\x9c\xf9\x1a\xb6\x47\x49\x04\x39\x26\x69\x38\xa8\xe3\xb1\x9a\x98\x05\xb6\xea\x8b\x32\xb8\x81\x9a\x44\x95\x0f\x3b\x08\x52\xde\x7b\x3f\xb7\xbb\x3d\xca\xaa\x26\x5a\x31\x59\x7d\xc8\x57\xa5\xcd\xe7\x07\x63\xce\xcf\x73\x90\x49\x3b\xba\xee\xf5\x2d\xc9\x6c\xa3\x7e\xaf\x38\x48\xa5\x26\x9d\x52\xca\x07\x16\xd6\x4c\x95\xfa\xfd\x61\x35\xb2\x91\x6d\xd0\x61\x8c\xab\xc0\x6b\xcc\x21\x77\x22\xfe\x03\x5e\xd4\x5d\x29\x0f\xb2\xab\xd4\xed\x3c\x11\xdf\xe2\x4f\x6d\x01\xc1\xd8\xff\x05\x27\x72\x1d\x34\xc6\x17\xa7\x65\xf0\xfc\xca\x9c\xa0\x6b\xc6\x79\x2e\xd4\x92\x78\x5b\x52\x7a\x17\xf2\xa0\x73\xaf\x96\xbf\x7e\xf8\x73\xf9\xfe\xe9\xfd\xf2\x82\x7c\x93\xe2\xf3\xf1\x1b\xe7\x5a\x5c\x40\x66\xf2\x36\x91\x4e\xc7\x06\x1f\xfb\x47\xf6\xec\x36\x3e\x9d\x57\xa8\x79\x21\x7a\x19\x0f\x36\x1b\xeb\xac\x8c\x26\xf6\x74\xf2\xdf\xe2\x08\x53\x7a\xcf\xf2\xf7\x60\x74\xf8\x2a\xb4\x61\x35\x1f\x5b\x3d\x9d\x51\xc9\xef\x52\xdf\x22\x39\x15\x9d\xd7\xcd\xbf\x01\x00\x00\xff\xff\x14\x3c\x86\xb2\x50\x0a\x00\x00"),
		},
		"/templates/imagepullsecret.yaml": &vfsgen۰CompressedFileInfo{
			name:             "imagepullsecret.yaml",
			modTime:          time.Time{},
			uncompressedSize: 652,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xc1\x6a\xe3\x30\x10\x86\xef\x7a\x8a\x41\xa7\x5d\x58\x69\x59\x08\x7b\xd0\xb9\xe7\x1e\x5a\xc8\x5d\x91\x27\xae\x12\x45\x32\x33\x52\x4b\x70\xf4\xee\x45\xb6\x53\x9a\x36\x86\x9e\x6c\x66\xfe\x6f\xbe\x19\x5b\x29\x25\xec\xe0\xb7\x48\xec\x53\x34\xf0\xfa\x4f\x1c\x7d\xec\x0c\x3c\xa3\x23\xcc\xe2\x84\xd9\x76\x36\x5b\x23\x00\xa2\x3d\xa1\x81\x71\x04\x1f\x5d\x28\x1d\x82\xf4\x27\xdb\xe3\x50\x42\xe0\x29\xad\x5c\x8a\x99\x52\x08\x48\x7a\x5f\x42\x68\x80\x04\x0d\xb5\xaa\x16\x52\x73\x6a\x99\xc4\x83\x75\xf3\x38\xfd\x84\x01\x2d\xa3\x7e\xbc\x96\xa1\x56\x01\x10\xec\x0e\x03\x37\x33\xc0\x38\xaa\x1f\x69\x67\xa6\x49\x2f\x10\x7d\xec\x30\x66\xd8\xb4\x71\x9c\xc9\xc7\xfe\x61\x39\x45\x77\xc9\x1d\x91\x5c\x8a\x7b\xdf\x1f\xb8\x5d\x7e\x99\x3c\x00\xe3\xf2\x04\x90\xb6\xe4\x17\x96\xe6\xa3\xf0\xb9\x39\x05\xda\xf2\x5b\x1b\x0a\xb2\x26\xec\x3d\x67\x3a\xeb\x42\x01\x6a\x95\xe6\x36\xba\x00\x85\x91\xa6\xaf\x62\xee\xb3\x8c\xd4\xe0\x3f\xf7\xd8\xc1\x32\xbf\x25\xea\x56\xd8\x6b\x7b\x95\x6f\xd7\xcc\xec\x21\xf9\x08\x52\xc2\xaf\xe0\x39\xaf\x6c\x21\x8d\x5c\x77\xfc\x86\x0b\xec\xfe\x6f\x30\xba\x66\xbb\x91\x55\xf1\xfd\xbd\x8a\x7c\x1e\xd0\xc0\xb1\xec\x90\x22\x66\x64\xed\xd3\xdf\xaf\xbf\x40\xbc\x07\x00\x00\xff\xff\x20\xeb\xaa\x4f\x8c\x02\x00\x00"),
		},
		"/templates/poddistruptionbudget.yaml": &vfsgen۰CompressedFileInfo{
			name:             "poddistruptionbudget.yaml",
			modTime:          time.Time{},
			uncompressedSize: 431,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xbd\x6a\x2b\x31\x10\x46\xfb\x7d\x8a\xc1\xbd\x75\x31\x5c\x52\x6c\x97\x90\xd2\x84\x90\xc2\xfd\xac\xf4\xd9\x19\x32\x3b\x12\xfa\x31\x84\x8d\xdf\x3d\xc8\x8b\x21\x90\x14\x6e\x85\xce\x39\x23\xcd\xb2\x6c\x49\x8e\xe4\x0e\xac\x0d\xc5\xa5\x18\x9e\xa5\xe4\x96\xaa\x44\x7b\x6a\xe1\x84\xea\x60\x3c\x29\x02\x5d\x2e\x03\x27\x39\x20\x17\x89\x36\x52\x8a\x2a\xfe\xf3\xdf\x79\x37\xa1\xf2\x6e\xf8\x10\x0b\x23\xbd\xfe\xe6\x87\x19\x95\x03\x57\x1e\x07\x22\xe3\x19\x23\x2d\x0b\x89\x79\x6d\x01\xb4\x91\x99\x4f\x48\x4d\xb5\xc0\x67\xd4\xad\x8f\x56\x73\x54\x45\x76\xc7\xa6\xda\x81\x0d\xb9\x1e\x5f\xe9\x92\xd8\xaf\x0a\xf7\x06\x05\x17\xb8\x97\xdb\xf1\x7a\x4b\x79\x82\x96\x5e\x23\xba\x3e\xef\x8e\xd4\xca\xf4\xd0\x17\x99\x58\x80\x55\xfa\xdf\x75\x25\xc1\x77\xd5\x2c\xf6\x78\x66\xd1\xfe\x15\x23\xed\x06\xa2\x02\x85\xaf\x31\xaf\xa1\x99\xab\x7f\xdf\xff\x28\xdf\xdf\xbe\x89\xf6\x7f\xcc\xf0\xd0\x67\xe8\x22\xd8\x75\x01\xdf\x01\x00\x00\xff\xff\x4c\xf7\xc1\xf2\xaf\x01\x00\x00"),
		},
		"/templates/rbac.yaml": &vfsgen۰CompressedFileInfo{
			name:             "rbac.yaml",
			modTime:          time.Time{},
			uncompressedSize: 1259,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x94\xcf\x8a\x14\x41\x0c\xc6\xef\xfd\x14\x61\x8e\x42\x97\x08\x1e\xa4\x6f\xea\xc1\x9b\x87\x11\x16\x44\x3c\xa4\xab\x32\x33\x71\xd3\xa9\xa6\x92\x9a\x85\x1d\xf7\xdd\xa5\xba\x5d\x1d\xd7\x3f\x2c\xec\x22\x7b\xea\x74\x48\x7e\x5f\xea\x4b\x51\x38\xf3\x05\x15\xe3\xac\x03\x1c\x5f\x74\x97\xac\x69\x80\x0f\x54\x8e\x1c\xe9\x75\x8c\xb9\xaa\x77\x13\x39\x26\x74\x1c\x3a\x00\xc5\x89\x06\x38\x9d\x80\x35\x4a\x4d\x04\x1b\x9e\x70\x4f\x73\x15\x31\x8a\x85\xbc\x8f\x59\xbd\x64\x11\x2a\x61\x57\x45\x5a\xc3\x06\x02\xdc\xdc\x7c\xef\xb6\x19\xe3\x8a\x08\x5b\x12\x42\xa3\xf0\xfe\x36\xbd\x56\x09\x8e\x24\xd6\xd4\x00\x4e\xa7\xfe\x5e\x52\x6b\x4f\x13\xfa\x0a\xca\x9a\x48\x1d\x5e\xae\xb8\x86\xb8\x62\x3f\x40\xb8\x40\xa9\x64\xc1\x7e\x39\x5e\x40\xd5\xec\xe8\x9c\xd5\xd6\x86\xb3\xc4\xcf\x21\x3c\x7f\xc4\x49\xfe\xc6\x27\x4d\xed\xa7\xeb\xfb\xbe\x3b\xb7\xb4\x8c\x18\x03\x56\x3f\xe4\xc2\xd7\x0b\x32\x5c\xbe\xb2\xc0\xf9\xf9\x0f\xb3\xdf\x4a\x35\xa7\xb2\xcd\x42\x8f\xe8\xf4\xa3\x7a\x58\xaa\x90\x0d\x5d\x0f\x38\xf3\xbb\x92\xeb\x6c\x03\x7c\x5a\x49\x16\x46\xd4\x6b\xe4\x28\xb9\xa6\xc0\x79\xf3\xb9\x03\x28\x64\xb9\x96\x48\x4b\xd9\xb3\x25\x75\xa4\x32\x2e\xe3\xf4\xb0\x27\x5f\xbe\xc2\xb6\x06\x57\xe8\xf1\xb0\x44\x75\x4e\xe8\x74\x57\xe8\x0e\x73\xa9\x8c\x59\x77\xbc\x9f\x70\xb6\x7b\xc3\x63\xa1\x06\x3f\xd3\x69\x61\x22\xa1\x26\xf9\xa0\xc5\xbd\x61\x4d\xac\xfb\x27\xbb\xbf\x2c\xb4\xa5\x5d\xa3\xfd\x7e\xe7\x1e\x3a\xe9\xed\xaa\xfe\xe1\x59\x67\x75\xfc\x42\xd1\x97\x3b\xf4\xc7\x27\xe6\x7f\x3c\x2c\xdf\x02\x00\x00\xff\xff\x2a\x14\xd1\x4e\xeb\x04\x00\x00"),
		},
		"/templates/scraping.yaml": &vfsgen۰CompressedFileInfo{
			name:             "scraping.yaml",
			modTime:          time.Time{},
			uncompressedSize: 848,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xcd\x8a\x14\x31\x10\x80\xef\xfd\x14\xc5\xdc\x27\x22\x88\x48\x5f\x3d\x09\xeb\x1e\x5c\x59\x10\xf1\x50\x93\xd4\x4c\x87\xad\xfc\x98\xaa\x0c\x68\xef\xbe\xbb\xa4\xd3\xad\x3b\xab\xab\xde\x42\xa8\xfa\xbe\xfa\x9b\xe7\x3d\xf8\x23\xd0\x57\x30\xb7\xc8\x95\xc4\x58\xae\xa2\x54\xae\x31\x10\xec\x02\xb6\xf7\x0e\x1e\x1e\x06\xcc\xfe\x96\x8a\xf8\x14\x47\x08\x29\x7a\x4d\xc5\xc7\x93\x39\xa0\xbd\xfb\x86\xc5\x89\x39\x60\xfc\x8e\xde\x72\xaa\xce\xf8\xf4\xe2\xfc\x72\xb8\xf3\xd1\x8d\x70\x43\xe5\xec\x2d\xbd\xef\x39\x43\x20\x45\x87\x8a\xe3\x00\x10\x31\xd0\x08\xf3\x0c\x3e\x5a\xae\x8e\x60\xe7\x03\x9e\x28\x57\x66\x21\x5b\x48\xf7\x36\x45\x2d\x89\x99\x8a\x39\x56\xe6\x96\xb0\x03\xd3\xea\xe9\xd9\x92\xd1\x76\x84\xf9\x40\x4c\x28\x64\xae\xb7\xef\x1e\xc5\x78\x20\x96\x66\x03\x58\xba\xfd\x0f\x55\xcf\x69\xa2\x7b\x88\x3e\x3a\x8a\x0a\xaf\x1a\x4e\x32\xd9\x86\x12\x62\xb2\x9a\x4a\xc7\x06\x54\x3b\x5d\x3d\xf2\x5c\x9a\xfe\xaa\xda\x48\x57\x7f\x50\xbe\x81\x27\x8d\xde\xfc\xae\x5d\xba\x1d\xe1\xf3\xb3\x23\x80\x2f\x03\x00\x45\x97\x93\x8f\xba\xd4\xd7\x2a\x53\x2a\x67\xe4\x3e\xb8\x75\xf1\x72\xb1\x27\x23\xb6\x60\xa6\x77\x6b\x64\x2f\x04\x20\xa7\xa2\x23\x4c\xaa\x79\x1f\x48\x8b\xb7\xb2\x7c\x4f\x29\xa6\xf2\xd1\x07\x12\xc5\x90\x65\x04\x2d\x95\x86\xf5\xb8\x9e\x11\x28\xcb\xdb\x14\x8f\xfe\xb4\xb1\xc5\x4e\xd4\xee\xa1\xd1\x3b\xf6\x67\xc8\x38\xcc\x33\x68\xfa\x84\x81\xff\x8d\xbb\x87\x75\x7e\xaf\x1b\xb9\x15\x41\xd1\x6d\x92\x42\xcb\x72\x7d\x3c\x3d\xda\xd5\xaf\xa3\xc8\x25\x05\xd2\x89\xaa\x98\xb5\x3f\xe3\xe8\x88\x95\x75\xeb\xf7\x72\x45\x4f\x15\x3f\x02\x00\x00\xff\xff\x7a\x21\x2f\xb5\x50\x03\x00\x00"),
		},
		"/templates/service.yaml": &vfsgen۰CompressedFileInfo{
			name:             "service.yaml",
			modTime:          time.Time{},
			uncompressedSize: 467,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x90\x4f\x4b\xc4\x30\x10\xc5\xef\xfd\x14\xc3\xde\x1b\x10\x3c\xf5\xea\x55\x64\x51\xd9\xfb\x98\x3e\xd7\xe0\x34\x09\x99\xe9\x82\xd4\xfd\xee\x92\x4d\x16\x2f\x0a\xda\x5b\xe7\xfd\xf9\xf1\xc2\x39\x1c\x50\x34\xa4\x38\xd1\xe9\x66\x78\x0f\x71\x9e\xe8\x09\xe5\x14\x3c\x86\x05\xc6\x33\x1b\x4f\x03\x51\xe4\x05\x13\x6d\x1b\x85\xe8\x65\x9d\x41\xbb\xb0\xf0\x11\x79\x15\x51\xf8\x02\x1b\x7d\x8a\x56\x92\x08\x8a\x7b\x5d\x45\x6a\x60\x47\x8e\xce\xe7\x9e\xd6\xcc\xbe\x55\xb8\x47\x08\x58\xe1\x1e\xae\xe7\xe6\x12\x7e\x81\x68\xa5\x11\x6d\xdb\xf8\x27\x54\xcb\x54\xd0\x27\xc5\x10\x67\x44\xa3\xdb\x5a\xa7\x19\xbe\x56\xd9\x47\xee\xd4\x03\xcb\x0a\x75\xda\xd6\xb9\x2a\x34\x6e\x4e\xc5\x3a\x76\xbc\xfc\xfc\xe8\xaf\x42\xf3\xd7\xcf\xb8\x1c\x61\xfb\x8b\x79\x81\x95\xe0\xb5\x2b\xb9\x24\x4b\x3e\xc9\x44\xcf\x77\xfb\x7e\x6b\xaf\xf7\x66\x96\xc7\x6f\xb3\x42\xe0\x2d\x95\xff\x0d\xbe\xa6\xee\x7f\x1b\xfe\x15\x00\x00\xff\xff\xc7\x19\xdc\x99\xd3\x01\x00\x00"),
		},
		"/values.yaml": &vfsgen۰CompressedFileInfo{
			name:             "values.yaml",
			modTime:          time.Time{},
			uncompressedSize: 1055,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x93\xcf\x6e\x1c\x37\x0c\xc6\xef\x7a\x0a\xc2\xbe\xae\xe3\x4d\x53\x17\x89\x6e\x5b\xa7\x28\x0c\xd4\xc1\x16\x69\x0f\x45\xd1\x03\x57\xc3\x19\x13\xd0\x88\x0a\x49\xad\xbb\x35\xfc\xee\x85\x66\x62\xc3\x48\x5a\x20\x97\xc5\xea\x9b\x1f\xc5\x3f\x1f\x75\x0e\xef\x69\xc4\x96\x1d\x8e\x98\x1b\x19\x8c\xa2\xc0\x33\x4e\x54\x5b\xce\x46\x49\xc9\x2f\x92\x14\x57\xc9\x99\x34\x9c\xc3\x6f\x77\x6c\xc0\x06\x08\x7f\xec\x6e\x7f\xb9\x18\x45\x67\x74\xa7\x01\x46\xce\xf4\x2a\xf4\x0b\x53\x46\x25\x38\xa2\x32\x1e\x32\x19\xb8\xc0\x81\xa0\xa2\x19\x0d\xc0\xc5\x05\x4e\xd2\x14\x9c\xe6\x9a\xd1\xc9\x5e\x85\x90\x72\x33\x27\xfd\x80\x33\x45\x98\xb1\xff\x0f\x55\x65\x26\xbf\xa3\x66\x31\x00\x60\xe5\xdd\x30\x28\x99\xd1\x72\x06\xb8\x80\x3b\xf7\x1a\x2f\x2f\x5f\x80\xef\xb6\xef\xb6\x01\xe0\x53\x23\x3d\x7d\xc6\xe3\xff\x63\x03\x2b\x25\xff\xf5\x5b\xe0\xa0\x54\x33\x27\xb4\x08\xaf\x43\x60\x73\x96\x5e\x85\xd2\x91\x8d\xa5\x44\x38\x3b\x0b\xa1\xca\xb0\x2b\x45\x1c\x9d\xa5\x58\x84\x3f\xff\x5a\xb4\x8f\x94\x9a\xb2\x9f\xae\xa5\x38\xfd\xed\x4b\x5c\x2b\x3b\xfb\xdd\x48\x23\xfc\x70\x75\xf5\xe6\xfb\x27\xe9\x67\x95\x56\x9f\x34\xfb\x3a\x0e\x73\x96\xfb\xbd\xf2\x91\x33\x4d\xf4\x93\x25\xcc\x4b\xb2\x08\x23\x66\xa3\xb0\x58\xb7\x16\x56\xc5\xd8\x45\x4f\x11\x0e\x58\xfe\x41\x4e\x59\xda\x70\xf9\x85\xb7\x16\x00\x1c\xa7\xe7\x99\x03\xf4\x6f\x7b\xc9\x9c\x4e\x11\x6e\xc6\x0f\xe2\x7b\x25\xa3\xe2\xa1\xc8\x40\x1f\x29\x53\x72\xd1\x08\x0f\x8f\x01\xc7\x91\x0b\xfb\x69\x39\xb8\x64\xd2\x17\x7d\x2b\x99\x34\x4d\xab\x57\x4a\x9f\x1a\x99\x7f\xf6\x6d\xa6\x79\x29\xeb\xec\xf5\x76\x7b\xcb\x67\x8b\x96\x6a\x5b\x85\xb9\x9f\x33\xcf\xfc\x15\xfd\xdd\x97\xf4\x9b\x85\x0e\x46\x7a\xe4\xb4\x34\xed\xa7\x4a\x11\xae\xd7\x5d\xba\xd9\xf7\x6e\x44\x3d\xc2\xdb\xed\xdb\xed\x33\xb8\x4b\x49\x5a\x59\xa7\xf9\xd2\xac\x87\xc7\x67\xe4\x56\x4a\x1f\x5d\x47\x2c\x29\x56\xba\x29\x4e\x7a\xc4\x1c\xe1\x6a\x99\x58\xb6\x6b\x29\x23\x4f\x6b\xd4\x40\x47\xca\x52\x67\x2a\x7e\x2b\xc3\x52\x09\x95\xbe\xfa\xc3\x93\x2d\x7d\x0b\xde\xb3\x69\xab\x3d\xd9\x8f\x6d\x98\xc8\xff\x0b\xcb\x32\x75\xf9\xe1\x31\x00\x9c\x43\xee\xf7\x46\xe0\x32\x0a\x9c\x03\xa6\x44\xb5\x3f\xb4\xf5\xa5\x46\xa8\x58\x38\x6d\x60\x44\xc7\xbc\x01\x52\x15\xdd\xc0\x3d\x6a\x59\x7f\xb9\x4c\x9b\x25\x76\x03\x03\x1d\xda\xb4\x01\x57\x4c\xd4\x37\x79\x62\x73\x3d\xf5\x4c\x6d\x59\x41\xec\x93\x42\xb3\x7b\xd1\x21\xc2\xa1\xeb\x9a\x23\xa4\xf0\x6f\x00\x00\x00\xff\xff\xd6\xbc\x50\x16\x1f\x04\x00\x00"),
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
		fs["/templates/deployment.yaml"].(os.FileInfo),
		fs["/templates/imagepullsecret.yaml"].(os.FileInfo),
		fs["/templates/poddistruptionbudget.yaml"].(os.FileInfo),
		fs["/templates/rbac.yaml"].(os.FileInfo),
		fs["/templates/scraping.yaml"].(os.FileInfo),
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
