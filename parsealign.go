package main

import (
	"fmt"
	"strings"

	"github.com/keyneston/mktable/table"
)

type ParseAlignments struct {
	alignments map[int]table.Alignment
}

func (p *ParseAlignments) String() string {
	list := []string{}
	for k, v := range p.alignments {
		arrow := ""
		switch v {
		case table.AlignCenter:
			arrow = "<>"
		case table.AlignRight:
			arrow = ">"
		case table.AlignLeft:
			arrow = "<"
		default:
			arrow = "="
		}

		list = append(list, fmt.Sprintf("%d%s", k, arrow))
	}
	return strings.Join(list, ",")
}

func (p *ParseAlignments) Set(input string) error {
	if p.alignments == nil {
		p.alignments = map[int]table.Alignment{}
	}

	sep := strings.Split(input, ",")

	for _, cur := range sep {
		var num int
		var arrow string

		if _, err := fmt.Sscanf(cur, "%d%s", &num, &arrow); err != nil {
			return fmt.Errorf("Error parsing %q: %v", cur, err)
		}

		align := table.AlignDefault
		switch arrow {
		case "<":
			align = table.AlignLeft
		case ">":
			align = table.AlignRight
		case "=":
			align = table.AlignCenter
		default:
			return fmt.Errorf("Unknown alignment: \"%d%s\"", num, arrow)
		}

		p.alignments[num] = align
	}

	return nil
}
