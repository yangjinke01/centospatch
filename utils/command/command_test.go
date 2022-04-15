package command

import (
	"testing"
)

func TestExecute(t *testing.T) {
	Execute("/bin/bash", "-c", "sleep 10 && echo a")
}
