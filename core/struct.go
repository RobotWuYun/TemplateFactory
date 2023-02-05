package core

import (
	"bytes"
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func MakeStructsFromFile(plugin *protogen.Plugin, file *protogen.File) {
	var structStrs []string

	for _, v := range file.Messages {
		structStrs = append(structStrs, MakeStructsFromMessage(v))
	}

	var buf bytes.Buffer
	pkg := fmt.Sprintf("package %s", file.GoPackageName)
	buf.Write([]byte(pkg))
	buf.Write([]byte(strings.Join(structStrs, "\n")))
	filename := file.GeneratedFilenamePrefix + ".foo.go"
	newfile := plugin.NewGeneratedFile(filename, ".")

	// 将内容写入插件文件内容
	newfile.Write(buf.Bytes())
}

func MakeStructsFromMessage(message *protogen.Message) string {
	var fields []string
	for _, field := range message.Fields {
		fields = append(fields, fmt.Sprintf("%s\t%s", field.GoName, field.Desc.Kind().GoString()))
		fmt.Println("TEST")
	}
	fieldStr := strings.Join(fields, "\n")
	return fmt.Sprintf("type %s struct {/n%s}", message.GoIdent.GoName, fieldStr)
}
