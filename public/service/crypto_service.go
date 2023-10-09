package service

import "golang.org/x/crypto/bcrypt"

type ICryptoService interface {
	HashPassword(password string) (string, error)
	CompareHashAndPassword(hashedPassword, password string) error
}

type CryptoService struct{}

func NewCryptoService() ICryptoService {
	return &CryptoService{}
}

func (s *CryptoService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *CryptoService) CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
