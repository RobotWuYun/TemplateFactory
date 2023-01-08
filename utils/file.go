package utils

import (
	"io/ioutil"
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func GetString(path string) (data string, err error) {
	var dataByte []byte
	dataByte, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}
	data = string(dataByte)
	return
}
