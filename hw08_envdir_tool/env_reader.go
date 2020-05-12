package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]string

const forbiddenSymbol = "="

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	envDir, err := ioutil.ReadDir(dir)
	resEnvList := Environment{}
	if err != nil {
		return nil, fmt.Errorf("error while try to open dir: %w", err)
	}

	var relativePath string
	for _, val := range envDir {
		relativePath = filepath.Join(dir, val.Name())
		if val.IsDir() {
			innerRes, err := ReadDir(relativePath)
			if err != nil {
				return nil, err
			}
			for envName, envVal := range innerRes {
				resEnvList[envName] = envVal
			}
		} else {
			if strings.Contains(val.Name(), forbiddenSymbol) {
				continue
			}
			envStr, err := processFile(relativePath)
			if err != nil {
				return nil, err
			}
			resEnvList[val.Name()] = envStr
		}
	}
	return resEnvList, err
}

func processFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("error while open file: %w", err)
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf("error while get stats of file: %w", err)
	}
	if fileStat.Size() == 0 {
		return "", nil
	}
	r := bufio.NewReader(file)
	envParam, _, err := r.ReadLine()

	if err != nil {
		return "", fmt.Errorf("error while reading file: %w", err)
	}

	resEnvStr := strings.TrimRight(string(bytes.ReplaceAll(envParam, []byte("\x00"), []byte("\n"))), "\t \n")
	return resEnvStr, nil
}
