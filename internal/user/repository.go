package user

import (
	"errors"
	"gorm.io/gorm"
)

var ErrUserAlreadyExists = errors.New("user with this email already exists")

type UserRepository interface {
	CreateUser(user *UserDTO) error
	GetUserByEmail(email string) (*UserDTO, error)
	GetAllUsers() ([]UserDTO, error)
	// GetUserByID(id uint) (*UserDTO, error)
	// DeleteUser(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetUserByEmail(email string) (*UserDTO, error) {
	var user UserDTO
	result := r.db.Where(&UserDTO{Email:email}).First(&user);
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *repository) CreateUser(user *UserDTO) error {
	result := r.db.Create(user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return ErrUserAlreadyExists
		}
		return result.Error
	}

	return nil
}

func (r *repository) GetAllUsers() ([]UserDTO, error) {
	var user []UserDTO
	result := r.db.Find(&user);
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}