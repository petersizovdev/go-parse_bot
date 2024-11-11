package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)


const configFile = "data/config.yaml"

type Config struct {
	Tocken string `yaml:"tocken"`
}

type Service struct {
	config Config
}

func New() (*Service, error){
	s:=&Service{}

	rawYAML, err := os.ReadFile(configFile)
	if err != nil {
		return nil, errors.Wrap(err, "reading config file")
	}

	err = yaml.Unmarshal(rawYAML, &s.config)
	if err != nil {
		return nil, errors.Wrap(err, "parsing yaml")
	}

	return s, nil
}

func (s *Service) Tocken() string {
	return s.config.Tocken
}