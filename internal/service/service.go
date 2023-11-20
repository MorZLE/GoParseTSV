package service

import "github.com/MorZLE/GoParseTSV/internal/model"

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=Service
type Service interface {
	Scan()
	GetAllGuid(guid model.RequestGetGuid) ([][]model.Guid, error)
}
