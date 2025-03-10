# Georefar Go

This is the *un*official Go SDK Client for the *Servicio de Normalización de Datos Geográficos de Argentina*.

## Installation

Install from GitHub using ```go get```:

```
go get github.com/ctomatis/georefar-go
```

## Make your first call
Create a client, then call commands on it.
```go
package main

import (
	"fmt"
	"log"

	geoar "github.com/ctomatis/georefar-go"
	"github.com/ctomatis/georefar-go/resources"
)

func main() {
	client := geoar.New(&geoar.Config{})

	// Find an address and filter by provincia ID
	address := resources.NewDireccion("balcarce 50")
	address.SetProvincia("02") // CABA

	res, err := client.Send(address).Json()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}
```
## Client / Call configuration specifics

### Base URL
The default base URL for the API is ***https://apis.datos.gob.ar/georef/api***. You can modify this base URL by adding a different URL in the client configuration:

```go
config := &geoar.Config{
	BaseUrl:    "<BASE_URL>",
}

client := geoar.New(config)
```
### Authentication
Credentials for authentication are optional. The API is public, but it also uses Public and Secret keys for authentication (see [JWT Authentication](https://datosgobar.github.io/georef-ar-api/jwt-token) for more info). You must initialize a ``geoar.Config{}`` struct and pass it to the ``geoar.New()`` function.

It's recommended that you set your credentials as environment variables. Alternatively, you can pass it in your code directly.

```bash
export GEOAR_KEY='API key'
export GEOAR_SECRET='API secret'
```
Then initialize the client using configuration settings:
```go
publicKey := os.Getenv("GEOAR_KEY")
secretKey := os.Getenv("GEOAR_SECRET")

config := &geoar.Config{
	Key:    publicKey,
	Secret: secretKey,
}

client := geoar.New(config)
```
## Request examples
### Use filtering

```go
package main

import (
	"fmt"
	"log"

	geoar "github.com/ctomatis/georefar-go"
	"github.com/ctomatis/georefar-go/resources"
)

// List all provinces and order by `nombre`. 
// Return all `completo` fields.
func main() {
	client := geoar.New(&geoar.Config{})

	filters := geoar.NewFilters()
	filters.Set("campos", []string{"completo"}).
	Set("orden", "nombre")

	provincia := resources.NewProvincia()

	res, err := client.Send(provincia, filters).Json()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}
```

### Retrieve a specific format

```go
// Gets csv output
func main() {
	client := geoar.New(&geoar.Config{})

	filters := geoar.NewFilters()
	filters.Set("campos", []string{"completo"}).
	Set("orden", "nombre")

	provincia := resources.NewProvincia()

	res, err := client.Send(provincia, filters).Csv() // Call Csv() method.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}
```

### Download a dataset file

```go
func main() {
	client := geoar.New(&geoar.Config{})

	provincia := resources.NewProvincia()

	_, err := client.Download(provincia, geoar.CSV).Save()
	if err != nil {
		log.Fatal(err)
	}
	// To save to a custom file, you can pass `filename` in Save() method.
	// fname := "output.csv"
	// _, err := client.Download(provincia, geoar.CSV).Save(fname)
	// ...
}
```

### Bulk request
```go
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
	fmt.Printf("Data: %+v\n", res)
}
```

## Request with a Context

```go
func main() {
	client := geoar.New(&geoar.Config{})
	provincia := resources.NewProvincia("54")

	// cancel request if time exceeded 
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	res, err := client.WithContext(ctx).Send(provincia).Json()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}
// Outputs: Get "https://apis.datos.gob.ar/georef/api/provincias?id=54": context deadline 
```

## Running tests

To execute tests, run the following command:

```bash
make test
```

## Contributing
Feel free to contribute:

- Fork the project.
- Create a new branch.
- Improve documentation to it.
- Report an issue.
- All pull requests are welcomed.

## Documentation

- [Georef API Reference](https://www.argentina.gob.ar/datos-abiertos/georef/openapi)
- [Documentation and resources](https://www.argentina.gob.ar/georef/documentacion-y-recursos-georef)
- [Github Repo](https://github.com/datosgobar/georef-ar-api)