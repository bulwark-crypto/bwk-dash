#!/bin/bash

rm -fR build
mkdir build 

tar -zcf build/bwk-dash-1.0.0-linux-arm.tar.gz bin/linux-arm/*
tar -zcf build/bwk-dash-1.0.0-linux-amd64.tar.gz bin/linux-amd64/*