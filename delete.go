package config

import (
	"errors"
	"os"
)

//Removes complete config and disregarding any saved configurations within
func (config Config) DeleteConfig() (conErr error) {
	err := os.Remove(config.params.workdir + config.params.name)
	if err != nil {
		conErr = errors.New("Deletion failed")
	}
	conErr = nil
	return
}
