//+build unit

package users

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestAddUser(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	mockRepo := NewMockRepository(ctrl)

	s := New()
	s.SetRepo(mockRepo)
	u := User{"jack", 23}
	mockRepo.EXPECT().AddUser(User{"jack", 23})
	if err := s.AddUser(u); err != nil {
		t.Fatalf("add user failed unexpectedly, err: %v", err)
	}
}

func TestGetUserEmptyResponse(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	mockRepo := NewMockRepository(ctrl)

	mockRepo.EXPECT().GetUsers().Return([]*User{}, nil)

	s := New()
	s.SetRepo(mockRepo)

	_, err := s.GetUsers()
	if err == nil {
		t.Fatalf("Test case 'get user with empty response' failed, got %s, want %s", err, ErrNotFound)
	}
	if err != ErrNotFound {
		t.Fatalf("Test case 'get user with empty response' failed, got %v, want %s", err, ErrNotFound)
	}
}

func TestGetUserNonEmptyResponse(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	mockRepo := NewMockRepository(ctrl)

	mockRepo.EXPECT().GetUsers().Return([]*User{&User{"thanh", 22}}, nil)

	s := New()
	s.SetRepo(mockRepo)

	users, err := s.GetUsers()
	if err != nil {
		t.Fatalf("Test case 'get user with non empty response' failed, got err %v, want %s", err, "nil")
	}
	if len(users) < 1 {
		t.Fatalf("Test case 'get user with non empty response' failed, got len=%v, want len=1", len(users))
	}
}
