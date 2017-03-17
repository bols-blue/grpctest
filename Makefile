

export GOPATH=$(PWD)
#PKG=$(shell GOPATH=$(GOPATH) go list ./...)
PKG=client server
all:
	env |grep GO
	go list ./...
	go build $(PKG)
