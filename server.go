package main

import (
	"flag"

	"github.com/derekdowling/pantry-api/kernel"
)

var production = flag.Bool("prod", false, "Starts the server in production mode")

func main() {
	// Parse in flags
	flag.Parse()
	kernel.Start(*production)
}
