package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"protoc-gen-foo/config"
	"protoc-gen-foo/constants"
	"protoc-gen-foo/core"
	errs "protoc-gen-foo/error"
	"protoc-gen-foo/utils"

	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func Start() {
	config, err := config.GetConf()
	if err != nil {
		panic(err.(errs.SelfError).Str())
	}

	plugin, err, _ := getPlugin()
	if err != nil {
		panic(err.(errs.SelfError).Str())
	}

	for _, file := range plugin.Files {
		fileName := utils.GetFileName(file.GeneratedFilenamePrefix)

		if strings.HasPrefix(fileName, constants.MessageFilePre) {
			err := core.MakeMessageFile(plugin, file, config)
			if err != nil {
				panic(err.(*errs.SelfError).Str)
			}
		} else if strings.HasPrefix(fileName, constants.ServiceFilePre) {
			continue
		}
	}

	// 生成响应
	stdout := plugin.Response()
	out, err := proto.Marshal(stdout)
	if err != nil {
		panic(err)
	}

	// 将响应写回标准输入, protoc会读取这个内容
	fmt.Fprintf(os.Stdout, string(out))
}

func getPlugin() (p *protogen.Plugin, err error, input []byte) {
	input, err = ioutil.ReadAll(os.Stdin)
	if err != nil {
		err = errs.ErrInput(err)
		return
	}
	var req pluginpb.CodeGeneratorRequest
	err = proto.Unmarshal(input, &req)
	if err != nil {
		err = errs.ErrGeneral(err)
		return
	}

	opts := protogen.Options{}
	p, err = opts.New(&req)
	if err != nil {
		err = errs.ErrGeneral(err)
		return
	}
	return
}
