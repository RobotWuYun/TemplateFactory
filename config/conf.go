package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Ent struct {
		Make bool
	}
	Sql struct {
		Make bool
	}
	Struct struct {
		Make bool
	}
}

func GetConf(path string) Config {
	var conf Config // 加载文件
	yamlFile, err := ioutil.ReadFile(path + "../config/config.yaml")
	if err != nil {
		panic(err)
	} // 将读取的yaml文件解析为响应的 struct
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		panic(err)
	}
	return conf
}
