package main

import (
	"log"
	"os"

	"github.com/ChankyDEV/emotion_server/handlers"
	"github.com/ChankyDEV/emotion_server/repositories"
	"github.com/ChankyDEV/emotion_server/routers"
	"github.com/ChankyDEV/emotion_server/services"
)

var (
	address = "127.0.0.1"
	port    = ":3000"
)

func main() {
	logger := log.New(os.Stdout, "my-api", log.LstdFlags)

	authRepository := repositories.NewMongoRepository()
	authService := services.NewAuthService(authRepository)
	users := handlers.NewUsersHandler(authService)

	routers.StartUp(logger, users)
}
