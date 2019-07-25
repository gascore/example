#!/usr/bin/env bash

find . -name '*_gas.go' -exec rm {} \;
rm -rf dist/
rm log.txt
rm build
