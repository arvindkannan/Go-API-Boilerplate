#!/bin/bash

# Set strict mode
set -eu

# Get the version
# VERSION=${VERSION:-$(git describe --tags --first-parent --abbrev=7 --long --dirty --always | sed -e "s/^v//g")}
# get version from git tag if environment variable is not set
if [ -z "$VERSION" ]; then
    VERSION=$(git describe --tags --first-parent --abbrev=7 --long --dirty --always | sed -e "s/^v//g")
fi

# Set Go flags
export GOFLAGS='-ldflags=-w -s -X=go-api-boilerplate/version.Version='"$VERSION"' -X=go-api-boilerplate/server.mode=release'

# Define build architectures
BUILD_ARCH=${BUILD_ARCH:-"amd64 arm64"}

# Set AMDGPU_TARGETS
AMDGPU_TARGETS=${AMDGPU_TARGETS:=""}

# Create the dist directory
mkdir -p dist

# Build for each target architecture
for TARGETARCH in ${BUILD_ARCH}; do
    GOARCH=$TARGETARCH go build -o dist/$VERSION/windows/go-api-boilerplate-$TARGETARCH.exe "$GOFLAGS"
done

echo "Build completed."
