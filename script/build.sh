#!/bin/bash

echo "Moving to source folder..."
cd $GOPATH/src/github.com/dustinengle/bwk-dash

echo "Buidling BWK-Dash..."
go build -o $GOPATH/bin/bwk-dash cmd/bwk-dash/*.go

echo "Buidling BWK-Clean..."
go build -o $GOPATH/bin/bwk-clean cmd/bwk-clean/*.go

echo "Buidling BWK-Cron..."
go build -o $GOPATH/bin/bwk-cron cmd/bwk-cron/*.go

echo "Finished!"
