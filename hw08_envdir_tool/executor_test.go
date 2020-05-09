package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	notFoundErrorCode = 127
)

func TestRunCmd(t *testing.T) {
	t.Run("should return non zero code for empty cmd slice", func(t *testing.T) {
		var tmpCmd []string
		var tmpEnv Environment
		code := RunCmd(tmpCmd, tmpEnv)
		assert.Equal(t, defaultErrorCode, code)
		code = RunCmd(nil, tmpEnv)
		assert.Equal(t, defaultErrorCode, code)
	})
	t.Run("should succesfully run cmd", func(t *testing.T) {
		tmpCmd := []string{"echo"}
		code := RunCmd(tmpCmd, nil)
		assert.Equal(t, defaultSuccessCode, code)
	})
	t.Run("should return code from command", func(t *testing.T) {
		tmpCmd := []string{"/bin/bash", "wrong_command_name"}
		code := RunCmd(tmpCmd, nil)
		assert.Equal(t, notFoundErrorCode, code)
	})
	t.Run("should succesfully set env variable during execution", func(t *testing.T) {
		tmpCmd := []string{"echo"}
		tmpEnv := Environment{"TEST": "test"}
		code := RunCmd(tmpCmd, tmpEnv)
		assert.Equal(t, defaultSuccessCode, code)
		assert.Equal(t, "test", os.ExpandEnv("$TEST"))
	})
	t.Run("should unset env var if empty", func(t *testing.T) {
		os.Setenv("TEST", "test")
		tmpCmd := []string{"echo"}
		tmpEnv := Environment{"TEST": ""}
		code := RunCmd(tmpCmd, tmpEnv)
		assert.Equal(t, defaultSuccessCode, code)
		remEnv, isPresent := os.LookupEnv("$TEST")
		assert.Equal(t, false, isPresent)
		assert.Equal(t, "", remEnv)
	})
}
