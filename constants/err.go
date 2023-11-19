package constants

import "errors"

var (
	ErrGetAllGuid  = errors.New("err get guid")
	ErrEnabledData = errors.New("enabled data")
	ErrNotFound    = errors.New("data not found")
)
