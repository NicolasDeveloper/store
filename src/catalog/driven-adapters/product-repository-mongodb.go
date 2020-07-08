package drivenadapters

import (
	"context"

	"github.com/NicolasDeveloper/store/src/catalog/core/entities"
	"github.com/NicolasDeveloper/store/src/catalog/core/ports"
	"github.com/NicolasDeveloper/store/src/catalog/dbcontext"
)

type productsRepositoryAdapter struct {
	ctx dbcontext.DbContext
	ports.IProductRepositoryPort
}

//NewProductsRepositoryAdapter constructor
func NewProductsRepositoryAdapter(ctx dbcontext.DbContext) (ports.IProductRepositoryPort, error) {
	return &productsRepositoryAdapter{
		ctx: ctx,
	}, nil
}

//Save search products in database
func (r *productsRepositoryAdapter) Save(product entities.Product) error {
	collection, error := r.ctx.GetCollection(product)
	collection.InsertOne(context.TODO(), product)
	return error
}
