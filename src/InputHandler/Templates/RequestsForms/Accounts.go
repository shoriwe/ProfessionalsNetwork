package RequestsForms

type RequestTeamMembersForm struct {
	APIKey       string
	Cookie       string
	ContractorID int
	TeamName     string
}

type RemoveProfessionalFromTeamForm struct {
	APIKey         string
	Cookie         string
	ProfessionalID int
	TeamName       string
}

type RequestProfileForm struct {
	APIKey string
	Cookie string
}

type ChangeDescriptionForm struct {
	APIKey      string
	Cookie      string
	Description string
}

type ChangeLocationForm struct {
	APIKey   string
	Cookie   string
	Location string
}

type RequestChangeEmailPhoneNumberForm struct {
	APIKey      string
	Cookie      string
	Email       string
	CountryCode string
	PhoneNumber string
}

type RequestAccountForm struct {
	APIKey string
	Cookie string
}

type SearchSkillsForm struct {
	APIKey      string
	Cookie      string
	SearchQuery string
	ObjectID    int
}

type SearchLanguagesForm struct {
	APIKey      string
	Cookie      string
	SearchQuery string
	ObjectID    int
}

type RequestAccountTypeForm struct {
	APIKey string
	Cookie string
}

type ChangePhoneNumberForm struct {
	APIKey      string
	Cookie      string
	Password    string
	PhoneNumber string
}

type ChangeEmailForm struct {
	APIKey   string
	Cookie   string
	Password string
	Email    string
}

type ChangeOldNewPasswordForm struct {
	APIKey                  string
	Cookie                  string
	OldPassword             string
	NewPassword             string
	NewPasswordConfirmation string
}

type RegistrationForm struct {
	APIKey               string
	Name                 string
	Username             string
	Password             string
	PasswordConfirmation string
	Email                string
	EmailConfirmation    string
	AccountType          string
	CountryCode          string
	PhoneNumber          string
}

type LoginForm struct {
	APIKey   string
	Username string
	Password string
}

type EmailPhoneNumberValidationForm struct {
	APIKey       string
	UsernameHash string
	EmailKey     string
	PhoneKey     string
}

type RequestPasswordResetForm struct {
	APIKey   string
	Username string
}

type ResetPasswordForm struct {
	APIKey                  string
	UsernameHash            string
	ResetKey                string
	NewPassword             string
	NewPasswordConfirmation string
}
