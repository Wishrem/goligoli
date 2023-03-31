package util

import (
	"io"
	"os"
)

func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	defer f.Close()

	n, err := f.Write(data)
	if err == nil && n != len(data) {
		err = io.ErrShortWrite
	}

	return err
}
