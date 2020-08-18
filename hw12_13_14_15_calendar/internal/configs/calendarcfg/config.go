package calendarcfg

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host     string `yaml:"host"`
	RestPort string `yaml:"rest_port"`
	GrpcPort string `yaml:"grpc_port"`
	Log      struct {
		Path  string `yaml:"path"`
		Level string `yaml:"level"`
	} `yaml:"log"`
	Database struct {
		InMemory      bool   `yaml:"in_memory"`
		DSN           string `yaml:"dsn"`
		MigrationsDir string `yaml:"migrations_dir"`
	} `yaml:"database"`
}

func NewConfig(file string) (*Config, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error while open config file %w", err)
	}
	config := Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error while parse configs %w", err)
	}
	return &config, nil
}
