package main

import (
	"fmt"
	"log"
	"os"

	geoar "github.com/ctomatis/georefar-go"
	"github.com/ctomatis/georefar-go/resources"
)

func main() {
	publicKey := os.Getenv("GEOAR_KEY")
	secretKey := os.Getenv("GEOAR_SECRET")

	config := &geoar.Config{
		Key:    publicKey,
		Secret: secretKey,
	}

	client := geoar.New(config)

	provincia := resources.NewProvincia()

	res, err := client.Send(provincia).Json()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}
