package AccountCore

import (
	"encoding/hex"
	"errors"
	"github.com/shoriwe/ProNet/src/CryptoTools"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/RedisDatabase/Functionalities"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates"
	"golang.org/x/crypto/sha3"
	"log"
)

func handleEmailPhoneNumberUpdate(usernameHash *string, changeEmailPhoneNumber *Templates.ChangeEmailPhoneNumber) bool {
	if UpdateEmail(&changeEmailPhoneNumber.AccountID, &changeEmailPhoneNumber.Username, &changeEmailPhoneNumber.Email) && UpdatePhoneNumber(&changeEmailPhoneNumber.AccountID, &changeEmailPhoneNumber.Username, &changeEmailPhoneNumber.PhoneNumber) {
		Functionalities.DeletePendingChangeEmailPhoneNumber(usernameHash)
		return true
	}
	return false
}

func ConfirmChangeEmailPhoneNumber(usernameHash *string, emailKey *string, phoneKey *string) bool {
	if Functionalities.IsPendingEmailPhoneNumberChange(usernameHash) {
		changeEmailPhoneNumber := new(Templates.ChangeEmailPhoneNumber)
		userInformationResult := Functionalities.GetPendingEmailPhoneNumberChange(usernameHash)
		scanError := userInformationResult.Scan(changeEmailPhoneNumber)
		if scanError == nil {
			if changeEmailPhoneNumber.Tries < 3 {
				if changeEmailPhoneNumber.EmailKey == *emailKey {
					if changeEmailPhoneNumber.PhoneKey == *phoneKey {
						return handleEmailPhoneNumberUpdate(usernameHash, changeEmailPhoneNumber)
					}
				}
			} else {
				Functionalities.DeletePendingChangeEmailPhoneNumber(usernameHash)
				return false
			}
		} else {
			log.Print(scanError)
		}
		changeEmailPhoneNumber.Tries += 1
		expiration, _ := userInformationResult.Time()
		Functionalities.UpdatePendingChangeEmailPhoneNumber(usernameHash, changeEmailPhoneNumber, expiration)
	}
	return false
}

func PrepareChangeEmailPhoneNumberData(accountID *int, username *string, emailAddress *string, countryCode *string, phoneNumber *string) (string, string, string, error) {

	usernameHashHandler := sha3.New256()
	usernameHashHandler.Write([]byte(*username))
	usernameHash := hex.EncodeToString(usernameHashHandler.Sum(nil))
	if !Functionalities.IsPendingEmailPhoneNumberChange(&usernameHash) {
		changeEmailPhoneNumber := new(Templates.ChangeEmailPhoneNumber)
		phoneKey := CryptoTools.GeneratePhoneKey()
		emailKey := CryptoTools.GenerateEmailKey()
		changeEmailPhoneNumber.AccountID = *accountID
		changeEmailPhoneNumber.Tries = 0
		changeEmailPhoneNumber.Username = *username
		changeEmailPhoneNumber.Email = *emailAddress
		changeEmailPhoneNumber.PhoneNumber = *countryCode + " " + *phoneNumber
		changeEmailPhoneNumber.PhoneKey = phoneKey
		changeEmailPhoneNumber.EmailKey = emailKey
		Functionalities.NewPendingChangeEmailPhoneNumber(&usernameHash, changeEmailPhoneNumber)
		return usernameHash, emailKey, phoneKey, nil
	}
	return "", "", "", errors.New("username is already in the pending change email and phone number")
}
