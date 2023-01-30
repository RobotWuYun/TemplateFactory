# TemplateFactory

make go code from proto

## Function

make

- proto
- dao (change by config)
  - ent
  - sqlmock
- biz
  from go struct

## 如何安装 proto

1. 下载 proto https://github.com/protocolbuffers/protobuf/releases
2. 解压，并将 bin 目录配置到环境变量中
3. protoc --version 测试安装

4. go get -u "google.golang.org/protobuf"
5. go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
6. 将 GOPATH/bin 下的 protoc-gen-go 复制到 GOROOT（注意：需要将 bin 目录配置到环境变量中）
