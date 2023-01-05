package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Overwrite  bool
	Production struct {
		Ent   bool `yaml:"ent"`
		Proto bool `yaml:"proto"`
		Data  bool `yaml:"data"`
		Biz   bool `yaml:"biz"`
	}
	Path struct {
		Ent   string `yaml:"ent"`
		Proto string `yaml:"proto"`
		Data  string `yaml:"data"`
		Biz   string `yaml:"biz"`
	}
	Test struct {
		Data bool `yaml:"data"`
		Biz  bool `yaml:"biz"`
	}
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
