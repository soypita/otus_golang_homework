package main

import (
	"os"
	"os/exec"
)

const (
	defaultErrorCode   = 1
	defaultSuccessCode = 0
	minimalCmdLen      = 1
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	var resCmd *exec.Cmd
	cmdLen := len(cmd)
	switch {
	case cmdLen < minimalCmdLen:
		return defaultErrorCode
	case cmdLen == minimalCmdLen:
		resCmd = exec.Command(cmd[0]) //nolint:gosec
	default:
		resCmd = exec.Command(cmd[0], cmd[1:]...) //nolint:gosec
	}

	resCmd.Stdout = os.Stdout
	resCmd.Stdin = os.Stdin
	resCmd.Stderr = os.Stderr

	for envName, envVal := range env {
		if envVal == "" {
			err := os.Unsetenv(envName)
			if err != nil {
				return defaultErrorCode
			}
		}
		os.Setenv(envName, envVal)
	}
	resCmd.Env = os.Environ()
	if err := resCmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode()
		}
		return defaultErrorCode
	}
	return defaultSuccessCode
}
