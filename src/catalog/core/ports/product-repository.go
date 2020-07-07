package ports

import "github.com/NicolasDeveloper/store/src/catalog/core/entities"

//IProductRepositoryPort database port
type IProductRepositoryPort interface {
	Save(product entities.Product) error
}
