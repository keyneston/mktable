package table

import (
	"bytes"
	"strings"
	"testing"

	"github.com/go-test/deep"
)

func TestReadFormatRE(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		sep      string
		expected [][]string
	}

	cases := []testCase{
		{name: "basic", input: "a\nb\nc\n",
			sep:      `\t+`,
			expected: [][]string{{"a"}, {"b"}, {"c"}},
		},
		{name: "multi-column", input: "a\t1\nb\t2\nc\t3\n",
			sep:      `\t+`,
			expected: [][]string{{"a", "1"}, {"b", "2"}, {"c", "3"}},
		},
		{
			name: "missing final newline", input: "a\tb\ta\t\n1\t2\t3",
			sep:      `\t+`,
			expected: [][]string{{"a", "b", "a"}, {"1", "2", "3"}},
		},
	}

	for _, c := range cases {
		tb := NewTable(NewConfig().SetSeperator(c.sep))
		if err := tb.Read(bytes.NewBufferString(c.input)); err != nil {
			t.Errorf("Error doing read: %v", err)
			continue
		}

		if diff := deep.Equal(c.expected, tb.data); diff != nil {
			t.Errorf("Table.Read(%q) =\n%v", c.name, strings.Join(diff, "\n"))
		}
	}
}
