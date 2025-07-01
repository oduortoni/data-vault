package models

import (
	"errors"

	"dv/internal/users"
)

// Users is a slice of UserDTOs that implements the IUserRepository
type InternalUserRepository struct {
	users           []users.UserDTO
	autoIncrementID int
}

// Statically check that * InternalUserRepository  implements the UserRepository interface.
var _ users.IUserRepository = (*InternalUserRepository)(nil)

func NewInternalUserRepository() *InternalUserRepository {
	return &InternalUserRepository{
		users:           []users.UserDTO{},
		autoIncrementID: 1,
	}
}

// Create adds a new user to the list
func (u *InternalUserRepository) Create(user users.UserDTO) error {
	if err := user.Validate(); err != nil {
		return err
	}
	if _, exists := u.Exists(user.Email); exists {
		return errors.New("user with this email already exists")
	}
	user.ID = u.autoIncrementID
	u.autoIncrementID++
	u.users = append(u.users, user)

	return nil
}

// Read returns a user by ID
func (u *InternalUserRepository) Read(id int) (*users.UserDTO, error) {
	for _, user := range u.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

// Update modifies an existing user by ID
func (u *InternalUserRepository) Update(updated users.UserDTO) error {
	if err := updated.Validate(); err != nil {
		return err
	}
	for i, user := range u.users {
		if user.ID == updated.ID {
			u.users[i] = updated
			return nil
		}
	}
	return errors.New("user not found")
}

// Delete removes a user by ID
func (u *InternalUserRepository) Delete(id int) error {
	for i, user := range u.users {
		if user.ID == id {
			u.users = append(u.users[:i], u.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

// Exists checks if a user exists by email
func (u *InternalUserRepository) Exists(email string) (*users.UserDTO, bool) {
	for _, user := range u.users {
		if user.Email == email {
			return &user, true
		}
	}
	return nil, false
}

// List returns all users
func (u *InternalUserRepository) List() ([]users.UserDTO, error) {
	return u.users, nil
}
