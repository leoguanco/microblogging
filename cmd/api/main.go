package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"microblogging/internal/infrastructure"
	"microblogging/internal/infrastructure/incoming/rest"
	"microblogging/pkg/logging"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Starting microblogging service...")

	appDeps := infrastructure.BuildDependencies()

	handlers := rest.NewHandlers(
		appDeps.TweetCreator,
		appDeps.UserFollower,
		appDeps.UserUnfollower,
		appDeps.TimelineGetter,
	)

	router := mux.NewRouter()

	router.HandleFunc("/healthz", ping).Methods("GET")

	router.HandleFunc("/v1/tweets", handlers.TweetHandler.PostTweet).Methods("POST")
	router.HandleFunc("/v1/users/{userId}/follow", handlers.UserHandler.FollowUser).Methods("POST")
	router.HandleFunc("/v1/users/{userId}/unfollow", handlers.UserHandler.UnfollowUser).Methods("POST")
	router.HandleFunc("/v1/users/{userId}/timeline", handlers.TimelineHandler.GetTimeline).Methods("GET")

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.WithError(err).Error("Server forced to shutdown")
	}

	logger.Info("Server shutdown complete")
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}
