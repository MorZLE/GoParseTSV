package service

import (
	"fmt"
	"github.com/MorZLE/GoParseTSV/constants"
	"github.com/MorZLE/GoParseTSV/internal/model"
	"github.com/MorZLE/GoParseTSV/internal/repository"
	"github.com/MorZLE/GoParseTSV/internal/workers"
	"github.com/MorZLE/GoParseTSV/logger"
)

func NewServiceImpl(rep repository.Repository, watcher *workers.Watcher, writer *workers.Writer, parser *workers.Parser) Service {
	return &serviceImpl{r: rep, wSkaner: watcher, wWriter: writer, wParser: parser}
}

type serviceImpl struct {
	r       repository.Repository
	wSkaner *workers.Watcher
	wWriter *workers.Writer
	wParser *workers.Parser
}

func (s *serviceImpl) Scan() {
	filesCheck, err := s.r.GetFileName()
	if err != nil {
		logger.Fatal("", err)
		return
	}
	s.wSkaner.InitFileCheck(filesCheck)

	out := make(chan string)
	defer close(out)

	go s.wSkaner.Scan(out)
	for file := range out {
		parsGuid, slGuid, err := s.wParser.Parse(file)
		if err != nil {
			logger.Error("ошибка:", err)
			err := s.r.SetError(file, err)
			if err != nil {
				logger.Error("ошибка при записи ошибки:", err)
			}
			continue
		}
		err = s.r.Set(parsGuid[1:])
		if err != nil {
			logger.Error("ошибка:", err)
			err := s.r.SetError(file, err)
			if err != nil {
				logger.Error("ошибка при записи ошибки:", err)
			}
			continue
		}
		err = s.wWriter.WriteFilePDF(parsGuid, slGuid)
		if err != nil {
			logger.Error("ошибка:", err)
			err := s.r.SetError(file, err)
			if err != nil {
				logger.Error("ошибка при создании файла:", err)
			}
			continue
		}
		err = s.r.SetFileName(file)
		if err != nil {
			logger.Error("ошибка внесения имени обработанного файла в базу:", err)
		} else {
			logger.Info(fmt.Sprintf("файл: %s обработан", file))
		}
	}
}

func (s *serviceImpl) GetAllGuid(req model.RequestGetGuid) ([][]model.Guid, error) {
	if req.UnitGUID == "" || req.Limite <= 0 || req.Page < 0 {
		return nil, constants.ErrEnabledData
	}

	guids, err := s.r.Get(req.UnitGUID)
	if err != nil {
		return nil, err
	}

	var guidsRes [][]model.Guid
	var gd []model.Guid
	cPage := 1

	for i, guid := range guids {
		if i < req.Page {
			continue
		}
		if cPage > req.Limite {
			guidsRes = append(guidsRes, gd)
			gd = nil
			gd = append(gd, guid)
			cPage = 2
			continue
		}
		if cPage <= req.Limite {
			gd = append(gd, guid)
		}
		cPage++
	}
	if gd != nil {
		guidsRes = append(guidsRes, gd)
	}

	return guidsRes, nil
}
