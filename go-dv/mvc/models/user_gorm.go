/*
*
* Project: go-dv
* File: go-dv/mvc/models/user_gorm.go
* Description: GORM implementation of the IUserRepository (go-dv/internal/users/user_interface.go)
*
 */
package models

import (
	"errors"

	"dv/internal/users"

	"gorm.io/gorm"
)

type GormUserModel struct {
	gorm.Model
	Username string `gorm:"not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Status   bool
}

type GormUserRepository struct {
	DB *gorm.DB
}

var _ users.IUserRepository = (*GormUserRepository)(nil)

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	// Run auto-migration
	_ = db.AutoMigrate(&GormUserModel{})
	return &GormUserRepository{DB: db}
}

func (s *GormUserRepository) Create(user users.UserDTO) error {
	if err := user.Validate(); err != nil {
		return err
	}

	_, exists := s.Exists(user.Email)
	if exists {
		return errors.New("email already in use")
	}

	pwdHash, err := user.HashPassword()
	if err != nil {
		return err
	}

	user.Password = pwdHash

	model := GormUserModel{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Status:   user.Status,
	}

	return s.DB.Create(&model).Error
}

func (s *GormUserRepository) Read(id int) (*users.UserDTO, error) {
	var model GormUserModel
	err := s.DB.First(&model, id).Error
	if err != nil {
		return nil, err
	}

	return &users.UserDTO{
		ID:       int(model.ID),
		Username: model.Username,
		Email:    model.Email,
		Status:   model.Status,
	}, nil
}

func (s *GormUserRepository) Update(user users.UserDTO) error {
	if err := user.Validate(); err != nil {
		return err
	}

	var model GormUserModel
	err := s.DB.First(&model, user.ID).Error
	if err != nil {
		return err
	}

	model.Username = user.Username
	model.Email = user.Email
	model.Status = user.Status

	return s.DB.Save(&model).Error
}

func (s *GormUserRepository) Delete(id int) error {
	return s.DB.Delete(&GormUserModel{}, id).Error
}

func (s *GormUserRepository) Exists(email string) (*users.UserDTO, bool) {
	var count int64
	user := GormUserModel{}

	s.DB.Model(&GormUserModel{}).Where("email = ?", email).Count(&count)
	if count > 0 {
		s.DB.Where("email = ?", email).First(&user)
		return &users.UserDTO{
			Email:    user.Email,
			Username: user.Username,
			Password: user.Password,
			Status:   user.Status,
		}, true
	} else {
		return nil, false
	}
}

func (s *GormUserRepository) List() ([]users.UserDTO, error) {
	var models []GormUserModel
	err := s.DB.Find(&models).Error
	if err != nil {
		return nil, err
	}

	var dtos []users.UserDTO
	for _, m := range models {
		dtos = append(dtos, users.UserDTO{
			ID:       int(m.ID),
			Username: m.Username,
			Email:    m.Email,
			Status:   m.Status,
		})
	}

	return dtos, nil
}
