# install.ps1 — OperaTree installer for Windows
# Usage: irm https://raw.githubusercontent.com/hanymamdouh82/operatree/main/install.ps1 | iex

$ErrorActionPreference = "Stop"

$repo    = "hanymamdouh82/operatree"
$binary  = "operatree.exe"
$installDir = "$env:USERPROFILE\bin"

# Create install dir if it doesn't exist
if (-not (Test-Path $installDir)) {
    New-Item -ItemType Directory -Path $installDir | Out-Null
}

# Fetch latest release tag
Write-Host "Fetching latest release..."
$release = Invoke-RestMethod -Uri "https://api.github.com/repos/$repo/releases/latest" `
    -Headers @{ "User-Agent" = "operatree-installer" }
$version = $release.tag_name

# Download the Windows binary
$url = "https://github.com/$repo/releases/download/$version/operatree-windows-amd64.exe"
Write-Host "Installing operatree $version (windows/amd64)..."
Invoke-WebRequest -Uri $url -OutFile "$installDir\$binary"

# Add to user PATH if not already there
$userPath = [Environment]::GetEnvironmentVariable("PATH", "User")
if ($userPath -notlike "*$installDir*") {
    [Environment]::SetEnvironmentVariable("PATH", "$userPath;$installDir", "User")
    Write-Host "Added $installDir to PATH — restart your terminal for it to take effect."
}

Write-Host ""
Write-Host "Done. Run: operatree version"
