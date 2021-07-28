package handlers

import (
	"net/http"

	"github.com/ChankyDEV/emotion_server/services"
)

type Users struct {
	AuthService services.IAuthService
}

// CONSTRUCTOR
func NewUsersHandler(authService services.IAuthService) *Users {
	return &Users{authService}
}

func (u *Users) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

}
