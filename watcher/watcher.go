package watcher

import (
	"github.com/MorZLE/ParseTSVBiocad/config"
	"github.com/MorZLE/ParseTSVBiocad/service"
)

func NewWatcher(cnf *config.Config, s service.Service) *Watcher {
	return &Watcher{tickerTime: cnf.Timer, dirIN: cnf.RepIN, s: s}
}

type Watcher struct {
	tickerTime int
	dirIN      string
	s          service.Service
}

func (w *Watcher) Scan() {

}
