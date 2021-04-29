package table

import (
	"bytes"
	"strings"
	"testing"

	"github.com/go-test/deep"
)

func TestReadFormatCSV(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected [][]string
	}

	cases := []testCase{
		{
			name: "basic", input: "a\nb\nc\n",
			expected: [][]string{{"a"}, {"b"}, {"c"}},
		},
		{
			name: "multi-column", input: "a,1\nb,2\nc,3\n",
			expected: [][]string{{"a", "1"}, {"b", "2"}, {"c", "3"}},
		},
		{
			name: "quotations", input: "a,\",\"\nb,2\nc,3\n",
			expected: [][]string{{"a", ","}, {"b", "2"}, {"c", "3"}},
		},
	}

	for _, c := range cases {
		tb := NewTable(NewConfig().SetFormat(FormatCSV))
		if err := tb.Read(bytes.NewBufferString(c.input)); err != nil {
			t.Errorf("Error doing read: %v", err)
			continue
		}

		if diff := deep.Equal(c.expected, tb.data); diff != nil {
			t.Errorf("Table.Read(%q) =\n%v", c.name, strings.Join(diff, "\n"))
		}
	}
}
