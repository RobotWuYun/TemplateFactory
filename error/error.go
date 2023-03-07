package errs

import (
	"fmt"
	"time"

	"github.com/spf13/cast"
)

type SelfError struct {
	source error
	desc   string
	date   time.Time
}

func (e SelfError) SimpleStr() (str string) {
	return fmt.Sprintf("[Error] [%s] [%s]", cast.ToString(e.date), e.desc)
}

func (e SelfError) Str() (str string) {
	desc := ""
	if e.source != nil {
		desc = e.source.Error()
	}
	return fmt.Sprintf("[Error] [%s] [%s][sorurce error : %s]", cast.ToString(e.date), desc, e.source.Error())
}

// General
var ErrGeneral = func(err error) (selferr *SelfError) {
	return &SelfError{source: err, desc: err.Error(), date: time.Now()}
}

// input
var ErrInput = func(err error) (selferr *SelfError) {
	return &SelfError{source: err, desc: "Input err", date: time.Now()}
}

// file
var ErrFileNotFound = func(err error) (selferr *SelfError) {
	return &SelfError{source: err, desc: "File not found", date: time.Now()}
}

var ErrStructNameExist = func(err error) (selferr *SelfError) {
	return &SelfError{source: err, desc: "struct name is exist,please check your source", date: time.Now()}
}

// field
var ErrFieldNameHasUppper = func(err error) (selferr *SelfError) {
	return &SelfError{source: err, desc: "field name has upper", date: time.Now()}
}

// config
