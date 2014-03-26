package version

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var NumericRegexp *regexp.Regexp
var VersionRegexp *regexp.Regexp

// Version represents a single version.
type Version struct {
	original   string
	metadata   string
	preVersion string
	segments   []int
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

	segmentsStr := strings.Split(matches[1], ".")
	segments := make([]int, len(segmentsStr))
	for i, str := range segmentsStr {
		val, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			return nil, fmt.Errorf(
				"Error parsing version: %s", err)
		}

		segments[i] = int(val)
	}

	return &Version{
		original:   v,
		metadata:   matches[7],
		preVersion: matches[4],
		segments: segments,
	}, nil
}

// Compare compares this version to another version. This
// returns -1, 0, or 1 if this version is smaller, equal,
// or larger than the other version, respectively.
//
// If you want boolean results, use the LessThan, Equal,
// or GreaterThan methods.
func (v *Version) Compare(other *Version) int {
	// A quick, efficient equality check
	if v.String() == other.String() {
		return 0
	}

	return 0
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

// Segments returns the numeric segments of the version as a slice.
//
// This excludes any metadata or pre-release information. For example,
// for a version "1.2.3-beta", segments will return a slice of
// 1, 2, 3.
func (v *Version) Segments() []int {
	return v.segments
}

// String returns the full version string included pre-release
// and metadata information.
func (v *Version) String() string {
	return v.original
}
