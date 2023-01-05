package config

import (
	"TemplateFactory/config/utils"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type product struct {
	Production bool
	Path       string
}
type safeProduct struct {
	product
	Test bool
}

type Config struct {
	Source    string
	IsDir     bool
	Overwrite bool
	Ent       product
	Proto     product
	Data      safeProduct
	Biz       safeProduct
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
