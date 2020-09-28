#!/bin/bash
BINARY=webhash
BIN_PATH=$HOME/bin

UNAME=$(uname)
ARCH=$(uname -m)
if [ "$UNAME" = "Darwin" ]; then
	if [ "$ARCH" = "x86_64" ]; then
		PLATFORM="darwin_amd64"
	fi
elif [ "$UNAME" = "Linux" ]; then
	if [ "$ARCH" = "i686" ]; then
		PLATFORM="linux_386"
	elif [ "$ARCH" = "x86_64" ]; then
		PLATFORM="linux_amd64"
	fi
fi

DEST=$BIN_PATH/$BINARY
# latest
get_latest_release() {
	curl --silent "https://api.github.com/repos/xuender/$BINARY/releases/latest" | # Get latest release from GitHub api
		grep '"tag_name":' |                                                          # Get tag line
		sed -E 's/.*"([^"]+)".*/\1/'                                                  # Pluck JSON value
}
# 安装
install() {
	echo "Installing $BINARY..."
	if [ "$UNAME" != "Linux" ] && [ "$UNAME" != "Darwin" ] && [ "$ARCH" != "x86_64" ] && [ "$ARCH" != "i686" ]; then
		echo "Sorry, OS/Architecture not supported: ${UNAME}/${ARCH}. Download binary from https://github.com/xuender/${BINARY}/releases"
		exit 1
	fi
	mkdir -p $BIN_PATH
	LATEST=$(get_latest_release)
	echo "$LATEST"
	URL="https://github.com/xuender/webhash/releases/download/$LATEST/${BINARY}_${PLATFORM}"
	echo "download $URL"
	rm -rf $DEST
	curl -L $URL -o $DEST
	chmod +x $DEST
}

run() {
	echo "Running webhash..."
	webhash
}

if [ -f $DEST ]; then
	echo "Webhash is already installed."
else
	echo "Webhash is not installed."
fi

install
run
