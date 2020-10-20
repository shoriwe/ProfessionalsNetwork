package General

const (
	ChangeAccountDescriptionQuery = "MATCH (account {id: $AccountID, ownerID: $AccountID}) SET (account).description = $Description RETURN account"
	ChangeAccountLocationQuery    = "MATCH (account {id: $AccountID, ownerID: $AccountID}) SET (account).location = $Location RETURN account"
)
