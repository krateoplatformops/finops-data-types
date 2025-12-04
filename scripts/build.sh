#!/bin/bash

LOCALBIN=$(pwd)/bin
mkdir -p ${LOCALBIN}

CONTROLLER_TOOLS_VERSION=v0.19.0
CONTROLLER_GEN=${LOCALBIN}/controller-gen-${CONTROLLER_TOOLS_VERSION}

go-install-tool() {
    [ -f $1 ] || {
    set -e
    package=$2@$3
    echo "Downloading ${package}"
    GOBIN=$1 go install ${package}
    mv $(echo $4 | sed "s/-$3//") $4
    }
}

go-install-tool "${LOCALBIN}" "sigs.k8s.io/controller-tools/cmd/controller-gen" "${CONTROLLER_TOOLS_VERSION}" "${CONTROLLER_GEN}"

$(echo $CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."