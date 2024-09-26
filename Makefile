SHELL := /bin/bash

export GOPROXY=https://goproxy.cn,direct
export GO111MODULE=on
export GOSUMDB=off
export OUTPUT_DIR=bin

# default APPVER or from $VAR or command-line
APPVER ?= 0.0.0

ci: gdk verbose
	

%:
	go build -buildmode=plugin -o $(OUTPUT_DIR)/$@.so filters/plugins/$@/*.go

vendor:
	go mod tidy && go mod vendor

gdk:
	go generate
	go build -v -ldflags " \
		-X 'main.GitRemote=$(shell git remote -v | xargs)' \
		-X 'main.GitTag=$(shell git tag --sort=version:refname | tail -n 1)' \
		-X 'main.GitCommit=$(shell git log --pretty=oneline -n 1)' \
		-X 'main.BuildTime=$(shell date +'%Y.%m.%d.%H%M%S')' \
		-X 'main.GoVersion=$(shell go version)' \
		-X 'github.com/cocktail828/gdk/v1/cmd/status.AppVersion=$(APPVER)' \
	" -o $(OUTPUT_DIR)/$@ main.go

clean:
	@echo Cleaning...
	rm -rf $(OUTPUT_DIR)
