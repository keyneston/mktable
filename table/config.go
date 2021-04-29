package table

import "regexp"

type TableConfig struct {
	MaxPadding  int
	SkipHeaders bool
	Format      Format
	Alignments  map[int]Alignment
	Seperator   *regexp.Regexp
	NewLine     NewLine
}

func NewConfig() TableConfig {
	return TableConfig{}
}

func (t TableConfig) SetSeperator(in string) TableConfig {
	t.Seperator = regexp.MustCompile(in)
	t.Format = FormatRE

	return t
}

func (t TableConfig) SetFormat(format Format) TableConfig {
	t.Format = format
	return t
}
