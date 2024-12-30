package controller

import (
	"errors"
	"net/http"

	"github.com/Cattle0Horse/url-shortener/internal/schema"
	"github.com/Cattle0Horse/url-shortener/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Login(c *gin.Context) {
	var req schema.LoginRequest

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, schema.ErrorResponse{Message: err.Error()})
		return
	}

	resp, err := uc.userService.Login(c, req)

	if err != nil {
		c.JSON(http.StatusBadRequest, schema.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
func (uc *UserController) Register(c *gin.Context) {
	var req schema.RegisterReqeust

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, schema.ErrorResponse{Message: err.Error()})
		return
	}

	resp, err := uc.userService.Register(c, req)

	if err != nil {
		c.JSON(http.StatusBadRequest, schema.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
func (uc *UserController) ResetPassword(c *gin.Context) {
	var req schema.ResetPasswordReqeust
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, schema.ErrorResponse{Message: err.Error()})
		return
	}

	avaliable, err := uc.userService.IsEmailAvailable(c, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schema.ErrorResponse{Message: err.Error()})
		return
	}
	if !avaliable {
		c.SecureJSON(http.StatusBadRequest, schema.ErrorResponse{Message: schema.ErrEmailAleadyExist.Error()})
	}

	resp, err := uc.userService.ResetPassword(c, &req)
	if err != nil {
		if errors.Is(err, schema.ErrEmailCodeNotEqual) {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (uc *UserController) SendEmailCode(c *gin.Context) {
	email := c.Param("email")

	if err := uc.userService.SendEmailCode(c, email); err != nil {
		c.JSON(http.StatusInternalServerError, schema.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Email code sent successfully")
}
