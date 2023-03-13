package core

import (
	"protoc-gen-foo/config"

	"google.golang.org/protobuf/compiler/protogen"
)

func MakeMessageFile(plugin *protogen.Plugin, file *protogen.File, config config.Config) (err error) {
	if config.Struct.Make {
		err = MakeStructsFromFile(plugin, file, config)
		if err != nil {
			return
		}
	}
	if config.Sql.Make {
		err = MakeSQLsFromFile(plugin, file, config.Sql)
		if err != nil {
			return
		}
	}
	if config.Ent.Make {
		err = MakeEntsFromFile(plugin, file, config.Ent)
		if err != nil {
			return
		}
	}
	return
}
