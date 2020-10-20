package MessageSending

var (
	APIKey          = ""
	MessagingAPIURI = ""
)

var (
	FrontEndURI = ""
)

var (
	EmailPhoneNumberChangeValidationURL = ""
	AccountConfirmationURL              = ""
	ResetPasswordURL                    = ""
	InvitationURL                       = ""
	ProfessionalProfile                 = ""
)

func ConfigureMessagingAPIVariables(messagingAPIKey string, messagingAPIURI string, frontEndURI string) {
	APIKey = messagingAPIKey
	MessagingAPIURI = messagingAPIURI
	FrontEndURI = frontEndURI

	EmailPhoneNumberChangeValidationURL = FrontEndURI + "/email/phone/number/change/confirmation"
	AccountConfirmationURL = FrontEndURI + "/register/confirmation"
	ResetPasswordURL = FrontEndURI + "/reset/password"
	InvitationURL = FrontEndURI + "/dashboard/teams/invitation"
	ProfessionalProfile = FrontEndURI + "/dashboard/teams/view/professional"
}
