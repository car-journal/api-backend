package container

import (
	"github.com/car-journal/config"
	"github.com/car-journal/lib/clock"
	"github.com/car-journal/lib/crypt"
	"github.com/car-journal/lib/hash"
	"github.com/car-journal/lib/uuid"
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
	cryptLib := crypt.NewCrypt(config.Get(config.ENCRYPTION_SECRET))
	hashLib := hash.NewHash()
	return ClientContainer{
		Clock: clockLib,
		Crypt: cryptLib,
		Hash:  hashLib,
		UUID:  uuidLib,
	}
}
