package configs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicConfig(t *testing.T) {
	t.Run(`should successfully parse valid config`, func(t *testing.T) {
		validConfigPath := "../../tests/testdata/configs/valid_config.yml"
		config, err := NewConfig(validConfigPath)
		assert.NoError(t, err)
		assert.NotNil(t, config)
		assert.Equal(t, config.Host, "localhost")
		assert.Equal(t, config.Port, "8080")
		assert.Equal(t, config.Log.Path, "/app/logs/logfile.log")
		assert.Equal(t, config.Log.Level, "INFO")
		assert.Equal(t, config.Database.InMemory, false)
		assert.Equal(t, config.Database.DSN, "host=test port=test user=test password=test dbname=test sslmode=disable")
	})
	t.Run(`should successfully parse valid config with additional fields`, func(t *testing.T) {
		validConfigPath := "../../tests/testdata/configs/valid_config_additional.yml"
		config, err := NewConfig(validConfigPath)
		assert.NoError(t, err)
		assert.NotNil(t, config)
		assert.Equal(t, config.Host, "localhost")
		assert.Equal(t, config.Port, "8080")
		assert.Equal(t, config.Log.Path, "/app/logs/logfile.log")
		assert.Equal(t, config.Log.Level, "INFO")
		assert.Equal(t, config.Database.InMemory, false)
		assert.Equal(t, config.Database.DSN, "host=test port=test user=test password=test dbname=test sslmode=disable")
	})
	t.Run(`should return error for invalid path`, func(t *testing.T) {
		invalidPath := "invalidDir/configs"
		config, err := NewConfig(invalidPath)
		assert.Error(t, err)
		assert.Nil(t, config)
	})
	t.Run(`should successfully parse json representation of yml config`, func(t *testing.T) {
		invalidConfig := "../../tests/testdata/configs/config.json"
		config, err := NewConfig(invalidConfig)
		assert.NoError(t, err)
		assert.NotNil(t, config)
		assert.Equal(t, config.Host, "localhost")
		assert.Equal(t, config.Port, "8080")
		assert.Equal(t, config.Log.Path, "/app/logs/logfile.log")
		assert.Equal(t, config.Log.Level, "INFO")
		assert.Equal(t, config.Database.InMemory, false)
		assert.Equal(t, config.Database.DSN, "host=test port=test user=test password=test dbname=test sslmode=disable")
	})
}
