package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	geoar "github.com/ctomatis/georefar-go"
	"github.com/ctomatis/georefar-go/resources"
)

func main() {
	withContext()
	list()
	findByID()
}

func list() {
	client := geoar.New(&geoar.Config{})

	filters := geoar.NewFilters()
	filters.Set("campos", []string{"completo"}).
		Set("orden", "nombre").
		Set("max", 2)

	provincia := resources.NewProvincia()

	if res, err := client.Send(provincia, filters).Json(); err == nil {
		jbytes, _ := json.MarshalIndent(res, " ", " ")
		println(string(jbytes))
	}
}

func findByID() {
	client := geoar.New(&geoar.Config{})

	filters := geoar.NewFilters()
	filters.Set("campos", []string{"completo"})

	provincia := resources.NewProvincia("54")

	if res, err := client.Send(provincia, filters).Json(); err == nil {
		jbytes, _ := json.MarshalIndent(res, " ", " ")
		println(string(jbytes))
	}
}

func withContext() {
	client := geoar.New(&geoar.Config{})
	provincia := resources.NewProvincia("54")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	res, err := client.WithContext(ctx).Send(provincia).Json()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}
