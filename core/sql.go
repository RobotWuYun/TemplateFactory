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

func MakeSQLsFromFile(plugin *protogen.Plugin, file *protogen.File) (err error) {
	var buf bytes.Buffer

	var sqlStrs []string
	for _, v := range file.Messages {
		var sqlStr string
		if sqlStr, err = MakeSqlFromMessage(v); err == nil {
			sqlStrs = append(sqlStrs, fmt.Sprintf(`%s
		`, sqlStr))
		} else {
			return
		}
	}

	buf.Write([]byte(strings.Join(sqlStrs, "\r\n")))

	filename := utils.GetSelfFileName(constants.MessageFilePre, file.GeneratedFilenamePrefix) + ".sql"
	newfile := plugin.NewGeneratedFile(filename, ".")
	newfile.Write(buf.Bytes())

	return
}

func MakeSqlFromMessage(message *protogen.Message) (content string, err error) {
	notes := fmt.Sprintf("-- %s\r\n", message.GoIdent.GoName)
	dorpSQL := fmt.Sprintf("DROP TABLE IF EXISTS `%s`;\r\n", message.GoIdent.GoName)

	var hasFieldID bool
	var fields []string
	for _, field := range message.Fields {
		if utils.StringHasUpper(field.GoName) {
			err = errs.ErrFieldNameHasUppper
			return
		}
		if field.GoName == "id" {
			hasFieldID = true
		}
		var fileType string
		switch field.Desc.Kind().String() {
		case "string":
			fileType = "varchar"
		}
		fields = append(fields, fmt.Sprintf(" `%s` %s DEFAULT NULL COMMENT '%s',", field.GoName, fileType, field.GoName))
	}

	if !hasFieldID {
		fields = append([]string{"`id` bigint(20) NOT NULL AUTO_INCREMENT,"}, fields...)
	}

	fieldStr := strings.Join(fields, "\r\n")

	content = notes + dorpSQL + fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (\r\n%s PRIMARY KEY (`id`),\r\n) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '%s';", message.GoIdent.GoName, fieldStr, message.GoIdent.GoName)
	return
}
