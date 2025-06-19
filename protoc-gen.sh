#!/bin/bash
set -e

echo "Generating .proto → Go..."

protoc \
  --proto_path=proto \
  --go_out=gen/go \
  --go-grpc_out=gen/go \
  $(find proto -name "*.proto")

echo "Done!"
