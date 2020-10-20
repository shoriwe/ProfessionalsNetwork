package Templates

import "encoding/json"

type UserResetInformation struct {
	ResetKey string
	Username string
	UserID   int
}

func (userResetInformation *UserResetInformation) MarshalBinary() ([]byte, error) {
	return json.Marshal(userResetInformation)
}

func (userResetInformation *UserResetInformation) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &userResetInformation)
}

type UserInformation struct {
	Tries        int
	Name         string
	Username     string
	PasswordHash string
	Email        string
	PhoneNumber  string
	AccountType  int
	PhoneKey     string
	EmailKey     string
}

func (userInformation *UserInformation) MarshalBinary() ([]byte, error) {
	return json.Marshal(userInformation)
}

func (userInformation *UserInformation) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &userInformation)
}

type ChangeEmailPhoneNumber struct {
	Tries       int
	AccountID   int
	Username    string
	Email       string
	PhoneNumber string
	PhoneKey    string
	EmailKey    string
}

func (changeEmailPhoneNumber *ChangeEmailPhoneNumber) MarshalBinary() ([]byte, error) {
	return json.Marshal(changeEmailPhoneNumber)
}

func (changeEmailPhoneNumber *ChangeEmailPhoneNumber) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &changeEmailPhoneNumber)
}

type UserInvitation struct {
	Username       string
	ProfessionalID int
	ContractorID   int
	TeamName       string
	TeamID         int
}

func (userInvitation *UserInvitation) MarshalBinary() ([]byte, error) {
	return json.Marshal(userInvitation)
}

func (userInvitation *UserInvitation) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &userInvitation)
}

type UserAppliance struct {
	Username       string
	ProfessionalID int
	ContractorID   int
	TeamName       string
	TeamID         int
}

func (userAppliance *UserAppliance) MarshalBinary() ([]byte, error) {
	return json.Marshal(userAppliance)
}

func (userAppliance *UserAppliance) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &userAppliance)
}
