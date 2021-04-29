package main

import "github.com/keyneston/mktable/table"

type FormatValue struct {
	format table.Format
}

func (fv *FormatValue) Get() interface{} {
	return fv.GetFormat()
}

func (fv *FormatValue) GetFormat() table.Format {
	if fv.format == "" {
		return table.FormatRE
	}

	return fv.format
}

func (fv *FormatValue) String() string {
	return string(fv.format)
}

func (fv *FormatValue) Set(in string) error {
	f, err := table.ParseFormat(in)
	if err != nil {
		return err
	}

	fv.format = f
	return nil
}
