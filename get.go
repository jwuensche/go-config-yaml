package config

import (
	"errors"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

//Get implements a simple getter for config files
func (config Config) Get(key string) (value string, conErr error) {
	f, err := ioutil.ReadFile(config.params.workdir + config.params.name)
	if err != nil {
		conErr = errors.New("Config failed to open")
		return
	}
	var yml map[string]string
	err = yaml.Unmarshal([]byte(f), &yml)
	if err != nil {
		conErr = errors.New("Config file failed to read")
		return
	}

	value = yml[key]

	conErr = nil
	return
}
