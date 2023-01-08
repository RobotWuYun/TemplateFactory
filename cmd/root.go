package cmd

import (
	"TemplateFactory/config"
	"TemplateFactory/utils"
	"fmt"
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
	for _, v := range files {
		fmt.Println(v)
	}
}
