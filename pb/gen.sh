#!/usr/bin/env bash

#proto3语法规范参考：
#https://developers.google.com/protocol-buffers/docs/reference/proto3-spec

rm -rf ./_gen
mkdir -p ./_gen/go ./_gen/py ./_gen/java ./_gen/php ./_gen/cs

protoc -I ./ \
    --go_out=plugins=grpc:./_gen/go \
    --csharp_out=./_gen/cs/ \
    --java_out=./_gen/java/ \
    --php_out=./_gen/php/ \
    --python_out=./_gen/py/ \
    ./*.proto
