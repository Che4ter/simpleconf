package simpleconf

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// Loads configuration from path, creates empty config file if it not exists
func LoadConfig(configPath string, config interface{}) error {
	if !checkIfConfigExists(configPath) {
		if err := createEmptyConfig(configPath, config); err != nil {
			return err
		}
		return errors.New("Config file not found, initialized under " + configPath)
	} else {
		f, err := os.Open(configPath)
		if err != nil {
			return err
		}
		defer f.Close()

		decoder := yaml.NewDecoder(f)
		if err := decoder.Decode(config); err != nil {
			return err
		}
		return nil
	}
}

func checkIfConfigExists(configPath string) bool {
	info, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func createEmptyConfig(configPath string, config interface{}) error {
	file, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(configPath, file, 0664)

	return err
}
