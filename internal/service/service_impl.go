package service

import (
	"encoding/csv"
	"fmt"
	"github.com/MorZLE/ParseTSVBiocad/config"
	"github.com/MorZLE/ParseTSVBiocad/internal/model"
	"github.com/MorZLE/ParseTSVBiocad/internal/repository"
	"github.com/MorZLE/ParseTSVBiocad/internal/watcher"
	"github.com/MorZLE/ParseTSVBiocad/logger"
	"github.com/jung-kurt/gofpdf"
	"os"
)

func NewServiceImpl(cnf *config.Config, rep repository.Repository, watcher *watcher.Watcher) Service {
	return &serviceImpl{r: rep, w: watcher, dirIN: cnf.RepIN, dirOUT: cnf.RepOUT}
}

type serviceImpl struct {
	r      repository.Repository
	w      *watcher.Watcher
	dirIN  string
	dirOUT string
}

func (s *serviceImpl) Scan() {
	files := make(chan string)
	defer close(files)

	go func() {
		s.w.Scan(files)
	}()
	go func() {
		for {
			select {
			case file := <-files:
				parsGuid, slGuid, err := s.parse(file)
				if err != nil {
					logger.Error("ошибка:", err)
					err := s.r.SetError(file, err)
					if err != nil {
						logger.Error("ошибка при записи ошибки:", err)
					}
					continue
				}
				err = s.r.Set(parsGuid)
				if err != nil {
					logger.Error("ошибка:", err)
					err := s.r.SetError(file, err)
					if err != nil {
						logger.Error("ошибка при записи ошибки:", err)
					}
					continue
				}
				err = s.writeFilePDF(parsGuid, slGuid)
				if err != nil {
					logger.Error("ошибка:", err)
					err := s.r.SetError(file, err)
					if err != nil {
						logger.Error("ошибка при создании файла:", err)
					}
					continue
				}
			}
		}
	}()
	<-make(chan struct{})
}

func (s *serviceImpl) writeFilePDF(guid []model.Guid, filename []string) error {

	for _, f := range filename {
		pdf := gofpdf.New("P", "mm", "A4", f)
		pdf.SetFont("Arial", "B", 16)
		for _, g := range guid {
			if g.UnitGUID == f {
				pdf.Cell(30, 10, g.Number)
				pdf.Cell(60, 10, g.MQTT)
				pdf.Cell(40, 10, g.InventoryID)
				pdf.Cell(40, 10, g.UnitGUID)
				pdf.Cell(40, 10, g.MessageID)
				pdf.Cell(30, 10, g.MessageText)
				pdf.Cell(60, 10, g.Context)
				pdf.Cell(40, 10, g.MessageClass)
				pdf.Cell(40, 10, g.Level)
				pdf.Cell(40, 10, g.Area)
				pdf.Cell(30, 10, g.Address)
				pdf.Cell(60, 10, g.Block)
				pdf.Cell(40, 10, g.Type)
				pdf.Cell(40, 10, g.Bit)
				pdf.Cell(40, 10, g.InvertBit)
				pdf.Ln(-1)
			}
			pdf.OutputFileAndClose(s.dirOUT + f + ".pdf")
		}
	}
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
		return nil, nil, fmt.Errorf("ошибка при чтение файла %s", filename, err)
	}

	var guid []model.Guid

	for _, record := range records {
		unitGUID := record[4]

		if !mGuid[unitGUID] {
			mGuid[unitGUID] = true
			slGuid = append(slGuid, unitGUID)
		}

		g := model.Guid{
			ID:           record[0],
			Number:       record[1],
			MQTT:         record[2],
			InventoryID:  record[3],
			UnitGUID:     unitGUID,
			MessageID:    record[5],
			MessageText:  record[6],
			Context:      record[7],
			MessageClass: record[8],
			Level:        record[9],
			Area:         record[10],
			Address:      record[11],
			Block:        record[12],
			Type:         record[13],
			Bit:          record[14],
			InvertBit:    record[15],
		}
		guid = append(guid, g)
	}
	return guid, slGuid, nil
}
