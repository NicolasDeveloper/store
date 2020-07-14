package main

import (
	"github.com/NicolasDeveloper/store/src/inventory/driver-adapters/api/core"
)

func main() {
	app := core.NewApp()
	app.Initialize().StartupDatabase().ConfigEndpoints().Run(":3202")
}
