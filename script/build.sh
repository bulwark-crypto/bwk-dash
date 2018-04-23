#!/bin/bash
echo "Buidling BWK-Dash..."
go build -o $GOPATH/bin/bwk-dash $GOPATH/src/github.com/dustinengle/bwk-dash/cmd/bwk-dash/*.go
echo "Finished!"
