

export GOPATH=$(PWD)
#PKG=$(shell GOPATH=$(GOPATH) go list ./...)
PKG=client server
all:$(PKG)
	env |grep GO
	go list ./...
	go build $(PKG)

client:
	go build -o bin/$@ $@

server:
	go build -o bin/$@ $@
