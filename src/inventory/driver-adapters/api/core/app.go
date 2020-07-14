package core

import (
	"log"
	"net/http"

	"github.com/golobby/container"

	"go.uber.org/dig"

	adapters "github.com/NicolasDeveloper/store/src/inventory/driven-adapters"
	"github.com/NicolasDeveloper/store/src/inventory/driver-adapters/api/controllers"
	"github.com/NicolasDeveloper/store/src/inventory/ioc"

	"github.com/NicolasDeveloper/store/src/inventory/driver-adapters/api/common"
	"github.com/gorilla/mux"
)

//App initialize app
type App struct {
	router    *mux.Router
	container *dig.Container
}

//NewApp contructor
func NewApp() *App {
	return &App{}
}

//Initialize configure app
func (a *App) Initialize() *App {
	ioc.NewContainerIOC().RegisterDependencies()
	return a
}

//StartupDatabase connect to database
func (a *App) StartupDatabase() *App {
	var dbcontext *adapters.DbContext
	container.Make(&dbcontext)
	dbcontext.Connect()
	return a
}

//ConfigEndpoints endpoints
func (a *App) ConfigEndpoints() *App {
	a.router = mux.NewRouter()
	s := a.router.PathPrefix("/inventory-api/v1/").Subrouter()

	for _, b := range initBundles() {
		for _, route := range b.GetRoutes() {
			s.HandleFunc(route.Path, route.Handler).Methods(route.Method)
		}
	}

	http.Handle("/", a.router)
	return a
}

//Run startup app
func (a *App) Run(port string) *App {
	log.Fatal(http.ListenAndServe(port, a.router))
	return a
}

func initBundles() []common.Bundle {
	return []common.Bundle{
		controllers.NewProductRouter(),
	}
}
