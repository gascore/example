# !/bin/bash

cd ..
sh std.sh
cd example
go run main.go -w -p=gopherjs
# ./builder -w -p=gopherjs
