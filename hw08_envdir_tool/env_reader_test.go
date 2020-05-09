package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Check somple case
func TestReadDir(t *testing.T) {
	baseDir := "testdata/env"
	baseNumOfFiles := 4

	t.Run("should correctrly read env vars", func(t *testing.T) {
		envMap, err := ReadDir(baseDir)
		assert.Equal(t, nil, err)
		assert.Equal(t, baseNumOfFiles, len(envMap))
	})

	t.Run("should skip file with invalid name", func(t *testing.T) {
		invalidFileName := "test="
		tmpFile, _ := ioutil.TempFile(baseDir, invalidFileName)
		defer os.Remove(tmpFile.Name())
		envMap, err := ReadDir(baseDir)
		assert.Equal(t, nil, err)
		assert.Equal(t, baseNumOfFiles, len(envMap))
	})

	t.Run("should collect env vars from subdirectory", func(t *testing.T) {
		tmpDirPath := filepath.Join(baseDir, "tmpDir")
		tmpFileName := "tmpFile"
		os.Mkdir(tmpDirPath, os.ModePerm)
		defer os.Remove(tmpDirPath)
		tmpFile, _ := ioutil.TempFile(tmpDirPath, tmpFileName)
		defer os.Remove(tmpFile.Name())
		ioutil.WriteFile(tmpFile.Name(), []byte("test_data"), os.ModePerm)
		envMap, err := ReadDir(baseDir)
		assert.Equal(t, nil, err)
		assert.Equal(t, baseNumOfFiles+1, len(envMap))
	})

	t.Run("should collect empty env string from empty file", func(t *testing.T) {
		fileName := "test"
		tmpFile, _ := ioutil.TempFile(baseDir, fileName)
		defer os.Remove(tmpFile.Name())
		envMap, err := ReadDir(baseDir)
		assert.Equal(t, nil, err)
		assert.Equal(t, baseNumOfFiles+1, len(envMap))
		assert.Equal(t, "", envMap[filepath.Base(tmpFile.Name())])
	})

	t.Run("should collect only first line for env var from file", func(t *testing.T) {
		fileName := "test"
		tmpFile, _ := ioutil.TempFile(baseDir, fileName)
		defer os.Remove(tmpFile.Name())
		ioutil.WriteFile(tmpFile.Name(), []byte("first line\nsecond line"), os.ModePerm)
		envMap, err := ReadDir(baseDir)
		assert.Equal(t, nil, err)
		assert.Equal(t, baseNumOfFiles+1, len(envMap))
		assert.Equal(t, "first line", envMap[filepath.Base(tmpFile.Name())])
	})

	t.Run("should remove whitespace characters from env var", func(t *testing.T) {
		fileName := "test"
		tmpFile, _ := ioutil.TempFile(baseDir, fileName)
		defer os.Remove(tmpFile.Name())
		ioutil.WriteFile(tmpFile.Name(), []byte("first line\t \t \nsecond line"), os.ModePerm)
		envMap, err := ReadDir(baseDir)
		assert.Equal(t, nil, err)
		assert.Equal(t, baseNumOfFiles+1, len(envMap))
		assert.Equal(t, "first line", envMap[filepath.Base(tmpFile.Name())])
	})
}
