package resources

type localidades struct {
	base
	localidadesCensales
	LocalidadCensal []string `url:"localidad_censal,comma,omitempty"`
}

func NewLocalidad() *localidades {
	return &localidades{}
}

func (l *localidades) SetLocalidadCensal(value ...string) {
	l.LocalidadCensal = value
}
