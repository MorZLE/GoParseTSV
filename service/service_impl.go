package service

import (
	"github.com/MorZLE/ParseTSVBiocad/config"
	"github.com/MorZLE/ParseTSVBiocad/repository"
)

func NewServiceImpl(cnf *config.Config, rep repository.Repository) Service {
	return &serviceImpl{r: rep}
}

type serviceImpl struct {
	r repository.Repository
}
