# Set strict mode
$ErrorActionPreference = "Stop"
$ProgressPreference = "SilentlyContinue"

# $env:VERSION = if (-not $env:VERSION) { & git describe --tags --first-parent --abbrev=7 --long --dirty --always | ForEach-Object { $_ -replace '^v', '' } }
if (-not $env:VERSION) {
    $VERSION = git describe --tags --first-parent --abbrev=7 --long --dirty --always | ForEach-Object { $_ -replace '^v' }
} else {
    $VERSION = $env:VERSION
}

# $env:GOFLAGS = '-ldflags=-w -s -X=go-api-boilerplate/version.Version=' + $env:VERSION + ' -X=go-api-boilerplate/server.mode=release'
$GOFLAGS = "-ldflags=-w -s -X=go-api-boilerplate/version.Version=$($VERSION) -X=go-api-boilerplate/server.mode=release"


$BUILD_ARCH = if (-not $env:BUILD_ARCH) { "amd64", "arm64" } else { $env:BUILD_ARCH -split ' ' }

$env:AMDGPU_TARGETS = if (-not $env:AMDGPU_TARGETS) { "" } else { $env:AMDGPU_TARGETS }

New-Item -ItemType Directory -Force -Path dist | Out-Null

foreach ($TARGETARCH in $BUILD_ARCH) {
    $env:GOARCH = $TARGETARCH
    & go build -o "dist\$VERSION\linux\go-api-boilerplate-$TARGETARCH" $GOFLAGS
}

Write-Output "Linux Build completed."
