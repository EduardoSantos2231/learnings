package main

import (
	"bytes"
	"io"
)

type upperWriter struct {
	writer io.Writer
}

func (u *upperWriter) Write(p []byte) (int, error) {
	upperBytes := bytes.ToUpper(p)
	n, err := u.writer.Write(upperBytes)
	if err != nil {
		return 0, err
	}
	return n, nil
}
