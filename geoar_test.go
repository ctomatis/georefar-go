package geoar

import (
	"context"
	"testing"
	"time"

	"github.com/ctomatis/georefar-go/internal"
	"github.com/ctomatis/georefar-go/resources"
	"github.com/go-playground/assert/v2"
)

func TestBuildURL(t *testing.T) {
	expected := "https://apis.datos.gob.ar/georef/api/provincias?id=02&max=1&orden=nombre"

	provincia := resources.NewProvincia("02")
	filters := NewFilters().
		Set("orden", "nombre").
		Set("max", 1)

	url := buildUrl(BASE_URL, provincia, filters)
	assert.Equal(t, url, expected)
}

func TestBuildPayload(t *testing.T) {
	provincia := []any{resources.NewProvincia()}
	filters := NewFilters().
		Set("orden", "nombre").
		Set("max", 1)

	path := internal.Resource(provincia[0])

	payload := buildPayload(path, provincia, filters)

	if _, ok := payload["provincias"]; !ok {
		t.Fatal("Fail to build request body")
	}
}

func TestRequestWithContext(t *testing.T) {
	client := New(&Config{})
	provincia := resources.NewProvincia("54")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.WithContext(ctx).Send(provincia).Json()
	if err != nil {
		t.Fatal(err)
	}
}

func TestJsonResponse(t *testing.T) {
	client := New(&Config{})
	provincia := resources.NewProvincia("54")

	res, err := client.Send(provincia).Json()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, res.Total)
}

func TestCsvResponse(t *testing.T) {
	client := New(&Config{})
	provincia := resources.NewProvincia()

	res, err := client.Send(provincia).Csv()
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, 0, len(res))
}

func TestDownloadDataset(t *testing.T) {
	client := New(&Config{})
	provincia := resources.NewProvincia()

	b, err := client.Download(provincia, CSV).Save()
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, 0, b)
}
