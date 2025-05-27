package rest

import (
	"encoding/json"
	"microblogging/internal/domain/usecases"
	"microblogging/pkg/logging"
	"net/http"
	"time"
)

type TweetHandler struct {
	tweetCreator usecases.TweetCreator
	logger       *logging.Logger
}

type CreateTweetRequest struct {
	TweetID string `json:"tweetId"`
	Content string `json:"content"`
}

type CreateTweetResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewTweetHandler(tweetCreator usecases.TweetCreator) *TweetHandler {
	return &TweetHandler{
		tweetCreator: tweetCreator,
		logger:       logging.GetLogger().WithField("component", "TweetHandler"),
	}
}

func (h *TweetHandler) PostTweet(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		h.logger.Error("X-User-ID header is required")
		http.Error(w, "X-User-ID header is required", http.StatusBadRequest)
		return
	}

	var req CreateTweetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.WithError(err).Error("Failed to decode request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.Content == "" {
		h.logger.Error("Content is required")
		http.Error(w, "Content is required", http.StatusBadRequest)
		return
	}

	createdAt := time.Now().UTC()

	err := h.tweetCreator.Create(r.Context(), req.TweetID, userID, req.Content, createdAt)
	if err != nil {
		h.logger.WithError(err).Error("Failed to create tweet")
		http.Error(w, "Failed to create tweet", http.StatusInternalServerError)
		return
	}

	response := CreateTweetResponse{
		ID:        req.TweetID,
		UserID:    userID,
		Content:   req.Content,
		CreatedAt: createdAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logger.WithError(err).Error("Failed to encode response")
	}
}
