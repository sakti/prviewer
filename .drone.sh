#!/bin/sh
set -e
set -x

# disable CGO for cross-compiling
export CGO_ENABLED=0

# compile for all architectures
GOOS=linux   GOARCH=amd64 go build -ldflags "-X main.Version=${DRONE_TAG##v}" -o release/linux/amd64/prviewer   github.com/sakti/prviewer
GOOS=linux   GOARCH=arm64 go build -ldflags "-X main.Version=${DRONE_TAG##v}" -o release/linux/arm64/prviewer   github.com/sakti/prviewer
GOOS=linux   GOARCH=arm   go build -ldflags "-X main.Version=${DRONE_TAG##v}" -o release/linux/arm/prviewer     github.com/sakti/prviewer
GOOS=windows GOARCH=amd64 go build -ldflags "-X main.Version=${DRONE_TAG##v}" -o release/windows/amd64/prviewer.exe github.com/sakti/prviewer
GOOS=darwin  GOARCH=amd64 go build -ldflags "-X main.Version=${DRONE_TAG##v}" -o release/darwin/amd64/prviewer  github.com/sakti/prviewer

# tar binary files prior to upload
tar -cvzf release/prviewer_linux_amd64.tar.gz   -C release/linux/amd64   prviewer
tar -cvzf release/prviewer_linux_arm64.tar.gz   -C release/linux/arm64   prviewer
tar -cvzf release/prviewer_linux_arm.tar.gz     -C release/linux/arm     prviewer
tar -cvzf release/prviewer_windows_amd64.tar.gz -C release/windows/amd64 prviewer.exe
tar -cvzf release/prviewer_darwin_amd64.tar.gz  -C release/darwin/amd64  prviewer

# generate shas for tar files
sha256sum release/*.tar.gz > release/drone_checksums.txt