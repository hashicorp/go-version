package version

import (
	"testing"
)

func TestNewConstraint(t *testing.T) {
	cases := []struct {
		input string
		count int
		err   bool
	}{
		{">= 1.2", 1, false},
		{">= 1.x", 0, true},
		{">= 1.2, < 1.0", 2, false},
	}

	for _, tc := range cases {
		v, err := NewConstraint(tc.input)
		if tc.err && err == nil {
			t.Fatalf("expected error for input: %s", tc.input)
		} else if !tc.err && err != nil {
			t.Fatalf("error for input %s: %s", tc.input, err)
		}

		if len(v) != tc.count {
			t.Fatalf("input: %s\nexpected len: %d\nactual: %d",
				tc.input, tc.count, len(v))
		}
	}
}
