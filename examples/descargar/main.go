package main

import (
	"log"

	geoar "github.com/ctomatis/georefar-go"
	"github.com/ctomatis/georefar-go/resources"
)

func main() {

	client := geoar.New(&geoar.Config{})

	provincia := resources.NewProvincia()

	_, err := client.Download(provincia, geoar.CSV).Save()
	if err != nil {
		log.Fatal(err)
	}

	/*
		fname := "output.csv"
		_, err := client.Download(provincia, geoar.CSV).Save(fname)
		if err != nil {
			log.Fatal(err)
		}
	*/
}
