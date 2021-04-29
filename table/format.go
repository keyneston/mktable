package table

import (
	"fmt"
	"strings"
)

type Format string

const (
	FormatRE      Format = "regexp"
	FormatCSV     Format = "csv"
	FormatTSV     Format = "tsv"
	FormatMK      Format = "mk"
	FormatUnknown Format = "unknown"
)

func AllFormats() []string {
	return []string{
		string(FormatRE),
		string(FormatMK),
		string(FormatCSV),
	}
}

func ParseFormat(in string) (Format, error) {
	switch strings.ToLower(in) {
	case "re", "regexp", "regex":
		return FormatRE, nil
	case "csv":
		return FormatCSV, nil
		//	case "tsv":
		//		return FormatTSV, nil
	case "mk", "markdown":
		return FormatMK, nil
	default:
		return FormatUnknown, fmt.Errorf("Unknown format %q", in)
	}
}
