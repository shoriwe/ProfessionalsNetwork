package General

const (
	SearchSkillsQuery    = "MATCH (object_ {ownerID: $OwnerID, id: $ObjectID}) WITH object_ MATCH (skill:Skill) WHERE NOT (object_)-->(skill) WITH skill WHERE toLower(skill.name) CONTAINS $SearchQuery WITH COLLECT(skill.name) AS skills RETURN skills"
	SearchLanguagesQuery = "MATCH (object_ {ownerID: $OwnerID, id: $ObjectID}) WITH object_ MATCH (language:Language) WHERE NOT (object_)-->(language) WITH language WHERE toLower(language.name) CONTAINS $SearchQuery WITH COLLECT(language.name) AS languages RETURN languages"
)

const (
	GetAccountProfileNodeByIDQuery = "MATCH (account {id: $AccountID, ownerID: $AccountID}) RETURN account"
	GetAccountSubNodesByIDQuery    = "MATCH (account {id: $AccountID, ownerID: $AccountID})-[rawRelationship]->(subNode) WITH subNode, LABELS(subNode)[0] as nodeFamily, TYPE(rawRelationship) AS relationship_ WITH { relationships: COLLECT(relationship_), node_: subNode, nodeFamily: nodeFamily } AS results WITH results.relationships AS relationships, results.nodeFamily AS nodeFamily, results.node_ AS subNode WITH COLLECT([subNode, nodeFamily, relationships]) AS results RETURN results"
)

const (
	GetTeamMembersQuery = "MATCH (account)-[rawRelationship]->(team:Team {id:$TeamID, ownerID:$ContractorID, name:$TeamName}) WITH account, LABELS(account)[0] as accountType, TYPE(rawRelationship) AS relationship_ WITH { relationships: COLLECT(relationship_), account: account, accountType: accountType } AS results WITH results.relationships AS relationships, results.accountType AS accountType, results.account AS account WITH COLLECT([account, accountType, relationships]) AS results RETURN results"
)
