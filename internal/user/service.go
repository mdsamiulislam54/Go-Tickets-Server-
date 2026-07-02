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
	err := s.repo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	response := &dto.Response{
		Message: "User created successfully",
		Data: UserDTO{
			Name:  user.Name,
			Email: user.Email,
		},
	}

	return response, nil
}
