.PHONY: build install init

init:
	go mod init motdGO

build:
	export GO11MODULE="auto" \
		go mod download; \
		go mod vendor; \
		CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo -o motd-larry cmd/figlet4go/main.go

install:
	export GO11MODULE="on"; \
		go mod tidy; \
		go mod download

