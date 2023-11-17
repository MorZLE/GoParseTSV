package config

import (
	"flag"
	"github.com/MorZLE/ParseTSVBiocad/logger"
	"os"
	"strconv"
)

type Config struct {
	RepIN  string
	RepOUT string
	Timer  int
}

func NewConfig() *Config {
	cnf := &Config{}
	return ParseFlags(cnf)
}

func ParseFlags(p *Config) *Config {

	flag.StringVar(&p.RepIN, "a", "", "directory with .tsv files")
	flag.StringVar(&p.RepIN, "b", "", "directory output .pdf files")
	flag.StringVar(&p.RepIN, "t", "1", "time to scan directory")

	flag.Parse()

	if RepIN := os.Getenv("REPIN"); RepIN != "" {
		p.RepIN = RepIN
	}
	if p.RepIN == "" {
		logger.Fatal("directory with .tsv files is empty", nil)
	}
	if RepOUT := os.Getenv("REPOUT"); RepOUT != "" {
		p.RepOUT = RepOUT
	}
	if p.RepOUT == "" {
		logger.Fatal("directory output .tsv files is empty", nil)
	}
	if Timer := os.Getenv("TIMERSCAN"); Timer != "" {
		Timer, err := strconv.Atoi(Timer)
		if err != nil {
			p.Timer = 1
			logger.Error("err get env TIMERSCAN:", err)
			logger.Info("default TIMERSCAN = 1")
		}
		p.Timer = Timer
	}
	return p

}
