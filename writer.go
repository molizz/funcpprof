package funcpprof

import (
	"os"
	"strings"
)

type eachWriter interface {
	Each(func(string))
}

type Write struct {
	w eachWriter
}

func NewWrite(s eachWriter) *Write {
	return &Write{w: s}
}

func (w *Write) Flush(path string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	var buf strings.Builder

	w.w.Each(func(s string) {
		buf.WriteString(s)
		buf.Write([]byte("\r\n"))
	})

	_, err = f.WriteString(buf.String())
	return err
}
