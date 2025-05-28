package outgoing

import (
	"context"
	"reflect"
	"testing"
)

func TestInMemoryFolloweeRepository_Get(t *testing.T) {
	repo := NewInMemoryFolloweeRepository()
	ctx := context.Background()

	_ = repo.Save(ctx, "user1", []string{"follower1", "follower2"})

	tests := []struct {
		name     string
		id       string
		expected []string
		wantErr  bool
	}{
		{name: "existing user", id: "user1", expected: []string{"follower1", "follower2"}, wantErr: false},
		{name: "non-existent user", id: "user2", expected: []string{}, wantErr: true},
		{name: "empty ID", id: "", expected: []string{}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := repo.Get(ctx, tt.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("unexpected error status, got err: %v, wantErr: %v", err, tt.wantErr)
			}

			if !tt.wantErr && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("unexpected result, got: %v, expected: %v", result, tt.expected)
			}
		})
	}
}

func TestInMemoryFolloweeRepository_Save(t *testing.T) {
	repo := NewInMemoryFolloweeRepository()
	ctx := context.Background()

	tests := []struct {
		name       string
		id         string
		followers  []string
		checkID    string
		wantSaved  []string
		expectSave bool
	}{
		{name: "save valid user", id: "user1", followers: []string{"follower1"}, checkID: "user1", wantSaved: []string{"follower1"}, expectSave: true},
		{name: "overwrite existing user", id: "user1", followers: []string{"follower2", "follower3"}, checkID: "user1", wantSaved: []string{"follower2", "follower3"}, expectSave: true},
		{name: "empty followers list", id: "user2", followers: []string{}, checkID: "user2", wantSaved: []string{}, expectSave: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.Save(ctx, tt.id, tt.followers)

			if (err == nil) != tt.expectSave {
				t.Errorf("unexpected save status, got err: %v, expectSave: %v", err, tt.expectSave)
			}

			if tt.expectSave {
				result, err := repo.Get(ctx, tt.checkID)
				if err != nil {
					t.Fatalf("Get after Save failed, err: %v", err)
				}

				if !reflect.DeepEqual(result, tt.wantSaved) {
					t.Errorf("unexpected saved result, got: %v, want: %v", result, tt.wantSaved)
				}
			}
		})
	}
}
