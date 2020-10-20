package SessionHandling

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Functionalities/AccountCore"
)

func CheckCookie(cookie *string, cookieType int) (int, string, int, bool) {
	rawCookie, _ := base64.StdEncoding.DecodeString(*cookie)
	splitCookie := bytes.Split(rawCookie, []byte(";"))
	if len(splitCookie) == 3 {
		hexUsername := string(splitCookie[0])
		hexSalt := string(splitCookie[1])
		hexCookieSum := string(splitCookie[2])
		decodedUsername, decodingError := hex.DecodeString(hexUsername)
		decodedUsernameString := string(decodedUsername)
		usernameExists, accountID, accountType := AccountCore.UsernameExists(&decodedUsernameString)
		if usernameExists && accountType == cookieType {
			if decodingError == nil {
				decodedSalt, decodingError := hex.DecodeString(hexSalt)
				if decodingError == nil {
					decodedCookieSum, decodingError := hex.DecodeString(hexCookieSum)
					if decodingError == nil {
						return accountID, string(decodedUsername), accountType, bytes.Equal(cookieCheckSum(decodedUsername, decodedSalt), decodedCookieSum)
					}
				}
			}
		}
	}
	return -1, "", -1, false
}
