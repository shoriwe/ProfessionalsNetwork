package SessionHandling

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"github.com/shoriwe/ProNet/src/CryptoTools"
	"golang.org/x/crypto/sha3"
	"log"
)

var SECRET []byte

func CreateSecret() {
	chunk := make([]byte, 64)
	_, readingError := rand.Read(chunk)
	if readingError == nil {
		SECRET = chunk
	} else {
		log.Fatal(readingError)
	}
}

func cookieCheckSum(username []byte, salt []byte) []byte {
	hashHandler := sha3.New256()
	hashHandler.Write(append(salt, append(username, SECRET...)...))
	return hashHandler.Sum(nil)
}

func CreateCookie(username *string) []byte {
	salt := CryptoTools.GenerateSalt()
	hexSalt := hex.EncodeToString(salt)
	rawCookie := []byte(hex.EncodeToString([]byte(*username)) + ";" + hexSalt + ";" + hex.EncodeToString(cookieCheckSum([]byte(*username), salt)))
	rawJsonCookie := map[string]string{
		"Cookie": base64.StdEncoding.EncodeToString(rawCookie),
	}
	jsonCookie, _ := json.Marshal(rawJsonCookie)
	return jsonCookie
}
