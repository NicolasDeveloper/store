package usecases

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/NicolasDeveloper/store/src/catalog/core/entities"

	"github.com/NicolasDeveloper/store/src/catalog/core/dtos"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(product entities.Product) error {
	product.ID = ""
	args := r.Called(product)
	return args.Error(0)
}

func TestRegisterUseCase(t *testing.T) {
	t.Run("Should register a new product in catalog", func(t *testing.T) {
		repository := new(repositoryMock)

		product := entities.Product{
			Name:        "Teste",
			Price:       19.99,
			Description: "Teste de descrição",
		}

		repository.On("Save", product).Return(nil)

		productDTO := dtos.ProductDTO{
			Name:        "Teste",
			Price:       19.99,
			Description: "Teste de descrição",
		}

		useCase := NewRegisterUseCase(repository)
		useCase.Execute(productDTO)

		repository.AssertExpectations(t)
	})
}
