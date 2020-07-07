package usecases

import (
	"github.com/NicolasDeveloper/store/src/catalog/core/dtos"
	"github.com/NicolasDeveloper/store/src/catalog/core/entities"
	"github.com/NicolasDeveloper/store/src/catalog/core/ports"
)

//RegisterUseCase use case
type RegisterUseCase struct {
	repository ports.IProductRepositoryPort
}

//NewRegisterUseCase constructor
func NewRegisterUseCase(repository ports.IProductRepositoryPort) *RegisterUseCase {
	return &RegisterUseCase{
		repository: repository,
	}
}

//Execute execute use case
func (u *RegisterUseCase) Execute(productDTO dtos.ProductDTO) error {
	entity, error := entities.NewProduct(
		productDTO.Name,
		productDTO.Price,
		productDTO.Description,
	)

	u.repository.Save(entity)

	return error
}
