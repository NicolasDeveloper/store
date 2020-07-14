package ioc

import (
	"github.com/NicolasDeveloper/store/src/inventory/core/ports"
	usecases "github.com/NicolasDeveloper/store/src/inventory/core/use-cases"
	adapters "github.com/NicolasDeveloper/store/src/inventory/driven-adapters"
	"github.com/golobby/container"
)

//ContainerIOC typing
type ContainerIOC struct {
}

//NewContainerIOC constructor
func NewContainerIOC() *ContainerIOC {
	return &ContainerIOC{}
}

//RegisterDependencies register dependency injection
func (i *ContainerIOC) RegisterDependencies() error {
	container.Singleton(func() *adapters.DbContext {
		return adapters.NewContext()
	})

	container.Transient(func(dbcontext *adapters.DbContext) ports.IProductRepositoryPort {
		adapter, _ := adapters.NewProductsRepositoryAdapter(dbcontext)
		return adapter
	})

	container.Transient(func(port ports.IProductRepositoryPort) *usecases.RegisterProductQuantity {
		return usecases.NewRegisterProductQuantity(port)
	})

	return nil
}
