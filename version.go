package version

import (
	"fmt"
	"regexp"
)

var NumericRegexp *regexp.Regexp
var VersionRegexp *regexp.Regexp

// Version represents a single version.
type Version struct {
	metadata   string
	version    string
	preVersion string
}

func init() {
	NumericRegexp = regexp.MustCompile(`^[0-9]+$`)
	VersionRegexp = regexp.MustCompile(
		`^([0-9]+(\.[0-9]+)*)` +
			`(-([0-9A-Za-z]+(\.[0-9A-Za-z]+)*))?` +
			`(\+([0-9A-Za-z]+(\.[0-9A-Za-z]+)*))?` +
			`?$`)
}

// NewVersion parses the given version and returns a new
// Version.
func NewVersion(v string) (*Version, error) {
	matches := VersionRegexp.FindStringSubmatch(v)
	if matches == nil {
		return nil, fmt.Errorf("Malformed version: %s", v)
	}

	return &Version{
		metadata:   matches[7],
		version:    matches[1],
		preVersion: matches[4],
	}, nil
}

// Metadata returns any metadata that was part of the version
// string.
//
// Metadata is anything that comes after the "+" in the version.
// For example, with "1.2.3+beta", the metadata is "beta".
func (v *Version) Metadata() string {
	return v.metadata
}

// Prerelease returns any prerelease data that is part of the version,
// or blank if there is no prerelease data.
//
// Prerelease information is anything that comes after the "-" in the
// version (but before any metadata). For example, with "1.2.3-beta",
// the prerelease information is "beta".
func (v *Version) Prerelease() string {
	return v.preVersion
}
