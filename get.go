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
		conErr = err
		return
	}
	var yml map[string]string
	err = yaml.Unmarshal([]byte(f), &yml)
	if err != nil {
		conErr = err
		return
	}
	_, exists := yml[key]
	if exists != true {
		conErr = errors.New("Config contains no key called " + key)
		return
	}
	value = yml[key]

	conErr = nil
	return
}
