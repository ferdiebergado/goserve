.PHONY: run build

MAIN:=cmd/goserve/main.go

run:
	go run ${MAIN}

build:
	go build -o bin/goserve ${MAIN}

install:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${GOPATH}/bin/goserve ${MAIN}