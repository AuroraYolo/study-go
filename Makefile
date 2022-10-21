BINARY="demo"
VERSION=0.0.1
BUILD=`date +%F`
SHELL := /bin/bash

versionDir="github.com/iswbm/demo/utils"
gitTag=$(shell git log --pretty=format:'%h' -n 1)
gitBranch=$(shell git rev-parse --abbrev-ref HEAD)
buildDate=$(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit=$(shell git rev-parse --short HEAD)

ldflags="-s -w -X ${versionDir}.version=${VERSION} -X ${versionDir}.gitBranch=${gitBranch} -X '${versionDir}.gitTag=${gitTag}' -X '${versionDir}.gitCommit=${gitCommit}' -X '${versionDir}.buildDate=${buildDate}'"

default:
    @echo "build the ${BINARY}"
    @GOOS=linux GOARCH=amd64 go build -ldflags ${ldflags} -o  build/${BINARY}.linux  -tags=jsoniter
    @go build -ldflags ${ldflags} -o  build/${BINARY}.mac  -tags=jsoniter
    @echo "build done."