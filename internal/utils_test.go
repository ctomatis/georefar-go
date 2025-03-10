package internal

import (
	"testing"

	"github.com/ctomatis/georefar-go/resources"
	"github.com/go-playground/assert/v2"
)

func TestDashCase(t *testing.T) {
	expected := "test-dash-case"
	dashed := toDashCase("TestDashCase")

	assert.Equal(t, dashed, expected)
}

func TestResourceName(t *testing.T) {
	expected := "localidades-censales"
	rcs := Resource(resources.NewLocalidadCensal())

	assert.Equal(t, rcs, expected)
}
