package container

import (
	"github.com/car-journal/api-backend/config"
	"github.com/car-journal/api-backend/lib/clock"
	"github.com/car-journal/api-backend/lib/crypt"
	"github.com/car-journal/api-backend/lib/flag"
	"github.com/car-journal/api-backend/lib/hash"
	"github.com/car-journal/api-backend/lib/logger"
	"github.com/car-journal/api-backend/lib/uuid"
)

type ClientContainer struct {
	Clock clock.ClockInterface
	Crypt crypt.CryptInterface
	Hash  hash.HashInterface
	UUID  uuid.UUIDInterface
}

func CreateClientContainer() ClientContainer {
	// config
	featureFlagCfg := config.GetFlagConfig()

	// library
	clockLib := clock.New()
	cryptLib := crypt.NewCrypt(config.Get(config.EncryptionSecret))
	flagLib := flag.NewFlag(featureFlagCfg)
	hashLib := hash.NewHash()
	logger.Init(flagLib.Flag(config.FeatureLog))

	uuidLib := uuid.New()
	return ClientContainer{
		Clock: clockLib,
		Crypt: cryptLib,
		Hash:  hashLib,
		UUID:  uuidLib,
	}
}
