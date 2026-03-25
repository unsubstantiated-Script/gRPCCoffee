PROTO_FILE := coffee_shop.proto
PROTO_OUT_DIR := coffeeshop_proto
APP_NAME := grpc-coffee
MAIN_FILE := main.go

.PHONY: help check-tools proto build_proto build run deps clean

help:
	@echo "Available targets:"
	@echo "  make proto       Generate Go files from $(PROTO_FILE)"
	@echo "  make build_proto Alias for proto"
	@echo "  make build       Build binary $(APP_NAME)"
	@echo "  make run         Run the application"
	@echo "  make deps        Run go mod tidy"
	@echo "  make clean       Remove build and generated artifacts"

check-tools:
	@command -v protoc >/dev/null 2>&1 || { echo "Error: protoc is not installed"; exit 1; }
	@command -v protoc-gen-go >/dev/null 2>&1 || { echo "Error: protoc-gen-go is not installed"; exit 1; }
	@command -v protoc-gen-go-grpc >/dev/null 2>&1 || { echo "Error: protoc-gen-go-grpc is not installed"; exit 1; }

proto: check-tools
	@mkdir -p $(PROTO_OUT_DIR)
	protoc --go_out=$(PROTO_OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_OUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_FILE)

build_proto: proto

build:
	go build -o $(APP_NAME) $(MAIN_FILE)

run:
	go run $(MAIN_FILE)

deps:
	go mod tidy

clean:
	rm -f $(APP_NAME)
	rm -f $(PROTO_OUT_DIR)/*.pb.go
