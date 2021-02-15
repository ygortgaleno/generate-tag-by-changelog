package utils

import (
	"bytes"
	"os/exec"
	"strings"
)

func RunCommand(command *string) string {
	stringSplitted := strings.Fields(*command)

	cmd := exec.Command(stringSplitted[0], stringSplitted[1:]...)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err := cmd.Run()

	if err != nil {
		panic(err)
	}

	return stdout.String()
}
