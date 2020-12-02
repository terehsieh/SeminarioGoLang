package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// DbConfig ...
type DbConfig struct {
	Driver string `yaml:"driver"`
}

// Config ...
type Config struct {
	DB      DbConfig `yaml:"db"`
	Version string   `yaml:"version"`
}

// LoadConfig ...
func LoadConfig(filename string) (*Config, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		//	panic(err)
		return nil, err
	}

	var c = &Config{} // puntero a una estructura config crear y devolver el puntero
	err = yaml.Unmarshal(file, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

//github.com/terehsieh/SeminarioGoLang.git
