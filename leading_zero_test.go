package version

import "testing"

// Regression test for https://github.com/hashicorp/go-version/issues/86:
// leading zeros are forbidden in semantic versioning numeric segments, so
// NewSemver must reject them.
func TestNewSemverRejectsLeadingZeros(t *testing.T) {
	invalid := []string{"1.010.2", "01.2.3", "1.2.03", "1.2.00", "00.1.2"}
	for _, s := range invalid {
		if _, err := NewSemver(s); err == nil {
			t.Errorf("NewSemver(%q): expected error, got none", s)
		}
	}

	valid := []string{"1.10.2", "0.1.2", "0.0.0", "10.20.30", "1.2"}
	for _, s := range valid {
		if _, err := NewSemver(s); err != nil {
			t.Errorf("NewSemver(%q): unexpected error %v", s, err)
		}
	}

	// The lax NewVersion parser intentionally keeps accepting and
	// canonicalizing leading zeros (see Version.String docs).
	if v, err := NewVersion("1.010.2"); err != nil {
		t.Errorf("NewVersion(%q): unexpected error %v", "1.010.2", err)
	} else if v.String() != "1.10.2" {
		t.Errorf("NewVersion(%q) canonical form: got %q, want %q", "1.010.2", v.String(), "1.10.2")
	}
}
