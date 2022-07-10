package util

import (
	"fmt"
	"os/exec"

	"trino.com/trino-connectors/util/log"
)

func RunCommand(name string, arg ...string) (string, error) {

	cmd := exec.Command(name, arg...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Printf("[RunCommand] [cmd.Output] error: %s", err)
		log.Logger().Errorf("[RunCommand] [cmd.Output] error: %s", err)
		return "", err
	}

	return string(stdout), nil
}
