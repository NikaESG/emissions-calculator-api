package util

import (
	"fmt"
	"gotest.tools/v3/assert"
	"strings"
	"testing"
)

type User struct {
	Name       string
	Occupation string
}

func TestParseYaml(t *testing.T) {
	d := make(map[string]User)
	err := ParseYaml("../testFiles/users.yaml", &d)
	fmt.Println(d)
	assert.Equal(t, nil, err)
}

func TestParseYaml2(t *testing.T) {
	d := make(map[interface{}]interface{})
	err := ParseYaml("../testFiles/test.yaml", &d)

	fmt.Println(d["data"])

	v := fmt.Sprint(d["data"])
	vs := strings.Split(v[4:len(v)-1], "\n")
	for _, s := range vs {
		d := strings.Split(s, "=")
		if len(d) < 2 {
			break
		}
		fmt.Printf("i: %v, v : %v", d[0], d[1])
	}

	assert.Equal(t, nil, err)
}

func TestWriteYaml(t *testing.T) {
	users := map[string]User{"user 1": {"John Doe", "gardener"},
		"user 2": {"Lucy Black", "teacher"}}
	err := WriteYaml("../testFiles/users1.yaml", users)
	assert.Equal(t, err, nil)
}

func TestWriteYaml2(t *testing.T) {
	users := map[string]map[string]string{"data": {"sheets.properties": "connector.name=gsheets\nconnector.name=gsheets"}}
	err := WriteYaml("../testFiles/test2.yaml", users)
	assert.Equal(t, err, nil)
}
