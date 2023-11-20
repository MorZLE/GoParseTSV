package logger

import (
	"go.uber.org/zap"
	"log"
	"os"
)

var mylog *zap.Logger = zap.NewNop()

// Initialize инициализирует собственный zap logger
func Initialize() {
	lvl, err := zap.ParseAtomicLevel("info")
	if err != nil {
		log.Fatal(err)
	}
	cfg := zap.NewProductionConfig()
	cfg.Level = lvl
	zl, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
	mylog = zl
}

func Info(info string) {
	mylog.Info("INFO", zap.String("info", info))
}

func Error(info string, err error) {
	mylog.Info(info, zap.Error(err))
}

func Fatal(info string, err error) {
	mylog.Error(info, zap.Error(err))
	os.Exit(1)
}
