package config

import (
	"flag"
)

type Config struct {
	ServAddr string
	DBAddr   string
}

func NewConfig() *Config {

	var cfg Config

	flag.StringVar(&cfg.ServAddr, "a",
		"127.0.0.1:8080", "server address",
	)

	flag.StringVar(&cfg.DBAddr, "d",
		"127.0.0.1:8080", "database address",
	)

	flag.Parse()

	return &cfg
}
