package table

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type NewLine byte

const (
	NewLineUnix    NewLine = '\n'
	minHeaderWidth         = 3
)

type Table struct {
	data    [][]string
	sep     *regexp.Regexp
	newLine NewLine

	MaxPadding  int
	SkipHeaders bool

	rowCount    int
	columnChars []int
}

func NewTable(sep *regexp.Regexp) *Table {
	return &Table{
		sep:        sep,
		newLine:    NewLineUnix,
		MaxPadding: -1,
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

		res := t.sep.Split(string(line), -1)
		for i := range res {
			res[i] = prepareContent(res[i])
		}

		// Handle line here
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
	startingRow := 0

	if !t.SkipHeaders && len(t.data) >= 1 {
		startingRow += 1
		t.writeRow(sw, t.data[0])
		t.writeRow(sw, t.genHeaderBreaks())
	} else {
		// We need to write an empty header that way some markdown engines
		// still recognise it as a table.
		t.writeRow(sw, []string{})
		t.writeRow(sw, t.genHeaderBreaks())
	}

	for i := startingRow; i < len(t.data); i++ {
		if err := t.writeRow(sw, t.data[i]); err != nil {
			return sw.Sum, err
		}
	}

	return sw.Sum, nil
}

func (t *Table) writeRow(w io.Writer, row []string) error {
	fmt.Fprintf(w, "|")
	for i := 0; i < t.rowCount; i++ {
		column := ""
		if i < len(row) {
			column = row[i]
		}

		padding := t.genPadding(t.getColumnWidth(i), len(column))
		_, err := fmt.Fprintf(w, " %s%s |", column, padding)
		if err != nil {
			return err
		}
	}

	_, err := fmt.Fprintf(w, string(t.newLine))
	return err
}

func (t *Table) genPadding(spaceSize, content int) string {
	padding := spaceSize - content
	if t.MaxPadding > 0 && padding > t.MaxPadding {
		padding = t.MaxPadding
	}

	return strings.Repeat(" ", padding)
}

func (t *Table) getColumnWidth(i int) int {
	if i > len(t.columnChars) {
		return 0
	}
	width := t.columnChars[i]
	if !t.SkipHeaders && width < minHeaderWidth {
		return minHeaderWidth
	}

	return width
}

func (t *Table) genHeaderBreaks() []string {
	breaks := []string{}
	for i := 0; i < t.rowCount; i++ {
		breaks = append(breaks, strings.Repeat("-", t.getColumnWidth(i)))
	}

	return breaks
}
