package resources

type direcciones struct {
	Direccion       string   `url:"direccion"`
	Provincia       []string `url:"provincia,comma,omitempty"`
	Departamento    []string `url:"departamento,comma,omitempty"`
	LocalidadCensal []string `url:"localidad_censal,comma,omitempty"`
	Localidad       []string `url:"localidad,comma,omitempty"`
}

func NewDireccion(value string) *direcciones {
	return &direcciones{
		Direccion: value,
	}
}

func (b *direcciones) SetProvincia(value ...string) {
	b.Provincia = value
}
