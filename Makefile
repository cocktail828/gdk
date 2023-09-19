BIN := bin
GitTag := $(shell git tag --sort=version:refname | tail -n 1)
GitCommitLog := $(shell git log --pretty=oneline -n 1)
BuildTime := $(shell date +'%Y.%m.%d.%H%M%S')
BuildGoVersion := $(shell go version)

gdk:
	go build -v -ldflags "\
		-X 'main.GitTag=${GitTag}' \
		-X 'main.GitCommitLog=${GitCommitLog}' \
		-X 'main.BuildTime=${BuildTime}' \
		-X 'main.BuildGoVersion=${BuildGoVersion}' " \
		-o ${BIN}/gdk main.go

clean:
	rm -rf ${BIN}