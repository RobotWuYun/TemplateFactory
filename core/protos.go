package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

const (
	MessageFilePre = "message_"
	ServiceFilePre = "service_"
)

func GetMessage() {
	plugin := getPlugin()

	for _, file := range plugin.Files {
		// 通过文件名前缀区分服务和结构
		if strings.HasPrefix(file.GeneratedFilenamePrefix, MessageFilePre) {
			MakeStructsFromFile(plugin, file)
		} else if strings.HasPrefix(file.GeneratedFilenamePrefix, ServiceFilePre) {
			continue
		}
	}

	// 生成响应
	stdout := plugin.Response()
	out, err := proto.Marshal(stdout)
	if err != nil {
		panic(err)
	}

	// 将响应写回标准输入, protoc会读取这个内容
	fmt.Fprintf(os.Stdout, string(out))
}

func getPlugin() (p *protogen.Plugin) {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	var req pluginpb.CodeGeneratorRequest
	err = proto.Unmarshal(input, &req)
	if err != nil {
		panic(err)
	}

	opts := protogen.Options{}
	p, err = opts.New(&req)
	if err != nil {
		panic(err)
	}
	return
}
