package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"trino.com/trino-connectors/util/log"
)

func WriteFile(path string, data string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err = f.WriteString(string(data))
	if err != nil {
		fmt.Printf("[WriteYaml] [ioutil.WriteFile] error: %s", err)
		log.Logger().Errorf("[WriteYaml] [ioutil.WriteFile] error: %s", err)
		return err
	}
	return nil
}

func ReadFile(path string) (string, error) {
	yfile, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Printf("[ParseYaml] [ioutil.ReadFile] error: %s", err)
		log.Logger().Errorf("[ParseYaml] [ioutil.ReadFile] error: %s", err)
		return "", err
	}

	return string(yfile), nil
}
