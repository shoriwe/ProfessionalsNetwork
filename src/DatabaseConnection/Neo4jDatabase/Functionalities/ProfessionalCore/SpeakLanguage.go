package ProfessionalCore

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/Professionals"
	"github.com/shoriwe/ProNet/src/InputHandler"
	"log"
)

func findLanguage(languages *[]string, wantedLanguage *string) int {
	for index, language := range *languages {
		if language == *wantedLanguage {
			return index
		}
	}
	return -1
}

func professionalLanguagesFilter(result *neo4j.Result) interface{} {
	if (*result).Next() {
		record := (*result).Record()
		languages, found := record.Get("languages")
		if found {
			switch languages.(type) {
			case []interface{}:
				return InputHandler.ToStringSlice(languages.([]interface{}))
			}
		}
	}
	return nil
}

func getProfessionalLanguages(professionalID *int, session neo4j.Session) []string {
	data := map[string]interface{}{
		"ProfessionalID": *professionalID,
	}
	result := Neo4jDatabase.ExecuteQuery(session, Professionals.GetProfessionalLanguagesQuery, data, professionalLanguagesFilter)
	if result != nil {
		return result.([]string)
	}
	return []string{}
}

func doesntSpeakLanguageFilter(_ *neo4j.Result) interface{} {
	return []byte("{\"Success\":\"Professional now doesn't speaks the wanted language\"}")
}

func doesntSpeakLanguage(professionalID *int, language *string, session neo4j.Session) []byte {
	data := map[string]interface{}{
		"ProfessionalID": *professionalID,
		"Language":       *language,
	}
	result := Neo4jDatabase.ExecuteQuery(session, Professionals.DoesntSpeakLanguageQuery, data, doesntSpeakLanguageFilter)
	if result != nil {
		return result.([]byte)
	}
	return []byte("{\"Error\":\"Something goes wrong\"}")
}

func speakLanguageFilter(_ *neo4j.Result) interface{} {
	return []byte("{\"Success\":\"Professional now speaks the wanted language\"}")
}

func speakLanguage(professionalID *int, language *string, session neo4j.Session) []byte {
	data := map[string]interface{}{
		"ProfessionalID": *professionalID,
		"Language":       *language,
	}
	result := Neo4jDatabase.ExecuteQuery(session, Professionals.SpeaksLanguageQuery, data, speakLanguageFilter)
	if result != nil {
		return result.([]byte)
	}
	return []byte("{\"Error\":\"Something goes wrong\"}")
}

func handleSpeakLanguage(professionalID *int, language *string, session neo4j.Session) []byte {
	languages := getProfessionalLanguages(professionalID, session)
	wantedLanguageIndex := findLanguage(&languages, language)
	if wantedLanguageIndex == -1 { // When is -1, means the language doesn't exists
		return speakLanguage(professionalID, language, session)
	} else {
		return []byte("{\"Error\":\"The language is already in the skill set of this professional\"}")
	}
}

func handleDoesntSpeakLanguage(professionalID *int, language *string, session neo4j.Session) []byte {
	languages := getProfessionalLanguages(professionalID, session)
	wantedLanguageIndex := findLanguage(&languages, language)
	if wantedLanguageIndex != -1 { // When is -1, means the language doesn't exists
		languages = append(languages[0:wantedLanguageIndex], languages[wantedLanguageIndex+1:]...)
		return doesntSpeakLanguage(professionalID, language, session)
	} else {
		return []byte("{\"Error\":\"The language is not in the skill set of this professional\"}")
	}
}

func SpeakLanguageBackend(professionalID *int, language *string, speaks *bool) []byte {
	session, connectionError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
	defer Neo4jDatabase.CloseSession(session)
	if connectionError == nil {
		if *speaks {
			return handleSpeakLanguage(professionalID, language, session)
		}
		return handleDoesntSpeakLanguage(professionalID, language, session)
	}
	log.Print(connectionError)
	return []byte("{\"Error\":\"Semething goes wrong\"}")
}
