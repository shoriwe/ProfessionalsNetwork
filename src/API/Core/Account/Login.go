package Account

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/InputHandler/TemplateValidation"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"net/http"
)

func loginBackend(loginForm *RequestsForms.LoginForm) []byte {
	isValid, message := TemplateValidation.IsLoginFormValid(loginForm, false)
	if isValid {
		message = SessionHandling.CreateCookie(&loginForm.Username)
	}
	return message
}

func Login(writer http.ResponseWriter, request *http.Request) {
	loginForm := new(RequestsForms.LoginForm)
	if _, _, _, _, isRequestValid := SessionHandling.RequestHandler(writer, request, loginForm, true); isRequestValid {
		_, _ = writer.Write(loginBackend(loginForm))
	}
}
