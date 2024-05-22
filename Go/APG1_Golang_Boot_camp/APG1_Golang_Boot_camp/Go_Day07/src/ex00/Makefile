BUILD_DIR := ../build
PROJECT   := ex01

all: run

run: build
	$(BUILD_DIR)/$(PROJECT)

build: go.mod
	go mod tidy
#	swag init
	go build -o $(BUILD_DIR)/$(PROJECT) .

go.mod:
	go mod init $(PROJECT)

test:
	go test .

dev: go.mod
	go get github.com/githubnemo/CompileDaemon
	go install github.com/githubnemo/CompileDaemon
#	export PATH=$PATH:$(go env GOPATH)/bin
	CompileDaemon -directory=. \
		-build="make build" \
		-command="make run" \
		-polling \
		-exclude-dir="docs" \
		-include="*.html" \
		-graceful-kill \
		-color
		