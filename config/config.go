package config

import (
	"os"
	"encoding/json"
)

type Configuration struct {
	Host           string
	Port           string
	SmtpServer     string
	SmtpServerPort int
	Username       string
	Password       string
}

var Config Configuration

func Init(env string) Configuration {

	dir, _ := os.Getwd()
	conf, err := os.Open(dir + "/config/" + env + ".json")
	defer conf.Close()

	if err != nil {
		panic("Configuration file can not opened")
	}

	err = json.NewDecoder(conf).Decode(&Config)
	if err != nil {
		panic("Configuration can not initial")
	}

	return Config
}
