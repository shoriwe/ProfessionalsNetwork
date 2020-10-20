package Professional

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/ProfessionalCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/InputHandler"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"net/http"
)

func SpeakLanguage(writer http.ResponseWriter, request *http.Request) {
	languageForm := new(RequestsForms.ChangeLanguageForm)
	if professionalID, _, _, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, languageForm, true, SQLDatabase.ProfessionalAccount); isAccountTypeCorrect && isRequestValid {
		if InputHandler.IsLanguageValid(&languageForm.Language) {
			_, _ = writer.Write(ProfessionalCore.SpeakLanguageBackend(&professionalID, &languageForm.Language, &languageForm.Speaks))
		} else {
			_, _ = writer.Write([]byte("{\"Error\":\"Language field is invalid\"}"))
		}
	}
}
