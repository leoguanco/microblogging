package rest

import (
	"encoding/json"
	"microblogging/internal/domain/usecases"
	"microblogging/pkg/logging"
	"net/http"
)

type UserHandler struct {
	userFollower   usecases.UserFollower
	userUnfollower usecases.UserUnfollower
	logger         *logging.Logger
}

type FollowRequest struct {
	FolloweeID string `json:"followeeId"`
}

type UnfollowRequest struct {
	FolloweeID string `json:"followeeId"`
}

type FollowResponse struct {
	Message string `json:"message"`
}

func NewUserHandler(userFollower usecases.UserFollower, userUnfollower usecases.UserUnfollower) *UserHandler {
	return &UserHandler{
		userFollower:   userFollower,
		userUnfollower: userUnfollower,
		logger:         logging.GetLogger().WithField("component", "UserHandler"),
	}
}

func (h *UserHandler) FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID := r.Header.Get("X-User-ID")
	if followerID == "" {
		h.logger.Error("X-User-ID header is required")
		http.Error(w, "X-User-ID header is required", http.StatusBadRequest)
		return
	}

	var req FollowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.WithError(err).Error("Failed to decode request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if followerID == req.FolloweeID {
		h.logger.Error("Cannot follow yourself")
		http.Error(w, "Cannot follow yourself", http.StatusBadRequest)
		return
	}

	err := h.userFollower.Follow(r.Context(), followerID, req.FolloweeID)
	if err != nil {
		h.logger.WithError(err).Error("Failed to follow user")
		http.Error(w, "Failed to follow user", http.StatusInternalServerError)
		return
	}

	response := FollowResponse{
		Message: "Successfully followed user",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logger.WithError(err).Error("Failed to encode response")
	}
}

func (h *UserHandler) UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerID := r.Header.Get("X-User-ID")
	if followerID == "" {
		h.logger.Error("X-User-ID header is required")
		http.Error(w, "X-User-ID header is required", http.StatusBadRequest)
		return
	}

	var req UnfollowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.WithError(err).Error("Failed to decode request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if followerID == req.FolloweeID {
		h.logger.Error("Cannot unfollow yourself")
		http.Error(w, "Cannot unfollow yourself", http.StatusBadRequest)
		return
	}

	err := h.userUnfollower.Unfollow(r.Context(), followerID, req.FolloweeID)
	if err != nil {
		h.logger.WithError(err).Error("Failed to unfollow user")
		http.Error(w, "Failed to unfollow user", http.StatusInternalServerError)
		return
	}

	response := FollowResponse{
		Message: "Successfully unfollowed user",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logger.WithError(err).Error("Failed to encode response")
	}
}
