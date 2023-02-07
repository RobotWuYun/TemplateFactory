package utils

import "strings"

func GetFileName(path string) (name string) {
	if path == "" {
		return
	}
	fileNameIndex := strings.Split(path, "/")
	return fileNameIndex[len(fileNameIndex)-1]
}

func GetFileNameWithoutFormat(path string) (name string) {
	if path == "" {
		return
	}
	name = GetFileName(path)
	fileNameIndex := strings.Split(name, ".")
	return fileNameIndex[0]
}

func GetSelfFileName(pre, path string) (name string) {
	if path == "" {
		return
	}
	name = GetFileNameWithoutFormat(path)
	return strings.ReplaceAll(name, pre, "")
}
