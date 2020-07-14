package dtos

//FilterProducts filter
type FilterProducts struct {
	Filter interface{} `json:"filter"`
	Skip   int64       `json:"skip"`
	Take   int64       `json:"take"`
}
