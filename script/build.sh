#!/bin/bash

echo "Configuring system..."
sudo apt-get update -y
sudo apt-get install -y gcc vim git
sudo apt-get install -y libc6-armel-cross libc6-dev-armel-cross
sudo apt-get install -y binutils-arm-linux-gnueabi
sudo apt-get install -y libncurses5-dev
sudo apt-get install -y gcc-arm-linux-gnueabi

echo "Downloading golang..."
sudo wget https://dl.google.com/go/go1.10.2.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.10.2.linux-amd64.tar.gz
sudo rm go1.10.2.linux-amd64.tar.gz

echo "Configuring golang..."
mkdir -p $HOME/go/bin
sudo echo "GOPATH=$HOME/go" >> /etc/profile
sudo echo "PATH=/usr/local/go/bin:$GOPATH/bin:$PATH" >> /etc/profile
source /etc/profile

echo "Getting source code..."
go get github.com/dustinengle/bwk-dash
cd $GOPATH/src/github.com/dustinengle/bwk-dash

echo "Removing old builds..."
rm -fR bin
mkdir bin

echo "Building binaries..."
GOOS=linux GOARCH=arm CGO_ENABLED=1 CC=arm-linux-gnueabi-gcc go build -o bin/linux-arm/bwk-cron cmd/bwk-cron/*.go
GOOS=linux GOARCH=arm CGO_ENABLED=1 CC=arm-linux-gnueabi-gcc go build -o bin/linux-arm/bwk-dash cmd/bwk-dash/*.go
echo "- Linux ARM done!"
GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64/bwk-cron cmd/bwk-cron/*.go
GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64/bwk-dash cmd/bwk-dash/*.go
echo "- Linux AMD64 done!"
