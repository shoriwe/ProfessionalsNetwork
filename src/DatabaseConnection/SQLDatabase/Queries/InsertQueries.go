package Queries

const (
	CreateUserQuery = "INSERT INTO users (name, username, password, email, phone_number, account_type) VALUES (?, ?, ?, ?, ?, ?)"
)

const (
	CreateTeamQuery = "INSERT INTO teams (owner_id, team_name, active) VALUES (?, ?, ?)"
)
