package repository

import "github.com/MorZLE/ParseTSVBiocad/internal/model"

type Repository interface {
	Set(guid []model.Guid) error
	Get(interface{}) error
	SetError(filename string, err error) error
}
