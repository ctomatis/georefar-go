package resources

type localidadesCensales struct {
	base
	Provincia    []string `url:"provincia,comma,omitempty"`
	Departamento []string `url:"departamento,comma,omitempty"`
	Municipio    []string `url:"municipio,comma,omitempty"`
}

func NewLocalidadCensal() *localidadesCensales {
	return &localidadesCensales{}
}

func (d *localidadesCensales) SetProvincia(value ...string) {
	d.Provincia = value
}

func (d *localidadesCensales) SetDepartamento(value ...string) {
	d.Departamento = value
}

func (d *localidadesCensales) SetMunicipio(value ...string) {
	d.Municipio = value
}
