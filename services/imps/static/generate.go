// Copyright (c) 2019 Banzai Cloud Zrt. All Rights Reserved.

package main

import (
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/shurcooL/vfsgen"
)

// ZeroModTimeFileSystem is an http.FileSystem wrapper.
// It exposes a filesystem exactly like Source, except
// all file modification times are changed to zero.
type ZeroModTimeFileSystem struct {
	Source http.FileSystem
}

func (fs ZeroModTimeFileSystem) Open(name string) (http.File, error) {
	f, err := fs.Source.Open(name)

	return file{f}, err
}

type file struct {
	http.File
}

func (f file) Stat() (os.FileInfo, error) {
	fi, err := f.File.Stat()

	return fileInfo{fi}, err
}

type fileInfo struct {
	os.FileInfo
}

func (fi fileInfo) ModTime() time.Time { return time.Time{} }

func getRepoRoot() string {
	//nolint
	_, filename, _, _ := runtime.Caller(0)

	dir := filepath.Dir(filename)

	return filepath.Dir(path.Join(dir, "."))
}

func main() {
	var err error
	err = vfsgen.Generate(
		ZeroModTimeFileSystem{
			http.Dir(path.Join(getRepoRoot(), "deploy/charts/imagepullsecrets-controller"))},
		vfsgen.Options{
			Filename:     "static/charts/imagepullsecrets_controller/chart.gogen.go",
			PackageName:  "imagepullsecrets_controller",
			VariableName: "Chart",
		})
	if err != nil {
		panic(err)
	}
}
