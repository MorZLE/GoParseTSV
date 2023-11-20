package workers

import (
	"github.com/MorZLE/GoParseTSV/config"
	"github.com/MorZLE/GoParseTSV/internal/model"
	"github.com/MorZLE/GoParseTSV/logger"
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

func (w *Watcher) InitFileCheck(files []model.ParseFile) {
	for _, file := range files {
		w.fileCheck[file.File] = true
	}
}

func (w *Watcher) Scan(out chan string) {
	tick := time.NewTicker(time.Duration(w.tickerTime) * time.Second)
	defer tick.Stop()

	for range tick.C {
		files, err := os.ReadDir(w.dirIN)
		if err != nil {
			logger.Fatal("err read dir", err)
		}

		for _, file := range files {
			if !file.IsDir() && len(file.Name()) >= 4 && strings.HasSuffix(file.Name(), ".tsv") {
				w.mutex.Lock()
				if w.fileCheck[file.Name()] {
					w.mutex.Unlock()
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
