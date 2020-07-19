package drivenadapters

import (
	"context"

	"github.com/NicolasDeveloper/store/src/catalog/core/entities"
	"github.com/NicolasDeveloper/store/src/catalog/core/ports"
)

type productsRepositoryAdapter struct {
	ctx *DbContext
	ports.IProductRepositoryPort
}

//NewProductsRepositoryAdapter constructor
func NewProductsRepositoryAdapter(ctx *DbContext) (ports.IProductRepositoryPort, error) {
	return &productsRepositoryAdapter{
		ctx: ctx,
	}, nil
}

//Save search products in database
func (r *productsRepositoryAdapter) Save(product entities.Product) error {
	collection, error := r.ctx.GetCollection(product)
	_, error = collection.InsertOne(context.TODO(), product)
	return error
}
