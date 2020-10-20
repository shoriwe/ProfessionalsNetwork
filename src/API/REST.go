package API

import (
	"github.com/shoriwe/ProNet/src/API/Core/Account"
	"github.com/shoriwe/ProNet/src/API/Core/Administrator"
	"github.com/shoriwe/ProNet/src/API/Core/Contractor"
	"github.com/shoriwe/ProNet/src/API/Core/Professional"
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"log"
	"net/http"
)

type RequestData struct {
	Cookie string
	Data   interface{}
}

func Serve(bindAddress string) {
	// First create a SECRET to handle the Cookie creation
	SessionHandling.CreateSecret()

	// General Core
	http.HandleFunc("/general/change/description", Account.ChangeDescription)
	http.HandleFunc("/general/change/location", Account.ChangeLocation)
	http.HandleFunc("/general/request/professional/information", Account.RequestProfessionalInformation)
	http.HandleFunc("/general/request/team/members", Account.RequestTeamMembers)
	http.HandleFunc("/general/request/account", Account.RequestAccount)
	http.HandleFunc("/general/request/profile", Account.RequestProfile)
	http.HandleFunc("/general/request/change/email/phone/number", Account.RequestChangeEmailPhoneNumber) // Not Implemented
	http.HandleFunc("/general/confirm/new/email/phone/number", Account.ConfirmNewEmailPhoneNumber)       // Not implemented
	http.HandleFunc("/general/search/languages", Account.SearchLanguages)
	http.HandleFunc("/general/search/skills", Account.SearchSkills)
	http.HandleFunc("/general/account/type", Account.RequestAccountType)
	http.HandleFunc("/general/change/password", Account.ChangePassword)
	http.HandleFunc("/general/confirm/registration", Account.ConfirmRegistration)
	http.HandleFunc("/general/register", Account.Register)
	http.HandleFunc("/general/login", Account.Login)
	http.HandleFunc("/general/reset/password", Account.PasswordReset)
	http.HandleFunc("/general/request/password/reset", Account.RequestPasswordReset)

	// Contractor functionality
	http.HandleFunc("/contractor/request/owned/teams", Contractor.RequestTeamsOwned)
	http.HandleFunc("/contractor/accept/professional/in/team", Contractor.AcceptProfessionalInTeam)
	http.HandleFunc("/contractor/dissolve/team", Contractor.DissolveTeam)
	http.HandleFunc("/contractor/remove/professional/from/team", Contractor.RemoveProfessionalFromTeam)
	http.HandleFunc("/contractor/invite/professional/to/team", Contractor.InviteProfessionalToTeam)
	http.HandleFunc("/contractor/create/team", Contractor.CreateTeam)
	http.HandleFunc("/contractor/find/professionals", Contractor.FindProfessionals)
	// http.HandleFunc("/contractor/change/description", Contractor.ChangeDescription)
	// http.HandleFunc("/contractor/change/location", Contractor.ChangeLocation)

	// Administrator functionality
	http.HandleFunc("/admin/add/skill", Administrator.AddSkill)
	http.HandleFunc("/admin/login", Administrator.Login)

	// Professional functionality
	http.HandleFunc("/professional/search/teams", Professional.SearchTeams)
	http.HandleFunc("/professional/apply/to/team", Professional.ApplyToTeam)
	http.HandleFunc("/professional/exit/from/team", Professional.ExitFromTeam)
	http.HandleFunc("/professional/accept/team/invitation", Professional.AcceptTeamInvitation)
	// http.HandleFunc("/professional/change/description", Account.ChangeDescription)
	http.HandleFunc("/professional/change/gender", Professional.ChangeGender)
	http.HandleFunc("/professional/change/available", Professional.ChangeAvailable)
	http.HandleFunc("/professional/change/remote", Professional.ChangeRemote)
	http.HandleFunc("/professional/change/nationality", Professional.ChangeNationality)
	http.HandleFunc("/professional/speak/language", Professional.SpeakLanguage)
	http.HandleFunc("/professional/change/liked/only", Professional.ChangeLikedOnly)
	http.HandleFunc("/professional/know/skill", Professional.KnownSkill)
	http.HandleFunc("/professional/like/skill", Professional.LikeSkill)

	runtimeError := http.ListenAndServe(bindAddress, nil)

	if runtimeError != nil {
		log.Fatal(runtimeError.Error())
	}
}
