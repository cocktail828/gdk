SHELL := /bin/bash

all:
	go build -buildmode=plugin -o $(BIN)/$(shell basename `pwd`).so *.go
