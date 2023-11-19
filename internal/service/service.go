package service

import "github.com/MorZLE/ParseTSVBiocad/internal/model"

type Service interface {
	Scan()
	GetAllGuid(guid model.RequestGetGuid) ([][]model.Guid, error)
}
