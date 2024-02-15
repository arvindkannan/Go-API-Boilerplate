# Set strict mode
$ErrorActionPreference = "Stop"
$ProgressPreference = "SilentlyContinue"

# Get the version
# $VERSION = if ($env:VERSION) { $env:VERSION } else { git describe --tags --first-parent --abbrev=7 --long --dirty --always | ForEach-Object { $_ -replace '^v' } }
# if env version is not present, then use git describe to get the version
if (-not $env:VERSION) {
    $VERSION = git describe --tags --first-parent --abbrev=7 --long --dirty --always | ForEach-Object { $_ -replace '^v' }
} else {
    $VERSION = $env:VERSION
}

# Set Go flags
$GOFLAGS = "-ldflags=-w -s -X=go-api-boilerplate/version.Version=$($VERSION) -X=go-api-boilerplate/server.mode=release"

# Define build architectures
$BUILD_ARCH = if ($env:BUILD_ARCH) { $env:BUILD_ARCH } else { "amd64", "arm64" }

# Set AMDGPU_TARGETS
$AMDGPU_TARGETS = if ($env:AMDGPU_TARGETS) { $env:AMDGPU_TARGETS } else { "" }

# Create the dist directory
New-Item -ItemType Directory -Force -Path dist | Out-Null

# Build for each target architecture
foreach ($TARGETARCH in $BUILD_ARCH) {
    # set GOOS to windows and GOARCH to the current target architecture
    $env:GOOS = "windows"
    $env:GOARCH = $TARGETARCH
    & go build -o "dist\$VERSION\windows\go-api-boilerplate-$TARGETARCH.exe" $GOFLAGS
}

Write-Host "Windows Build completed."
