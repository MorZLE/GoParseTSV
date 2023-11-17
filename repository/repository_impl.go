package repository

import "github.com/MorZLE/ParseTSVBiocad/config"

func NewRepositoryImpl(cnf *config.Config) Repository {
	return &repositoryImpl{}
}

type repositoryImpl struct {
}

func (r *repositoryImpl) Get(interface{}) error {
	return nil
}
func (r *repositoryImpl) Set(interface{}) (interface{}, error) {
	return nil, nil
}
