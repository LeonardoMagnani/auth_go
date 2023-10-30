package helpers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
)

func GenerateHash() (string, int) {
	value := rand.Intn(9000) + 1000
	secret := os.Getenv("SECRET_HASH")

	combined := fmt.Sprintf("%d%s", value, secret)
	hash := sha256.New()
	hash.Write([]byte(combined))
	hashInBytes := hex.EncodeToString(hash.Sum(nil))

	return hashInBytes, value
}

func Hash(code int) string {
	secret := os.Getenv("SECRET_HASH")

	combined := fmt.Sprintf("%d%s", code, secret)
	hash := sha256.New()
	hash.Write([]byte(combined))
	return hex.EncodeToString(hash.Sum(nil))
}

func HashFromString(value string) string {
	secret := os.Getenv("PASSWORD_HASH")

	combined := value + secret
	hash := sha256.New()
	hash.Write([]byte(combined))
	return hex.EncodeToString(hash.Sum(nil))
}
