package util

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"trino.com/trino-connectors/util/log"
)

func ParseYaml(path string, data interface{}) error {
	yfile, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Printf("[ParseYaml] [ioutil.ReadFile] error: %s", err)
		log.Logger().Errorf("[ParseYaml] [ioutil.ReadFile] error: %s", err)
		return err
	}

	err = yaml.Unmarshal(yfile, data)
	if err != nil {
		fmt.Printf("[ParseYaml] [yaml.Unmarshal] error: %s", err)
		log.Logger().Errorf("[ParseYaml] [yaml.Unmarshal] error: %s", err)
		return err
	}
	return nil
}

func WriteYaml(path string, data interface{}) error {
	d, err := yaml.Marshal(&data)

	if err != nil {
		fmt.Printf("[WriteYaml] [yaml.Marshal] error: %s", err)
		log.Logger().Errorf("[WriteYaml] [yaml.Marshal] error: %s", err)
		return err
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err = f.WriteString(string(d))
	if err != nil {
		fmt.Printf("[WriteYaml] [ioutil.WriteFile] error: %s", err)
		log.Logger().Errorf("[WriteYaml] [ioutil.WriteFile] error: %s", err)
		return err
	}
	return nil
}
