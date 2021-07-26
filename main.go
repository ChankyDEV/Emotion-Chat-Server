package main

import (
	"fmt"

	"github.com/ChankyDEV/emotion_server/controllers"
	"github.com/ChankyDEV/emotion_server/repositories"
	"github.com/ChankyDEV/emotion_server/routers"
	"github.com/ChankyDEV/emotion_server/services"
)

func main() {
	fmt.Println("Server is starting")
	authRepository := repositories.NewMongoRepository()
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)
	routers.StartUp(authController)
}
