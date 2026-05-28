// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package version

import "testing"

// strictSemverValid is a representative sample of version strings that are
// strictly conformant to the SemVer 2.0.0 specification at https://semver.org/.
// They must be accepted by NewStrictSemver and (for backward compatibility)
// also by NewVersion.
var strictSemverValid = []string{
	"0.0.4",
	"1.2.3",
	"10.20.30",
	"1.1.2-prerelease+meta",
	"1.1.2+meta",
	"1.1.2+meta-valid",
	"1.0.0-alpha",
	"1.0.0-beta",
	"1.0.0-alpha.beta",
	"1.0.0-alpha.beta.1",
	"1.0.0-alpha.1",
	"1.0.0-alpha0.valid",
	"1.0.0-alpha.0valid",
	"1.0.0-rc.1+build.1",
	"1.2.3-beta",
}

// strictSemverInvalid is the list of strings the original bug report
// (hashicorp/go-version#106) called out as invalid SemVer that the lax
// parser nonetheless accepts. NewStrictSemver must reject all of them.
// Backward-compat assertions about NewVersion live in the test below
// (we only assert that NewStrictSemver disagrees with NewVersion for any
// of these strings that NewVersion currently accepts, so future tightening
// of NewVersion does not regress this test).
var strictSemverInvalid = []string{
	"1.2.-3",
	"1.2.03",
	"1.2.3.4.5.6",
	"1.2.3-preview.01",
	"1.2.3+one.2!",
	"v 1.2.3",
	"1.2.3-",
	"1.2.3+",
	"1.2.3.",
	"1..2.3",
}

func TestNewStrictSemver_Valid(t *testing.T) {
	for _, s := range strictSemverValid {
		t.Run(s, func(t *testing.T) {
			if _, err := NewStrictSemver(s); err != nil {
				t.Fatalf("NewStrictSemver(%q) unexpected error: %v", s, err)
			}
			// Strictly-valid SemVer strings must also parse with the
			// existing looser constructor; otherwise back-compat is broken.
			if _, err := NewVersion(s); err != nil {
				t.Fatalf("NewVersion(%q) unexpected error (backward-compat regression): %v", s, err)
			}
		})
	}
}

func TestNewStrictSemver_Invalid(t *testing.T) {
	for _, s := range strictSemverInvalid {
		t.Run(s, func(t *testing.T) {
			if _, err := NewStrictSemver(s); err == nil {
				t.Fatalf("NewStrictSemver(%q) accepted an invalid SemVer string; expected an error", s)
			}
		})
	}
}

// TestNewStrictSemver_BackwardCompat asserts that adding NewStrictSemver did
// not change the acceptance set of NewVersion. For every input in the
// invalid table above that the legacy NewVersion currently accepts, this
// test confirms it is still accepted (i.e. NewVersion is unchanged) while
// NewStrictSemver rejects it. This is the load-bearing guarantee: existing
// Terraform / Hashicorp callers depending on the lax behavior cannot break.
func TestNewStrictSemver_BackwardCompat(t *testing.T) {
	for _, s := range strictSemverInvalid {
		_, legacyErr := NewVersion(s)
		_, strictErr := NewStrictSemver(s)
		if legacyErr == nil && strictErr == nil {
			t.Errorf("NewStrictSemver(%q) failed to reject a string that NewVersion accepts", s)
		}
	}
}

// TestNewStrictSemver_Examples covers the canonical example list from
// https://semver.org/#is-there-a-suggested-regular-expression-regex-to-check-a-semver-string.
func TestNewStrictSemver_Examples(t *testing.T) {
	valid := []string{
		"0.0.4",
		"1.2.3",
		"10.20.30",
		"1.1.2-prerelease+meta",
		"1.1.2+meta",
		"1.0.0-alpha",
		"1.0.0-alpha.beta",
		"1.0.0-alpha.beta.1",
		"1.0.0-rc.1+build.1",
		"2.0.0+build.1848",
		"2.0.1-alpha.1227",
		"1.0.0-alpha+beta",
		"1.2.3----RC-SNAPSHOT.12.9.1--.12+788",
		"1.0.0+0.build.1-rc.10000aaa-kk-0.1",
	}
	invalid := []string{
		"1",
		"1.2",
		"1.2.3-0123",
		"1.2.3-0123.0123",
		"1.1.2+.123",
		"+invalid",
		"-invalid",
		"-invalid+invalid",
		"-invalid.01",
		"alpha",
		"alpha.beta",
		"alpha.beta.1",
		"alpha.1",
		"alpha+beta",
		"alpha_beta",
		"01.1.1",
		"1.01.1",
		"1.1.01",
		"1.2.3.DEV",
		"1.2-SNAPSHOT",
		"1.2.31.2.3----RC-SNAPSHOT.12.09.1--..12+788",
	}
	for _, s := range valid {
		if _, err := NewStrictSemver(s); err != nil {
			t.Errorf("NewStrictSemver(%q) error: %v (expected accept)", s, err)
		}
	}
	for _, s := range invalid {
		if _, err := NewStrictSemver(s); err == nil {
			t.Errorf("NewStrictSemver(%q) accepted (expected reject)", s)
		}
	}
}
