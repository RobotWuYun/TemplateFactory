package core

import (
	"bytes"
	"fmt"
	"protoc-gen-foo/config"
	"protoc-gen-foo/constants"
	errs "protoc-gen-foo/error"
	"protoc-gen-foo/utils"

	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func MakeEntsFromFile(plugin *protogen.Plugin, file *protogen.File, config config.EntConf) (err error) {
	var buf bytes.Buffer

	var entStrs []string
	for _, v := range file.Messages {
		var entStr string
		if entStr, err = MakeEntFromMessage(v); err == nil {
			entStrs = append(entStrs, fmt.Sprintf(`%s
		`, entStr))
		} else {
			return
		}
	}

	buf.Write([]byte(strings.Join(entStrs, "\r\n")))

	if len(config.FilePrefix) == 0 {
		config.FilePrefix = `ent/schema/`
	}

	filename := utils.GetSelfFileName(constants.MessageFilePre, file.GeneratedFilenamePrefix) + ".go"
	newfile := plugin.NewGeneratedFile(config.FilePrefix+filename, "./schema")
	newfile.Write(buf.Bytes())

	return
}

func MakeEntFromMessage(message *protogen.Message) (content string, err error) {
	head := `package schema
	
	import (
		"entgo.io/ent"
		"entgo.io/ent/dialect/entsql"
		"entgo.io/ent/schema"
		"entgo.io/ent/schema/field"
	)

	`

	typeStr := fmt.Sprintf(`// %s holds the schema definition for the %s entity.
	type %s struct {
		ent.Schema
	}
	`, message.GoIdent.GoName, message.GoIdent.GoName, message.GoIdent.GoName)

	annotationsStr := fmt.Sprintf(`// Annotations .
	func (%s) Annotations() []schema.Annotation {
		return []schema.Annotation{
			entsql.Annotation{Table: "%s"},
		}
	}
	`, message.GoIdent.GoName, utils.ToSnakeCase(message.GoIdent.GoName))

	EdgesStr := fmt.Sprintf(`// Edges of the %s.
	func (%s) Edges() []ent.Edge {
		return []ent.Edge{}
	}
	`, message.GoIdent.GoName, message.GoIdent.GoName)

	//var hasFieldID bool
	var fields []string
	for _, field := range message.Fields {
		if utils.StringHasUpper(string(field.Desc.Name())) {
			err = errs.ErrFieldNameHasUppper(nil)
			return
		}
		// if field.GoName == "id" {
		// 	hasFieldID = true
		// }
		fields = append(fields, fmt.Sprintf(`field.%s("%s"),`, getEntType(field.Desc.Kind().String()), field.Desc.Name()))
	}

	// if !hasFieldID {
	// 	fields = append([]string{"`id` bigint(20) NOT NULL AUTO_INCREMENT,"}, fields...)
	// }

	fieldStr := strings.Join(fields, "\r\n")

	content =
		head +
			typeStr +
			annotationsStr +
			fmt.Sprintf(`// Fields of the %s.
func (%s) Fields() []ent.Field {
	return []ent.Field{
		%s
	}
}
`, message.GoIdent.GoName, message.GoIdent.GoName, fieldStr) +
			EdgesStr
	return
}

func getEntType(source string) (result string) {
	if data, ok := constants.PbField2EntMap[source]; ok {
		result = data
	} else {
		return ""
	}
	return
}
