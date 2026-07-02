package user

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserDTO struct {
	gorm.Model
	Name     string `json:"name" validate:"required" `
	Email    string `json:"email" validate:"required,email" gorm:"unique"`
	Password string `json:"password" validate:"required,min=6"`
}

func (u *UserDTO) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *UserDTO) checkPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
