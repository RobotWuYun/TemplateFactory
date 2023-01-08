package core

import (
	"TemplateFactory/config"
	"TemplateFactory/utils"
)

var Protos protoModel

type protoModel struct{}

func (c *protoModel) MakeFileProto(product config.Product) error {
	data, err := utils.GetString(product.Path)
	if err != nil {
		return err
	}
	return nil
}

func (c *protoModel) MakeDirProto(product config.Product) {
}
