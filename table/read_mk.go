package table

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strings"
)

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
