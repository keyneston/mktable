package table

import "strings"

// prepareContent does a number of things in order to prepare the content for
// display.
// 1. It strips any additional whitespace on either the front or end of the content.
// 2. It escapes any excess `|` characters which would get confused with table endings.
func prepareContent(in string) string {
	return strings.Replace(
		strings.TrimSpace(in),
		"|",
		`\|`,
		-1,
	)
}
