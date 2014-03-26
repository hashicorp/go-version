package version

import (
	"reflect"
	"testing"
)

func TestNewVersion(t *testing.T) {
	cases := []struct {
		version string
		err     bool
	}{
		{"1.2.3", false},
		{"1.0", false},
		{"1", false},
		{"1.2.beta", true},
		{"foo", true},
		{"1.2-5", false},
		{"1.2-beta.5", false},
		{"\n1.2", true},
		{"1.2.0-x.Y.0+metadata", false},
	}

	for _, tc := range cases {
		_, err := NewVersion(tc.version)
		if tc.err && err == nil {
			t.Fatalf("expected error for version: %s", tc.version)
		} else if !tc.err && err != nil {
			t.Fatalf("error for version %s: %s", tc.version, err)
		}
	}
}

func TestVersionMetadata(t *testing.T) {
	cases := []struct {
		version  string
		expected string
	}{
		{"1.2.3", ""},
		{"1.2-beta", ""},
		{"1.2.0-x.Y.0", ""},
		{"1.2.0-x.Y.0+metadata", "metadata"},
	}

	for _, tc := range cases {
		v, err := NewVersion(tc.version)
		if err != nil {
			t.Fatalf("err: %s", err)
		}

		actual := v.Metadata()
		expected := tc.expected
		if actual != expected {
			t.Fatalf("expected: %s\nactual: %s", expected, actual)
		}
	}
}

func TestVersionPrerelease(t *testing.T) {
	cases := []struct {
		version  string
		expected string
	}{
		{"1.2.3", ""},
		{"1.2-beta", "beta"},
		{"1.2.0-x.Y.0", "x.Y.0"},
		{"1.2.0-x.Y.0+metadata", "x.Y.0"},
	}

	for _, tc := range cases {
		v, err := NewVersion(tc.version)
		if err != nil {
			t.Fatalf("err: %s", err)
		}

		actual := v.Prerelease()
		expected := tc.expected
		if actual != expected {
			t.Fatalf("expected: %s\nactual: %s", expected, actual)
		}
	}
}

func TestVersionSegments(t *testing.T) {
	cases := []struct {
		version  string
		expected []int
	}{
		{"1.2.3", []int{1,2,3}},
		{"1.2-beta", []int{1,2}},
		{"1.2.0-x.Y.0", []int{1,2,0}},
		{"1.2.0-x.Y.0+metadata", []int{1,2,0}},
	}

	for _, tc := range cases {
		v, err := NewVersion(tc.version)
		if err != nil {
			t.Fatalf("err: %s", err)
		}

		actual := v.Segments()
		expected := tc.expected
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("expected: %#v\nactual: %#v", expected, actual)
		}
	}
}

func TestVersionString(t *testing.T) {
	cases := []string {
		"1.2.3",
		"1.2-beta",
		"1.2.0-x.Y.0",
		"1.2.0-x.Y.0+metadata",
	}

	for _, tc := range cases {
		v, err := NewVersion(tc)
		if err != nil {
			t.Fatalf("err: %s", err)
		}

		actual := v.String()
		expected := tc
		if actual != expected {
			t.Fatalf("expected: %s\nactual: %s", expected, actual)
		}
	}
}
