package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

var Cfg *Config

type Config struct {
	Server string
	Port string
	User string
	Password string
	Channel string
	EventFrequencyInMs int64
	Usetls bool
	Debug bool
	Keysbinding []*Keysbinding
}

type Keysbinding struct {
	Word string
	Event string
}

func LoadConfig(filepath string) error {
	source, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(source, &Cfg)
	if err != nil {
		return err
	}
	return nil
}