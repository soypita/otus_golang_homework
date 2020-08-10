package sendercfg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicConfig(t *testing.T) {
	t.Run(`should successfully parse valid config`, func(t *testing.T) {
		validConfigPath := "../../../tests/testdata/configs/sender_valid_config.yml"
		config, err := NewConfig(validConfigPath)
		assert.NoError(t, err)
		assert.NotNil(t, config)
		assert.Equal(t, config.Log.Path, "/app/logs/logfile.log")
		assert.Equal(t, config.Log.Level, "INFO")
		assert.Equal(t, config.AMPQ.URI, "amqp://guest:guest@localhost:5672/")
		assert.Equal(t, config.AMPQ.QueueName, "events")
		assert.Equal(t, config.AMPQ.ExchangeName, "ev_exchange")
		assert.Equal(t, config.AMPQ.ExchangeType, "direct")
	})
}
