SHELL := /bin/bash

export GOPROXY=https://goproxy.cn,direct
export GO111MODULE=on
export GOSUMDB=off
export BIN=$(shell pwd)/build

GitRemote := $(shell git remote -v | xargs)
GitTag := $(shell git tag --sort=version:refname | tail -n 1)
GitCommit := $(shell git log --pretty=oneline -n 1)
BuildTime := $(shell date +'%Y.%m.%d.%H%M%S')
GoVersion := $(shell go version)

ifeq ($(MAKECMDGOALS),)
	plugins := $(wildcard zplugin/native/*)
else
	plugins := $(MAKECMDGOALS)
endif

.PHONY: all clean $(plugins)

all: clean gdk $(plugins)

$(plugins):
	$(MAKE) -C $@

gdk:
	go generate
	go build -v -ldflags "\
		-X 'main.GitRemote=${GitRemote}' \
		-X 'main.GitTag=${GitTag}' \
		-X 'main.GitCommit=${GitCommit}' \
		-X 'main.BuildTime=${BuildTime}' \
		-X 'main.GoVersion=${GoVersion}' " \
		-o $(BIN)/$@ main.go

clean:
	@echo Cleaning...
	rm -rf $(BIN)
