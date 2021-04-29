package table

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/go-test/deep"
)

func TestReadRE(t *testing.T) {
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

func TestFindColumnCount(t *testing.T) {
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
		tb := NewTable(NewConfig().SetSeperator(c.sep))
		if err := tb.Read(bytes.NewBufferString(c.input)); err != nil {
			t.Errorf("Error doing read: %v", err)
			continue
		}

		tb.findColumnCount()
		if c.expected != tb.columnCount {
			t.Errorf("tb[%q].columnCount = %d; want %d", c.name, tb.columnCount, c.expected)
		}
	}
}

func TestPrint(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		sep      string
		expected string
	}

	cases := []testCase{
		{name: "basic", input: "a\nb\nc\n", sep: `\t+`, expected: ""},
		{name: "differing", input: "a\te\nb\nc\n", sep: `\t+`, expected: ""},
		{name: "snake", input: "a\te\nb\tab\tabc\nc\n", sep: `\t+`, expected: ""},
	}

	for _, c := range cases {
		buf := &bytes.Buffer{}
		tb := NewTable(NewConfig().SetSeperator(c.sep))
		if err := tb.Read(bytes.NewBufferString(c.input)); err != nil {
			t.Errorf("Error doing read: %v", err)
			continue
		}
		tb.Write(buf)
		fmt.Fprintln(buf, "")

		// TODO add expected and actually test it
	}
}
