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
	"strings"
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
	out := make(chan string)
	defer close(out)

	go s.w.Scan(out)
	for file := range out {
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
		logger.Info(fmt.Sprintf("файл: %s обработан", file))
	}

}

func (s *serviceImpl) writeFilePDF(guid []model.Guid, filename []string) error {
	for _, f := range filename {
		pdf := gofpdf.New("P", "mm", "A4", f)
		pdf.SetFont("Arial", "B", 8)

		for _, g := range guid {
			if g.UnitGUID == f {
				pdf.AddPage()
				pdf.Text(10, 10, strings.TrimSpace(guid[0].Number)+":")
				pdf.Text(10, 20, strings.TrimSpace(guid[0].MQTT)+":")
				pdf.Text(10, 30, strings.TrimSpace(guid[0].InventoryID)+":")
				pdf.Text(10, 40, strings.TrimSpace(guid[0].MessageID)+":")
				pdf.Text(10, 50, strings.TrimSpace(guid[0].MessageText)+":")
				pdf.Text(10, 60, strings.TrimSpace(guid[0].Context)+":")
				pdf.Text(10, 70, strings.TrimSpace(guid[0].MessageClass)+":")
				pdf.Text(10, 80, strings.TrimSpace(guid[0].Level)+":")
				pdf.Text(10, 90, strings.TrimSpace(guid[0].Area)+":")
				pdf.Text(10, 100, strings.TrimSpace(guid[0].Address)+":")
				pdf.Text(10, 120, strings.TrimSpace(guid[0].Block)+":")
				pdf.Text(10, 130, strings.TrimSpace(guid[0].Type)+":")
				pdf.Text(10, 140, strings.TrimSpace(guid[0].Bit)+":")
				pdf.Text(10, 150, strings.TrimSpace(guid[0].InvertBit)+":")

				pdf.Text(30, 10, strings.TrimSpace(g.Number))
				pdf.Text(30, 20, strings.TrimSpace(g.MQTT))
				pdf.Text(30, 30, strings.TrimSpace(g.InventoryID))
				pdf.Text(30, 40, strings.TrimSpace(g.MessageID))
				pdf.Text(30, 50, strings.TrimSpace(g.MessageText))
				pdf.Text(30, 60, strings.TrimSpace(g.Context))
				pdf.Text(30, 70, strings.TrimSpace(g.MessageClass))
				pdf.Text(30, 80, strings.TrimSpace(g.Level))
				pdf.Text(30, 90, strings.TrimSpace(g.Area))
				pdf.Text(30, 100, strings.TrimSpace(g.Address))
				pdf.Text(30, 120, strings.TrimSpace(g.Block))
				pdf.Text(30, 130, strings.TrimSpace(g.Type))
				pdf.Text(30, 150, strings.TrimSpace(g.InvertBit))
			}
		}

		trimmedFilename := strings.TrimSpace(f)
		outputFilename := s.dirOUT + trimmedFilename + ".pdf"
		pdf.OutputFileAndClose(outputFilename)
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

	for i, record := range records {

		unitGUID := record[3]
		if i != 0 {
			if !mGuid[unitGUID] {
				mGuid[unitGUID] = true
				slGuid = append(slGuid, unitGUID)
			}
		}

		g := model.Guid{
			Number:       record[0],
			MQTT:         record[1],
			InventoryID:  record[2],
			UnitGUID:     unitGUID,
			MessageID:    record[4],
			MessageText:  record[5],
			Context:      record[6],
			MessageClass: record[7],
			Level:        record[8],
			Area:         record[9],
			Address:      record[10],
			Block:        record[11],
			Type:         record[12],
			Bit:          record[13],
			InvertBit:    record[14],
		}
		guid = append(guid, g)

	}
	return guid, slGuid, nil
}
