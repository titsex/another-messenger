#!/bin/bash
set -e

echo "Installing protoc-gen-go..."
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

echo "Installing protoc-gen-doc..."
go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest