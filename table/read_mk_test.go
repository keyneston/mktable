package table

import (
	"bytes"
	"strings"
	"testing"

	"github.com/go-test/deep"
)

func TestReadFormatMK(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected [][]string
	}

	cases := []testCase{
		{
			name: "basic", input: "| a | b | c |\n",
			expected: [][]string{{"a", "b", "c"}},
		},
		{
			name: "multiline", input: "| a | b | c |\n| --- | ---| ---|\n|1 | 2|3|\n",
			expected: [][]string{{"a", "b", "c"}, {"1", "2", "3"}},
		},
		{
			name: "trailing space", input: "| a | \n| --- |\n|1 |\n",
			expected: [][]string{{"a"}, {"1"}},
		},
		{
			name: "header", input: "| --- | --- | --- |\n",
			expected: nil,
		},
		{
			name: "trailing pipe", input: "| a | b | \\| |\n",
			expected: [][]string{{"a", "b", `\|`}},
		},
	}

	for _, c := range cases {
		tb := NewTable(TableConfig{
			Format: FormatMK,
		})
		if err := tb.Read(bytes.NewBufferString(c.input)); err != nil {
			t.Errorf("Error doing read: %v", err)
			continue
		}

		if diff := deep.Equal(c.expected, tb.data); diff != nil {
			t.Errorf("Table.Read(%q) =\n%v", c.name, strings.Join(diff, "\n"))
		}
	}
}
