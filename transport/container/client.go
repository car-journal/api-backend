package container

import (
	"github.com/car-journal/api-backend/config"
	"github.com/car-journal/api-backend/lib/clock"
	"github.com/car-journal/api-backend/lib/crypt"
	"github.com/car-journal/api-backend/lib/hash"
	"github.com/car-journal/api-backend/lib/uuid"
)

type ClientContainer struct {
	Clock clock.ClockInterface
	Crypt crypt.CryptInterface
	Hash  hash.HashInterface
	UUID  uuid.UUIDInterface
}

func CreateClientContainer() ClientContainer {
	// library

	clockLib := clock.New()
	uuidLib := uuid.New()
	cryptLib := crypt.NewCrypt(config.Get(config.EncryptionSecret))
	hashLib := hash.NewHash()
	return ClientContainer{
		Clock: clockLib,
		Crypt: cryptLib,
		Hash:  hashLib,
		UUID:  uuidLib,
	}
}
