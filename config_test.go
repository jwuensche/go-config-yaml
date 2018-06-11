package config_test

import (
	"fmt"
	"os"
	"testing"

	config "github.com/jwuensche/go-config-yaml"
)

//TestInit tests
func TestInit(t *testing.T) {
	os.MkdirAll("config", 0722)
	os.Remove("config/test.yml")
	configF, err := config.NewConfig("config/", "test", false, 0722)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	configF.Set("foo", "bar")
	_, _ = config.NewConfig("config/", "test", false, 0722)
	if _, err = os.Stat("config/test.yml"); err != nil {
		fmt.Println("Existing File ignored")
		t.FailNow()
	}
}

func TestSet(t *testing.T) {
	confFile, err := config.NewConfig("config", "test", true, 0722)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	confFile.Set("stuff", "something")
}

func TestGet(t *testing.T) {
	configFile, err := config.NewConfig("config", "test", false, 0722)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	value, err := configFile.Get("stuff")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	if value != "something" {
		fmt.Println("Wrong value entered in key")
		t.FailNow()
	}
}

func TestDelete(t *testing.T) {
	configFile, err := config.NewConfig("config", "test", false, 0722)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	if err = configFile.DeleteConfig(); err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}
