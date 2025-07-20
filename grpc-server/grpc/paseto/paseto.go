package paseto

import (
	"crypto/rand"
	"github.com/o1egl/paseto"
	"grpc-server/app/config"
	auth "grpc-server/grpc/proto"
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

func (m *Maker) CreateToken(auth *auth.AuthData) (string, error) {
	randomBytes := make([]byte, 16)
	n, err := rand.Read(randomBytes)
	if err != nil || n != len(randomBytes) {
		return "", err
	}
	return m.Paseto.Encrypt(m.Key, auth, randomBytes)
}

func (m *Maker) VerifyToken(token string) error {
	var authData *auth.AuthData
	return m.Paseto.Decrypt(token, m.Key, &authData, nil)
}
