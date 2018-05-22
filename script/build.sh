#!/bin/bash

rm -fR bin
mkdir bin

echo "Building:"

GOOS=linux GOARCH=arm go build -o bin/linux-arm/bwk-cron cmd/bwk-cron/*.go
GOOS=linux GOARCH=arm go build -o bin/linux-arm/bwk-dash cmd/bwk-dash/*.go
echo "- Linux ARM done!"

GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64/bwk-cron cmd/bwk-cron/*.go
GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64/bwk-dash cmd/bwk-dash/*.go
echo "- Linux AMD64 done!"
