package controllers

import (
	"net/http"

	usecases "github.com/NicolasDeveloper/store/src/catalog/core/use-cases"
	"github.com/golobby/container"

	"github.com/NicolasDeveloper/store/src/catalog/core/dtos"

	"github.com/NicolasDeveloper/store/src/catalog/driver-adapters/api/common"
)

//ProductController controller
type ProductController struct {
	common.Controller
}

//NewProductController contructor
func NewProductController() ProductController {
	return ProductController{}
}

//Index get products in database
func (c *ProductController) Index(w http.ResponseWriter, r *http.Request) {

}

//Post insert product
func (c *ProductController) Post(w http.ResponseWriter, r *http.Request) {
	var usecase *usecases.RegisterUseCase
	container.Make(&usecase)

	dto := &dtos.ProductDTO{}
	c.GetContent(dto, r)

	usecase.Execute(dto)

	c.SendJSON(
		w,
		dto,
		http.StatusOK,
	)
}
