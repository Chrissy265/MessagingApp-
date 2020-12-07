package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

var EnvironmentConfiguration Config

type Config struct {
	Connection struct {
		MySQL string
		Redis string
	}
}

func InitializeConfiguration(filename string) {

	readFile(filename)
	fmt.Printf("%+v", EnvironmentConfiguration)
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(filename string) {
	filename, _ = filepath.Abs("configs/" + filename + ".yml")
	yamlFile, err := ioutil.ReadFile(filename)
	err = yaml.Unmarshal(yamlFile, &EnvironmentConfiguration)
	if err != nil {
		processError(err)
	}
}
