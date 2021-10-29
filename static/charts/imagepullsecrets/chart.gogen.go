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
			uncompressedSize: 275,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x5c\x8f\x41\x4e\xc5\x30\x0c\x44\xf7\x3e\x85\x2f\x40\x0b\x42\x6c\xb2\x42\xc0\x19\xfe\xde\x4d\xfd\x5b\x8b\xc4\x89\xe2\xa4\x52\x39\x3d\x4a\x4b\x25\x60\x3b\x7e\x33\xe3\xa1\x2c\x37\x2e\x26\x49\x1d\x6e\x4f\xa0\x14\xd9\xa1\x44\x5a\x38\xb7\x10\x8c\x7d\xe1\x6a\x30\xb3\xf9\x22\xb9\x1e\xd8\xc7\xae\x14\xc5\x53\x08\x3b\x46\x52\x5a\x18\x3b\x8b\x3f\x30\xde\x53\x39\x13\x0c\xd6\xd4\xe3\xd6\x5a\xb3\xb9\x71\x9c\x48\xbf\x48\x7c\x48\x6d\x1e\x7c\x8a\x00\x91\x44\x2b\x89\x72\x31\x07\x0f\xc8\x91\x24\x38\x14\xbd\xa7\xd7\xff\x2c\xe2\xf9\xda\xdb\xa1\xe3\x7b\x3f\x80\xa5\x56\x3c\x1f\xde\xab\x63\x91\xba\xb6\xa9\x5b\x7e\xd7\x8d\x13\xf9\xcf\x9d\xca\x6c\x00\xdb\x35\xf7\x71\x78\x1e\x5e\x80\x72\xbe\xfd\x55\xbe\x03\x00\x00\xff\xff\xe3\x04\xb6\x1b\x13\x01\x00\x00"),
		},
		"/crds": &vfsgen۰DirInfo{
			name:    "crds",
			modTime: time.Time{},
		},
		"/crds/crds.yaml": &vfsgen۰CompressedFileInfo{
			name:             "crds.yaml",
			modTime:          time.Time{},
			uncompressedSize: 18528,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x5c\x5b\x6f\x1b\xb9\x15\x7e\xd7\xaf\x38\x70\x1f\xd2\x02\xd6\xb8\xd9\x05\x8a\x42\x6f\x86\x13\xa0\x46\xb7\xae\x61\x19\xd9\x87\xa2\x0f\xd4\xf0\x68\xc4\x35\x87\x9c\x25\x39\x4a\xb4\x8b\xfc\xf7\x82\x97\xb9\x5f\x44\x39\x0e\x5a\x14\xe3\x97\x40\x1c\xf2\xf0\xdc\xcf\x77\x86\xe1\x90\x82\x7d\x42\xa5\x99\x14\x1b\x20\x05\xc3\x2f\x06\x85\xfd\xa5\x93\x97\xbf\xea\x84\xc9\x9b\xe3\xfb\xd5\x0b\x13\x74\x03\x77\xa5\x36\x32\x7f\x42\x2d\x4b\x95\xe2\x07\xdc\x33\xc1\x0c\x93\x62\x95\xa3\x21\x94\x18\xb2\x59\x01\x10\x21\xa4\x21\x76\x58\xdb\x9f\x00\xa9\x14\x46\x49\xce\x51\xad\x33\x14\xc9\x4b\xb9\xc3\x5d\xc9\x38\x45\xe5\x88\x57\x5b\x1f\xff\x9c\xfc\x25\xf9\x61\x05\x90\x2a\x74\xcb\x9f\x59\x8e\xda\x90\xbc\xd8\x80\x28\x39\x5f\x01\x08\x92\xe3\x06\x58\x4e\x32\x2c\x4a\xce\x35\xa6\x0a\x8d\x4e\xdc\x80\x4e\x76\x44\xfc\x46\x58\xca\x65\x49\x13\x26\x57\xba\xc0\xd4\x32\x90\x29\x59\x16\x61\xd5\x60\x92\xa7\x19\x18\xf5\x42\xde\xdb\x89\x8f\x25\xe7\x5b\x47\xde\x3d\xe1\x4c\x9b\xbf\x8f\x3d\xfd\x89\x69\x3f\xa3\xe0\xa5\x22\x7c\xc8\x9c\x7b\xa8\x0f\x52\x99\x87\x66\xa3\x35\xb0\xbc\x08\x8f\x98\xc8\x4a\x4e\xd4\x60\xe5\x0a\x40\xa7\xb2\xc0\x0d\xdc\xf1\x52\x1b\x54\x2b\x80\xa0\x2b\x47\x65\x0d\x84\x52\xa7\x7d\xc2\x1f\x15\x13\x06\xd5\x9d\xe4\x65\x2e\xea\x3d\x28\xea\x54\xb1\xc2\x38\xed\x3e\x61\xa1\x50\xa3\x30\x1a\xd8\x1e\xcc\x01\x41\xee\x7e\xc1\xd4\xc0\x81\x68\xd8\x21\x0a\xd0\x65\x9a\xa2\xd6\xfb\x92\xf3\x13\x28\x4c\xa5\x48\x19\x47\xea\x88\x01\xec\xa5\xca\x89\xd9\xc0\xee\x64\x30\x0c\xfd\xa2\xa5\x78\x24\xe6\xb0\x81\x44\x1b\x62\x4a\x1d\xfe\x09\x8f\xbd\xb1\xb6\x86\xd4\x0b\xcc\xc9\x8a\xa3\x8d\x62\x22\x1b\xe3\xf1\xe7\x03\x8a\x6f\x60\x8d\x92\x19\xd6\x38\xd1\x66\x5b\x93\x79\x0a\x34\x98\x73\xb4\x0e\xc3\x4f\x7d\xf2\x9e\xeb\x9a\x78\x8f\xe7\xbf\xc9\xcf\xc0\xa5\xc8\x1c\xdf\x19\x0a\x54\xc4\x20\xb5\x3e\x4c\x51\x18\x46\x38\x30\x0d\x47\xc2\x19\xb5\x7c\x02\xd9\x1b\x54\x6e\xae\x65\xa8\x96\xa5\xc3\x47\x25\x0f\x13\xe6\xc7\x1f\x26\x05\x72\x34\x99\x39\x6d\x2d\x09\xda\x55\xfa\xa7\xf0\x0c\x74\xe7\xa1\x97\xc4\xba\x4a\xe6\xdc\x69\x20\x8c\x75\x51\x90\xde\x3d\xbc\x17\x36\x22\xc5\xf9\x41\x81\x69\x62\x88\xca\xd0\x24\x9e\x40\x62\x39\xea\x7a\x84\x27\xfc\xd0\x8c\x9f\xf3\x8b\x36\x5b\x2e\x5c\x0b\x92\xa2\x6e\x73\xc9\x74\x4b\xf7\x4c\x5c\xe2\xb3\x39\x11\x24\x43\xfa\x50\xd3\xed\x30\x3b\x18\x1e\xf0\xea\xe7\x1d\xdf\x13\x5e\x1c\xc8\x7b\x1f\xd3\xe9\x01\x73\xb2\x09\x2b\x64\x81\xe2\xf6\xf1\xfe\xd3\x8f\xdb\xce\x30\x74\x65\xec\xe5\x15\x2b\x91\x95\xcf\xaf\x71\xae\x63\x7f\xf6\xb3\x0b\xdc\x3e\xde\xd7\xf4\x0a\x25\x0b\x54\x86\x55\x59\xc6\xff\xb5\xb2\x7b\x6b\xb4\xb7\xfb\x3b\xcb\xa0\x9f\x05\xd4\xa6\xf5\xa0\xde\x90\x70\x90\x06\x99\xbc\x15\x98\x06\x55\xa5\x93\xb6\xe3\x56\x7f\x72\x0f\x44\x84\x20\x4e\x60\x8b\xca\x92\xb1\x49\xb0\xe4\xd4\x56\x83\x23\x2a\xef\xfa\x99\x60\xbf\xd5\xb4\x35\x18\x19\x22\xc3\x60\xc8\xaa\xcd\x9f\x4b\x70\x82\x70\x1b\x4b\x25\x5e\x03\x11\x14\x72\x62\xb3\x81\x4b\x15\xa5\x68\xd1\x73\x53\x74\x02\xff\x90\x0a\x81\x89\xbd\xdc\xc0\xc1\x98\x42\x6f\x6e\x6e\x32\x66\xaa\xaa\x96\xca\x3c\x2f\x05\x33\xa7\x1b\x57\xa0\xd8\xae\x34\x52\xe9\x1b\x8a\x47\xe4\x37\x9a\x65\x6b\xa2\xd2\x03\x33\x98\x9a\x52\xe1\x0d\x29\xd8\xda\xb1\x2e\x5c\x65\x4b\x72\xfa\x07\x15\xea\xa0\x7e\xd7\xe1\x75\xe0\x21\xfe\xcf\x15\x97\x19\x0b\xd8\xf2\x62\xad\x4e\xc2\x52\x2f\x45\xa3\x68\x16\x72\xcc\xd3\xc7\xed\x33\x54\x5b\x3b\x63\xf4\xb5\xef\x93\x67\xbd\x50\x37\x26\xb0\x0a\x63\x62\xef\x32\x10\xd3\xb0\x57\x32\x77\x34\x51\xd0\x42\x32\x61\xdc\x8f\x94\x33\x14\x7d\xf5\xeb\x72\x97\x33\x63\xed\xfe\x6b\x89\xda\x58\x5b\x25\x70\xe7\x4a\x3d\xec\x10\xca\xc2\xe6\x47\x9a\xc0\xbd\x80\x3b\x92\x23\xbf\x23\x1a\xbf\xbb\x01\xac\xa6\xf5\xda\x2a\x36\xce\x04\x6d\x94\xd2\x9f\xec\xb5\xd6\x7a\x50\xc1\x87\x09\x7b\xf5\xe2\x75\x5b\x60\xda\x89\x1c\x8a\x9a\x29\xeb\xdb\xb6\xfe\xd9\x88\x18\x83\x15\xf3\x91\x6b\xff\x14\x66\x4c\x1b\x75\xea\x8f\x43\xbf\xba\xfb\x69\x0e\x6b\x11\x26\x2a\x1e\x0c\x61\x5c\xf7\x32\xba\x91\xd6\x62\x0e\x64\xd5\x59\xbd\x1b\x69\x80\x24\x3d\x34\x99\x76\x30\x67\x9a\x5d\x07\xf7\xea\xd2\x37\xfa\xb8\xc7\xf9\x5d\x33\xdb\xa9\x9c\xed\x19\x6a\xf8\x7c\x60\xe9\xa1\xcb\x70\xa9\x91\x02\xd1\xa3\x14\xc1\xcb\xe7\x43\xc2\xe6\x4a\x2a\xd3\x17\x54\xc0\x65\xc6\x44\x9b\xa1\xd1\xd5\xcc\x60\x3e\xc1\xea\x39\x59\x5b\xa5\x62\xf2\xe9\xa4\x3b\x0e\x89\x38\x75\x7f\x23\x25\x1b\xa1\xd6\xf3\xa6\xc8\xac\xa1\x55\x91\xc7\x1f\x8e\x5b\x7d\x36\x5a\xfa\x13\x88\x52\xe4\xb4\x8a\xe7\x6d\x3d\x6b\xa5\x99\x4d\x3d\xd8\x38\x13\x1e\xcf\x6e\x52\xc7\xbf\x88\xa9\x0a\xd2\x0e\x6b\x60\xd1\x8d\x94\x71\xdf\x0e\x10\x03\x08\x64\xec\x88\xe2\xd5\x61\xd2\x00\x99\x88\x28\x69\x60\x48\x10\xc2\x45\xba\x87\xff\x1a\xa4\x18\x42\x23\x62\x26\x0c\x18\xa4\x3e\x90\xa3\x97\xdb\xeb\x66\x3b\x8e\xf6\x2e\x8d\x85\x41\xef\x77\x56\xb2\x2d\x72\x4c\x6d\xde\xaf\x05\x73\x35\xa8\x16\xee\x3a\xe4\x02\xa2\x10\x72\x62\xd2\xc3\x04\x6f\x61\xff\xcc\x26\x3f\xd3\x57\x07\x27\x3b\xe4\x0e\x5f\x50\x4c\x19\x45\xdf\x02\x31\x3d\x9b\x94\x47\xd5\xb6\x43\x20\x45\xc1\x19\xd2\x0a\xad\xf4\xbc\xe0\xda\xd2\xce\x4b\x6e\x58\xc1\xad\x23\x05\xe9\xe6\x78\x56\x58\x7b\x26\xb5\xab\xa5\x40\x0b\x04\xbc\xb4\x1d\x98\xfb\x99\x71\x6e\x59\x08\x98\x75\x86\xe8\x1f\xff\xf9\xf4\xa7\xc9\xc7\xb3\x09\x2f\x36\xe9\xb9\x82\x6a\x59\xbc\x8d\xb1\x79\x10\xb5\xd5\xb1\xc6\xed\x10\x9d\x3e\x23\xb3\x53\x87\xf3\x8f\x5f\x2c\x44\xd2\x31\x9c\x9f\x55\xd9\x88\x6b\xdf\x7a\xbf\xab\x9d\xa0\x4a\x7f\x39\x0a\x13\x90\x5e\x78\x72\x96\x2e\xb8\x88\x6e\xaa\xbb\x47\xb8\xd7\x40\xe0\x05\x4f\x1e\x0c\x5b\xc4\x5d\xd8\xe0\x75\xad\x02\x31\x11\x34\x15\x3a\xa8\xed\x3c\xec\x05\x4f\x8e\x4c\xc0\xce\x67\x57\x17\x17\x18\x10\x2c\xf5\x98\x69\x3d\x05\x5a\x9e\x42\x1f\xe4\x35\x69\x07\x9c\x22\x7c\x4c\x44\x2b\xcf\xf7\x41\x36\x6c\x1d\x70\x8d\x5a\x71\x81\xd7\x85\x2e\xcf\xe9\xfe\x15\x62\xd6\x66\x6b\x20\xbb\x37\xec\x3b\xed\x4d\x64\x1d\xf4\xc0\x8a\x48\x41\x8d\x74\x9e\x65\x6c\x31\xab\x3a\x21\xf7\x42\xa0\xde\x48\xbb\x9c\x73\x2f\xae\x23\x29\x3e\x48\x73\x2f\xae\xe1\xe3\x17\x66\x91\xbf\xf5\x92\x0f\x12\xf5\x83\x34\x6e\xe4\xbb\xa8\xd3\x33\xfe\x0a\x65\xfa\x85\x2e\xbc\x84\x47\x21\x56\x0f\xed\x96\x4a\x27\x91\x62\xdf\x87\xf7\x64\x95\x79\x6c\xbd\x10\x20\x55\xa5\x0f\xd7\x20\xfb\xed\xc6\xe1\xce\x44\xea\x29\xb5\xeb\x99\x84\x14\x6b\xcc\x0b\x73\x4a\xc6\x76\xf2\xca\x8e\x24\x29\x55\xc7\x22\x43\xd6\xea\x4d\xfd\x86\x91\x64\x9f\x6d\x89\xf4\xeb\x7d\xcb\xcf\x49\x8a\x14\x68\xe9\x94\xe9\x1a\x55\x62\x30\x63\x29\xe4\xa8\x32\x8c\xa4\x5a\xd8\xd4\x1b\xc7\x42\x64\xd6\x7d\x95\x87\xcd\x23\xd5\x4b\x51\x75\x1b\xc9\xbe\xe0\x29\x62\x56\x65\xec\xc8\xf2\x17\x51\xd3\x62\x25\x8a\x22\x78\x9e\x94\x07\x55\x91\x30\xef\x27\x8f\xc0\x16\x8c\xf7\x3f\x83\xf1\xe6\xa1\x8a\x83\x27\x7e\xec\xd7\x12\xd5\x09\xe4\x11\x95\xaf\x2b\xb3\x3e\x28\xf7\xf5\xbb\x29\x9d\xc0\xf3\x01\xed\xcf\x92\xbb\x62\xe4\x84\x0d\x8e\xe0\x5f\xdf\x75\x31\xd8\x2c\x61\xab\xc1\xdb\x87\x0f\x48\x13\xb8\x15\x3e\x8d\xf5\x79\xae\x94\x49\x38\x0f\xde\x7d\x26\xd5\xdf\xba\x73\xa5\x29\x32\x42\xc6\x50\xb9\x08\x2a\x5f\x00\x38\x3b\xe6\xe9\x2f\x0e\xe6\x61\xda\xe9\xb5\xcb\x7f\x6c\x2a\xcb\xfd\xab\x41\x6f\xa2\x66\xa4\xa5\xe8\x05\x12\x2f\x90\x78\x81\xc4\x0b\x24\x5e\x20\xf1\x02\x89\xff\xcf\x21\x31\xb4\xc1\xc9\x7f\xfd\x2d\xd6\xb0\xf6\x07\xd4\xe4\x2a\x64\x4e\x0a\x1b\xbf\xbf\xdb\x32\xe7\x1c\xfa\x2b\x14\x84\xa9\x88\x18\xbe\x75\xff\xdb\x84\x63\x67\x2d\xf3\x2f\x8f\xdb\xdb\xd8\x1d\x98\x06\x6b\xdf\x23\xe1\xc3\xc3\xc0\xf1\x04\x2b\x00\xb9\x2f\xe4\x15\xd6\x6b\x21\x16\x0b\xf1\xa5\xf6\x35\x75\xcf\x90\xd3\xb3\x24\x99\x86\xab\x17\x3c\x5d\x5d\x0f\xf2\xc0\xd5\xbd\xb8\xf2\x05\xfe\xe2\x74\x53\xa3\x05\x29\xf8\x09\xae\xdc\xda\xab\x6f\x01\x41\x91\x9e\xf8\x46\x2d\x57\xeb\xff\x2a\x5d\x7e\x64\x60\x2b\x7d\xe3\xba\xed\xd6\x69\x37\xa7\x38\x77\x2e\x62\x64\x7d\x34\xd0\xf9\x6f\x17\xc2\xc8\xd7\x77\x20\x51\x11\x71\x4e\x29\x67\x34\xdb\x48\xf9\x33\x33\x87\x47\x49\x63\x4e\x5b\xec\xb4\xf8\x56\xb5\x6a\x48\x67\x4e\x26\x0b\x4b\x70\x70\xae\xfa\xea\x5e\xf5\x4c\x9f\x5a\x48\xfa\x4e\x4f\x74\xab\x53\x14\xeb\x53\x98\xcb\x3b\xd5\xd5\x54\x42\x75\xfd\xeb\x74\x97\xba\x9c\x26\x2d\xa7\x49\xcb\x69\xd2\xd2\x3a\x2f\xad\xf3\xd2\x3a\x2f\xad\xf3\xd2\x3a\x2f\xad\xf3\x72\x9a\xb4\x9c\x26\x2d\xa7\x49\xcb\x69\xd2\x72\x9a\xb4\x40\xe2\x05\x12\x2f\x90\x78\x81\xc4\x0b\x24\x5e\x20\xf1\x72\x9a\xb4\x9c\x26\x2d\xa7\x49\x6f\xde\x72\x9d\x21\xe2\xfb\x87\x88\xc3\x92\xf6\xfd\x91\x3b\x29\xf6\x2c\x0b\xcf\x77\x01\xf8\x34\x40\x26\x5c\xaf\x99\x3c\x83\xf0\x57\x5a\x5b\x77\x6c\xdc\x79\x89\x87\x21\x48\xcf\x5c\x4d\x7a\xa3\xe3\x84\x4b\xc3\x31\x2a\x10\xbb\xa8\xb5\x61\x63\xfc\x64\xae\xc5\xe7\x6a\x36\x48\x76\x08\x45\x69\xaa\xeb\x3f\xe1\x0a\x4f\x94\x5b\xcc\x4e\x3a\xd7\x8a\x7f\x7f\x15\x75\xfa\xfb\x9e\x76\x9a\xce\xdc\x8b\x3f\xb3\xf1\x5b\x2b\x66\xfe\x9e\xdf\xf0\xca\x7a\xc3\xff\xf8\x05\xb3\xf3\x91\xce\x84\x66\x14\x5d\x0f\xda\x20\xf2\xb3\xa1\x10\xa5\xf4\xf9\x82\x3b\x73\x4b\x70\x56\x51\x73\x37\xfc\x26\x6e\xd5\x4d\xd2\x1b\xa7\xb5\xae\xef\xc4\xf6\x86\xfd\x4d\xc0\xa8\xab\xbd\xee\xe6\xff\x05\x97\x7b\xdd\xfc\xce\xf5\x5e\xb9\xd3\xa8\x8e\xdf\x7e\xbf\x77\xee\xb3\x14\x43\x1d\xb6\xbf\x74\xb1\x36\x6c\xc4\x40\x33\x66\x1f\x7c\xe8\x60\x48\x7f\x12\x0e\xce\x7a\xd3\x74\xa5\x51\x48\xf4\x98\x1c\x33\xe4\xfc\xdb\x9e\xb6\xde\x87\xcb\x63\x13\x50\x04\xd7\xa3\x6e\xac\x27\xf6\x9d\xa1\xd7\xfb\x1a\xc7\xb4\xe9\xda\x1f\xf5\xe8\x53\x6e\x7f\x99\x63\x96\xcd\xc1\xa0\x77\xc6\x0d\x18\x55\x7a\x9f\xd0\x46\x2a\x92\x61\x7b\xa4\xdc\xd5\xef\xd2\x2a\xf6\x82\xa0\xf0\xfb\xd7\x55\x23\x33\x49\x53\x2c\x4c\x70\x94\xf6\x87\x71\xae\xae\x3a\x5f\xba\x71\x3f\x9b\xd7\xae\x1b\xf8\xd7\xbf\x57\x7e\x63\xa4\x9f\xaa\x0f\xd5\xd8\xc1\xff\x04\x00\x00\xff\xff\x51\x14\xab\xbe\x60\x48\x00\x00"),
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
			uncompressedSize: 2021,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xac\x54\x4d\x6f\x13\x31\x10\xbd\xe7\x57\x8c\x56\x54\x82\xa2\x75\x0e\x48\x1c\x22\xf5\x80\x0a\x07\x04\x2a\x12\x95\xca\x11\x79\xbd\xe3\xee\x08\xaf\xd7\xf8\x23\x34\x6a\xf3\xdf\x91\xed\xcd\x66\x93\x26\x61\x83\x7a\xb3\xd6\x6f\xde\x3c\xbf\x79\xb3\x8f\x8f\xf3\x4b\x58\x52\xbb\x00\x87\x1e\x24\x29\xf4\x2b\x83\x57\x6d\x70\x9e\x8b\x06\x17\x70\x39\x5f\xaf\x67\x11\x35\xfb\xf4\x60\xb8\xae\xc1\x37\x08\x9a\xb7\x08\x9d\x4c\x67\xd1\x70\xeb\xd9\xac\xc7\x95\x50\xa3\x24\x8d\x50\x50\xcb\xef\xd1\x04\xa5\x1c\x0a\x8b\x9e\xc5\x9a\x02\xca\x2d\x8a\x07\xe5\x81\x5d\xa7\xf2\x9b\x48\xc8\xee\xb8\x0a\xe8\x12\xf2\xdb\x12\xad\xa5\x1a\xe1\x09\xbc\x0d\x5a\xc0\xfb\x77\xe9\x48\xed\x6d\x90\x92\x1e\xa0\x28\x0b\xe8\xb9\x50\xd7\xf1\x98\x55\x5e\x5b\xe4\x1e\x81\x0f\x1d\x64\x50\x6a\x05\xbf\x03\x57\x24\x09\x6b\xe0\xc6\x24\xfd\x6c\xf6\x03\x33\x77\xc2\xfb\xd8\x21\xbe\xc5\x41\x85\x82\x07\x87\xe0\xba\x16\xe1\x4b\xa8\xd0\x6a\xf4\xe8\xf2\xab\x25\xa1\xaa\x1d\x70\x8b\xa0\xa8\x25\x8f\x35\xf8\x0e\x7c\x43\x0e\x5e\x57\xab\xe4\xc8\xc7\x9b\xdb\x88\x25\x7d\x0f\xce\xa0\x78\xc3\x66\x9f\x25\x58\x54\xc8\x5d\x6f\x9d\xe8\xb4\xe7\xa4\x5d\x36\x2f\x7f\x23\x0f\x7f\x48\x29\xa8\x10\x82\x8b\x3a\x1d\xf0\x24\xbe\x57\xfb\x4f\x83\x23\x76\xd7\x64\x92\x83\xa7\x9b\xcb\xc1\xd7\x1e\x72\xf4\x7e\x8a\xef\xca\x0d\x3c\xaf\xd2\x1b\x16\x57\xd3\x07\xbb\xd5\x38\xb8\x91\x49\xd8\xf7\x6c\x55\xae\xdd\xe8\xdc\xf9\x78\xa6\x38\x63\x49\x7b\x09\xc5\x85\x2b\x2f\x5c\xb1\xc7\x95\x9b\x4e\x8f\xd9\xe1\xe3\x4e\xf8\x46\x53\x8d\x1b\xb3\x44\xeb\xa8\xd3\x71\xa2\x69\xb2\x7d\x4c\x32\x4a\xf1\x0a\xd5\x84\xe9\x26\xf4\x76\xb4\xfb\x4f\x1a\xbb\x9d\xcf\x77\x7d\xd7\x27\xb0\x68\x14\x17\x08\xc5\xdb\x02\x8a\x9f\xc5\x39\x2b\x75\x4a\x52\x19\xe7\x66\x3b\xa5\xd0\x3e\xcb\x1e\x90\x16\x2a\xd4\x27\x53\xca\x60\xbd\x1e\x71\xec\xb9\x39\xa9\xed\xb4\x96\x2f\xd6\x2e\xcd\xca\x25\x93\xb8\x31\x0b\x38\xd1\xf4\xb0\x39\xac\x2f\x65\xbf\x86\xbf\x0a\xa3\x6e\x1e\x2f\x27\xb3\x8d\x98\x1a\x54\x2d\x73\xcd\x3c\x65\xe3\x24\xc1\x26\x3e\x47\x04\xb4\x5c\xf3\x7b\xac\xcb\x6a\x95\x58\x86\x05\xb9\x45\xbb\x24\x81\x87\x8b\x48\x3b\xcf\xb5\xc0\xdd\x92\xcd\xd2\x3e\xc7\xf7\x6b\x90\xe1\x39\xa3\x1f\x8c\x39\x1e\xd3\x83\x24\xa2\x6b\x4d\xa7\x51\xfb\x05\x1c\xf7\xe8\x40\x9d\xe1\xd6\x97\x9d\x3c\x6d\xd2\xc8\xda\xff\x48\x87\x43\x85\xc2\x77\xf6\x6b\x9f\x92\xf2\x45\x67\x7d\xae\xff\xe3\x07\xfc\x0d\x00\x00\xff\xff\xc9\x4f\xc9\x50\xe5\x07\x00\x00"),
		},
		"/templates/default_imps_cr.yaml": &vfsgen۰CompressedFileInfo{
			name:             "default_imps_cr.yaml",
			modTime:          time.Time{},
			uncompressedSize: 906,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xac\x92\x4f\x4f\xe3\x40\x0c\xc5\xef\xf9\x14\x56\xef\x99\x55\xa5\x3d\xac\x72\xdd\x13\x97\xaa\x02\xa9\x88\xa3\x9b\x71\x52\x0b\x67\x26\x9a\x3f\x20\x08\xf9\xee\x68\x32\x0d\x0d\xa2\x20\x8a\xb8\x59\x13\xbf\x9f\x9f\x9d\x37\x0c\x25\x70\x03\x6a\x87\x12\xc9\x2b\x4d\x0d\x46\x09\xff\xad\x69\xb8\x55\x64\x70\x2f\xa4\x61\x1c\x8b\xb2\x2c\x0b\xec\x79\x47\xce\xb3\x35\x15\x70\x87\x2d\x79\xb5\x47\xf3\x8c\x5c\x8b\x8d\x5a\xb1\xfd\xf3\xb0\x46\xe9\x0f\xb8\x2e\xee\xd9\xe8\x0a\xae\x52\xd3\x36\x8a\xdc\x50\xed\x28\x14\x1d\x05\xd4\x18\xb0\x2a\x00\x0c\x76\x54\xc1\x30\x00\x9b\x5a\xa2\x26\x58\x4d\xc8\x3e\x8a\xf8\xa9\xbb\xac\xad\x09\xce\x8a\x90\x53\x4d\x14\x49\x82\x15\x28\x18\xc7\xf2\xe8\xb2\x00\x10\xdc\x93\xf8\xc4\x9b\x16\xf9\x06\x2a\x2b\x12\xe8\x05\x0c\x1b\x4d\x26\xc0\xdf\xb4\xa1\xef\xa9\x4e\xa0\x80\xae\xa5\x90\x2a\x80\x2c\xcf\xf5\xc2\xf2\xf9\x6b\x65\x61\x5e\x75\x83\x1d\x25\x68\x92\x25\x67\x8f\x1c\x0e\x9f\xc8\x12\xd5\xf7\x58\x93\x9f\x05\xa7\x97\xea\x08\x80\x60\xef\xb0\x93\x77\xa6\xff\x2d\xf9\x64\xf4\xc5\xe3\x6e\x39\x1c\xb6\x56\x9f\x19\x3b\x7f\xb9\x7c\xbc\xa3\x96\x7d\x70\x4f\x59\x59\x3b\x4a\xbd\x8c\xf2\x86\xfa\xd2\xdb\xa2\x7f\xa6\xff\x64\xfb\x8f\x79\xce\x3f\x65\x99\xe7\xcc\x2e\x7f\x29\x86\xa7\x7c\x4c\x07\xcc\x21\xb9\x26\x21\xf4\xa4\x36\xf3\xf3\x19\xcb\x8b\xf2\x35\x00\x00\xff\xff\xcc\x40\x6a\xa3\x8a\x03\x00\x00"),
		},
		"/templates/default_imps_secret.yaml": &vfsgen۰CompressedFileInfo{
			name:             "default_imps_secret.yaml",
			modTime:          time.Time{},
			uncompressedSize: 428,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x8c\x90\x41\x4b\xc4\x30\x10\x85\xef\xfd\x15\x8f\xbd\x27\x20\x78\xea\xd9\xb3\x07\x85\x05\x8f\xb3\xcd\x74\x0d\x4e\xa7\x25\x99\x28\x52\xfb\xdf\x25\x9b\x2e\x0a\x2a\x78\x9c\xe4\xbd\xf7\xbd\x99\x75\x75\x88\x23\xfc\x91\xa4\x70\xf6\x81\x47\x2a\x62\x8f\x3c\x24\x36\xcf\x4a\x27\xe1\x80\x6d\xeb\x9c\x73\x1d\x2d\xf1\xc8\x29\xc7\x59\x7b\xbc\xde\x74\x2f\x51\x43\x8f\x26\xed\x26\x36\x0a\x64\xd4\x77\x80\xd2\xc4\x3d\xd6\x15\x51\x07\x29\x81\x71\x88\x13\x9d\x79\x29\x22\xf9\xa2\x76\xc3\xac\x96\x66\x11\x4e\x7e\x2c\x22\xd5\x70\x80\xc7\xb6\xb9\xbd\xc0\x9e\x92\x17\x1a\x5a\x94\x7f\x60\x61\xca\xec\xef\xaf\xcf\xb5\x15\x20\x74\x62\xc9\x95\x0a\x5c\x76\xf9\x07\xb2\x79\x2a\xf0\x03\x1a\x35\xb0\x1a\x6e\x6b\x9c\xbd\x2f\x3b\xed\xd7\x73\xd4\xef\x2a\xcb\x96\xa2\x9e\xef\xf6\x6d\x1b\xf7\x2d\xda\xf3\x1f\xb6\x2f\x79\xab\x5c\x0d\xb0\xf9\x89\x26\xf9\x59\xe1\x9a\xc7\x1a\xda\xf8\x6d\xf8\x0c\x00\x00\xff\xff\x93\x90\xf6\xc5\xac\x01\x00\x00"),
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
			uncompressedSize: 2296,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xe4\x55\x4d\x8f\x13\x31\x0c\xbd\xcf\xaf\xb0\x7a\x44\x9a\x20\x24\x0e\x68\x6e\xc0\x81\x1b\x87\x22\xad\x84\x10\x07\x4f\xe2\xb6\x61\x33\xf6\x28\x1f\x5d\x69\xcb\xfe\x77\x94\xa4\xdb\x4e\x4b\x81\x65\x5b\xf1\xa1\x3d\xd5\x89\xec\xe7\x17\xbf\xe7\x0e\x8e\xf6\x8a\x7c\xb0\xc2\x1d\xac\x5f\x34\xd7\x96\x4d\x07\x1f\xc8\xaf\xad\xa6\xd7\x5a\x4b\xe2\xd8\x0c\x14\xd1\x60\xc4\xae\x01\x60\x1c\xa8\x83\xcd\x06\x2c\x6b\x97\x0c\xc1\xcc\x0e\xb8\xa4\x31\x39\x17\x48\x7b\x8a\xad\x16\x8e\x5e\x9c\x23\xaf\x16\xc9\xb9\x5c\x30\x03\x05\x77\x77\xdb\xea\x30\xa2\xae\x10\x6a\x4e\x8e\x30\x90\x7a\x7f\x7f\x5d\xb3\x1c\xf6\xe4\x42\xee\x06\xb0\xd9\xb4\x0f\x6a\x55\x6b\x72\xa3\xaf\xc0\x96\x0d\x71\x84\x97\x15\x2e\x43\xdc\xd8\xb8\x02\x75\x85\x2e\x51\x50\xe1\xe0\x79\x0a\x99\x25\x62\xb4\xc2\xa1\x16\x4c\x2e\xf6\x24\xa2\x7c\xc4\xc1\xfd\x08\x9f\xd8\xe4\x43\xd3\xb6\x6d\x33\x1d\xa9\xef\x51\x2b\x4c\x71\x25\xde\xde\x16\x48\x75\xfd\x2a\x28\x2b\xcf\x77\xc3\x7e\xeb\x52\x88\xe4\xe7\xe2\xe8\x82\x93\xbe\xe8\x0c\x7d\x72\x14\xba\xa6\x05\x1c\xed\x3b\x2f\x69\x0c\x1d\x7c\xaa\x48\x41\xf5\xc8\xb7\x68\xb5\x93\x64\x94\x95\xd9\xe7\x06\xc0\x53\x90\xe4\x35\x95\xb4\x67\xe5\x6a\x4d\xbe\x2f\x74\x5a\x58\x52\x2c\xbf\xce\x86\x1a\xdc\x60\xd4\xab\x12\xa5\xd1\x60\xa4\xe3\x46\x47\x98\x25\x53\x0b\x2f\xec\x72\xc0\x31\x94\x63\x7d\x4c\x8d\x69\x4d\x5c\xc2\x87\xf5\xd4\x9e\x72\xcf\x49\xfb\x1c\x1a\x72\xb4\x0d\xc7\x92\xfa\x4b\x4e\x39\x75\xe7\xef\xb0\xbd\x18\xc5\x1c\x30\xd9\x73\x99\xb0\xd9\xf3\x79\x84\x7d\xb2\x6f\xe0\xff\x5e\xd1\xd3\xf6\xd2\x22\xde\x58\x9e\xbe\xfa\x94\x11\x0a\xbd\xdf\x99\x71\x8d\x77\xaa\x1f\xe8\x3e\x51\xfe\xbc\x4d\x7e\x63\xd9\x58\x5e\xfe\xb3\x0b\x2d\x8e\xe6\xb4\xc8\x68\xdf\xff\x09\x9d\xcb\xf4\x5e\xc5\x9f\xcc\xac\x09\xa9\xff\x42\x3a\x16\xd5\x4f\x7e\x73\xfe\x84\x8d\x1f\xb9\x6d\x97\xd7\xf6\x6f\xec\xdc\xb1\x03\x9e\x96\xf4\xcd\xb7\x00\x00\x00\xff\xff\x8f\xb7\xff\xd4\xf8\x08\x00\x00"),
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
			uncompressedSize: 1156,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x6c\x53\xc1\x6e\xdc\x46\x0c\xbd\xcf\x57\x10\xde\xeb\x7a\xb3\xa9\xeb\x22\x98\x9b\x6b\x17\x85\x81\x3a\x58\xc0\x2d\x8a\x22\xc8\x81\x1e\x51\x32\x0b\x6a\xa8\x72\x38\x4a\xb7\x41\xfe\xbd\x18\x69\xb7\x05\x6c\x5f\x24\x0d\xf5\x48\x3e\xbe\x79\xdc\xc0\x1d\xf5\x58\xc5\x61\x46\xa9\x54\xa0\x57\x03\x1e\x71\xa0\xa9\x8a\x14\x4a\x46\x7e\x99\x34\xbb\xa9\x08\x59\xd8\xc0\xaf\xcf\x5c\x80\x0b\x20\xfc\x71\xf3\xf0\xcb\x65\xaf\x36\xa2\x3b\x75\xd0\xb3\xd0\x2e\xb4\x82\x49\xd0\x08\x66\x34\xc6\x27\xa1\x02\xae\xf0\x44\x30\x61\x29\xd4\x01\x67\x57\x38\x6a\x35\x70\x1a\x27\x41\xa7\xb2\x0b\xc1\x68\x12\x4e\x58\x22\xbc\x0f\x81\x8b\xb3\xc6\x00\x60\x34\x73\x61\xcd\x11\x2e\x2e\x42\x98\xb4\xbb\xc9\x59\x1d\x9d\x35\x97\x08\x5f\xbf\x2d\xb1\x47\x4a\xd5\xd8\x8f\xb7\x9a\x9d\xfe\xf6\x25\xaf\xe6\x9b\xf2\x5b\x21\x8b\xf0\xc3\xf5\xf5\xd5\xf7\xe7\xd0\xcf\xa6\x75\x3a\xc7\xca\xeb\x3c\x14\xd1\x2f\x07\xe3\x99\x85\x06\xfa\xa9\x24\x94\xa5\x59\x84\x1e\xa5\x50\x58\x84\x59\x89\x4d\x5a\xd8\xd5\x8e\x11\x86\xe7\x64\x3b\xd6\x77\x4f\x98\xff\x41\x4e\xa2\xb5\x7b\xf7\x42\xc1\x12\x00\x1c\x87\x08\xf3\x7e\x77\xb5\xbb\x0e\x00\xed\xdf\x41\x85\xd3\x31\xc2\x7d\xff\x51\xfd\x60\x54\x28\x7b\x58\x7b\x1c\xaa\xc8\xe3\x9a\x1a\xe1\xd3\xe7\x10\xb2\x76\xf4\x48\x42\xc9\xd5\x96\xc9\xb1\xef\x39\xb3\x1f\x97\x83\xab\x90\x9d\x65\xf9\xf4\x39\x18\x15\xad\x96\xa8\xac\x5c\xff\xaa\x54\x7c\xf9\x06\x18\x69\x5c\x58\x5f\xbc\xdf\xef\x1f\xf8\x62\x89\xa5\xa9\xae\x81\xb1\x9d\x85\x47\x7e\x85\xfe\xee\x25\xfa\x6a\x41\x87\x42\x36\x73\x5a\x34\xf1\xe3\x44\x11\x6e\xa5\x16\x27\xbb\x3f\xb4\x21\xd5\x3c\xc2\x87\xfd\x87\xfd\x7f\xc0\x9b\x94\xb4\xe6\x55\xec\x97\x77\x79\x82\x3c\x68\x6e\xca\x36\x48\x49\x86\x13\xdd\x67\x27\x9b\x51\x22\x5c\x2f\x42\x4a\xb9\xd5\xdc\xf3\xb0\x66\x75\x34\x93\xe8\x34\x52\xf6\x07\xed\x16\x26\x94\x9b\xef\xba\xf3\xad\x35\x93\xdc\x71\xb1\x3a\xb5\x66\x3f\xd6\x6e\x20\x7f\x0b\x26\x3a\xb4\xf0\xd7\x6f\x01\x60\x03\xd2\xea\x46\xe0\xdc\x2b\x6c\x00\x53\xa2\xa9\xb9\x7c\x5d\x93\x08\x13\x66\x4e\x5b\xe8\xd1\x51\xb6\x40\x66\x6a\x5b\xf8\x82\x96\xd7\x27\xe7\x61\xbb\xe4\x6e\xa1\xa3\xa7\x3a\x6c\xc1\x0d\x13\x85\xb0\x01\xec\x3a\x6e\x44\x50\x80\xf2\xcc\xa6\xb9\x71\x7f\xb5\x2e\x9c\xff\xa4\xe4\xe7\x85\xf1\x67\x82\xb6\x86\xc8\x99\x2c\x50\x9e\x4f\x4c\xdb\xfc\xcb\xfe\x9e\x24\x79\x3d\x56\x73\x9e\x0d\xe4\xab\x9d\x3e\xe2\x48\x11\x4e\x39\x97\xa7\xfd\xce\x38\x36\x5c\x7b\x95\x09\x9b\x6d\x56\x11\xfe\x0f\xfc\xce\xfe\x7c\xd0\x6e\x75\x17\x40\x32\xea\x28\x3b\xa3\x9c\xec\x79\x2a\xb8\xf6\x78\x93\x44\x71\xe3\x3c\xdc\xa1\xe3\xa9\xf8\x6a\x97\xb6\xd7\xff\x06\x00\x00\xff\xff\xab\x41\x1c\x85\x84\x04\x00\x00"),
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
		fs["/templates/default_imps_secret.yaml"].(os.FileInfo),
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
