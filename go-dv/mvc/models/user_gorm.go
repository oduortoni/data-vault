/*
*
* Project: go-dv
* File: go-dv/mvc/models/user_gorm.go
* Description: GORM implementation of the UserInterface (go-dv/internal/users/user_interface.go)
*
*/
package models

import (
	"dv/internal/users"
	"errors"
	"gorm.io/gorm"
)

type GormUserStore struct {
	DB *gorm.DB
}


type GormUserModel struct {
	gorm.Model
	Username   string `gorm:"not null"`
	Email  string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Status bool
}

func NewGormUserModel(db *gorm.DB) *GormUserStore {
	// Run auto-migration
	_ = db.AutoMigrate(&GormUserModel{})
	return &GormUserStore{DB: db}
}

func (s *GormUserStore) Create(user users.UserDTO) error {
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
		Username:   user.Username,
		Email:  user.Email,
		Password: user.Password,
		Status: user.Status,
	}

	return s.DB.Create(&model).Error
}

func (s *GormUserStore) Read(id int) (*users.UserDTO, error) {
	var model GormUserModel
	err := s.DB.First(&model, id).Error
	if err != nil {
		return nil, err
	}

	return &users.UserDTO{
		ID:     int(model.ID),
		Username:   model.Username,
		Email:  model.Email,
		Status: model.Status,
	}, nil
}

func (s *GormUserStore) Update(user users.UserDTO) error {
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

func (s *GormUserStore) Delete(id int) error {
	return s.DB.Delete(&GormUserModel{}, id).Error
}

func (s *GormUserStore) Exists(email string) (*users.UserDTO, bool) {
	var count int64
	user := GormUserModel{}

	s.DB.Model(&GormUserModel{}).Where("email = ?", email).Count(&count)
	if count > 0 {
		s.DB.Where("email = ?", email).First(&user)
		return &users.UserDTO{
			Email: user.Email,
			Username: user.Username,
			Password: user.Password,
			Status: user.Status,
		}, true
	} else {
		return nil, false
	}
}

func (s *GormUserStore) List() ([]users.UserDTO, error) {
	var models []GormUserModel
	err := s.DB.Find(&models).Error
	if err != nil {
		return nil, err
	}

	var dtos []users.UserDTO
	for _, m := range models {
		dtos = append(dtos, users.UserDTO{
			ID:     int(m.ID),
			Username:   m.Username,
			Email:  m.Email,
			Status: m.Status,
		})
	}

	return dtos, nil
}
