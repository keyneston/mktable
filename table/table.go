package table

import (
	"bufio"
	"bytes"
	"fmt"
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

	rowCount    int
	columnChars []int
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

func (t *Table) longestLine() int {
	longest := 0

	for _, row := range t.data {
		if len(row) > longest {
			longest = len(row)
		}
	}

	t.rowCount = longest
	t.columnChars = make([]int, t.rowCount)

	for _, row := range t.data {
		for i, column := range row {
			l := len(column)
			if l > t.columnChars[i] {
				t.columnChars[i] = l
			}
		}
	}

	return longest
}

func (t *Table) Write(w io.Writer) (int, error) {
	t.longestLine()

	charsWritten := 0

	for _, row := range t.data {
		fmt.Fprintf(w, "| ")
		for i := 0; i < t.rowCount; i++ {
			column := ""
			if i < len(row) {
				column = row[i]
			}

			paddingLen := t.columnChars[i] - len(column)
			padding := make([]byte, paddingLen)
			for j := range padding {
				padding[j] = ' '
			}

			count, err := fmt.Fprintf(w, " %s%s |", padding, column)
			charsWritten += count
			if err != nil {
				return charsWritten, err
			}
		}

		count, err := fmt.Fprintf(w, string(t.newLine))
		charsWritten += count
		if err != nil {
			return charsWritten, err
		}
	}

	return charsWritten, nil
}
