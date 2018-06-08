package config

import (
	"errors"
	"os"
)

//Config is the base type that config library uses to save and abstract usage for the user
type Config struct {
	params Parameters
	Get    func(string) string
	create func(string) error
}

//Parameters contains all parameters necessary to
type Parameters struct {
	workdir        string
	name           string
	ignoreExisting bool
	permissions    os.FileMode
}

//NewConfig implements the Initilization of a new Config File or object
func NewConfig(workdir string, name string, ignoreExisting bool, permissions os.FileMode) (file Config, conErr error) {
	param := Parameters{
		workdir:        workdir,
		name:           name + ".yml",
		ignoreExisting: ignoreExisting,
		permissions:    permissions,
	}

	file = Config{
		params: param,
	}

	// This is commented out until i fix the issue occuring on first init with empty field
	// if _, err := os.Stat(param.workdir + param.name); err != nil {
	// 	_, err := os.Create(param.workdir + param.name)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	if param.ignoreExisting == false {
		conErr = errors.New("File already exists")
		return
	}
	conErr = nil
	return
}
