package main

import (
	"athenify/app/services"
	"athenify/config"
	"athenify/persistence"
	"athenify/presentation"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	postgresUsername := config.GetEnv("POSTGRES_USERNAME", "")
	postgresPassword := config.GetEnv("POSTGRES_PASSWORD", "")
	postgresDB := config.GetEnv("POSTGRES_DB", "")
	postgresIP := config.GetEnv("POSTGRES_IP", "")
	postgresPort := config.GetEnv("POSTGRES_PORT", "")

	log.Println("Connecting to database")
	db, err := persistence.InitDB(postgresUsername, postgresPassword, postgresIP, postgresPort, postgresDB)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	userRepository := persistence.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := presentation.NewUserHandler(userService)

	r := mux.NewRouter()
	s := r.PathPrefix("/users").Subrouter()
	s.HandleFunc("/", userHandler.Create).Methods("POST")
	s.HandleFunc("/", userHandler.Get).Methods("GET")

	server := &http.Server{
		Addr:    "localhost:8000",
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Shutting down")
	server.Shutdown(ctx)
	os.Exit(0)
}
