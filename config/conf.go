package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	errs "protoc-gen-foo/error"
)

type pubConf struct {
	Make       bool
	FilePrefix string
}

type EntConf struct {
	pubConf
}

type SqlConf struct {
	pubConf
}

type StructConf struct {
	pubConf
	StructSuffix string
}

type Config struct {
	Ent    EntConf
	Sql    SqlConf
	Struct StructConf
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
