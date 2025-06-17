package hash

import (
	"golang.org/x/crypto/bcrypt"
)

type HashInterface interface {
	Make(value string) (string, error)
	Check(hashedPassword string, stringPassword string) error
}

type Hash struct{}

func (h Hash) Make(value string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	return string(hashed), err
}

func (h Hash) Check(hashedPassword string, stringPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(stringPassword))
}

func NewHash() HashInterface {
	return &Hash{}
}
