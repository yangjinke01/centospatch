package cmd

import (
	"errors"
	"os"
	"testing"
)

func Test_unpackTarball(t *testing.T) {
	if _, err := os.Stat("/tmp/"); errors.Is(err, os.ErrNotExist) {
		// file does not exist
	} else {
		// file exists
	}
}
