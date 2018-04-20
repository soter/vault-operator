#!/bin/bash

set -x
set -eou pipefail

GOPATH=$(go env GOPATH)
REPO_ROOT="$GOPATH/src/github.com/soter/vault-operator"

pushd $REPO_ROOT

export APPSCODE_ENV=prod
rm -rf dist
./hack/docker/setup.sh
./hack/docker/setup.sh release
rm -rf dist/.tag

popd