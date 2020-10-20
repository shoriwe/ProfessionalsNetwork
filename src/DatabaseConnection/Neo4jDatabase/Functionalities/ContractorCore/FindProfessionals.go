package ContractorCore

import (
	"encoding/json"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/Contractors"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"math"
	"net/http"
)

func findTransactionFilter(result *neo4j.Result) interface{} {
	if (*result).Next() {
		record := (*result).Record()
		results, _ := record.Get("results")
		numberOfResults, _ := record.Get("numberOfResults")
		professionals := make([]interface{}, 0)
		for _, professionalEntry := range results.([]interface{}) {
			professionalNode := professionalEntry.([]interface{})[0].(neo4j.Node).Props()
			professionals = append(professionals, []interface{}{professionalNode, professionalEntry.([]interface{})[1]})
		}
		encodedResult, _ := json.Marshal(map[string]interface{}{"Results": professionals, "NumberOfResults": numberOfResults})
		return encodedResult
	}
	return nil
}

func handleFindTransaction(searchFilter *RequestsForms.SearchFilterForm, session neo4j.Session) []byte {
	if searchFilter.Page == 0 {
		searchFilter.Page = 1
	}
	query := Contractors.ProfessionalSearchIncludeNotLikedQuery
	switch searchFilter.LikedOnly.(type) {
	case bool:
		if searchFilter.LikedOnly.(bool) {
			query = Contractors.ProfessionalSearchLikedOnlyQuery
		}
	}
	data := map[string]interface{}{
		"Languages":     searchFilter.Languages,
		"Locations":     searchFilter.Locations,
		"Nationalities": searchFilter.Nationalities,
		"Available":     searchFilter.Available,
		"Gender":        searchFilter.Gender,
		"Remote":        searchFilter.Remote,
		"Skills":        searchFilter.Skills,
		"LikedOnly":     searchFilter.LikedOnly,
		"Page":          math.Abs(float64(searchFilter.Page)),
	}
	result := Neo4jDatabase.ExecuteQuery(session, query, data, findTransactionFilter)
	if result != nil {
		return result.([]byte)
	}
	return []byte("[]")
}

func FindProfessionalsBackend(searchFilter *RequestsForms.SearchFilterForm) ([]byte, int) {
	session, sessionCreationError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
	defer Neo4jDatabase.CloseSession(session)
	if sessionCreationError == nil {
		defer session.Close()
		return handleFindTransaction(searchFilter, session), http.StatusOK
	}
	return []byte("Something goes wrong"), http.StatusInternalServerError

}
