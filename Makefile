PREFIX ?= /usr/local

default: build

build:
	go build -o pkg/path_helper path_helper.go

install:
	mkdir -p $(PREFIX)/bin
	cp pkg/path_helper $(PREFIX)/bin
