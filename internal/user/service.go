package user

import "gotickets/internal/user/dto"

type service struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateUser(req *dto.CreateUserRequest) (*dto.Response, error) {
	user := UserDTO{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password, 
	}

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
