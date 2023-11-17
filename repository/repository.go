package repository

type Repository interface {
	Set(interface{}) (interface{}, error)
	Get(interface{}) error
}
