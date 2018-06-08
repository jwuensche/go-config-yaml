package config_test

import (
	"os"
	"testing"

	config "github.com/jwuensche/go-config-yaml"
)

//TestInit tests
func TestInit(t *testing.T) {
	os.MkdirAll("config", 0722)
	confFile, err := config.NewConfig("config/", "test", true, 0722)
	if err != nil {
		t.FailNow()
	}
	confFile.Set("foo", "bar")
	confFile.Set("stuff", "something")
}

//
// func TestSet(t *testing.T) {
// 	confFile, err := config.NewConfig("config/", "test", false, 0722)
// 	if err != nil {
// 		t.FailNow()
// 	}
//
// 	confFile.Set("stuff", "something")
// }
