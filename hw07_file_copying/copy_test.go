package main

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	t.Run("try to copy dir", func(t *testing.T) {
		err := Copy("testdata", "/tmp", 0, 0)
		require.Equal(t, ErrIsDir, err)
	})

	t.Run("try to copy not existing file", func(t *testing.T) {
		err := Copy("asdasd.txt", "/tmp", 0, 0)
		require.Equal(t, ErrUnsupportedFile, err)
	})

	t.Run("try to copy file with exceed offset", func(t *testing.T) {
		fileFrom := "testdata/input.txt"
		fileStat, _ := os.Stat(fileFrom)
		fileTo := "testdata/tmp.txt"
		err := Copy(fileFrom, fileTo, fileStat.Size()+100, 0)
		require.Equal(t, ErrOffsetExceedsFileSize, err)
	})

	t.Run("try to copy file without read permission", func(t *testing.T) {
		tmpFile, _ := ioutil.TempFile("testdata", "test")
		os.Chmod(tmpFile.Name(), 0700)
		defer os.Remove(tmpFile.Name())
		err := Copy(tmpFile.Name(), "/tmp", 0, 0)
		require.Equal(t, ErrReadPermissions, err)
	})

	t.Run("try to copy file with unknown system size", func(t *testing.T) {
		fileName := "/dev/urandom"
		err := Copy(fileName, "testdata", 0, 0)
		require.Equal(t, ErrUnsupportedFile, err)
	})

	t.Run("try to copy from file to directory", func(t *testing.T) {
		fileName := "testdata/input.txt"
		testDir := "testdata/testDir"
		os.Mkdir(testDir, os.ModePerm)
		defer os.Remove(testDir)
		err := Copy(fileName, testDir, 0, 0)
		log.Println("Error is ", err)
		require.NotNil(t, err)
	})
}
