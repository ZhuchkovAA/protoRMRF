#!/bin/bash
set -e

PROTO_DIR="proto"
OUT_DIR_GO="gen/go"
OUT_DIR_TS="gen/ts"

PROTO_FILES=$(find "$PROTO_DIR" -name "*.proto")

echo "Generating .proto → Go..."
protoc \
  --proto_path="$PROTO_DIR" \
  --go_out="$OUT_DIR_GO" --go_opt=paths=source_relative \
  --go-grpc_out="$OUT_DIR_GO" --go-grpc_opt=paths=source_relative \
  $PROTO_FILES


echo "Generating .proto → TypeScript..."
mkdir -p "$OUT_DIR_TS"
protoc \
  --plugin=/app/node_modules/.bin/protoc-gen-ts_proto \
  --proto_path="$PROTO_DIR" \
  --ts_proto_out="$OUT_DIR_TS" \
  --ts_proto_opt=outputServices=grpc-js,env=node,esModuleInterop=true,forceLong=long \
  $PROTO_FILES


echo "Done!"
