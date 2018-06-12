package config

import (
	"errors"
	"os"
)

//Create creates a config based on the existence of current files and ignore flags
func (config Config) Create() (conErr error) {
	if _, err := os.Stat(config.params.workdir + config.params.name); err != nil {
		_, err := os.Create(config.params.workdir + config.params.name)
		if err != nil {
			conErr = errors.New("Error creating file")
			return
		}
	}

	conErr = nil
	return
}
