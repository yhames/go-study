package paseto

import (
	"github.com/o1egl/paseto"
	"grpc-server/app/config"
)

type Maker struct {
	Paseto *paseto.V2
	Key    []byte
}

func NewPasetoMaker(config *config.Config) *Maker {
	return &Maker{
		Paseto: paseto.NewV2(),
		Key:    []byte(*config.Paseto.Key),
	}
}

func (m *Maker) CreateToken(userID string, duration int64) (string, error) {
	return "", nil
}

func (m *Maker) VerifyToken(token string) error {
	return nil
}
