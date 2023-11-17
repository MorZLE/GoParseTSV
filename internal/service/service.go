package service

import "github.com/MorZLE/ParseTSVBiocad/internal/model"

type Service interface {
	Scan()
	parse(filename string) ([]model.Guid, []string, error)
}
