package users

import "errors"

//ErrNotFound represent a not found error.
var ErrNotFound = errors.New("not found")

//User hold user information.
type User struct {
	Name string
	Age  int
}

// Repository user data store.
type Repository interface {
	GetUsers() ([]*User, error)
	AddUser(u User) error
}

type service interface {
	GetUsers() ([]*User, error)
	AddUser(u User) error
}

var _ service = &UserService{}

// UserService hold all operations related to user.
type UserService struct {
	repo Repository
}

// New create a new/empty user service
func New() *UserService {
	return &UserService{}
}

// SetRepo set the repository to UserService for querying db
func (s *UserService) SetRepo(r Repository) {
	s.repo = r
}

// GetUsers return list of users. return error if no user found.
func (s *UserService) GetUsers() ([]*User, error) {
	users, err := s.repo.GetUsers()
	if len(users) == 0 {
		return []*User{}, ErrNotFound
	}
	return users, err
}

// AddUser add a new user
func (s *UserService) AddUser(u User) error {
	return s.repo.AddUser(u)
}
