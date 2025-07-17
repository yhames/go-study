package config

import (
	"github.com/naoina/toml"
	"os"
)

type Config struct {
	DB struct {
		Database string
		Url      string
	}

	Kafka struct {
		Url      string
		ClientId string
	}

	Info struct {
		Port string
	}
}

func NewConfig(path string) *Config {
	c := new(Config)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	err = toml.NewDecoder(f).Decode(c)
	if err != nil {
		panic(err)
	}

	return c
}
