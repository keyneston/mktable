package table

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
)

type NewLine byte

const (
	NewLineUnix NewLine = '\n'
)

type Table struct {
	data    [][]string
	sep     *regexp.Regexp
	newLine NewLine

	maxPadding int

	rowCount    int
	columnChars []int
}

func NewTable(sep *regexp.Regexp) *Table {
	return &Table{
		sep:        sep,
		newLine:    NewLineUnix,
		maxPadding: -1,
	}
}

func (t *Table) SetMaxPadding(i int) {
	t.maxPadding = i
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
		log.Printf("Seperator: %s", t.sep)
		res := t.sep.Split(string(line), -1)
		t.data = append(t.data, res)
	}
}

// longestLine finds the line with the most number of columns, and figures out
// how wide each column needs to be.
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
		for i := range row {
			row[i] = strings.TrimSpace(row[i])
			l := len(row[i])
			if l > t.columnChars[i] {
				t.columnChars[i] = l
			}
		}
	}

	return longest
}

func (t *Table) Write(w io.Writer) (int, error) {
	t.longestLine()
	sw := newSumWriter(w)
	w = sw // alias to prevent any accidental mis-writes

	for _, row := range t.data {
		fmt.Fprintf(w, "|")
		for i := 0; i < t.rowCount; i++ {
			column := ""
			if i < len(row) {
				column = row[i]
			}

			padding := t.genPadding(t.columnChars[i], len(column))
			_, err := fmt.Fprintf(sw, " %s%s |", column, padding)
			if err != nil {
				return sw.Sum, err
			}
		}

		_, err := fmt.Fprintf(sw, string(t.newLine))
		if err != nil {
			return sw.Sum, err
		}
	}

	return sw.Sum, nil
}

func (t *Table) genPadding(spaceSize, content int) string {
	padding := spaceSize - content
	if t.maxPadding > 0 && padding > t.maxPadding {
		padding = t.maxPadding
	}

	return strings.Repeat(" ", padding)
}
