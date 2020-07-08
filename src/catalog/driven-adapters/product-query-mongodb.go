package drivenadapters

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/NicolasDeveloper/store/src/catalog/core/entities"
	"github.com/NicolasDeveloper/store/src/catalog/core/ports"
	"github.com/NicolasDeveloper/store/src/catalog/dbcontext"
)

type productQuery struct {
	ctx dbcontext.DbContext
	ports.IProductQueryPort
}

//NewProductQuery constructor
func NewProductQuery(ctx dbcontext.DbContext) (ports.IProductQueryPort, error) {
	return &productQuery{
		ctx: ctx,
	}, nil
}

//Search search products in database
func (r *productQuery) Search(skip int64, take int64, filter interface{}) ([]entities.Product, error) {
	collection, error := r.ctx.GetCollection(entities.Product{})
	options := options.Find()
	options.SetLimit(take)

	cursor, error := collection.Find(
		context.TODO(),
		filter,
		options,
	)

	products := &[]entities.Product{}
	cursor.All(context.TODO(), products)

	return *products, error
}
