package config

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type PasswordHash struct {
}

func NewPasswordHash() *PasswordHash {
	return &PasswordHash{}
}

func (pss *PasswordHash) Generate(defaultPassword string) (*[]byte, bool) {
	salt := []byte{
		0x8a, 0xe3, 0x0b, 0x12, 0x4f, 0xf7, 0x6d, 0x2a,
		0xa0, 0x56, 0x9c, 0x33, 0xd4, 0x71, 0x8e, 0xbb,
	}
	passwordWithSalt := append([]byte(defaultPassword), salt...)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordWithSalt, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}

	return &hashedPassword, true
}

func (pss *PasswordHash) Compare(hash, defaultPass string) bool {
	salt := []byte{
		0x8a, 0xe3, 0x0b, 0x12, 0x4f, 0xf7, 0x6d, 0x2a,
		0xa0, 0x56, 0x9c, 0x33, 0xd4, 0x71, 0x8e, 0xbb,
	}
	passwordWithSalt := append([]byte(defaultPass), salt...)
	err := bcrypt.CompareHashAndPassword([]byte(hash), passwordWithSalt)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
