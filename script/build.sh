#!/bin/bash

echo "Building:"

GOOS=linux GOARCH=arm go build -o bin/bwk-cron cmd/bwk-cron/*.go
GOOS=linux GOARCH=arm go build -o bin/bwk-dash cmd/bwk-dash/*.go
echo "- Linux ARM done!"

GOOS=linux GOARCH=amd64 go build -o bin/bwk-cron cmd/bwk-cron/*.go
GOOS=linux GOARCH=amd64 go build -o bin/bwk-dash cmd/bwk-dash/*.go
echo "- Linux AMD64 done!"
