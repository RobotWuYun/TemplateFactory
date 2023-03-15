package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	errs "protoc-gen-foo/error"
)

type pubConf struct {
}

type EntConf struct {
	Make       bool   `yaml:"make"`
	FilePrefix string `yaml:"filePrefix"`
}

type SqlConf struct {
	Make       bool   `yaml:"make"`
	FilePrefix string `yaml:"filePrefix"`
}

type StructConf struct {
	Make         bool   `yaml:"make"`
	FilePrefix   string `yaml:"filePrefix"`
	StructSuffix string `yaml:"structSuffix"`
}

type Config struct {
	Ent    EntConf    `yaml:"ent"`
	Sql    SqlConf    `yaml:"sql"`
	Struct StructConf `yaml:"struct"`
}

func GetConf() (conf Config, err error) {
	yamlFile, err := ioutil.ReadFile("config/proto.yaml")
	if err != nil {
		err = errs.ErrFileNotFound(err)
		return
	}
	// 将读取的yaml文件解析为响应的 struct
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		err = errs.ErrGeneral(err)
		return
	}
	return
}
