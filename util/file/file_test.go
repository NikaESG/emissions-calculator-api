package file

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestReadFile(t *testing.T) {
	_, err := ReadFile("../../testFiles/users.yaml")
	assert.Equal(t, err, nil)
}

func TestWriteFile(t *testing.T) {
	err := WriteFile("../../testFiles/users1.yaml", "test")
	assert.Equal(t, err, nil)
}
