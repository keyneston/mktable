package table

import (
	"fmt"
	"regexp"
	"strings"
)

type Alignment string

const (
	AlignDefault Alignment = ""
	AlignLeft    Alignment = "left"
	AlignRight   Alignment = "right"
	AlignCenter  Alignment = "center"
)

func (a Alignment) header(size int) string {
	if size < 3 {
		size = 3
	}

	switch a {
	case AlignRight:
		return fmt.Sprintf("%s:", strings.Repeat("-", size-1))
	case AlignLeft:
		return fmt.Sprintf(":%s", strings.Repeat("-", size-1))
	case AlignCenter:
		return fmt.Sprintf(":%s:", strings.Repeat("-", size-2))
	default:
		return strings.Repeat("-", size)
	}
}

var (
	reAlignLeft   = regexp.MustCompile("^:-{2,}$")
	reAlignRight  = regexp.MustCompile("^-{2,}:$")
	reAlignCenter = regexp.MustCompile("^:-+:$")
)

func parseAlignmentHeader(input string) Alignment {
	switch {
	case reAlignCenter.MatchString(input):
		return AlignCenter
	case reAlignLeft.MatchString(input):
		return AlignLeft
	case reAlignRight.MatchString(input):
		return AlignRight
	default:
		return AlignDefault
	}
}
