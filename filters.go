package geoar

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/ctomatis/georefar-go/internal"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type filter struct {
	Campos  []string `url:"campos,comma,omitempty" json:"campos"`
	Orden   string   `url:"orden,omitempty" json:"orden,omitempty"`
	Aplanar bool     `url:"aplanar,omitempty" json:"aplanar,omitempty"`
	Max     int      `url:"max,omitempty" json:"max,omitempty"`
	Inicio  int      `url:"inicio,omitempty" json:"inicio"`
	Exacto  bool     `url:"exacto,omitempty" json:"exacto,omitempty"`
	Formato string   `url:"-" json:"formato,omitempty"`
}

func NewFilters() *filter {
	return &filter{}
}

func (f *filter) Set(name string, value any) *filter {
	fieldName := cases.Title(language.Und).String(name)

	field := reflect.ValueOf(f).Elem().FieldByName(fieldName)

	if field.IsValid() {
		val := reflect.ValueOf(value)

		if field.Kind() == reflect.String {
			field.SetString(val.String())
		}
		if field.Kind() == reflect.Int {
			field.SetInt(val.Int())
		}

		if field.Kind() == reflect.Slice {
			field.Set(val.Slice(0, val.Len()))
		}

		if field.Kind() == reflect.Bool {
			field.SetBool(val.Bool())
		}

	}
	return f
}

func (f *filter) toMap() (m internal.Map) {
	return internal.ToMap(f)
}

func (f *filter) MarshalJSON() ([]byte, error) {
	type Alias filter
	return json.Marshal(&struct {
		Campos string `json:"campos,omitempty"`
		*Alias
	}{
		Campos: strings.Join(f.Campos, ","),
		Alias:  (*Alias)(f),
	})
}
