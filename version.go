package version

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var NumericRegexp *regexp.Regexp
var VersionRegexp *regexp.Regexp

// Version represents a single version.
type Version struct {
	metadata string
	pre      string
	segments []int
}

func init() {
	NumericRegexp = regexp.MustCompile(`^[0-9]+$`)
	VersionRegexp = regexp.MustCompile(
		`^([0-9]+(\.[0-9]+){0,2})` +
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
	segments := make([]int, len(segmentsStr), 3)
	for i, str := range segmentsStr {
		val, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			return nil, fmt.Errorf(
				"Error parsing version: %s", err)
		}

		segments[i] = int(val)
	}
	for i := len(segments); i < 3; i++ {
		segments = append(segments, 0)
	}

	return &Version{
		metadata: matches[7],
		pre:      matches[4],
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

	segmentsSelf := v.Segments()
	segmentsOther := other.Segments()

	// If the segments are the same, we must compare on prerelease info
	if reflect.DeepEqual(segmentsSelf, segmentsOther) {
		preSelf := v.Prerelease()
		preOther := other.Prerelease()
		if preSelf == "" && preOther == "" {
			return 0
		}
		if preSelf == "" {
			return 1
		}
		if preOther == "" {
			return -1
		}

		panic("proper prerelase compare not done yet")
	}

	// Compare the segments
	for i := 0; i < len(segmentsSelf); i++ {
		lhs := segmentsSelf[i]
		rhs := segmentsOther[i]

		if lhs == rhs {
			continue
		} else if lhs < rhs {
			return -1
		} else {
			return 1
		}
	}

	panic("should not be reached")
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
	return v.pre
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
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%d.%d.%d", v.segments[0], v.segments[1], v.segments[2])
	if v.pre != "" {
		fmt.Fprintf(&buf, "-%s", v.pre)
	}
	if v.metadata != "" {
		fmt.Fprintf(&buf, "+%s", v.metadata)
	}

	return buf.String()
}
