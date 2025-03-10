package resources

type ubicacion struct {
	Lat float64 `url:"lat,omitempty"`
	Lon float64 `url:"lon,omitempty"`
}

func NewUbicacion(lat, lon float64) *ubicacion {
	return &ubicacion{
		Lat: lat,
		Lon: lon,
	}
}
