#!/bin/bash

set -eu

# export VERSION=${VERSION:-$(git describe --tags --first-parent --abbrev=7 --long --dirty --always | sed -e "s/^v//g")}
# if env version is not set, get version from git tag
if [ -z "$VERSION" ]; then
    VERSION=$(git describe --tags --first-parent --abbrev=7 --long --dirty --always | sed -e "s/^v//g")
fi
export GOFLAGS='-ldflags=-w -s -X=go-api-boilerplate/version.Version='"$VERSION"' -X=go-api-boilerplate/server.mode=release'

BUILD_ARCH=${BUILD_ARCH:-"amd64 arm64"}
export AMDGPU_TARGETS=${AMDGPU_TARGETS:=""}
mkdir -p dist

for TARGETARCH in ${BUILD_ARCH}; do
    GOARCH=$TARGETARCH go build -o dist/$VERSION/linux/go-api-boilerplate-$TARGETARCH "$GOFLAGS"
done

echo "Build completed."
