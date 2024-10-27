package config

import (
	"os"

	"github.com/go-logr/logr"
	"gopkg.in/yaml.v2"
)

type Reader interface {
	Read(path string) (Config, error)
}

type Config struct {
	X int `yaml:"x"`
}

type FileReader struct {
	logger logr.Logger
}

func NewFileReader(logger logr.Logger) Reader {
	return &FileReader{
		logger: logger,
	}
}

func (r *FileReader) Read(path string) (Config, error) {
	var config Config

	r.logger.V(1).Info("reading config file", "path", path)

	data, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	r.logger.V(1).Info("config loaded successfully", "x", config.X)
	return config, nil
}
