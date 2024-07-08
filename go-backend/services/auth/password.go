package auth

import (
	model "github.com/mmvergara/biscord/go-backend/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (model.HashedPassword, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return model.HashedPassword(hash), nil
}

func ComparePasswords(hashed model.HashedPassword, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}
