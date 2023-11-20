package workers

import (
	"encoding/csv"
	"fmt"
	"github.com/MorZLE/GoParseTSV/config"
	"github.com/MorZLE/GoParseTSV/internal/model"
	"os"
	"strings"
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

	filename = fmt.Sprintf("%s\\%s", s.dirIN, filename)
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		return nil, nil, fmt.Errorf("ошибка при открытии файла: %v", err)
	}

	reader := csv.NewReader(file)
	reader.Comma = '\t' // Устанавливаем разделитель как табуляцию

	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, fmt.Errorf("ошибка при чтение файла %s: %v", filename, err)
	}

	var guid []model.Guid

	for i, record := range records {

		unitGUID := strings.TrimSpace(record[3])
		if i != 0 {
			if !mGuid[unitGUID] {
				mGuid[unitGUID] = true
				slGuid = append(slGuid, unitGUID)
			}
		}

		g := model.Guid{
			Number:       strings.TrimSpace(record[0]),
			MQTT:         strings.TrimSpace(record[1]),
			InventoryID:  strings.TrimSpace(record[2]),
			UnitGUID:     unitGUID,
			MessageID:    strings.TrimSpace(record[4]),
			MessageText:  strings.TrimSpace(record[5]),
			Context:      strings.TrimSpace(record[6]),
			MessageClass: strings.TrimSpace(record[7]),
			Level:        strings.TrimSpace(record[8]),
			Area:         strings.TrimSpace(record[9]),
			Address:      strings.TrimSpace(record[10]),
			Block:        strings.TrimSpace(record[11]),
			Type:         strings.TrimSpace(record[12]),
			Bit:          strings.TrimSpace(record[13]),
			InvertBit:    strings.TrimSpace(record[14]),
		}
		guid = append(guid, g)

	}
	return guid, slGuid, nil
}
