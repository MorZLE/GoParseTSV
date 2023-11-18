package model

type Guid struct {
	ID           string `gorm:"primary_key" gorm:"AUTO_INCREMENT"`
	Number       string `tsv:"n"`
	MQTT         string `tsv:"mqtt"`
	InventoryID  string `tsv:"invid"`
	UnitGUID     string `tsv:"unit_guid" gorm:"not null"`
	MessageID    string `tsv:"msg_id"`
	MessageText  string `tsv:"text"`
	Context      string `tsv:"context"`
	MessageClass string `tsv:"class"`
	Level        string `tsv:"level"`
	Area         string `tsv:"area"`
	Address      string `tsv:"addr"`
	Block        string `tsv:"block"`
	Type         string `tsv:"type"`
	Bit          string `tsv:"bit"`
	InvertBit    string `tsv:"invert_bit"`
}

type Err struct {
	File string
	Err  error
}