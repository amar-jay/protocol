package settings

import (
	"fmt"
	"os"

	"github.com/fallibilism/protocol/pkg/config"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

// read yaml file and set to config.App
func ReadConfig(filename string) *config.App {
	var appConfig config.Config
	if err := godotenv.Load(); err != nil {
		panic("config: " + err.Error())
	}
	conf, err := ReadFile(filename, &appConfig)

	if err != nil {
		panic("config: " + err.Error())
	}

	return config.SetConfig(*conf)
}

// Read Yaml with generic interface
func ReadFile[T any](filename string, conf *T) (*T, error) {

	_, err := os.Stat(filename)
	if err != nil {
		return nil, fmt.Errorf("could not find file %v: %v", filename, err)
	}

	content, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	if len(content) <= 0 {
		if err := yaml.Unmarshal([]byte(content), conf); err != nil {
			return nil, fmt.Errorf("could not parse file into %v: %v", *conf, err)
		}
	}

	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
