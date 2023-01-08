package config

import (
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
}
