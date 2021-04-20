package table

import (
	"bytes"
	"fmt"
	"os"
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

func TestLongestRow(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		sep      string
		expected int
	}

	cases := []testCase{
		{name: "basic", input: "a\nb\nc\n", sep: `\t+`, expected: 1},
		{name: "differing", input: "a\te\nb\nc\n", sep: `\t+`, expected: 2},
		{name: "snake", input: "a\te\nb\tab\tabc\nc\n", sep: `\t+`, expected: 3},
	}

	for _, c := range cases {
		tb := NewTable(regexp.MustCompile(c.sep))
		tb.Read(bytes.NewBufferString(c.input))

		if c.expected != tb.longestLine() {
			t.Errorf("tb[%q].longestLine() = %d; want %d", c.name, tb.longestLine(), c.expected)
		}
	}
}

func TestPrint(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		sep      string
		expected int
	}

	cases := []testCase{
		{name: "basic", input: "a\nb\nc\n", sep: `\t+`, expected: 1},
		{name: "differing", input: "a\te\nb\nc\n", sep: `\t+`, expected: 2},
		{name: "snake", input: "a\te\nb\tab\tabc\nc\n", sep: `\t+`, expected: 3},
	}

	for _, c := range cases {
		tb := NewTable(regexp.MustCompile(c.sep))
		tb.Read(bytes.NewBufferString(c.input))
		tb.Write(os.Stdout)
		fmt.Fprintln(os.Stdout, "")

		t.Errorf("testing")
		//if c.expected != tb.longestLine() {
		//	t.Errorf("tb[%q].longestLine() = %d; want %d", c.name, tb.longestLine(), c.expected)
		//}
	}
}
