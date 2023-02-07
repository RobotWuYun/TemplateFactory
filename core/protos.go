package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"protoc-gen-foo/constants"
	"protoc-gen-foo/utils"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func GetMessage() {
	plugin := getPlugin()

	for _, file := range plugin.Files {
		fileName := utils.GetFileName(file.GeneratedFilenamePrefix)

		if strings.HasPrefix(fileName, constants.MessageFilePre) {
			MakeStructsFromFile(plugin, file)
		} else if strings.HasPrefix(fileName, constants.ServiceFilePre) {
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
	input, _ := ioutil.ReadAll(os.Stdin)

	var req pluginpb.CodeGeneratorRequest
	proto.Unmarshal(input, &req)

	opts := protogen.Options{}
	p, err := opts.New(&req)
	if err != nil {
		panic(err)
	}
	return
}
