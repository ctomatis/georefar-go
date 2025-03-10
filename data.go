package geoar

type (
	Ubicacion struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	}

	BaseItem struct {
		Categoria string `json:"categoria,omitempty"`
		ID        string `json:"id,omitempty"`
		Nombre    string `json:"nombre,omitempty"`
	}

	FullItem struct {
		BaseItem
		Centroide      *Ubicacion `json:"centroide,omitempty"`
		Fuente         string     `json:"fuente,omitempty"`
		IsoID          string     `json:"iso_id,omitempty"`
		IsoNombre      string     `json:"iso_nombre,omitempty"`
		NombreCompleto string     `json:"nombre_completo,omitempty"`
		Interseccion   float64    `json:"interseccion,omitempty"`
	}

	Item struct {
		FullItem
		Provincia       *FullItem    `json:"provincia,omitempty"`
		Departamento    *FullItem    `json:"departamento,omitempty"`
		Municipio       *FullItem    `json:"municipio,omitempty"`
		LocalidadCensal *FullItem    `json:"localidad_censal,omitempty"`
		CalleAltura     *CalleAltura `json:"altura,omitempty"`
	}

	Direccion struct {
		Altura struct {
			Unidad string `json:"unidad"`
			Valor  int    `json:"valor"`
		} `json:"altura"`
		Calle           BaseItem  `json:"calle"`
		CalleCruce1     BaseItem  `json:"calle_cruce_1"`
		CalleCruce2     BaseItem  `json:"calle_cruce_2"`
		Departamento    BaseItem  `json:"departamento"`
		Fuente          string    `json:"fuente"`
		LocalidadCensal BaseItem  `json:"localidad_censal"`
		Nomenclatura    string    `json:"nomenclatura"`
		Provincia       BaseItem  `json:"provincia"`
		Piso            string    `json:"piso"`
		Ubicacion       Ubicacion `json:"ubicacion"`
	}

	CalleAltura struct {
		Fin struct {
			Derecha   int `json:"derecha"`
			Izquierda int `json:"izquierda"`
		} `json:"fin"`
		Inicio struct {
			Derecha   int `json:"derecha"`
			Izquierda int `json:"izquierda"`
		} `json:"inicio"`
	}

	Json struct {
		Cantidad            int         `json:"cantidad"`
		Inicio              int         `json:"inicio"`
		Total               int         `json:"total"`
		Parametros          filter      `json:"parametros"`
		Provincias          []Item      `json:"provincias,omitempty"`
		Departamentos       []Item      `json:"departamentos,omitempty"`
		Municipios          []Item      `json:"municipios,omitempty"`
		Direcciones         []Direccion `json:"direcciones,omitempty"`
		LocalidadesCensales []Item      `json:"localidades_censales,omitempty"`
		Localidades         []Item      `json:"localidades,omitempty"`
		Asentamientos       []Item      `json:"asentamientos,omitempty"`
		Calles              []Item      `json:"calles,omitempty"`
		Ubicacion           *Item       `json:"ubicacion,omitempty"`
	}

	GeoJson struct {
		Features []struct {
			Geometry struct {
				Coordinates []float64 `json:"coordinates"`
				Type        string    `json:"type"`
			} `json:"geometry"`
			Properties struct {
				Item
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"features"`
		Type string `json:"type"`
	}

	BulkJson struct {
		Resultados []Json `json:"resultados"`
	}
)
