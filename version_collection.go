package version

// VersionCollection is a type that implements the sort.Interface interface
// so that versions can be sorted.
type VersionCollection []*Version

func (v VersionCollection) Len() int {
	return len(v)
}

func (v VersionCollection) Less(i, j int) bool {
	return v[i].LessThan(v[j])
}

func (v VersionCollection) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
