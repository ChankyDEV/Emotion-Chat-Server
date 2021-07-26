package routers

import (
	"net/http"

	"github.com/ChankyDEV/emotion_server/controllers"
	"github.com/gorilla/mux"
)

func StartUp(authController controllers.IAuthController) {
	serveMux := mux.NewRouter()
	serveMux.HandleFunc("/auth/signup", authController.SignUp)

	http.ListenAndServe(":3000", serveMux)
}
