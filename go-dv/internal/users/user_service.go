package users

import (
	"errors"
)

// UserService handles user-related operations
type UserService struct {
	IUser UserInterface
}

// NewUserService creates a new UserService instance
func NewUserService(i_user UserInterface) *UserService {
	return &UserService{
		IUser: i_user,
	}
}

// log in a user is exists
func (s *UserService) Login(user UserDTO) (bool, error) {
	if err := user.Validate(); err != nil {
		return false, err
	}

	existingUser, err := s.IUser.Read(user.ID)
	if err != nil {
		return false, err
	}

	if existingUser.Email == user.Email && existingUser.Status {
		// Assuming password validation is handled elsewhere, e.g., hashing and comparing
		return true, nil
	}

	return false, errors.New("invalid credentials or user inactive")
}

// RegisterUser registers a new user
func (s *UserService) Register(user UserDTO) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if _, exists := s.IUser.Exists(user.Email); exists {
		return errors.New("email already registered")
	}

	return s.IUser.Create(user)
}

func (s *UserService) List() ([]UserDTO, error) {
	users, err := s.IUser.List()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Exists checks if a user exists by email
func (s *UserService) Exists(email string) (IUser *UserDTO, exists bool) {
	return s.IUser.Exists(email)
}

// GetUser retrieves a user by ID
func (s *UserService) Get(id int) (*UserDTO, error) {
	user, err := s.IUser.Read(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser updates an existing user
func (s *UserService) Update(user UserDTO) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if _, exists := s.IUser.Exists(user.Email); !exists {
		return errors.New("user does not exist")
	}

	return s.IUser.Update(user)
}

// DeleteUser deletes a user by ID
func (s *UserService) Delete(id int) bool {
	if err := s.IUser.Delete(id); err != nil {
		return false
	}
	return true
}
