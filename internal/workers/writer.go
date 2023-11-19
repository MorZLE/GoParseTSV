package workers

import (
	"fmt"
	"github.com/MorZLE/ParseTSVBiocad/config"
	"github.com/MorZLE/ParseTSVBiocad/internal/model"
	"github.com/jung-kurt/gofpdf"
	"strings"
)

func NewWriter(cnf *config.Config) *Writer {
	return &Writer{dirOUT: cnf.RepOUT}
}

type Writer struct {
	dirOUT string
}

func (w *Writer) WriteFilePDF(guid []model.Guid, filename []string) error {
	//TODO исправь эту ошибку (EXTRA *json.SyntaxError=invalid character '\\x00' looking for beginning of value)"

	for _, f := range filename {
		pdf := gofpdf.New("P", "mm", "A4", f)
		//pdf.AddFont("liberation", "", "../resources/LiberationSerif-Regular.ttf")
		pdf.SetFont("arial", "I", 12)

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
		outputFilename := w.dirOUT + trimmedFilename + ".pdf"
		err := pdf.OutputFileAndClose(outputFilename)
		if err != nil {
			return fmt.Errorf("ошибка при записи файла %s", outputFilename, err)
		}
	}
	return nil
}
