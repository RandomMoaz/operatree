#!/bin/sh
# install.sh — OperaTree installer
# Usage: curl -fsSL https://raw.githubusercontent.com/hanymamdouh82/operatree/main/install.sh | sh

set -e

REPO="hanymamdouh82/operatree"
BINARY="operatree"
INSTALL_DIR="/usr/local/bin"

# Detect OS
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
case "$OS" in
  linux)  OS="linux" ;;
  darwin) OS="darwin" ;;
  *)
    echo "Unsupported OS: $OS"
    echo "Please download manually from https://github.com/$REPO/releases"
    exit 1
    ;;
esac

# Detect architecture
ARCH=$(uname -m)
case "$ARCH" in
  x86_64)  ARCH="amd64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  *)
    echo "Unsupported architecture: $ARCH"
    echo "Please download manually from https://github.com/$REPO/releases"
    exit 1
    ;;
esac

# Fetch latest release tag from GitHub API
echo "Fetching latest release..."
VERSION=$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" \
  | grep '"tag_name"' \
  | sed 's/.*"tag_name": *"\(.*\)".*/\1/')

if [ -z "$VERSION" ]; then
  echo "Could not determine latest version. Check your internet connection."
  exit 1
fi

ASSET="${BINARY}-${OS}-${ARCH}"
URL="https://github.com/$REPO/releases/download/$VERSION/$ASSET"

echo "Installing $BINARY $VERSION ($OS/$ARCH)..."
curl -fsSL "$URL" -o "/tmp/$BINARY"
chmod +x "/tmp/$BINARY"

# Install — try without sudo first, fall back to sudo
if mv "/tmp/$BINARY" "$INSTALL_DIR/$BINARY" 2>/dev/null; then
  echo "Installed to $INSTALL_DIR/$BINARY"
else
  echo "Needs elevated permissions to install to $INSTALL_DIR..."
  sudo mv "/tmp/$BINARY" "$INSTALL_DIR/$BINARY"
  echo "Installed to $INSTALL_DIR/$BINARY"
fi

echo ""
echo "Done. Run: operatree version"
