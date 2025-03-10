package resources

type provincias struct {
	base
}

func NewProvincia(id ...string) *provincias {
	p := &provincias{}
	p.SetID(id...)
	return p
}
