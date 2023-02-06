package core

import (
	"bytes"
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
		var fileContent string
		var buf bytes.Buffer
		pkg := fmt.Sprintf(`package %s`, file.GoPackageName)

		fileName := strings.Split(file.GeneratedFilenamePrefix, "/")

		if strings.HasPrefix(fileName[len(fileName)-1], MessageFilePre) {
			fileContent = MakeStructsFromFile(plugin, file)
		} else if strings.HasPrefix(fileName[len(fileName)-1], ServiceFilePre) {
			continue
		}
		buf.Write([]byte(pkg))

		buf.Write([]byte(fileContent))

		filename := file.GeneratedFilenamePrefix + ".foo.go"
		file := plugin.NewGeneratedFile(filename, ".")
		file.Write(buf.Bytes())
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
