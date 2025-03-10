package resources

type asentamientos struct {
	localidades
}

func NewAsentamiento() *asentamientos {
	return &asentamientos{}
}
