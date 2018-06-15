package config

import (
	"os"
)

//DeleteConfig complete config and disregarding any saved configurations within
func (config Config) DeleteConfig() (conErr error) {
	err := os.Remove(config.params.workdir + config.params.name)
	if err != nil {
		conErr = err
		return
	}
	conErr = nil
	return
}
