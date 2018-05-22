#!/bin/bash

rm -fR build
mkdir build 

tar -zcf build/bwk-dash-1.0.0-linux-arm.tar.gz bin/linux-arm/bwk-cron bin/linux-arm/bwk-dash
tar -zcf build/bwk-dash-1.0.0-linux-amd64.tar.gz bin/linux-amd64/bwk-cron bin/linux-amd64/bwk-dash