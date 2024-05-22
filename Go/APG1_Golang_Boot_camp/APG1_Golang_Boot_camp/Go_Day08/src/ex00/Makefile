BUILD_DIR 	:= ../build
PROJECT 	:= example.com/ex00

all: test

run: build
	./$(BUILD_DIR)/$(PROJECT)

.PHONY: build
build: go.mod
	go mod tidy
	go build -o $(BUILD_DIR)/$(PROJECT) .

go.mod:
	go mod init $(PROJECT)

test: go.mod
	go test