package ports

import "github.com/NicolasDeveloper/store/src/catalog/core/entities"

//IProductQueryPort database port
type IProductQueryPort interface {
	Search(skip int64, take int64, filter interface{}) ([]entities.Product, error)
}
