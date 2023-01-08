package config

import (
	"TemplateFactory/core"
	"TemplateFactory/utils"

	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Product struct {
	Production bool
	Path       string
	Overwrite  bool
}
type SafeProduct struct {
	Product
	Test bool
}

type Config struct {
	Source string
	IsDir  bool
	Ent    Product
	Proto  Product
	Data   SafeProduct
	Biz    SafeProduct
}

func (c *Config) InitConfig() {
	config, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatalf("Loding config err : %v", err)
	}
	err = yaml.Unmarshal(config, &c)
	if err != nil {
		log.Fatalf("Unmarshal config err : %v", err)
	}
	c.checkSource()
}

func (c *Config) checkSource() {
	if len(c.Source) == 0 {
		log.Fatalf("source is null")
		return
	}
	if !utils.Exists(c.Source) {
		log.Fatalf("source is not exists")
		return
	}
	c.IsDir = utils.IsDir(c.Source)
}

func (c *Config) MakeProto() error {
	if c.Proto.Production {
		if c.IsDir {
			return core.Protos.MakeDirProto(c.Proto)
		} else {
			core.Protos.MakeFileProto(c.Proto)
		}
	}
	return nil
}
