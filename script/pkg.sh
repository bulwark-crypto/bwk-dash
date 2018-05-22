#!/bin/bash

rm -fR build
mkdir build 

cd bin/linux-arm
tar -zcf ../../build/bwk-dash-1.0.0-linux-arm.tar.gz bwk-cron bwk-dash

cd ../linux-amd64
tar -zcf ../../build/bwk-dash-1.0.0-linux-amd64.tar.gz bwk-cron bwk-dash