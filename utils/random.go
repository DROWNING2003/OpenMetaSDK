package utils

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/google/uuid"
)

func NonceRandom() string {
	token := make([]byte, 24)
	rand.Read(token)
	return base64.StdEncoding.EncodeToString(token)
}
func GenerateOperationId() string {
	return uuid.New().String()
}
