package core

import (
	"regexp"
)

func GetStructs(data string) (structs map[string]string, err error) {
	structs = make(map[string]string)
	re := regexp.MustCompile("(?i)fox(es)?")
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
		}
	}
	return
}
