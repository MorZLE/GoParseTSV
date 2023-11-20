package constants

import "errors"

var (
	ErrEnabledData = errors.New("enabled data")   // ErrEnabledData ошибка не правильных данных
	ErrNotFound    = errors.New("data not found") // ErrNotFound ошибка нет данных для представления
)
