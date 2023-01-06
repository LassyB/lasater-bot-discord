package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var (
	Token  string
	config *configStruct
)

type configStruct struct {
	Token string `json:"token"`
}

func ReadConfig() error {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
		return err
	}
	Token = config.Token

	return nil
}
