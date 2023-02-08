package core

import (
	"bytes"
	"fmt"
	"protoc-gen-foo/constants"
	"protoc-gen-foo/utils"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func MakeStructsFromFile(plugin *protogen.Plugin, file *protogen.File) {
	var buf bytes.Buffer
	buf.Write([]byte(fmt.Sprintf(`package %s
	`, file.GoPackageName)))

	var structStrs []string
	for _, v := range file.Messages {
		structStrs = append(structStrs, fmt.Sprintf(`%s
		`, MakeStructsFromMessage(v)))
	}

	buf.Write([]byte(strings.Join(structStrs, "\r\n")))

	filename := utils.GetSelfFileName(constants.MessageFilePre, file.GeneratedFilenamePrefix) + ".go"
	newfile := plugin.NewGeneratedFile(filename, ".")
	newfile.Write(buf.Bytes())

	return
}

func MakeStructsFromMessage(message *protogen.Message) string {
	var fields []string
	for _, field := range message.Fields {
		fields = append(fields, fmt.Sprintf("%s\t%s", field.GoName, field.Desc.Kind().String()))
	}
	fieldStr := strings.Join(fields, "\r\n")

	return fmt.Sprintf("type %s struct {\r\n%s}", message.GoIdent.GoName, fieldStr)
}
