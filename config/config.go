package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type AppConfig struct {
	Es    string
	Mysql string
}

func GetConfig() AppConfig {
	var config = AppConfig{}
	file, err := ioutil.ReadFile("./config/app.yaml")
	if err != nil {
		log.Fatal("fail to read file:", err)
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("fail to yaml unmarshal:", err)
	}
	return config
}
