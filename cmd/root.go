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
	var structMap = make(map[string]core.StrcutModel)

	// 循环处理每一个文件
	for _, path := range files {
		fileMap[path], err = utils.GetString(path)
		if err != nil {
			log.Default().Println("Read file err :", path)
		}
	}

	//var structMap = make(map[string]map[string]string)
	for fileName, data := range fileMap {
		if len(data) == 0 {
			return
		}
		var structs map[string]string
		structs, err = core.GetStructs(data)
		if err != nil {
			log.Default().Println("parse fail :", fileName)
			continue
		}
		for key, modStr := range structs {
			if _, ok := structMap[key]; ok {
				log.Default().Println("strcut is exist :", key)
			} else {
				structMap[key] = core.StrcutModel{
					FilePath: fileName,
					Data:     modStr,
				}
			}
		}

	}

}
