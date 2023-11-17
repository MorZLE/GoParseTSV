package repository

import (
	"github.com/MorZLE/ParseTSVBiocad/config"
	"github.com/MorZLE/ParseTSVBiocad/internal/model"
)

func NewRepositoryImpl(cnf *config.Config) Repository {
	return &repositoryImpl{}
}

type repositoryImpl struct {
}

func (r *repositoryImpl) Get(interface{}) error {
	return nil
}
func (r *repositoryImpl) Set(guid []model.Guid) error {
	return nil
}

func (r *repositoryImpl) SetError(filename string, err error) error {
	return nil

}
