param (
    [string]$VERSION
)

function Usage {
    Write-Host "usage: $(Get-Command $MyInvocation.MyCommand.Definition -CommandType Cmdlet).Name VERSION"
    exit 1
}

if (-not $VERSION) {
    Usage
}

$env:VERSION = $VERSION

# build universal Windows binary
& "$(Split-Path $MyInvocation.MyCommand.Path)\build_windows.ps1"

# build arm64 and amd64 Linux binaries
& "$(Split-Path $MyInvocation.MyCommand.Path)\build_linux.ps1"
