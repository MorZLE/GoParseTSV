package service

import (
	"encoding/csv"
	"fmt"
	"github.com/MorZLE/ParseTSVBiocad/config"
	"github.com/MorZLE/ParseTSVBiocad/internal/model"
	"github.com/MorZLE/ParseTSVBiocad/internal/repository"
	"github.com/MorZLE/ParseTSVBiocad/internal/watcher"
	"github.com/MorZLE/ParseTSVBiocad/logger"
	"os"
	"strconv"
)

func NewServiceImpl(cnf *config.Config, rep repository.Repository, watcher *watcher.Watcher) Service {
	return &serviceImpl{r: rep, w: watcher, dirIN: cnf.RepIN}
}

type serviceImpl struct {
	r     repository.Repository
	w     *watcher.Watcher
	dirIN string
}

func (s *serviceImpl) Scan() {
	files := make(chan string)

	defer close(files)
	go s.w.Scan(files)

	select {
	case <-files:
		file := <-files
		parsGuid, slGuid, err := s.parse(file)
		if err != nil {
			logger.Error("ошибка:", err)
			err := s.r.SetError(file, err)
			if err != nil {
				logger.Error("ошибка при записи ошибки:", err)
			}
		}
		err = s.r.Set(parsGuid)
		if err != nil {
			logger.Error("ошибка:", err)
			err := s.r.SetError(file, err)
			if err != nil {
				logger.Error("ошибка при записи ошибки:", err)
			}
		}
		err = s.writeFilePDF(parsGuid, slGuid)
		if err != nil {
			logger.Error("ошибка:", err)
			err := s.r.SetError(file, err)
			if err != nil {
				logger.Error("ошибка при создании файла:", err)
			}
		}
	}

}

func (s *serviceImpl) writeFilePDF(guid []model.Guid, filename []string) error {

	return nil
}

func (s *serviceImpl) parse(filename string) ([]model.Guid, []string, error) {

	var slGuid []string
	mGuid := make(map[string]bool)

	filename = fmt.Sprintf("%s/%s", s.dirIN, filename)
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		return nil, nil, fmt.Errorf("ошибка при открытии файла", err)
	}

	reader := csv.NewReader(file)
	reader.Comma = '\t' // Устанавливаем разделитель как табуляцию

	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, fmt.Errorf("ошибка при чтение файла", err)
	}

	var guid []model.Guid

	for _, record := range records {
		number, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, nil, fmt.Errorf("ошибка типа поля int", err)
		}
		level, err := strconv.Atoi(record[9])
		if err != nil {
			return nil, nil, fmt.Errorf("ошибка типа поля int", err)
		}
		block, err := strconv.ParseBool(record[12])
		if err != nil {
			return nil, nil, fmt.Errorf("ошибка поля Block", err)
		}
		bit, err := strconv.Atoi(record[14])
		if err != nil {
			return nil, nil, fmt.Errorf("ошибка типа поля int", err)
		}
		invertBit, err := strconv.Atoi(record[15])
		if err != nil {
			return nil, nil, fmt.Errorf("ошибка типа поля int", err)
		}
		unitGUID := record[4]

		if !mGuid[unitGUID] {
			mGuid[unitGUID] = true
			slGuid = append(slGuid, unitGUID)
		}

		g := model.Guid{
			ID:           record[0],
			Number:       number,
			MQTT:         record[2],
			InventoryID:  record[3],
			UnitGUID:     unitGUID,
			MessageID:    record[5],
			MessageText:  record[6],
			Context:      record[7],
			MessageClass: record[8],
			Level:        level,
			Area:         record[10],
			Address:      record[11],
			Block:        block,
			Type:         record[13],
			Bit:          bit,
			InvertBit:    invertBit,
		}
		guid = append(guid, g)
	}
	return guid, slGuid, nil
}
