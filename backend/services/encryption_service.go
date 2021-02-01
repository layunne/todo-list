package services

import (
	"golang.org/x/crypto/bcrypt"
)

type EncryptionService interface {
	Check(hashedPassword string, password string) bool
	GetEncryption(password string) string
}

func NewEncryptionService() EncryptionService {
	return &encryptionService{}
}

type encryptionService struct {
}

func (s *encryptionService) Check(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}

func (s *encryptionService) GetEncryption(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}



