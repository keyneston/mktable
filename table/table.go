package table

import (
	"bufio"
	"bytes"
	"io"
	"regexp"
)

type NewLine byte

const (
	NewLineUnix NewLine = '\n'
)

type Table struct {
	data    [][]string
	sep     *regexp.Regexp
	newLine NewLine
}

func NewTable(sep *regexp.Regexp) *Table {
	return &Table{
		sep:     sep,
		newLine: NewLineUnix,
	}
}

func (t *Table) Read(r io.Reader) error {
	reader := bufio.NewReader(r)
	for {
		line, err := reader.ReadSlice(byte(t.newLine))
		if err != nil {
			return err
		}
		if line == nil {
			return nil
		}

		line = bytes.TrimSpace(line)
		if len(line) < 1 {
			continue
		}

		// Handle line here
		res := t.sep.Split(string(line), -1)
		t.data = append(t.data, res)
	}
}
