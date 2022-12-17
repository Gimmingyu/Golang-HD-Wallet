BUILD_DIR = build

BUILD_TARGET = ./cmd

BUILD_OUTPUT = outfile

all:
	@make create
	@make build

create:
	protoc --proto_path=proto proto/*.proto --go_out=gen
	protoc --proto_path=proto proto/*.proto --go-grpc_out=gen

build:
	@echo "\033[92mBuilding HD Wallet Server...\033[0m"
	@go build -o $(BUILD_DIR)/$(BUILD_OUTPUT) $(BUILD_TARGET)


clean:
	rm gen/*/*.go
	rm $(BUILD_DIR)/$(BUILD_OUTPUT)

fclean:
	rm -rf gen/*
	rm -rf build

.PHONY: hdwallet clean all test build