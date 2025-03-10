package resources

type municipios struct {
	departamentos
}

func NewMunicipio() *municipios {
	return &municipios{}
}
