package RequestsForms

type RequestTeamsOwnedForm struct {
	APIKey string
	Cookie string
}

type AcceptProfessionalApplicationForm struct {
	APIKey         string
	Cookie         string
	ApplicationKey string
	Accept         bool
}

type DissolveTeamForm struct {
	APIKey   string
	Cookie   string
	TeamName string
}

type InviteProfessionalToTeamForm struct {
	APIKey         string
	Cookie         string
	TeamName       string
	ProfessionalID int
}
type CreateTeamForm struct {
	APIKey   string
	Cookie   string
	TeamName string
}

// This template is vulnerable to object injection
type SearchFilterForm struct {
	APIKey        string
	Cookie        string
	Languages     []interface{}
	LikedOnly     interface{}
	Skills        []interface{}
	Locations     []interface{} // Typically is an String slice
	Nationalities []interface{} // Typically is an String slice
	Gender        interface{}   // Typically is String
	Remote        interface{}   // Typically is Boolean
	Available     interface{}   // Typically is Boolean
	Page          int
}

type AccountInformationForm struct {
	APIKey    string
	Cookie    string
	AccountID int
}
