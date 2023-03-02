package main

import (
	"fmt"
	"protoc-gen-foo/config"
)

func main() {
	//cmd.Start()
	fmt.Println(config.GetConf())
}
