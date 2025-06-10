package config

import (
	"github.com/naoina/toml"
	"os"
)

type Config struct {
	Server struct {
		Port string
	}
}

func NewConfig(filePath string) *Config {
	config := new(Config)

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close() // Ensure the file is closed after reading

	err = toml.NewDecoder(file).Decode(config)
	if err != nil {
		panic(err)
	}
	return config
}
