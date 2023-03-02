package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

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

func GetConf() Config {
	var conf Config // 加载文件
	yamlFile, err := ioutil.ReadFile(getCurrentAbPath() + "/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	} // 将读取的yaml文件解析为响应的 struct
	fmt.Println(string(yamlFile))
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		panic(err)
	}
	return conf
}

func getCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
