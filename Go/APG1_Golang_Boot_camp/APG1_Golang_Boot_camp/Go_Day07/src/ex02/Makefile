BUILD_DIR := ../build
PROJECT   := example.com/ex02

all: doc

go.mod:
	go mod init $(PROJECT)

test: go.mod
	go test .

doc: test
	godoc -http=:6060 &
	sleep 5
	-wget -r -np -k http://localhost:6060/pkg/ -P ./docs
	zip -r docs.zip ./docs
	killall godoc
	rm -rf ./docs