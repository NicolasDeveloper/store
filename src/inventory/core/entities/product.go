package entities

//Product entity
type Product struct {
	ID       string `bson:"_id" json:"id"`
	Quantity int    `bson:"quantity" json:"quantity"`
	Infinity bool   `bson:"infinity" json:"infinity"`
}

//NewProduct constructor
func NewProduct(id string, qtde int) (*Product, error) {
	return &Product{
		ID:       id,
		Quantity: qtde,
		Infinity: false,
	}, nil
}

//AddQuantity add quantity of products inventory
func (i *Product) AddQuantity(qtde int) {
	i.Quantity += qtde
}

//RemoveQuantity add quantity of products inventory
func (i *Product) RemoveQuantity(qtde int) {
	i.Quantity -= qtde
}
