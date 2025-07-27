// cmd/nanobyte/main.go
package main

import (
	"flag"
	"log"
	"os"

	"nanobyte/internal/nanobyte"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := nanobyte.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
