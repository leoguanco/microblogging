package outgoing

import (
	"context"
	"microblogging/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryUserRepository_Save(t *testing.T) {
	type input struct {
		initialData map[string]domain.User
		user        domain.User
	}

	type want struct {
		err bool
	}

	tests := []struct {
		name  string
		input input
		want  want
	}{
		{
			name: "Save new user",
			input: input{
				initialData: map[string]domain.User{},
				user: domain.User{
					UserID:  "user1",
					Follows: []string{},
				},
			},
			want: want{
				err: false,
			},
		},
		{
			name: "Update existing user",
			input: input{
				initialData: map[string]domain.User{
					"user1": {
						UserID:  "user1",
						Follows: []string{},
					},
				},
				user: domain.User{
					UserID:  "user1",
					Follows: []string{"user2"},
				},
			},
			want: want{
				err: false,
			},
		},
		{
			name: "Save user with empty ID",
			input: input{
				initialData: map[string]domain.User{},
				user: domain.User{
					UserID:  "",
					Follows: []string{},
				},
			},
			want: want{
				err: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &InMemoryUserRepository{
				users: make(map[string]domain.User),
			}

			for k, v := range tt.input.initialData {
				repo.users[k] = v
			}

			err := repo.Save(context.Background(), tt.input.user)

			if tt.want.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				storedUser, exists := repo.users[tt.input.user.UserID]
				assert.True(t, exists)
				assert.Equal(t, tt.input.user, storedUser)
			}
		})
	}
}

func TestInMemoryUserRepository_GetByID(t *testing.T) {
	type input struct {
		initialData map[string]domain.User
		userID      string
	}

	type want struct {
		user  domain.User
		found bool
		err   bool
	}

	user1 := domain.User{
		UserID:  "user1",
		Follows: []string{"user2"},
	}

	tests := []struct {
		name  string
		input input
		want  want
	}{
		{
			name: "Get existing user",
			input: input{
				initialData: map[string]domain.User{
					"user1": user1,
				},
				userID: "user1",
			},
			want: want{
				user:  user1,
				found: true,
				err:   false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &InMemoryUserRepository{
				users: make(map[string]domain.User),
			}

			for k, v := range tt.input.initialData {
				repo.users[k] = v
			}

			user, err := repo.Get(context.Background(), tt.input.userID)

			if tt.want.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.user, user)
			}
		})
	}
}

func TestNewInMemoryUserRepository(t *testing.T) {
	repo := NewInMemoryUserRepository()

	assert.NotNil(t, repo)

	user := domain.User{
		UserID:  "test-user",
		Follows: []string{},
	}

	err := repo.Save(context.Background(), user)
	assert.NoError(t, err)

	retrievedUser, err := repo.Get(context.Background(), "test-user")
	assert.NoError(t, err)
	assert.Equal(t, user, retrievedUser)
}
