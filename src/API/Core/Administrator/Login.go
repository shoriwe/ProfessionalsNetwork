package Administrator

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/InputHandler/TemplateValidation"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"net/http"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	loginForm := new(RequestsForms.LoginForm)
	if _, _, _, _, isRequestValid := SessionHandling.RequestHandler(writer, request, loginForm, true); isRequestValid {
		isValid, message := TemplateValidation.IsLoginFormValid(loginForm, true)
		if isValid {
			message = SessionHandling.CreateCookie(&loginForm.Username)
		}
		_, _ = writer.Write(message)
	}
}
