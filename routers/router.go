package routers

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ChankyDEV/emotion_server/handlers"
	"github.com/gorilla/mux"
)

func StartUp(logger *log.Logger, usersHandler *handlers.Users) {

	router := mux.NewRouter()

	// SET UP HANDLERS
	router.Handle("/users", usersHandler)
	address := "127.0.0.1:3000"
	server := http.Server{
		Handler:      router,
		Addr:         address,
		ErrorLog:     logger,            // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// RUN SERVER
	go runServer(server, logger)
	listenOnServerStatus(server)
}

func runServer(server http.Server, logger *log.Logger) {

	logger.Println("Starting server on port 9090")

	err := server.ListenAndServe()
	if err != nil {
		logger.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}

func listenOnServerStatus(server http.Server) {
	serverStatus := make(chan os.Signal, 1)
	signal.Notify(serverStatus, os.Interrupt)
	signal.Notify(serverStatus, os.Kill)

	status := <-serverStatus
	log.Println("Got signal:", status)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
