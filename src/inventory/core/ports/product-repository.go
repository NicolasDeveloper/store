package ports

import (
	"github.com/NicolasDeveloper/store/src/inventory/core/entities"
)

//IProductRepositoryPort port
type IProductRepositoryPort interface {
	Save(product *entities.Product) error
}
