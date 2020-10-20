package Queries

const (
	UpdatePasswordQuery    = "UPDATE users SET password = ? WHERE id = ? AND username = ?"
	UpdateEmailQuery       = "UPDATE users SET email = ? WHERE id = ? AND username = ?"
	UpdatePhoneNumberQuery = "UPDATE users SET phone_number = ? WHERE id = ? AND username = ?"
)

const (
	UpdateTeamStatusQuery = "UPDATE teams SET active = false WHERE id = ? AND team_name = ? AND owner_id = ?"
)
