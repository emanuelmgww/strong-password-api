package services

import (
	"crypto/rand"
	"errors"
	"math/big"
)

func GeneratePassword() (string, error) {

	var password string

	var character = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%&=-+?/")

	for i := 0; i < 12; i++ {
		max := big.NewInt(int64(len(character)))
		randomCharacter, err := rand.Int(rand.Reader, max)

		if err != nil {
			return "", errors.New("failed to generate password")
		}

		index := randomCharacter.Int64()
		password += string(character[index])
	}

	return password, nil
}
