package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/MorZLE/ParseTSVBiocad/logger"
	"io"
	"os"
)

type Config struct {
	RepIN      string `json:"directory_in"`
	RepOUT     string `json:"directory_out"`
	Timer      int    `json:"refresh_interval"`
	DB         string `json:"dsn"`
	configFile string
}

func NewConfig() *Config {
	return ParseFlags()
}

func ParseFlags() *Config {

	p := Config{}
	flag.StringVar(&p.configFile, "c", "config.json", "config file")

	flag.Parse()

	if err := ReadConfig(&p); err != nil {
		logger.Fatal("error read config file:", err)
	}
	if p.RepIN == "" || p.RepOUT == "" || p.Timer <= 0 {
		logger.Fatal("error empty config file:", nil)
	}
	return &p

}

func ReadConfig(cnf *Config) error {
	file, err := os.Open(cnf.configFile)
	if err != nil {
		return fmt.Errorf("can't open %s: %v", cnf.configFile, err)
	}
	defer file.Close()

	all, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("can't read %s: %v", cnf.configFile, err)
	}

	err = json.Unmarshal(all, &cnf)
	if err != nil {
		return fmt.Errorf("can't unmarshal %s: %v", cnf.configFile, err)
	}
	return nil
}
