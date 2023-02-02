package core

import "google.golang.org/protobuf/compiler/protogen"

func MakeStructsFromFile(file *protogen.File) {
	for _, v := range file.Messages {
		MakeStructsFromMessage(v)
	}
}

func MakeStructsFromMessage(message *protogen.Message) string {
	return ""
}
