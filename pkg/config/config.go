package config

import (
	"encoding/json"
	"fmt"
	"github.com/lasher23/cheetah/pkg/consoleutil"
	"github.com/lasher23/cheetah/pkg/ioutil"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"strings"
)

type Config struct {
	Shortcuts []Shortcut
}

type Shortcut struct {
	Name          string
	ExecutionPath string
	Command       string
}

const configFileName string = "settings.json"
const configFileLocation string = ".cheetah"

var NoConfigError error = errors.New("No configuration file found")

func CreateConfigFile() error {
	if ask("Should a config file be created? ") {
		dir, e := homedir.Dir()
		if e != nil {
			return errors.Wrap(e, "Home directory could not be found")
		}
		e = ioutil.CreateEmptyFileWithAllDirectories(dir+"/"+configFileLocation, configFileName)
		if e != nil {
			return errors.Wrap(e, "Error creating config file")
		}
		fmt.Println("Config file created")
	}
	return nil
}

func GetConfig() (Config, error) {
	content, e := readConfigFile()
	config := Config{}
	if e != nil {
		return config, NoConfigError
	}
	e = json.Unmarshal([]byte(content), &config)
	if e != nil {
		return config, errors.Wrap(e, "Config file has an invalid format")
	}
	return config, nil
}

func readConfigFile() (string, error) {
	dir, e := homedir.Dir()
	if e != nil {
		return "", errors.Wrap(e, "Home directory could not be found")
	}
	content, e := ioutil.ReadFile(dir + "/" + configFileLocation + "/" + configFileName)
	if e != nil {
		return "", errors.Wrap(e, "Error reading config file")
	}
	return content, nil
}

func ask(displayText string) bool {
	input := consoleutil.DisplayQuestion([]string{"y", "n"}, displayText)
	return strings.EqualFold(input, "y")

}
