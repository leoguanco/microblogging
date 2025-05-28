package rest

import (
	"encoding/json"
	"errors"
	"go.uber.org/mock/gomock"
	"microblogging/internal/domain"
	"microblogging/internal/domain/usecases/mocks"
	"microblogging/pkg/logging"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimelineHandler_GetTimeline(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name               string
		setupMock          func(mockGetter *mocks.MockTimelineGetter)
		headerUserID       string
		expectedStatusCode int
		expectedResponse   interface{}
	}{
		{
			name: "valid user ID",
			setupMock: func(mockGetter *mocks.MockTimelineGetter) {
				mockGetter.EXPECT().Get(gomock.Any(), "valid-user").
					Return(domain.Timeline{
						Tweets: []domain.Tweet{
							{
								TweetID:   "tweet1",
								UserID:    "valid-user",
								Content:   "Hello, world!",
								CreatedAt: time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC),
							},
						},
					}, nil)
			},
			headerUserID:       "valid-user",
			expectedStatusCode: http.StatusOK,
			expectedResponse: GetTimelineResponse{
				Tweets: []TimelineEntry{
					{
						ID:        "tweet1",
						UserID:    "valid-user",
						Content:   "Hello, world!",
						CreatedAt: time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC),
					},
				},
			},
		},
		{
			name: "missing X-User-ID header",
			setupMock: func(mockGetter *mocks.MockTimelineGetter) {
			},
			headerUserID:       "",
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   "X-User-ID header is required\n",
		},
		{
			name: "error retrieving timeline",
			setupMock: func(mockGetter *mocks.MockTimelineGetter) {
				mockGetter.EXPECT().Get(gomock.Any(), "user-with-error").
					Return(domain.Timeline{}, errors.New("database error"))
			},
			headerUserID:       "user-with-error",
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   "Failed to get timeline\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockGetter := mocks.NewMockTimelineGetter(ctrl)
			tt.setupMock(mockGetter)

			logger := logging.GetLogger()
			handler := NewTimelineHandler(mockGetter)
			handler.logger = logger

			req := httptest.NewRequest(http.MethodGet, "/timeline", nil)
			if tt.headerUserID != "" {
				req.Header.Set("X-User-ID", tt.headerUserID)
			}

			rec := httptest.NewRecorder()
			handler.GetTimeline(rec, req)

			assert.Equal(t, tt.expectedStatusCode, rec.Code)

			if rec.Code == http.StatusOK {
				var actualResponse GetTimelineResponse
				err := json.NewDecoder(rec.Body).Decode(&actualResponse)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResponse, actualResponse)
			} else {
				assert.Equal(t, tt.expectedResponse, rec.Body.String())
			}
		})
	}
}
