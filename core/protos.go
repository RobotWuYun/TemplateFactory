package core

import (
	errs "protoc-gen-foo/error"

	"google.golang.org/protobuf/compiler/protogen"
)

func MakeMessageFile(plugin *protogen.Plugin, file *protogen.File) (err *errs.SelfError) {
	err = MakeStructsFromFile(plugin, file)
	if err != nil {
		return
	}
	err = MakeSQLsFromFile(plugin, file)
	if err != nil {
		return
	}
	err = MakeEntsFromFile(plugin, file)
	if err != nil {
		return
	}
	return
}
