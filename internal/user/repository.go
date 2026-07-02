package user

import (
	"errors"

	"gorm.io/gorm"
)

var ErrUserAlreadyExists = errors.New("user with this email already exists")

type UserRepository interface {
	CreateUser(user *UserDTO) error
	// GetAllUsers() ([]UserDTO, error)
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
