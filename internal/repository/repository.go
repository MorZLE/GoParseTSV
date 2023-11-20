package repository

import "github.com/MorZLE/GoParseTSV/internal/model"

// Repository интерфейс репозитория
//
//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=Repository
type Repository interface {
	Set(guid []model.Guid) error
	Get(guid string) ([]model.Guid, error)
	SetError(filename string, err error) error
	SetFileName(filename string) error
	GetFileName() ([]model.ParseFile, error)
}
