package config

import (
	"errors"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

//Get implements a simple getter for config files
func Get(config Config, key string) (value string, conErr error) {
	f, err := ioutil.ReadFile(config.params.workdir + config.params.name)
	if err != nil {
		conErr = errors.New("Config failed to open")
		return
	}
	var yml interface{}

	err = yaml.Unmarshal([]byte(f), &yml)
	if err != nil {
		conErr = errors.New("Config file failed to read")
		return
	}

	//Asserting that the type of the read in file will be a string map and equally that the read in value is also a string
	assertYml := yml.(map[string]interface{})
	value = assertYml[key].(string)
	conErr = nil
	return
}
