BIN := bin
GitTag := $(shell git tag --sort=version:refname | tail -n 1)
GitCommitLog := $(shell git log --pretty=oneline -n 1)
BuildTime := $(shell date +'%Y.%m.%d.%H%M%S')
GoVersion := $(shell go version)

gdk:
	go build -v -ldflags "\
		-X 'main.GitTag=${GitTag}' \
		-X 'main.GitCommitLog=${GitCommitLog}' \
		-X 'main.BuildTime=${BuildTime}' \
		-X 'main.GoVersion=${GoVersion}' " \
		-o ${BIN}/$@ main.go

clean:
	rm -rf ${BIN}