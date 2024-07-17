package version

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollection(t *testing.T) {
	versionsRaw := []string{
		"1.1.1",
		"1.0",
		"1.2",
		"2",
		"0.7.1",
		"1.2.3",
		"1",
		"1.2-5",
		"1.2-beta.5",
		"1.2.0-x.Y.0+metadata",
		"1.2.0-x.Y.0+metadata-width-hypen",
		"1.2.3-rc1-with-hypen",
		"1.2.3.4",
		"1.2.0.4-x.Y.0+metadata",
		"1.2.0.4-x.Y.0+metadata-width-hypen",
		"1.2.0-X-1.2.0+metadata~dist",
		"1.2.3.4-rc1-with-hypen",
		"1.2.3.4",
		"V1.2.3",
		"1.7rc2",
		"v1.7rc2",
		"v1.0-",
		"2.28.0.618+gf4bc123cb7",
		"1.13.0+dev-545-gb3b1c081b",
		"2.28.0.618.gf4bc123cb7",
		"2.29.0.rc0.261.g7178c9af9c",
		"1.2.beta",
		"1.21.beta",
		"v1.13.0-rc1",
	}

	versions := make([]*Version, len(versionsRaw))
	for i, raw := range versionsRaw {
		t.Run(raw, func(t *testing.T) {
			v, err := NewVersion(raw)
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			versions[i] = v
		})
	}

	sort.Sort(Collection(versions))

	actual := make([]string, len(versions))
	for i, v := range versions {
		actual[i] = v.String()
	}

	expected := []string{
		"0.7.1",
		"1.0.0--",
		"1.0.0",
		"1.0.0",
		"1.1.1",
		"1.2.0-5",
		"1.2.0-X-1.2.0+metadata~dist",
		"1.2.0-beta",
		"1.2.0-beta.5",
		"1.2.0-x.Y.0+metadata",
		"1.2.0-x.Y.0+metadata-width-hypen",
		"1.2.0",
		"1.2.0.4-x.Y.0+metadata",
		"1.2.0.4-x.Y.0+metadata-width-hypen",
		"1.2.3-rc1-with-hypen",
		"1.2.3",
		"1.2.3",
		"1.2.3.4-rc1-with-hypen",
		"1.2.3.4",
		"1.2.3.4",
		"1.7.0-rc2",
		"1.7.0-rc2",
		"1.13.0-rc1",
		"1.13.0+dev-545-gb3b1c081b",
		"1.21.0-beta",
		"2.0.0",
		"2.28.0.618-gf4bc123cb7",
		"2.28.0.618+gf4bc123cb7",
		"2.29.0-rc0.261.g7178c9af9c",
	}

	if !reflect.DeepEqual(actual, expected) {
		assert.EqualValues(t, expected, actual)
	}
}
