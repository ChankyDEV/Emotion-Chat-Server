package controllers

import (
	"fmt"
	"net/http"

	"github.com/ChankyDEV/emotion_server/models"
	"github.com/ChankyDEV/emotion_server/services"
)

type IAuthController interface {
	SignUp(rw http.ResponseWriter, r *http.Request)
}

type AuthController struct {
	AuthService services.IAuthService
}

func NewAuthController(authService services.IAuthService) IAuthController {
	return &AuthController{AuthService: authService}
}

func (controller *AuthController) SignUp(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		user := &models.User{}
		parserErr := user.FromJSON(r.Body)
		if parserErr != nil {
			fmt.Fprint(rw, parserErr)
			return
		}

		user, serviceErr := controller.AuthService.SignUp(user.Email, user.Phone, user.Password)

		if serviceErr != nil {
			fmt.Fprint(rw, serviceErr)
			return
		}

		fmt.Fprint(rw, user)
	} else {
		fmt.Fprint(rw, http.StatusBadRequest)
	}

}
