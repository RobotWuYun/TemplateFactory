package core

import (
	"bytes"
	"fmt"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/compiler/protogen"
	"protoc-gen-foo/config"
	"protoc-gen-foo/constants"
	errs "protoc-gen-foo/error"
	"protoc-gen-foo/utils"
	"strings"
)

func MakeStructsFromFile(plugin *protogen.Plugin, file *protogen.File, config config.Config) (err error) {
	var buf bytes.Buffer
	buf.Write([]byte(fmt.Sprintf(`package %s
	`, file.GoPackageName)))

	var structStrs []string
	for _, v := range file.Messages {
		var initStr string
		initStr, err = makeInitsFromMessage(v, cast.ToString(file.GoPackageName), config.Struct.StructSuffix)
		if err != nil {
			return err
		}
		structStrs = append(structStrs, fmt.Sprintf(`%s
		`, initStr))
	}

	buf.Write([]byte(strings.Join(structStrs, "\r\n")))

	if len(config.Struct.FilePrefix) == 0 {
		config.Struct.FilePrefix = `data/`
	}

	filename := utils.GetSelfFileName(constants.MessageFilePre, file.GeneratedFilenamePrefix) + ".go"
	newfile := plugin.NewGeneratedFile(config.Struct.FilePrefix+filename, ".")
	newfile.Write(buf.Bytes())
	return
}
func makeInitsFromMessage(message *protogen.Message, packname, suffix string) (initStr string, err error) {
	if suffix == "" {
		suffix = "Repo"
	}
	structName := message.GoIdent.GoName + suffix
	firstLowName := utils.FirstLower(structName)

	structStr := fmt.Sprintf(constants.InitStructFormat, firstLowName)

	newFunc := fmt.Sprintf(constants.InitNewFuncFormat, structName, structName, structName, firstLowName, structName, fmt.Sprintf("%s:%s", packname, structName))

	initStr = fmt.Sprintf(`%s

%s`, structStr, newFunc)

	return
}

func makeConvertMessage(message *protogen.Message, conf config.Config) (str string, err error) {
	var fields []string
	for _, field := range message.Fields {
		if utils.StringHasUpper(string(field.Desc.Name())) {
			err = errs.ErrFieldNameHasUppper(nil)
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
