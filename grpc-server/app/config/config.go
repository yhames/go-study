package config

import (
	"github.com/naoina/toml"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Port *string
	}
	Paseto struct {
		Key *string
	}

	Grpc struct {
		Url *string
	}
}

func NewConfig(path string) *Config {
	config := new(Config)
	log.Println(path)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)

	err = toml.NewDecoder(file).Decode(config)
	if err != nil {
		panic(err)
	}

	return config
}
