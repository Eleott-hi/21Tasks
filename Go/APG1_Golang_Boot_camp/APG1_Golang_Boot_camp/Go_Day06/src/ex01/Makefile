BUILD_DIR := ../build
PROJECT   := main

all: run

run: build
	$(BUILD_DIR)/$(PROJECT)

build: go.mod
	go mod tidy
	swag init
	go build -o $(BUILD_DIR)/$(PROJECT) .

go.mod:
	go mod init $(PROJECT)
