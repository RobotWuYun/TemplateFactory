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
		// 通过文件名前缀区分服务和结构
		if strings.HasPrefix(file.GeneratedFilenamePrefix, MessageFilePre) {
			MakeStructsFromFile(file)
		} else if strings.HasPrefix(file.GeneratedFilenamePrefix, ServiceFilePre) {

		}

		//创建一个buf 写入生成的文件内容
		var buf bytes.Buffer

		// 写入go 文件的package名
		pkg := fmt.Sprintf("package %s", file.GoPackageName)
		buf.Write([]byte(pkg))
		content := ""
		//遍历消息,这个内容就是protobuf的每个消息
		for _, msg := range file.Messages {
			mapSrc := `
                         newMap:=make(map[%v]%v)
                     for k,v:=range x.%v {
                        newMap[k]=v
                     }
                            x.%v=newMap
                  `
			//遍历消息的每个字段
			for _, field := range msg.Fields {
				//只有map 才这样做
				if field.Desc.IsMap() {
					content += fmt.Sprintf(mapSrc, field.Desc.MapKey().Kind().String(), field.Desc.MapValue().Kind().String(), field.GoName, field.GoName)
				}
			}
			buf.Write([]byte(fmt.Sprintf(`
           func (x *%s) CheckMap() {
          %s
           }`, msg.GoIdent.GoName, content)))
		}
		//指定输入文件名,输出文件名为demo.foo.go
		filename := file.GeneratedFilenamePrefix + ".foo.go"
		file := plugin.NewGeneratedFile(filename, ".")

		// 将内容写入插件文件内容
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
