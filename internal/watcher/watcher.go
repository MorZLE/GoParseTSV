package watcher

import (
	"github.com/MorZLE/ParseTSVBiocad/config"
	"github.com/MorZLE/ParseTSVBiocad/logger"
	"os"
	"strings"
	"sync"
	"time"
)

func NewWatcher(cnf *config.Config) *Watcher {
	m := make(map[string]bool)
	return &Watcher{tickerTime: cnf.Timer, dirIN: cnf.RepIN, fileCheck: m}
}

type Watcher struct {
	tickerTime int
	dirIN      string
	fileCheck  map[string]bool
	mutex      sync.RWMutex
}

func (w *Watcher) Scan(out chan string) {
	//просканировать директорию и проверить на наличие файлов
	files, err := os.ReadDir(w.dirIN)
	if err != nil {
		logger.Error("err read dir", err)
	}

	tick := time.NewTicker(time.Duration(w.tickerTime) * time.Second)
	defer tick.Stop()

	select {
	case <-tick.C:
		for _, file := range files {
			//проверка, что файл не директория
			if file.IsDir() {
				continue
			}

			//проверка, что есть расширение
			if len(file.Name()) < 4 {
				continue
			}
			//проверка, что файл .tsv
			if strings.HasSuffix(file.Name(), ".tsv") {
				w.mutex.Lock()
				if w.fileCheck[file.Name()] {
					continue
				}
				w.fileCheck[file.Name()] = true
				w.mutex.Unlock()
				select {
				case out <- file.Name():
				}
			}
		}
	}

}
