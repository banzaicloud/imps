#!/usr/bin/env bash

# Copyright (c) 2021 Banzai Cloud Zrt.

set -euo pipefail

ROOT_DIR=${1:-}

if [ -z "$ROOT_DIR" ]; then
    echo "Usage $0 <service-root-dir>"
fi

readonly REPO_ROOT=$(git rev-parse --show-toplevel)
readonly ROOT_DIR="$(cd $ROOT_DIR; pwd)"
mkdir -p ${ROOT_DIR}/static
readonly RELATIVE_ROOT_DIR=$(echo $ROOT_DIR | sed "s?$REPO_ROOT??")

if [ -z "$RELATIVE_ROOT_DIR" ]; then
        GO_MODULE_NAME="github.com/banzaicloud/imps"
else
        GO_MODULE_NAME="github.com/banzaicloud/imps$RELATIVE_ROOT_DIR"
fi

if [ ! -d $ROOT_DIR/deploy/charts ]; then
        echo "No charts found, skipping generator"
        exit 0
fi

function generateStaticEntry() {
        DIRECTORY="$1"
        VARIABLE_NAME="$2"
        PATH_PREFIX="$3"

        VARIABLE_NAME_LOWER=$(echo $2 | tr '[A-Z]' '[a-z]')

        PACKAGE_NAME=$(basename $DIRECTORY | tr '-' '_')
        mkdir -p ${ROOT_DIR}/static/${PATH_PREFIX}/${PACKAGE_NAME}
        RELATIVE_DENTRY_PATH=$(echo $DIRECTORY | sed "s?$ROOT_DIR??;s?^/??")

        cat <<EOF
        err = vfsgen.Generate(
           ZeroModTimeFileSystem{
               http.Dir(path.Join(getRepoRoot(), "${RELATIVE_DENTRY_PATH}"))},
            vfsgen.Options{
                Filename:     "static/${PATH_PREFIX}/${PACKAGE_NAME}/${VARIABLE_NAME_LOWER}.gogen.go",
                PackageName:  "${PACKAGE_NAME}",
                VariableName: "${VARIABLE_NAME}",
        })
        if err != nil {
                panic(err)
        }
EOF
}

(
cat <<EOF

// Copyright © 2021 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
        var err error;
EOF

if [ -d $ROOT_DIR/deploy/charts ]; then
        for dentry in  $ROOT_DIR/deploy/charts/*; do
                if [ -d $dentry ]; then
                        generateStaticEntry $dentry "Chart" "charts"
                fi
        done
fi


if [ -d $ROOT_DIR/assets/charts ]; then
        for dentry in  $ROOT_DIR/assets/charts/*; do
                if [ -d $dentry ]; then
                        generateStaticEntry $dentry "Chart" "charts"
                fi
        done
fi

if [ -d $ROOT_DIR/assets ]; then
        for dentry in  $ROOT_DIR/assets/*; do
                if [ -d $dentry ] && [ $(basename $dentry) != "charts" ]; then
                        generateStaticEntry $dentry "Assets" "assets"
                fi
        done
fi

cat <<EOF
}
EOF
) > $ROOT_DIR/static/generate.go

go fmt $ROOT_DIR/static/generate.go > /dev/null


(

echo "module github.com/banzaicloud/imps$RELATIVE_ROOT_DIR/static

go 1.16
"

) > $ROOT_DIR/static/go.mod

(cd $ROOT_DIR;
go mod edit -replace=github.com/banzaicloud/imps$RELATIVE_ROOT_DIR/static=./static
)
