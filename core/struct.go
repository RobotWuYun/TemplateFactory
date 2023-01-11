package core

import (
	errs "TemplateFactory/error"
	"TemplateFactory/utils"
	"regexp"
)

type StrcutModel struct {
	FilePath string
	Data     string
	Exist    bool
}

func GetStructs(data string) (structs map[string]string, err error) {
	structs = make(map[string]string)

	re := regexp.MustCompile("type[\\s]+[\\w]+[\\s]+struct[\\s]*{")
	found := re.FindAllString(data, -1)
	if found == nil {
		return
	}

	for _, key := range found {
		if _, ok := structs[key]; ok {
			err = errs.ErrStructNameExist
			return
		}
		structs[key] = ""
	}

	for key := range structs {
		var structData []rune
		var index = utils.Utf8Index(data, key) + len(key)

		var single = 1
		for {
			if single == 0 || index == (len(data)-1) {
				break
			}

			charData := rune(data[index])
			if charData == '{' {
				single++
			} else if charData == '}' {
				single--
			}
			structData = append(structData, rune(data[index]))
			index++
		}
		structs[key] = key + string(structData)
	}
	return
}
