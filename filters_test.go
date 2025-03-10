package geoar

import (
	"encoding/json"
	"testing"

	"github.com/go-playground/assert/v2"
	qs "github.com/google/go-querystring/query"
)

func TestBuildFilters(t *testing.T) {
	filters := NewFilters().
		Set("campos", []string{"id", "nombre"}).
		Set("orden", "nombre").
		Set("max", 2)

	v, err := qs.Values(filters)
	if err != nil {
		t.Fatal("Fail to build filters:", err)
	}

	expected := "campos=id%2Cnombre&max=2&orden=nombre"
	assert.Equal(t, expected, v.Encode())
}

func TestJsonFilters(t *testing.T) {
	filters := NewFilters().
		Set("campos", []string{"id", "nombre"})

	jb, err := json.Marshal(filters)

	if err != nil {
		t.Fatal("Fail marshal filters:", err)
	}

	m := make(map[string]string)
	json.Unmarshal(jb, &m)

	if val, ok := m["campos"]; ok {
		assert.Equal(t, val, "id,nombre")
	} else {
		t.Fatal("Fail marshal filter: missing field")
	}
}
