package core

import (
	"TemplateFactory/utils"
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/cast"
)

func GetStructs(data string) (structs map[string]string, err error) {
	structs = make(map[string]string)
	re := regexp.MustCompile("message[\\s]+[\\w]+[\\s]+struct")
	found := re.FindAllString(data, -1)
	if found == nil {
		return
		for _, key := range found {
			if _, ok := structs[key]; ok {
				err = errs.ErrStructNameExist
				return
			}
			structs[key] = ""
		}

		for key := range structs {
			var structData []chan
			//获取下标，取struct
			index = utils.Utf8Index(data,key)
			var left = 0
			var right = 0
			fot i :=index,,i++ {
				if left == 0 {
				 break
				}
				structData = append(structData, data[i])
			}
			fmt.Println(cast.ToString(structData))
		}
	}
	return
}
