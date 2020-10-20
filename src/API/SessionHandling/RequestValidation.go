package SessionHandling

import (
	"encoding/json"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/RedisDatabase/Functionalities"
	"log"
	"net/http"
	"reflect"
)

func ValidateAPIKey(apiKey *string) bool {
	return Functionalities.IsAPIKeyValid(apiKey)
}

func ValidateAPIKeyAndCookie(apiKeyRequired bool, apiKey string, cookie string, accountTypes []int) (int, string, int, bool, bool) {
	if numberOfAccountTypesValid := len(accountTypes); numberOfAccountTypesValid > 0 {
		for index := 0; index < numberOfAccountTypesValid; index++ {
			if accountID, username, accountType, isValid := CheckCookie(&cookie, accountTypes[index]); isValid {
				if apiKeyRequired {
					return accountID, username, accountType, true, ValidateAPIKey(&apiKey)
				}
				return accountID, username, accountType, true, true
			}
		}
		if apiKeyRequired {
			return -1, "", -1, false, ValidateAPIKey(&apiKey)
		}
		return -1, "", -1, false, true
	}
	if apiKeyRequired {
		return -1, "", -1, true, ValidateAPIKey(&apiKey)
	}
	return -1, "", -1, true, true
}

func RequestHandler(writer http.ResponseWriter, request *http.Request, formTemplate interface{}, apiKeyRequired bool, availableAccountTypes ...int) (int, string, int, bool, bool) {
	statusCode := *new(int)
	message := make([]byte, 0)
	if request.Method == http.MethodPost {
		unmarshalError := json.NewDecoder(request.Body).Decode(formTemplate)
		if unmarshalError == nil {
			accountID, username, accountType, isAccountTypeCorrect, isRequestValid := ValidateAPIKeyAndCookie(
				apiKeyRequired,
				reflect.Indirect(reflect.ValueOf(formTemplate)).FieldByName("APIKey").String(),
				reflect.Indirect(reflect.ValueOf(formTemplate)).FieldByName("Cookie").String(),
				availableAccountTypes)
			if !isRequestValid {
				statusCode = http.StatusInternalServerError
				message = []byte("{\"Error\":\"Invalid API key\"}")
			} else if !isAccountTypeCorrect {
				statusCode = http.StatusInternalServerError
				message = []byte("{\"Error\":\"Invalid Cookies\"}")
			} else {
				return accountID, username, accountType, isAccountTypeCorrect, isRequestValid
			}
		} else {
			log.Print(unmarshalError)
			statusCode = http.StatusInternalServerError
			message = []byte("{\"Error\":\"Something goes wrong\"}")
		}
	} else {
		statusCode = http.StatusMethodNotAllowed
		message = []byte("{\"Error\":\"Method not allowed\"}")
	}
	writer.WriteHeader(statusCode)
	_, _ = writer.Write(message)
	return -1, "", -1, false, false
}
