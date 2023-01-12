package core

import (
	errs "TemplateFactory/error"
	"TemplateFactory/utils"
	"fmt"
	"reflect"
	"regexp"
	"strings"
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

func doStruct2pb(t reflect.Type) (res string) {
	newStructs := make([]string, 0)
	newFields := make([]string, 0)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name := f.Tag.Get("json")
		var newType string
		switch tp := f.Type.String(); tp {
		case "string", "interface {}":
			newType = "string"
		case "[]string":
			newType = "repeated string"
		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
			newType = "int64"
		case "time.Time":
			newType = "google.protobuf.Timestamp"
		case "float32", "float64":
			newType = "float64"
		case "bool":
			newType = "bool"
		default:
			switch k := f.Type.Kind(); k {
			case reflect.Struct:
				if strings.Contains(strings.ToLower(f.Name), "image") {
					newType = "Image"
				} else {
					sub := s.doStruct2pb(f.Type)
					title := strings.Title(name)
					newStructs = append(newStructs, fmt.Sprintf(`message %s {%s}`, title, sub))
					newType = title
				}
			case reflect.Slice:
				sub := s.doStruct2pb(f.Type.Elem())
				title := strings.Title(strings.TrimRight(name, "s"))
				newStructs = append(newStructs, fmt.Sprintf(`message %s {%s}`, title, sub))
				newType = "repeated " + title
			default:
				fmt.Println("Unrecognized kind", k, f.Type.Kind())
				continue
			}
		}
		newFields = append(newFields, newType+" "+name)
	}
	for _, v := range newStructs {
		res += v
	}
	for i, v := range newFields {
		res += fmt.Sprintf(`%s = %d;`, v, (i + 1))
	}
	return
}
