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

func (v *Version) Metadata() string {
	return v.metadata
}

func (v *Version) PrereleaseVersion() string {
	return v.preVersion
}
