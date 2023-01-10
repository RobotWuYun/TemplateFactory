package cmd

import (
	"TemplateFactory/config"
	"TemplateFactory/core"
	"TemplateFactory/utils"

	"log"
)

type data struct {
	Config config.Config
}

func Start() {
	log.Default().Println("template factory start")
	// config init
	var dataMod data
	dataMod.Config.InitConfig()
	makeFormConfig(dataMod.Config)
}

func makeFormConfig(c config.Config) {
	files, err := utils.GetFileNames(c.Source)
	if err != nil {
		log.Fatalf("GET fileNames err :", err)
	}
	var fileMap = make(map[string]string)

	// 循环处理每一个文件
	for _, path := range files {
		fileMap[path], err = utils.GetString(path)
		if err != nil {
			log.Default().Println("Read file err :", path)
		}
	}

	//var structMap = make(map[string]map[string]string)
	for _, data := range fileMap {
		if len(data) == 0 {
			return
		}
		core.GetStructs(data)
	}
}
