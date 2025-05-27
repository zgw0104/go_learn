.PHONY: all build run gotool clean help

BINARY="catcat"

all: gotool buld

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run:
	@go run ./main.go conf/config.yaml

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ]; then rm ${BINARY} ; fi

help:
	@echo "make - 格式化 GO 代码，并编译二进制文件"
	@echo "make build - 编译GO 代码 生成二进制文件"
	@echo "make run - 直接运行GO代码"
	@echo "make clean - 移除二进制文件 和 vim swap files"
	@echo "gotool - 运行GO工具 'fmt' and 'vet' "