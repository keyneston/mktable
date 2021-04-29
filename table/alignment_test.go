package table

import "testing"

func TestParseAlignmentHeader(t *testing.T) {
	type testCase struct {
		input    string
		expected Alignment
	}

	testCases := []testCase{
		{":--", AlignLeft},
		{":-:", AlignCenter},
		{":------:", AlignCenter},
		{"------:", AlignRight},
	}

	for _, c := range testCases {
		align := parseAlignmentHeader(c.input)
		if align != c.expected {
			t.Errorf("parseAlignmentHeader(%q) = %v; want %v", c.input, align, c.expected)
		}

	}

}
