package config

import (
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
}
