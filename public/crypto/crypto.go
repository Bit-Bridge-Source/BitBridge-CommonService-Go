package common_crypto

import "golang.org/x/crypto/bcrypt"

type ICrypto interface {
	GenerateFromPassword(password string) (string, error)
	CompareHashAndPassword(hashedPassword, password string) error
}

type Crypto struct{}

func NewCrypto() ICrypto {
	return &Crypto{}
}

func (s *Crypto) GenerateFromPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *Crypto) CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
