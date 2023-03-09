package errs

import (
	"fmt"
	"time"

	"github.com/spf13/cast"
)

type SelfError struct {
	error
	desc string
	date time.Time
}

func (e SelfError) SimpleStr() (str string) {
	return fmt.Sprintf("[Error] [%s] [%s]", cast.ToString(e.date), e.desc)
}

func (e SelfError) Str() (str string) {
	desc := ""
	if e.error != nil {
		desc = e.Error()
	}
	return fmt.Sprintf("[Error] [%s] [%s][sorurce error : %s]", cast.ToString(e.date), desc, e.error.Error())
}

// General
var ErrGeneral = func(err error) (selferr *SelfError) {
	return &SelfError{error: err, desc: err.Error(), date: time.Now()}
}

// input
var ErrInput = func(err error) (selferr *SelfError) {
	return &SelfError{error: err, desc: "Input err", date: time.Now()}
}

// file
var ErrFileNotFound = func(err error) (selferr *SelfError) {
	return &SelfError{error: err, desc: "File not found", date: time.Now()}
}

var ErrStructNameExist = func(err error) (selferr *SelfError) {
	return &SelfError{error: err, desc: "struct name is exist,please check your source", date: time.Now()}
}

// field
var ErrFieldNameHasUppper = func(err error) (selferr *SelfError) {
	return &SelfError{error: err, desc: "field name has upper", date: time.Now()}
}

// config
