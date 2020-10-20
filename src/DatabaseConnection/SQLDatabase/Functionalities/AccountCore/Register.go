package AccountCore

import (
	"encoding/hex"
	"errors"
	"github.com/shoriwe/ProNet/src/CryptoTools"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/AccountCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/RedisDatabase/Functionalities"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Queries"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates"
	"golang.org/x/crypto/sha3"
	"log"
)

func handleRegistration(userInformation *Templates.UserInformation) bool {
	_, connectionError := SQLDatabase.ExecuteInsertQuery(Queries.CreateUserQuery,
		userInformation.Name, userInformation.Username,
		userInformation.PasswordHash, userInformation.Email,
		userInformation.PhoneNumber, userInformation.AccountType)
	if connectionError == nil {
		usernameHashHandler := sha3.New256()
		usernameHashHandler.Write([]byte(userInformation.Username))
		usernameHash := hex.EncodeToString(usernameHashHandler.Sum(nil))

		_, userID, _ := UsernameExists(&userInformation.Username)
		switch userInformation.AccountType {
		case SQLDatabase.ProfessionalAccount:
			AccountCore.CreateProfessionalNode(&userInformation.Name, &userID)
		case SQLDatabase.ContractorAccount:
			AccountCore.CreateContractorNode(&userInformation.Name, &userID)
		}
		Functionalities.DeletePendingRegistration(&usernameHash)
		return true
	} else {
		log.Print(connectionError)
	}
	return false
}

func ConfirmAccountRegisterBackend(usernameHash *string, emailKey *string, phoneKey *string) bool {
	if pendingRegistrationAlreadyExists := Functionalities.IsPendingRegistration(usernameHash); pendingRegistrationAlreadyExists {
		userInformation := new(Templates.UserInformation)
		userInformationResult := Functionalities.GetPendingRegistration(usernameHash)
		scanError := userInformationResult.Scan(userInformation)
		if scanError == nil {
			if userInformation.Tries < 3 {
				if userInformation.EmailKey == *emailKey {
					if userInformation.PhoneKey == *phoneKey {
						return handleRegistration(userInformation)
					}
				}
			} else {
				Functionalities.DeletePendingRegistration(usernameHash)
				return false
			}
		} else {
			log.Print(scanError)
		}
		userInformation.Tries += 1
		expiration, _ := userInformationResult.Time()
		Functionalities.UpdatePendingRegistration(usernameHash, userInformation, expiration)
	}
	return false
}

func PrepareRegistrationData(name *string, username *string, password *string, emailAddress *string, countryCode *string, phoneNumber *string, accountType int) (string, string, string, error) {
	usernameHashHandler := sha3.New256()
	usernameHashHandler.Write([]byte(*username))
	usernameHash := hex.EncodeToString(usernameHashHandler.Sum(nil))

	if pendingRegistrationAlreadyExists := Functionalities.IsPendingRegistration(&usernameHash); !pendingRegistrationAlreadyExists {
		userInformation := new(Templates.UserInformation)
		phoneKey := CryptoTools.GeneratePhoneKey()
		emailKey := CryptoTools.GenerateEmailKey()
		passwordHashHandler := sha3.New256()
		passwordHashHandler.Write([]byte(*password))
		userInformation.Tries = 0
		userInformation.Name = *name
		userInformation.Username = *username
		userInformation.PasswordHash = hex.EncodeToString(passwordHashHandler.Sum(nil))
		userInformation.Email = *emailAddress
		userInformation.PhoneNumber = *countryCode + " " + *phoneNumber
		userInformation.AccountType = accountType
		userInformation.PhoneKey = phoneKey
		userInformation.EmailKey = emailKey
		Functionalities.NewPendingRegistration(&usernameHash, userInformation)
		return usernameHash, emailKey, phoneKey, nil
	}
	return "", "", "", errors.New("username is already in the registration queue")
}
