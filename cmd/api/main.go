package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"microblogging/pkg/logging"
)

func main() {
	logger := logging.GetLogger()

	router := mux.NewRouter()
	router.HandleFunc("/ping", ping).Methods("GET")

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		logger.Info("Server starting on port 8080...")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.WithError(err).Fatal("Error starting server")
		}
	}()

	<-stop
	logger.Info("Shutting down server...")
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}
