package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrIsDir                 = errors.New("source is directory")
	ErrReadPermissions       = errors.New("no read permissions for file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

const UserReadPermission = 4

func Copy(fromPath string, toPath string, offset, limit int64) error {
	fileFrom, err := os.Open(fromPath)

	if err != nil {
		return ErrUnsupportedFile
	}

	fileStat, err := fileFrom.Stat()
	if err != nil {
		return err
	}

	err = validateSourceFile(fileStat, offset)

	if err != nil {
		return err
	}

	fileSize := fileStat.Size()

	if (limit > fileSize-offset) || limit == 0 {
		limit = fileSize - offset
	}

	if offset != 0 {
		_, err := fileFrom.Seek(offset, 0)
		if err != nil {
			return fmt.Errorf("error while set offset: %w", err)
		}
	}

	fileTo, err := os.Create(toPath)

	if err != nil {
		return fmt.Errorf("error while creating new file: %w", err)
	}

	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(fileFrom)
	_, err = io.CopyN(fileTo, barReader, limit)
	if err != nil {
		return fmt.Errorf("error while copy data: %w", err)
	}
	bar.Finish()

	return nil
}

func validateSourceFile(fileStat os.FileInfo, offset int64) error {
	if fileStat.IsDir() {
		return ErrIsDir
	}

	mode := fileStat.Mode()

	if mode&(1<<2) < UserReadPermission {
		return ErrReadPermissions
	}

	fileSize := fileStat.Size()

	if fileSize == 0 {
		return ErrUnsupportedFile
	}

	if offset > fileSize {
		return ErrOffsetExceedsFileSize
	}

	return nil
}
