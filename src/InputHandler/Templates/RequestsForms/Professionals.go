package RequestsForms

type SearchTeamsForm struct {
	APIKey string
	Cookie string
	Query  string
}
type ApplyToTeamForm struct {
	APIKey       string
	Cookie       string
	ContractorID int
	TeamName     string
}

type ExitFromTeamForm struct {
	APIKey       string
	Cookie       string
	ContractorID int
	TeamName     string
}

type AcceptTeamInvitationForm struct {
	APIKey        string
	Cookie        string
	Accept        bool
	InvitationKey string
}

type LikedOnlyForm struct {
	APIKey    string
	Cookie    string
	LikedOnly bool
}

type AddSkillForm struct {
	APIKey string
	Cookie string
	Name   string
}

type LikeSkillForm struct {
	APIKey string
	Cookie string
	Skill  string
	Like   bool
}

type KnowSkillForm struct {
	APIKey string
	Cookie string
	Skill  string
	Know   bool
}

type ChangeLanguageForm struct {
	APIKey   string
	Cookie   string
	Language string
	Speaks   bool
}

type ChangeAvailableForm struct {
	APIKey    string
	Cookie    string
	Available bool
}

type ChangeRemoteForm struct {
	APIKey string
	Cookie string
	Remote bool
}

type ChangeGenderForm struct {
	APIKey string
	Cookie string
	Gender string
}

type ChangeNationalityForm struct {
	APIKey      string
	Cookie      string
	Nationality string
}
