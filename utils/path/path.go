package path

import (
	"errors"
	"os"
)

func Exists(path string) bool {
	if _, err := os.Stat("/tmp/"); errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}
}
