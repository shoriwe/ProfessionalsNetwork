package Account

import (
	"encoding/hex"
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/CryptoTools"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/RedisDatabase/Functionalities"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Functionalities/AccountCore"
	"github.com/shoriwe/ProNet/src/InputHandler"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"github.com/shoriwe/ProNet/src/MessageSending"
	"golang.org/x/crypto/sha3"
	"net/http"
)

func requestPasswordResetBackend(username *string, accountID *int) []byte {
	hashHandler := sha3.New256()
	hashHandler.Write([]byte(*username))
	usernameHash := hex.EncodeToString(hashHandler.Sum(nil))

	if !Functionalities.UserResetExists(&usernameHash) {
		resetKey := CryptoTools.GenerateResetKey()
		userResetInformation := new(Templates.UserResetInformation)

		userResetInformation.Username = *username
		userResetInformation.UserID = *accountID
		userResetInformation.ResetKey = resetKey

		Functionalities.SetUserResetInformation(&usernameHash, userResetInformation)

		go MessageSending.SendEmailResetURL(&usernameHash, &resetKey, &userResetInformation.UserID)

		return []byte("{\"Success\": \"An email with the steps need to reset the password was sent\"}")
	}
	return []byte("{\"Error\": \"A password reset was already requested for this account\"}")
}

func RequestPasswordReset(writer http.ResponseWriter, request *http.Request) {
	requestPasswordResetForm := new(RequestsForms.RequestPasswordResetForm)
	if _, _, _, _, isRequestValid := SessionHandling.RequestHandler(writer, request, requestPasswordResetForm, true); isRequestValid {
		if InputHandler.IsUsernameValid(&requestPasswordResetForm.Username) {
			if isValid, accountID, accountType := AccountCore.UsernameExists(&requestPasswordResetForm.Username); isValid {
				if accountType != SQLDatabase.AdministratorAccount {
					_, _ = writer.Write(requestPasswordResetBackend(&requestPasswordResetForm.Username, &accountID))
				} else {
					_, _ = writer.Write([]byte("{\"Error\": \"Can't change the password of an administrator from the web api\"}"))
				}
			} else {
				_, _ = writer.Write([]byte("{\"Error\": \"The username doesn't exists\"}"))
			}
		} else {
			_, _ = writer.Write([]byte("{\"Error\": \"The user contains illegal chars\"}"))
		}
	}
}
