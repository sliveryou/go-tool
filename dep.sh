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

# 需要 go1.21 以上版本构建安装
${download} github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56.1
${download} golang.org/x/tools/cmd/goimports@v0.18.0
${download} github.com/incu6us/goimports-reviser/v3@v3.6.4
${download} mvdan.cc/gofumpt@v0.6.0
${download} mvdan.cc/sh/v3/cmd/shfmt@v3.8.0
${download} mvdan.cc/sh/v3/cmd/gosh@v3.8.0

echo "done"
