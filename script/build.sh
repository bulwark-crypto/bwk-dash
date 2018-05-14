#!/bin/bash

echo "Buidling BWK-Dash..."
go build -o $GOPATH/bin/bwk-dash $GOPATH/src/github.com/dustinengle/bwk-dash/cmd/bwk-dash/*.go

echo "Buidling BWK-Clean..."
go build -o $GOPATH/bin/bwk-clean $GOPATH/src/github.com/dustinengle/bwk-dash/cmd/bwk-clean/*.go

echo "Buidling BWK-Cron..."
go build -o $GOPATH/bin/bwk-cron $GOPATH/src/github.com/dustinengle/bwk-dash/cmd/bwk-cron/*.go

echo "Finished!"
