package model

type Guid struct {
	ID           string
	Number       int    `tsv:"n"`
	MQTT         string `tsv:"mqtt"`
	InventoryID  string `tsv:"invid"`
	UnitGUID     string `tsv:"unit_guid"`
	MessageID    string `tsv:"msg_id"`
	MessageText  string `tsv:"text"`
	Context      string `tsv:"context"`
	MessageClass string `tsv:"class"`
	Level        int    `tsv:"level"`
	Area         string `tsv:"area"`
	Address      string `tsv:"addr"`
	Block        bool   `tsv:"block"`
	Type         string `tsv:"type"`
	Bit          int    `tsv:"bit"`
	InvertBit    int    `tsv:"invert_bit"`
}
