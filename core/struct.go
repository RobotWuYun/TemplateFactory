package core

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func MakeStructsFromFile(plugin *protogen.Plugin, file *protogen.File) string {
	var structStrs []string

	for _, v := range file.Messages {
		structStrs = append(structStrs, MakeStructsFromMessage(v))
	}

	return strings.Join(structStrs, "\n")
}

func MakeStructsFromMessage(message *protogen.Message) string {
	var fields []string
	for _, field := range message.Fields {
		fields = append(fields, fmt.Sprintf("%s\t%s", field.GoName, field.Desc.Kind().GoString()))
	}
	fieldStr := strings.Join(fields, "\n")

	return fmt.Sprintf("type %s struct {\n%s}", message.GoIdent.GoName, fieldStr)
}
