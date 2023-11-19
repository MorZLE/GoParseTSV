package repository

import "github.com/MorZLE/ParseTSVBiocad/internal/model"

type Repository interface {
	Set(guid []model.Guid) error
	Get(guid string) ([]model.Guid, error)
	SetError(filename string, err error) error
	SetFileName(filename string) error
	GetFileName() ([]model.ParseFile, error)
}
