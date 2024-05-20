BUILD_DIR := ../build
PROJECT   := ex01

all: dev

run: build
	$(BUILD_DIR)/$(PROJECT)

build: go.mod
	go mod tidy
	swag init
	go build -o $(BUILD_DIR)/$(PROJECT) .

go.mod:
	go mod init $(PROJECT)

dev: go.mod
	go get github.com/githubnemo/CompileDaemon
	go install github.com/githubnemo/CompileDaemon
	CompileDaemon -directory=. \
		-build="make build" \
		-command="make run" \
		-polling \
		-exclude-dir="docs" \
		-graceful-kill \
		-color