#!/bin/bash

# format code
go fmt

# build the go files in current directory
go build -o ../bin/main

# runs main
go run -exec ../bin/mian
