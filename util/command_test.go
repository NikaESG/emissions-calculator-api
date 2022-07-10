package util

import (
	"fmt"
	"gotest.tools/v3/assert"
	"testing"
)

func TestRunCommand(t *testing.T) {
	out, err := RunCommand("/bin/sh", "../shell/GetConfimap.sh", "test")
	fmt.Println(out)
	assert.Equal(t, err, nil)
}
