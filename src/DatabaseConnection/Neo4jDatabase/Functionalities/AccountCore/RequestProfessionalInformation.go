package AccountCore

import (
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
)

func RequestProfessionalInformationBackend(professionalId *int) []byte {
	return RequestProfileBackend(professionalId, SQLDatabase.ProfessionalAccount)
}
