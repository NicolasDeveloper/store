package usecases

import (
	"github.com/NicolasDeveloper/store/src/inventory/core/dtos"
	"github.com/NicolasDeveloper/store/src/inventory/core/entities"
	"github.com/NicolasDeveloper/store/src/inventory/core/ports"
)

//RegisterProductQuantity use case
type RegisterProductQuantity struct {
	repository ports.IProductRepositoryPort
}

//NewRegisterProductQuantity constructor
func NewRegisterProductQuantity(repository ports.IProductRepositoryPort) *RegisterProductQuantity {
	return &RegisterProductQuantity{
		repository: repository,
	}
}

//Execute execute use case
func (u *RegisterProductQuantity) Execute(productDTO *dtos.ProductDTO) error {
	entity, error := entities.NewProduct(
		productDTO.ID,
		productDTO.Quantity,
	)

	u.repository.Save(entity)
	productDTO.ID = entity.ID

	return error
}
