package main

import "io"

type countingReader struct {
	reader io.Reader
	count  int
}

func (c *countingReader) Read(p []byte) (int, error) {
	n, err := c.reader.Read(p)
	c.count += n
	return n, err
}

func (c *countingReader) BytesRead() int {
	return c.count
}
