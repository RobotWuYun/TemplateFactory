package utils

import (
	errs "TemplateFactory/error"
	"io/fs"
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

// 获取文件内容
func GetString(path string) (data string, err error) {
	var dataByte []byte
	dataByte, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}
	data = string(dataByte)
	return
}

// 读取目录下所有文件名
func GetFileNames(path string) (names []string, err error) {
	if len(path) == 0 {
		err = errs.ErrFileNotFound
		return
	}
	if !Exists(path) {
		err = errs.ErrFileNotFound
		return
	}
	if IsDir(path) {
		var files []fs.FileInfo
		files, err = ioutil.ReadDir(path)
		if err != nil {
			return
		}
		for _, v := range files {
			names = append(names, path+"\\"+v.Name())
		}
	} else {
		names = append(names, path)
	}
	return
}
