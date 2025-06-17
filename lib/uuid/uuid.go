package uuid

import (
	gofrsUUID "github.com/gofrs/uuid"
)

type UUIDInterface interface {
	GenerateNewUUID() UUID
}

type uuid struct{}

func New() UUIDInterface {
	return &uuid{}
}

func (u uuid) GenerateNewUUID() UUID {
	return UUID{
		UUID: gofrsUUID.Must(gofrsUUID.NewV4()),
	}
}
