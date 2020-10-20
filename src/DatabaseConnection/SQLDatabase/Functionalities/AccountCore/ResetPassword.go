package AccountCore

import (
	"github.com/shoriwe/ProNet/src/DatabaseConnection/RedisDatabase/Functionalities"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates"
	"log"
)

func ResetPassword(usernameHash *string, resetKey *string, newPassword *string) (bool, []byte, int, string) {
	userResetInformation := new(Templates.UserResetInformation)
	result := Functionalities.GetUserResetInformation(usernameHash)
	scanError := result.Scan(userResetInformation)
	if scanError == nil {
		if userResetInformation.ResetKey == *resetKey {
			updated, message := UpdatePassword(&userResetInformation.UserID, &userResetInformation.Username, newPassword)
			if updated {
				Functionalities.DeleteUserResetInformation(usernameHash)
			}
			return updated, message, userResetInformation.UserID, userResetInformation.Username
		}
		return false, []byte("{\"Error\": \"The Reset key is wrong\"}"), -1, ""
	}
	log.Print(scanError)
	return false, []byte("{\"Error\": \"Something goes wrong\"}"), -1, ""
}
