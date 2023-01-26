#!/usr/bin/env bash

set -euo pipefail

controller_gen_version=0.9.2

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
binpath=${script_dir}/../bin

mkdir -p "$binpath"

function ensure-binary-version() {
  local bin_name=$1
  local bin_version=$2
  local download_location=$3

  local target_name=${bin_name}-proper-${bin_version}
  local link_path=${binpath}/${bin_name}

  if [ ! -L "${link_path}" ]; then
    rm -f "${link_path}"
  fi

  if [ ! -e "${binpath}/${target_name}" ]; then

    BUILD_DIR=$(mktemp -d)
    pushd "${BUILD_DIR}"
    GOBIN=${BUILD_DIR} go install "${download_location}"
    mv "${bin_name}" "${binpath}/${target_name}"
    popd
    rm -rf "${BUILD_DIR}"
  fi

  ln -sf "${target_name}" "${link_path}"
}

ensure-binary-version controller-gen ${controller_gen_version} "sigs.k8s.io/controller-tools/cmd/controller-gen@v${controller_gen_version}"
