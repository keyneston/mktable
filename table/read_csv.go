package table

import (
	"encoding/csv"
	"io"
)

func (t *Table) readFormatCSV(r io.Reader) error {
	reader := csv.NewReader(r)

	var err error
	t.data, err = reader.ReadAll()
	return err
}
