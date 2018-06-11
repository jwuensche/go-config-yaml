package config

import (
	"errors"
	"os"
)

//DeleteConfig complete config and disregarding any saved configurations within
func (config Config) DeleteConfig() (conErr error) {
	err := os.Remove(config.params.workdir + config.params.name)
	if err != nil {
		conErr = errors.New("Deletion failed")
		return
	}
	conErr = nil
	return
}
