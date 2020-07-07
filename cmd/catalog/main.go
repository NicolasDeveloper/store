package main

import (
	"github.com/NicolasDeveloper/store/src/catalog/presentation/api/core"
)

func main() {
	app := core.NewApp()
	app.Initialize().ConfigEndpoints().Run(":3201")
}
