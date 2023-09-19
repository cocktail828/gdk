export GO111MODULE=on
export GOPROXY=https://goproxy.cn

gen:
	protoc \
		-I . \
		-I validate \
		--go_out=":." \
		--validate_out="lang=go:." \
		message.proto
env:
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go get google.golang.org/protobuf/proto
	go get google.golang.org/grpc
	go install github.com/envoyproxy/protoc-gen-validate
	go install google.golang.org/protobuf/cmd/protoc-gen-go