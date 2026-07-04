package dto

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Response struct {
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
	Data    any    `json:"data,omitempty"`
}
