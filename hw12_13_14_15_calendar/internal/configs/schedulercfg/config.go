package schedulercfg

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	EventAPI struct {
		Host     string `yaml:"host"`
		GrpcPort string `yaml:"grpc_port"`
	} `yaml:"event_api"`

	Log struct {
		Path  string `yaml:"path"`
		Level string `yaml:"level"`
	} `yaml:"log"`

	Schedule struct {
		Notify int64 `yaml:"notify"`
		Clean  int64 `yaml:"clean"`
	} `yaml:"schedule"`

	AMPQ struct {
		URI          string `yaml:"uri"`
		QueueName    string `yaml:"queue_name"`
		ExchangeName string `yaml:"exchange_name"`
		ExchangeType string `yaml:"exchange_type"`
	} `yaml:"ampq"`
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
