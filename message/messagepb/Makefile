export GO111MODULE=on
export GOPROXY=https://goproxy.cn

default:
	rm -rf messagepb && protoc -I . --go_out=":." message.proto