package querys

import (
	"github.com/NicolasDeveloper/store/src/catalog/core/dtos"
	"github.com/NicolasDeveloper/store/src/catalog/core/ports"
)

//SearchProductsByRangeQuery query
type SearchProductsByRangeQuery struct {
	query ports.IProductQueryPort
}

//NewSearchProductsByRangeQuery query
func NewSearchProductsByRangeQuery() (SearchProductsByRangeQuery, error) {
	return SearchProductsByRangeQuery{}, nil
}

//Execute query
func (q *SearchProductsByRangeQuery) Execute(dto dtos.FilterProducts) ([]dtos.ProductDTO, error) {
	// products, error := q.query.Search(dto.Skip, dto.Take, dto.Filter)

	return nil, nil
}
