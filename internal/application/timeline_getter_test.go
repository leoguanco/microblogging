package application

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"microblogging/internal/domain"
	"microblogging/internal/domain/ports/mocks"
	"testing"
	"time"
)

func TestTimelineGetter_Get(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type input struct {
		ctx    context.Context
		userID string
	}

	type want struct {
		err      error
		timeline domain.Timeline
	}

	timelineRepositoryMock := mocks.NewMockTimelineRepository(ctrl)
	timelineGetter := NewTimelineGetter(timelineRepositoryMock)

	timeline := domain.Timeline{
		UserID: "user-uuid",
		Limit:  10,
		Total:  1,
		Tweets: []domain.Tweet{
			{
				TweetID:   "tweet-uuid",
				UserID:    "user-uuid",
				Content:   "Lorem ipsum",
				CreatedAt: time.Now(),
			},
		},
	}

	tests := []struct {
		name  string
		input input
		mock  func()
		want  want
	}{
		{
			name: "Should return nil when get a timeline",
			input: input{
				ctx:    ctx,
				userID: "user-uuid",
			},
			mock: func() {
				timelineRepositoryMock.EXPECT().Get(ctx, "user-uuid").
					Return(timeline, nil)
			},
			want: want{
				err:      nil,
				timeline: timeline,
			},
		},
		{
			name: "Should return an error when get a timeline",
			input: input{
				ctx:    ctx,
				userID: "user-uuid",
			},
			mock: func() {
				timelineRepositoryMock.EXPECT().Get(ctx, "user-uuid").
					Return(domain.Timeline{}, errors.New("internal server error"))
			},
			want: want{
				err:      errors.New("internal server error"),
				timeline: domain.Timeline{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			timeline, err := timelineGetter.Get(tt.input.ctx, tt.input.userID)
			assert.Equal(t, tt.want.err, err)
			assert.Equal(t, tt.want.timeline, timeline)
		})
	}
}
