package dtos

//FilterProducts filter
type FilterProducts struct {
	Filter interface{}
	Skip   int64
	Take   int64
}
