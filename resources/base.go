package resources

import (
	"encoding/json"
	"strings"
)

type base struct {
	ID           []string `url:"id,comma,omitempty" json:"-"`
	Nombre       string   `url:"nombre,omitempty" json:"nombre,omitempty"`
	Interseccion []string `url:"interseccion,comma,omitempty" json:"-"`
}

func (b *base) SetID(id ...string) {
	b.ID = id
}

func (b *base) SetNombre(value string) {
	b.Nombre = value
}

func (b *base) SetInterseccion(value ...string) {
	b.Interseccion = value
}

func (u *base) MarshalJSON() ([]byte, error) {
	type Alias base
	return json.Marshal(&struct {
		ID           string `json:"id,omitempty"`
		Interseccion string `json:"interseccion,omitempty"`
		*Alias
	}{
		ID:           strings.Join(u.ID, ","),
		Interseccion: strings.Join(u.Interseccion, ","),
		Alias:        (*Alias)(u),
	})
}
