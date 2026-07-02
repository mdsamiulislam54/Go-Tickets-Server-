package user

import (
	"gotickets/internal/httpResponse"
	"gotickets/internal/user/dto"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewUserHandler(service *service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetAllUser(c *echo.Context) error {

	 response, err := h.service.GetAllUser()
	 if err != nil {
		 return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
			 Code:    http.StatusInternalServerError,
			 Message: "Failed to retrieved  user data",
			 Details: err.Error(),
		 })
	 }

	 return c.JSON(http.StatusCreated, response)
}
func (h *handler) loginUser(c *echo.Context) error {
	var req dto.LoginUserRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to bind request",
			Details: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to validate request",
			Details: err.Error(),
		})
	}

	 response, err := h.service.LoginUser(req)
	 if err != nil {
		 return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
			 Code:    http.StatusInternalServerError,
			 Message: "Failed to login user",
			 Details: err.Error(),
		 })
	 }

	 return c.JSON(http.StatusCreated, response)
}
func (h *handler) CreateUser(c *echo.Context) error {
	var req dto.CreateUserRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to bind request",
			Details: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to validate request",
			Details: err.Error(),
		})
	}

	 response, err := h.service.CreateUser(&req)
	 if err != nil {
		 return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
			 Code:    http.StatusInternalServerError,
			 Message: "Failed to create user",
			 Details: err.Error(),
		 })
	 }

	 return c.JSON(http.StatusCreated, response)
}