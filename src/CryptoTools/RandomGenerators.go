package CryptoTools

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateSalt() []byte {
	salt := make([]byte, 16)
	_, readingError := rand.Read(salt)
	if readingError == nil {
		return salt
	}
	return nil
}

func generateKey(length int) []byte {
	result := make([]byte, length)
	_, readingError := rand.Read(result)
	if readingError == nil {
		return result
	}
	return nil
}

func GenerateEmailKey() string {
	return hex.EncodeToString(generateKey(32))
}

func GeneratePhoneKey() string {
	return hex.EncodeToString(generateKey(6))
}

func GenerateResetKey() string {
	return hex.EncodeToString(generateKey(32))
}
