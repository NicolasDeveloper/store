package interfaces

//IAdapter use case interface
type IAdapter interface {
	Execute(dto interface{}) error
}
