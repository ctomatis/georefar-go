package main

import (
	"fmt"
	"log"

	geoar "github.com/ctomatis/georefar-go"
	"github.com/ctomatis/georefar-go/resources"
)

func main() {
	client := geoar.New(&geoar.Config{})

	filters := geoar.NewFilters()
	filters.Set("campos", []string{"completo"}).Set("max", 1)

	dir := resources.NewDireccion("balcarce 50")
	dir.SetProvincia("02")

	res, err := client.Send(dir, filters).Json()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}
