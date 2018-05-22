#!/bin/bash

sudo apt-get install -y libc6-armel-cross libc6-dev-armel-cross
sudo apt-get install -y binutils-arm-linux-gnueabi
sudo apt-get install -y libncurses5-dev
sudo apt-get install -y gcc-arm-linux-gnueabi

rm -fR bin
mkdir bin

echo "Building:"

GOOS=linux GOARCH=arm CGO_ENABLED=1 CC=arm-linux-gnueabi-gcc go build -o bin/linux-arm/bwk-cron cmd/bwk-cron/*.go
GOOS=linux GOARCH=arm CGO_ENABLED=1 CC=arm-linux-gnueabi-gcc go build -o bin/linux-arm/bwk-dash cmd/bwk-dash/*.go
echo "- Linux ARM done!"

GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64/bwk-cron cmd/bwk-cron/*.go
GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64/bwk-dash cmd/bwk-dash/*.go
echo "- Linux AMD64 done!"
