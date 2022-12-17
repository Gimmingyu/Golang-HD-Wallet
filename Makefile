BUILD_DIR = build

BUILD_TARGET_RPC = ./cmd/rpc
BUILD_TARGET_HTTP = ./cmd/http

BUILD_OUTPUT_RPC = rpcd
BUILD_OUTPUT_HTTP = httpd


all: build


create:
	protoc --proto_path=proto proto/*.proto --go_out=gen
	protoc --proto_path=proto proto/*.proto --go-grpc_out=gen

build:
	@echo "\033[92mBuilding gRPC Server...\033[0m"
	@go build -o $(BUILD_DIR)/$(BUILD_OUTPUT_RPC) $(BUILD_TARGET_RPC)
	@echo "\033[92mBuilding HTTP Server...\033[0m"
	@go build -o $(BUILD_DIR)/$(BUILD_OUTPUT_HTTP) $(BUILD_TARGET_HTTP)


clean:
	rm gen/*/*.go
	rm $(BUILD_DIR)/$(BUILD_OUTPUT_RPC)
	rm $(BUILD_DIR)/$(BUILD_OUTPUT_HTTP)

fclean:
	rm -rf gen/*
	rm -rf build

.PHONY: hdwallet clean all test build