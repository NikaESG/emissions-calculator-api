package service

import (
	"fmt"
	"strings"
	"trino.com/trino-connectors/util"
	"trino.com/trino-connectors/util/file"
	"trino.com/trino-connectors/util/log"
)

func AddConfigmap(ns string, catalogId string) error {
	_, err := util.RunCommand("/bin/sh", "../shell/RunConfimap.sh", ns, catalogId)
	if err != nil {
		fmt.Printf("[AddConfigmap] [util.RunCommand] error: %s", err)
		log.Logger().Errorf("[AddConfigmap] [util.RunCommand] error: %s", err)
		return err
	}
	return nil
}

func GetConfigmap(ns string, catalogId string) (map[string]string, error) {
	configMap := make(map[string]string)
	out, err := util.RunCommand("/bin/sh", "../shell/GetConfimap.sh", ns, catalogId)
	if err != nil {
		fmt.Printf("[GetConfigmap] [util.RunCommand] error: %s", err)
		log.Logger().Errorf("[GetConfigmap] [util.RunCommand] error: %s", err)
		return configMap, err
	}

	path := fmt.Sprintf("./testFiles/%v", catalogId)
	err = file.WriteFile(path, out)
	if err != nil {
		fmt.Printf("[GetConfigmap] [file.WriteFile] error: %s", err)
		log.Logger().Errorf("[GetConfigmap] [file.WriteFile] error: %s", err)
		return configMap, err
	}

	data := make(map[interface{}]interface{})
	err = util.ParseYaml(path, &data)
	if err != nil {
		fmt.Printf("[GetConfigmap] [util.ParseYaml] error: %s", err)
		log.Logger().Errorf("[GetConfigmap] [util.ParseYaml] error: %s", err)
		return configMap, err
	}

	config := fmt.Sprint(data["data"])
	config = config[4 : len(config)-1]
	detailConfig := strings.Split(config, "\n")
	for _, s := range detailConfig {
		d := strings.Split(s, "=")
		configMap[d[0]] = d[1]
	}
	return configMap, nil
}
