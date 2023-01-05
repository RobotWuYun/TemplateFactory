package cmd

import (
	"TemplateFactory/config"
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

	//

}
