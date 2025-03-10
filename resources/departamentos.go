package resources

type departamentos struct {
	base
	Provincia []string `url:"provincia,comma,omitempty"`
}

func NewDepartamento() *departamentos {
	return &departamentos{}
}

func (d *departamentos) SetProvincia(value ...string) {
	d.Provincia = value
}
