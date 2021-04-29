package table

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

type NewLine byte

const (
	NewLineUnix    NewLine = '\n'
	minHeaderWidth         = 3
)

type Table struct {
	TableConfig

	data        [][]string
	columnCount int
	columnChars []int
}

func NewTable(config TableConfig) *Table {
	if config.NewLine == 0x0 {
		config.NewLine = NewLineUnix
	}

	return &Table{
		TableConfig: config,
	}
}

func (t *Table) readFormatMK(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		start := 0
		for i := range data {
			switch data[i] {
			case '\n':
				token := bytes.TrimSpace(data[start:i])
				if len(token) > 0 {
					return i, token, nil
				}

				return i + 1, []byte{'\n'}, nil
			case '|':
				if i > 0 && data[i-1] == '\\' {
					continue
				}
				if i == 0 {
					start = i + 1
					continue
				}

				token := bytes.TrimSpace(data[start:i])
				return i + 1, token, nil
			}
		}

		return 0, nil, nil
	})

	current := []string{}
	alignments := map[int]Alignment{}
	isHeader := false
	column := 0

	for scanner.Scan() {
		token := scanner.Text()

		if token == "\n" {
			if isHeader {
				t.Alignments = alignments
			} else {
				t.data = append(t.data, current)
			}

			alignments = map[int]Alignment{}
			column = 0
			current = nil
			isHeader = false
			continue
		}

		// Check if it is likely part of a header row, by removing all header
		// row chars and seeing if we have nothing left
		if len(strings.Trim(token, ":-")) == 0 {
			isHeader = true
			alignments[column] = parseAlignmentHeader(token)
		}

		current = append(current, token)

		column++
	}
	if err := scanner.Err(); err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	return nil
}

func (t *Table) Read(r io.Reader) error {
	switch t.Format {
	case FormatMK:
		return t.readFormatMK(r)
	case FormatRE:
		return t.readFormatRE(r)
	default:
		return fmt.Errorf("Unable to read format: %q", string(t.Format))
	}
}

func (t *Table) readFormatRE(r io.Reader) error {
	reader := bufio.NewReader(r)
	for {
		line, err := reader.ReadSlice(byte(t.NewLine))
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		if line == nil {
			return nil
		}

		line = bytes.TrimSpace(line)
		if len(line) < 1 {
			continue
		}

		res := t.Seperator.Split(string(line), -1)
		for i := range res {
			res[i] = prepareContent(res[i])
		}
		t.data = append(t.data, res)

	}
}

// findColumnCount finds the line with the most number of columns, and figures out
// how wide each column needs to be.
func (t *Table) findColumnCount() {
	for _, row := range t.data {
		if len(row) > t.columnCount {
			t.columnCount = len(row)
		}

	}

	t.columnChars = make([]int, t.columnCount)
	for _, row := range t.data {
		for i := range row {
			l := len(row[i])
			if l > t.columnChars[i] {
				t.columnChars[i] = l
			}
		}
	}
}

func (t *Table) Write(w io.Writer) (int, error) {
	t.findColumnCount()

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
	for i := 0; i < t.columnCount; i++ {
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

	_, err := fmt.Fprintf(w, string(t.NewLine))
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
	for i := 0; i < t.columnCount; i++ {
		alignment := t.Alignments[i] // Get the alignment, if not set it will return AlignDefault
		breaks = append(breaks, alignment.header(t.getColumnWidth(i)))
	}

	return breaks
}
