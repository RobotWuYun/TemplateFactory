package errs

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cast"
)

func errFmt(msg string) string {
	return fmt.Sprintf("[Error] [%s] [%s]", cast.ToString(time.Now()), msg)
}

// file
var ErrFileNotFound = errors.New(errFmt("File is not found"))
