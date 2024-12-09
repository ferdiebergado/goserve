.PHONY: run build

run:
	go run main.go

build:
	go build -o goserve cmd/goserve/main.go

install:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${GOPATH}/bin/goserve cmd/goserve/main.go