package command

import (
	"bufio"
	"log"
	"os/exec"
	"strings"
)

func Execute(name string, arg ...string) {
	cmdSlice := append([]string{name}, arg...)
	log.Println("rum cmd: " + strings.Join(cmdSlice, " "))

	cmd := exec.Command(name, arg...)
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	_ = cmd.Start()
	scanner := bufio.NewScanner(stdoutPipe)
	scanner.Split(bufio.ScanLines)

	errScanner := bufio.NewScanner(stderrPipe)
	errScanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}
	for errScanner.Scan() {
		log.Println(errScanner.Text())
	}
	_ = cmd.Wait()
}
