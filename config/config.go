package config

import (
	"os"
	"encoding/json"
)

type Configuration struct {
	Host string
	Port string
}

func Init(env string) Configuration {

	var configuration Configuration

	dir, _ := os.Getwd()
	conf, err := os.Open(dir + "/config/" + env + ".json")
	defer conf.Close()

	if err != nil {
		panic("Configuration file can not opened")
	}

	err = json.NewDecoder(conf).Decode(&configuration)
	if err != nil {
		panic("Configuration can not initial")
	}

	return configuration
}
