package config

import (
	"errors"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type configContent struct {
	content map[string]string
}

//Set implements a basic set functionanality with one key and an according value
func (config Config) Set(key string, value string) (conErr error) {
	//Handling of simple setting first value or overwrite full config
	var yml map[string]string
	config.Create()

	rawRead, err := ioutil.ReadFile(config.params.workdir + config.params.name)
	if err != nil {
		conErr = err
		return
	}
	yml = map[string]string{}
	yaml.Unmarshal([]byte(rawRead), &yml)

	file, err := os.OpenFile(config.params.workdir+config.params.name, os.O_WRONLY, config.params.permissions)
	if err != nil {
		conErr = err
		return
	}
	yml[key] = value

	ymlByte, err := yaml.Marshal(yml)
	if err != nil {
		conErr = err
		return
	}
	file.Write(ymlByte)
	return
}

//SetByte for setting raw byte slice value into config
func (config Config) SetByte(key string, value []byte) (conErr error) {
	config.Create()

	rawRead, err := ioutil.ReadFile(config.params.workdir + config.params.name)
	if err != nil {
		conErr = errors.New("Reading in existing config failed")
		return
	}
	yml := make(map[string]interface{})
	yaml.Unmarshal([]byte(rawRead), &yml)

	//From here on arches should connect again otherwise i have to tidythis up again
	file, err := os.OpenFile(config.params.workdir+config.params.name, os.O_WRONLY, config.params.permissions)
	if err != nil {
		conErr = errors.New("Opening in File failed")
		return
	}
	yml[key] = value

	ymlByte, err := yaml.Marshal(yml)
	if err != nil {
		conErr = errors.New("Encoding of config failed")
		return
	}
	file.Write(ymlByte)
	return

}
