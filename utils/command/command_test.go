package command

import "testing"

func TestExecute(t *testing.T) {
	Execute("ls", "/")
	Execute("ls", "/asdfasdf")
}
