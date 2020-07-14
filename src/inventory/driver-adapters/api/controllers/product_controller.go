package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/store/src/inventory/core/dtos"
	usecases "github.com/NicolasDeveloper/store/src/inventory/core/use-cases"
	"github.com/golobby/container"

	"github.com/NicolasDeveloper/store/src/inventory/driver-adapters/api/common"
)

//ProductController controller
type ProductController struct {
	common.Controller
}

//NewProductController contructor
func NewProductController() ProductController {
	return ProductController{}
}

type response struct {
	Version string
}

//Index get products in database
func (c *ProductController) Index(w http.ResponseWriter, r *http.Request) {
	teste := response{
		Version: "v1",
	}

	c.SendJSON(
		w,
		teste,
		http.StatusOK,
	)
}

//Post insert product
func (c *ProductController) Post(w http.ResponseWriter, r *http.Request) {
	var usecase *usecases.RegisterProductQuantity
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
