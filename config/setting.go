package config

import (
	"encoding/json"
	"log"
	"os"
)

type Setting struct {
	Database struct {
		Host     string `json: "host"`
		User     string `json: "user"`
		Password string `json: "password"`
		Dbname   string `json: "Dbname"`
		Port     string `json: "port"`
	} `json: "database"`
}

func LoadSetting() (Setting, error) {
	file, err := os.Open("config/setting.json")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	var config Setting
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
	return config, err
}
