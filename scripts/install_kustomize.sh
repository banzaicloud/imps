#!/usr/bin/env bash

# Copyright (c) 2019 Banzai Cloud Zrt. All Rights Reserved.

set -euo pipefail

version=3.6.1
opsys=$(uname -s | awk '{print tolower($0)}')
versioned_path="bin/kustomize-${version}"

if [ ! -x $versioned_path ]; then
    # download the release
    curl -O -L "https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize/v${version}/kustomize_v${version}_${opsys}_amd64.tar.gz"

    tar xzf "./kustomize_v${version}_${opsys}_amd64.tar.gz"
    rm "./kustomize_v${version}_${opsys}_amd64.tar.gz"

    # move to bin
    mkdir -p bin
    mv kustomize bin/kustomize-${version}
    chmod u+x bin/kustomize-${version}
fi

ln -sf kustomize-${version} bin/kustomize
