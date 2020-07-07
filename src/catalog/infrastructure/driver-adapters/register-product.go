package driveradapters

import (
	"fmt"

	"github.com/NicolasDeveloper/store/src/catalog/core/usecases"
	"github.com/NicolasDeveloper/store/src/shared/interfaces"
)

//RegisterProductadAdapter adapter
type RegisterProductadAdapter struct {
	useCase usecases.RegisterUseCase
	interfaces.IAdapter
}

//NewRegisterProductAdapter constructor
func NewRegisterProductAdapter(
	useCase usecases.RegisterUseCase,
) (interfaces.IAdapter, error) {
	return &RegisterProductadAdapter{
		useCase: useCase,
	}, nil
}

//Execute execute adatper behavor
func (a *RegisterProductadAdapter) Execute(dto interface{}) error {
	fmt.Print(dto)
	return nil
}
