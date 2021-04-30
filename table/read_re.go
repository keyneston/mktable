package table

import (
	"bufio"
	"bytes"
	"errors"
	"io"
)

func (t *Table) readFormatRE(r io.Reader) error {
	reader := bufio.NewReader(r)

	done := false
	for !done {
		line, err := reader.ReadSlice(byte(t.NewLine))
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return err
			}

			// If we receive io.EOF then we need to finish processing the final
			// line, then stop looping.
			done = true
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

	return nil
}
