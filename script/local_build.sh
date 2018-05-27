#!/bin/bash

echo "Building website..."
cd client
yarn run build
cd ..

echo "Removing old builds..."
rm -fR bin
mkdir bin
rm -fR build
mkdir build

echo "Building binaries..."
GOOS=linux GOARCH=arm CGO_ENABLED=1 CC=arm-linux-gnueabi-gcc go build -o bin/linux-arm/bwk-cron cmd/bwk-cron/*.go
GOOS=linux GOARCH=arm CGO_ENABLED=1 CC=arm-linux-gnueabi-gcc go build -o bin/linux-arm/bwk-dash cmd/bwk-dash/*.go
echo "- Linux ARM done!"
GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64/bwk-cron cmd/bwk-cron/*.go
GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64/bwk-dash cmd/bwk-dash/*.go
echo "- Linux AMD64 done!"
