package main

import (
	"athenify/app/services"
	"athenify/config"
	"athenify/domain"
	"athenify/persistence"
	"athenify/presentation"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
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

	workers, _ := strconv.Atoi(config.GetEnv("WORKERS", "3"))
	jobs := make(chan domain.Job, 10)
	wg := &sync.WaitGroup{}
	workerPool := persistence.NewWorkerPool(workers, jobs, wg)
	workerPool.Start()

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

	log.Println("Closing jobs channel")
	close(jobs)

	wg.Wait()

	log.Println("Shutting down")
	server.Shutdown(ctx)
	os.Exit(0)
}
