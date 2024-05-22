BUILD_DIR 	:= ../build
PROJECT 	:= example.com/ex00

all: run

run: build
	./$(BUILD_DIR)/$(PROJECT)

.PHONY: build
build: go.mod tidy
	go build -o $(BUILD_DIR)/$(PROJECT) .

go.mod:
	go mod init $(PROJECT)

tidy: go.mod
	go mod tidy

test: go.mod tidy
	go test