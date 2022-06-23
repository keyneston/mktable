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
				// Ignore escaped pipe chars `|`
				if i > 0 && data[i-1] == '\\' {
					continue
				}
				// Ignore the first pipe char if it is the start of the line
				if i == 0 {
					start = i + 1
					continue
				}
				token := bytes.TrimSpace(data[start:i])
				if len(token) == 0 {
					token = []byte(" ")
				}

				return i + 1, token, nil
			}
		}

		trimmedData := bytes.TrimSpace(data)
		return 0, trimmedData, nil
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

	if len(current) > 0 {
		t.data = append(t.data, current)
	}

	return nil
}
