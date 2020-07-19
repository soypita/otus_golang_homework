package schedulercfg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicConfig(t *testing.T) {
	t.Run(`should successfully parse valid config`, func(t *testing.T) {
		validConfigPath := "../../../tests/testdata/configs/scheduler_valid_config.yml"
		config, err := NewConfig(validConfigPath)
		assert.NoError(t, err)
		assert.NotNil(t, config)
		assert.Equal(t, config.EventAPI.Host, "localhost")
		assert.Equal(t, config.EventAPI.GrpcPort, "8090")
		assert.Equal(t, config.Log.Path, "/app/logs/logfile.log")
		assert.Equal(t, config.Log.Level, "INFO")
		assert.Equal(t, config.Schedule.Notify, int64(3600))
		assert.Equal(t, config.Schedule.Clean, int64(36000))
		assert.Equal(t, config.AMPQ.URI, "amqp://guest:guest@localhost:5672/")
		assert.Equal(t, config.AMPQ.QueueName, "events")
		assert.Equal(t, config.AMPQ.ExchangeName, "ev_exchange")
		assert.Equal(t, config.AMPQ.ExchangeType, "direct")
	})
}
