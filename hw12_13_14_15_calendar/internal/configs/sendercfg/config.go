package sendercfg

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Log struct {
		Path  string `yaml:"path"`
		Level string `yaml:"level"`
	} `yaml:"log"`

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
