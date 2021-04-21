package table

import "testing"

func TestPrepareContent(t *testing.T) {
	type testCase struct {
		in       string
		expected string
	}

	testCases := []testCase{
		{"foo", "foo"},
		{"    foo     ", "foo"},
		{"    foo   bar    ", "foo   bar"},
		{"|", `\|`},
		{"foo | bar", `foo \| bar`},
	}

	for _, c := range testCases {
		out := prepareContent(c.in)
		if c.expected != out {
			t.Errorf("prepareContent(%q) = %q; want %q", c.in, out, c.expected)
		}
	}
}
