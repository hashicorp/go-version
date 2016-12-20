package version

type Builder struct {
	Version
}

func (b *Builder) ResetMetadata() {
	b.metadata = ""
}

func (b *Builder) SetMetadata(md string) {
	b.metadata = md
}

func (b *Builder) ResetPrerelease() {
	b.pre = ""
}

func (b *Builder) SetPrerelease(id string) {
	b.pre = id
}

func (b *Builder) NextPatch() {
	if b.pre != "" {
		b.pre = ""
	} else {
		b.segments[2]++
	}
}

func (b *Builder) NextMinor() {
	if b.pre != "" {
		b.pre = ""
	} else {
		b.segments[1]++
		b.segments[2] = 0
	}
}

func (b *Builder) NextMajor() {
	if b.pre != "" {
		b.pre = ""
	} else {
		b.segments[0]++
		b.segments[1] = 0
		b.segments[2] = 0
	}
}

func (b *Builder) Done() *Version {
	return &b.Version
}
