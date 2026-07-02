package user

import (
	"errors"
	"fmt"
	"gotickets/internal/user/dto"
)

var userCredentialError = errors.New("Invalid User Credential")

type service struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) LoginUser(req dto.LoginUserRequest) (*dto.Response, error) {

	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		fmt.Println("Database Error:", err)
		return nil, err
	}

	fmt.Println("========== LOGIN DEBUG ==========")
	fmt.Println("Email:", req.Email)
	fmt.Println("Request Password:", req.Password)
	fmt.Println("DB Password:", user.Password)

	user.checkPassword(req.Password)

	// if err != nil {
	// 	return nil, userCredentialError
	// }

	fmt.Println("Password Matched Successfully!")

	return &dto.Response{
		Message: "User Logging successfully",
		Data: &dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}
func (s *service) CreateUser(req *dto.CreateUserRequest) (*dto.Response, error) {
	user := UserDTO{
		Name:  req.Name,
		Email: req.Email,
	}
	user.HashPassword()

	if err := s.repo.CreateUser(&user); err != nil {
		return nil, err
	}

	return &dto.Response{
		Message: "User created successfully",
		Data: &dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}
func (s *service) GetAllUser() (*dto.Response, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	response := make([]dto.UserResponse, 0, len(users))

	for _, user := range users {
		response = append(response, dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return &dto.Response{
		Message: "User retrieved successfully",
		Data: response,
	}, nil
}
