package table

import (
	"fmt"
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
