#!/bin/bash
export GO111MODULE=on

download=""
go_version=$(go version | awk '{ print $3 }')
# go 1.17 之后下载编译成可执行文件要使用 go install
if ! (printf '%s\n%s\n' "go1.17" "${go_version}" | sort -V -C); then
  echo "go get dependencies..."
  download="go get"
else
  echo "go install dependencies..."
  download="go install"
fi

${download} github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2
${download} golang.org/x/tools/cmd/goimports@v0.16.1
${download} github.com/incu6us/goimports-reviser/v3@v3.6.2
${download} mvdan.cc/gofumpt@v0.5.0
${download} mvdan.cc/sh/v3/cmd/shfmt@v3.7.0
${download} mvdan.cc/sh/v3/cmd/gosh@v3.7.0

echo "done"
