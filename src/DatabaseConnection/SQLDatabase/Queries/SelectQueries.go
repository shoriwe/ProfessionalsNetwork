package Queries

const (
	GetAccountEmailPhoneNumberQuery                  = "SELECT email, phone_number  FROM users WHERE id = ? AND username = ?"
	GetAccountEmailPhoneNumberWithAccountIDQuery     = "SELECT email, phone_number  FROM users WHERE id = ?"
	CheckIfUsernamePasswordAreValidQuery             = "SELECT username, password FROM users WHERE username = ? AND password = ? LIMIT 1"
	CheckIfUsernameExistsAndGetIDAndAccountTypeQuery = "SELECT id, account_type FROM users WHERE username = ? LIMIT 1"
	GetUserEmailQuery                                = "SELECT email FROM users WHERE id = ? AND username = ? LIMIT 1"
	GetUserPhoneNumberQuery                          = "SELECT phone_number FROM users WHERE id = ? AND username = ? LIMIT 1"
)

const (
	CheckIfUserIDExistsQuery = "SELECT username, account_type FROM users WHERE id = ?"
)

const (
	CheckIfTeamExistsQuery = "SELECT id FROM teams WHERE owner_id = ? AND team_name = ? AND active = true"
)

const (
	GetOwnedTeamsByProfessional = "SELECT id, owner_id, team_name FROM teams WHERE owner_id = ? AND active = true"
)
