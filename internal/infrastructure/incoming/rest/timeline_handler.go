package rest

import (
	"encoding/json"
	"microblogging/internal/domain/usecases"
	"microblogging/pkg/logging"
	"net/http"
	"time"
)

type TimelineEntry struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type TimelineHandler struct {
	timelineGetter usecases.TimelineGetter
	logger         *logging.Logger
}

type GetTimelineResponse struct {
	Tweets []TimelineEntry `json:"tweets"`
}

func NewTimelineHandler(timelineGetter usecases.TimelineGetter) *TimelineHandler {
	return &TimelineHandler{
		timelineGetter: timelineGetter,
		logger:         logging.GetLogger().WithField("component", "TimelineHandler"),
	}
}

func (h *TimelineHandler) GetTimeline(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		h.logger.Error("X-User-ID header is required")
		http.Error(w, "X-User-ID header is required", http.StatusBadRequest)
		return
	}

	timeline, err := h.timelineGetter.Get(r.Context(), userID)
	if err != nil {
		h.logger.WithError(err).Error("Failed to get timeline")
		http.Error(w, "Failed to get timeline", http.StatusInternalServerError)
		return
	}

	response := GetTimelineResponse{}

	for _, tweet := range timeline.Tweets {
		response.Tweets = append(response.Tweets, TimelineEntry{
			ID:        tweet.TweetID,
			UserID:    tweet.UserID,
			Content:   tweet.Content,
			CreatedAt: tweet.CreatedAt,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logger.WithError(err).Error("Failed to encode response")
	}
}
