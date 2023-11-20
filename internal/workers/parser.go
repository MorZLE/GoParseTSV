package workers

import (
	"encoding/csv"
	"fmt"
	"github.com/MorZLE/GoParseTSV/config"
	"github.com/MorZLE/GoParseTSV/internal/model"
	"os"
)

func NewParser(cnf *config.Config) *Parser {
	return &Parser{dirIN: cnf.RepIN}
}

type Parser struct {
	dirIN string
}

func (s *Parser) Parse(filename string) ([]model.Guid, []string, error) {

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
