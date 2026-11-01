package version

import "testing"

// TestVersionComparePrereleaseDifferentSegments checks that the prerelease
// information is not ignored when the two versions have a different number of
// segments (see issue #216: 1.0.0-alpha used to compare equal to 1.0.0.0).
func TestVersionComparePrereleaseDifferentSegments(t *testing.T) {
	cases := []struct {
		v1       string
		v2       string
		expected int
	}{
		// Self has fewer segments and a prerelease; Other is zero-padded.
		{"1.0.0-alpha", "1.0.0.0", -1},
		{"1.0.0.0", "1.0.0-alpha", 1},
		{"1.2.3-rc.1", "1.2.3.0", -1},
		// Self has more segments (zero-padded) and a prerelease; Other is shorter.
		{"1.0.0.0-alpha", "1.0.0", -1},
		{"1.0.0", "1.0.0.0-alpha", 1},
		{"1.2.3.0-rc.1", "1.2.3", -1},
		// Both have prereleases, but a different number of segments.
		{"1.0.0-alpha", "1.0.0.0-beta", -1},
		{"1.0.0.0-beta", "1.0.0-alpha", 1},
		// Non-zero padding still dominates the comparison.
		{"1.0.0-alpha", "1.0.0.1", -1},
		{"1.0.0.1", "1.0.0-alpha", 1},
		// Zero-padded versions without prereleases stay equal.
		{"1.0.0", "1.0.0.0", 0},
		{"1.0.0.0", "1.0.0", 0},
	}

	for _, tc := range cases {
		v1, err := NewVersion(tc.v1)
		if err != nil {
			t.Fatalf("err: %v", err)
		}

		v2, err := NewVersion(tc.v2)
		if err != nil {
			t.Fatalf("err: %v", err)
		}

		actual := v1.Compare(v2)
		if actual != tc.expected {
			t.Fatalf("%v <=> %v: expected %v, got %v", tc.v1, tc.v2, tc.expected, actual)
		}
	}
}
