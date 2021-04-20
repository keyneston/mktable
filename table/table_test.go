package table

import (
	"bytes"
	"regexp"
	"strings"
	"testing"

	"github.com/go-test/deep"
)

func TestTableRead(t *testing.T) {

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
	}

	for _, c := range cases {
		tb := NewTable(regexp.MustCompile(c.sep))
		tb.Read(bytes.NewBufferString(c.input))

		if diff := deep.Equal(c.expected, tb.data); diff != nil {
			t.Errorf("Day.Parse(%q) =\n%v", c.name, strings.Join(diff, "\n"))
		}
	}
}
