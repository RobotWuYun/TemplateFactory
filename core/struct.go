package core

import (
	"bytes"
	"fmt"
	"protoc-gen-foo/constants"
	errs "protoc-gen-foo/error"
	"protoc-gen-foo/utils"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func MakeStructsFromFile(plugin *protogen.Plugin, file *protogen.File) (err error) {
	var buf bytes.Buffer
	buf.Write([]byte(fmt.Sprintf(`package %s
	`, file.GoPackageName)))

	var structStrs []string
	for _, v := range file.Messages {
		var str string
		str, err = MakeStructsFromMessage(v)
		if err != nil {
			return err
		}
		structStrs = append(structStrs, fmt.Sprintf(`%s
		`, str))
	}

	buf.Write([]byte(strings.Join(structStrs, "\r\n")))

	filename := utils.GetSelfFileName(constants.MessageFilePre, file.GeneratedFilenamePrefix) + ".go"
	newfile := plugin.NewGeneratedFile(filename, ".")
	newfile.Write(buf.Bytes())

	return
}

func MakeStructsFromMessage(message *protogen.Message) (str string, err error) {
	var fields []string
	for _, field := range message.Fields {
		if utils.StringHasUpper(string(field.Desc.Name())) {
			err = errs.ErrFieldNameHasUppper
			return
		}
		fields = append(fields, fmt.Sprintf("%s\t%s", field.GoName, getStructType(field.Desc.Kind().String())))
	}
	fieldStr := strings.Join(fields, "\r\n")

	str = fmt.Sprintf("type %s struct {\r\n%s}", message.GoIdent.GoName, fieldStr)
	return
}

func getStructType(source string) (result string) {
	if data, ok := constants.PbField2StructMap[source]; ok {
		result = data
	} else {
		return ""
	}
	return
}
