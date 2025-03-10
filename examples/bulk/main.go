package main

import (
	"encoding/json"
	"fmt"
	"log"

	geoar "github.com/ctomatis/georefar-go"
	"github.com/ctomatis/georefar-go/resources"
)

func main() {
	client := geoar.New(&geoar.Config{})

	provincia_1 := resources.NewProvincia("02", "42")

	filters_1 := geoar.NewFilters()
	filters_1.Set("campos", []string{"completo"})

	provincia_2 := resources.NewProvincia("54")

	filters_2 := geoar.NewFilters()
	filters_2.Set("campos", []string{"id", "nombre"})

	provincias := []any{provincia_1, provincia_2}

	res, err := client.Bulk(provincias, filters_1, filters_2)
	if err != nil {
		log.Fatal(err)
	}

	jb, _ := json.MarshalIndent(res, " ", " ")

	fmt.Printf("%s\n", jb)
}
