package DAO

import (
	"github.com/garyburd/redigo/redis"
	"trino.com/trino-connectors/util/log"
)

var Redis redis.Conn

func InitRedisConfig() error {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Logger().Errorf("[InitRedisConfig] conn redis failed %v", err)
		return err
	}

	Redis = c
	return nil
}

func SetValue(commandName string, args ...interface{}) (interface{}, error) {
	reply, err := Redis.Do(commandName, args)
	if err != nil {
		log.Logger().Errorf("[SetValue] redis failed %v", err)
		return reply, err
	}
	return reply, nil
}

func GetValue(commandName string, args ...interface{}) ([]string, error) {
	result, err := redis.Strings(Redis.Do(commandName, args))
	if err != nil {
		log.Logger().Errorf("[GetValue] redis failed %v", err)
		return []string{}, err
	}
	return result, nil
}
