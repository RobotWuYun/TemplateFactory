package core

import (
	"TemplateFactory/config"
)

var Protos protoModel

type protoModel struct{}

func (c *protoModel) MakeFileProto(product config.Product) (structName string, err error) {
	return
}

func (c *protoModel) MakeDirProto(product config.Product) (structNames []string, err error) {
	return
}
