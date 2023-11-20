package model

// Guid - структура guid
type Guid struct {
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

// Err - структура ошибок
type Err struct {
	File string
	Err  string
}

// ParseFile - структура обработанных файлов
type ParseFile struct {
	File string
}

// RequestGetGuid - структура запроса API
type RequestGetGuid struct {
	UnitGUID string `json:"unitguid"`
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
}
