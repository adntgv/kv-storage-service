SRC_DIR=proto
DST_DIR=gen

all: generate build

generate:
	protoc -I=${SRC_DIR} --go_out=${DST_DIR} --go_opt=paths=source_relative \
    --go-grpc_out=${DST_DIR} --go-grpc_opt=paths=source_relative \
    ${SRC_DIR}/service.proto

build:
	go build ./cmd/server/ 