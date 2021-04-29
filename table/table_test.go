package table

import (
	"bytes"
	"fmt"
	"testing"
)

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
