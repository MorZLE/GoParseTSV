package workers

import (
	"fmt"
	"github.com/MorZLE/GoParseTSV/config"
	"github.com/MorZLE/GoParseTSV/internal/model"
	"github.com/MorZLE/GoParseTSV/logger"
	"github.com/signintech/gopdf"
	"strings"
)

func NewWriter(cnf *config.Config) *Writer {
	return &Writer{dirOUT: cnf.RepOUT}
}

type Writer struct {
	dirOUT string
}

// WriteFilePDF записывает структуру в файл pdf
func (w *Writer) WriteFilePDF(guid []model.Guid, filename []string) error {
	for _, f := range filename {
		pdf := gopdf.GoPdf{}
		defer pdf.Close()

		pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

		err := pdf.AddTTFFont("LiberationSerif-Regular", "resources/LiberationSerif-Regular.ttf")
		if err != nil {
			logger.Error("AddTTFFont", err)
			return err
		}

		err = pdf.SetFont("LiberationSerif-Regular", "", 12)
		if err != nil {
			logger.Error("SetFont", err)
			return err
		}

		for _, g := range guid {
			if g.UnitGUID == f {
				pdf.AddPage()

				fields := []string{
					strings.TrimSpace(guid[0].Number) + ":" + strings.TrimSpace(g.Number),
					strings.TrimSpace(guid[0].MQTT) + ":" + strings.TrimSpace(g.MQTT),
					strings.TrimSpace(guid[0].InventoryID) + ":" + strings.TrimSpace(g.InventoryID),
					strings.TrimSpace(guid[0].MessageID) + ":" + strings.TrimSpace(g.MessageID),
					strings.TrimSpace(guid[0].MessageText) + ":" + strings.TrimSpace(g.MessageText),
					strings.TrimSpace(guid[0].Context) + ":" + strings.TrimSpace(g.Context),
					strings.TrimSpace(guid[0].MessageClass) + ":" + strings.TrimSpace(g.MessageClass),
					strings.TrimSpace(guid[0].Level) + ":" + strings.TrimSpace(g.Level),
					strings.TrimSpace(guid[0].Area) + ":" + strings.TrimSpace(g.Area),
					strings.TrimSpace(guid[0].Address) + ":" + strings.TrimSpace(g.Address),
					strings.TrimSpace(guid[0].Block) + ":" + strings.TrimSpace(g.Block),
					strings.TrimSpace(guid[0].Type) + ":" + strings.TrimSpace(g.Type),
					strings.TrimSpace(guid[0].Bit) + ":" + strings.TrimSpace(g.Bit),
					strings.TrimSpace(guid[0].InvertBit) + ":" + strings.TrimSpace(g.InvertBit),
				}

				for _, field := range fields {
					err := pdf.Text(field)
					if err != nil {
						logger.Error("Text", err)
						return err
					}
					pdf.Br(20)
				}
			}
		}

		trimmedFilename := strings.TrimSpace(f)
		outputFilename := w.dirOUT + trimmedFilename + ".pdf"

		err = pdf.WritePdf(outputFilename)
		if err != nil {
			return fmt.Errorf("ошибка при записи файла %s: %w", outputFilename, err)
		}
	}

	return nil
}
