package General

const (
	CheckIfIsOwnerOfObjectByIDsQuery = "MATCH (object_) WHERE object_.id = $ObjectID AND object_.ownerID = $OwnerID RETURN object_"
)
const (
	CheckIfSkillExistsQuery = "MATCH (skill:Skill {name: $Skill}) RETURN (skill).name"
)
