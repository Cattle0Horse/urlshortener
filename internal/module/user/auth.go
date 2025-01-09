package user

import (
	"github.com/Cattle0Horse/url-shortener/internal/global/database"
	"github.com/Cattle0Horse/url-shortener/internal/global/errs"
	"github.com/Cattle0Horse/url-shortener/internal/global/jwt"
	"github.com/Cattle0Horse/url-shortener/pkg/tools"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type LoginRequst struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}
type LoginResponse struct {
	AccessToken string `json:"access_token"`
	UserID      uint   `json:"user_id"`
	Email       string `json:"email"`
}

func Login(c *gin.Context) {
	var req LoginRequst
	if err := c.BindJSON(&req); err != nil {
		errs.Fail(c, errs.InvalidRequest.WithTips(err.Error()))
		return
	}
	u := database.Query.User
	userInfo, err := u.WithContext(c.Request.Context()).Where(u.Email.Eq(req.Email)).First()
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		errs.Fail(c, errs.NotFound)
		return
	case err != nil:
		errs.Fail(c, errs.ErrDatabase.WithOrigin(err))
		return
	}
	if !tools.PasswordCompare(req.Password, userInfo.Password) {
		errs.Fail(c, errs.InvalidPassword)
		return
	}
	token, err := jwt.CreateToken(jwt.Payload{UserId: userInfo.ID})
	if err != nil {
		errs.Fail(c, errs.FailedCreateToken.WithOrigin(err))
		return
	}
	errs.Success(c, LoginResponse{
		AccessToken: token,
		UserID:      userInfo.ID,
		Email:       req.Email,
	})
}

type MeResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

// TODO: token缓存
// func Me(c *gin.Context) {
// 	payload, ok := c.Get("payload")
// 	if !ok {
// 		errs.Fail(c, errs.InvaildToken.WithOrigin(errors.New("payload not found")))
// 		return
// 	}
// 	userId := payload.(*jwt.Claims).UserId

// 	u := database.Query.User
// 	userInfo, err := u.WithContext(c.Request.Context()).Where(u.ID.Eq(userId)).First()
// 	if err != nil {
// 		errs.Fail(c, errs.ErrDatabase.WithOrigin(err))
// 		return
// 	}

// 	errs.Success(c)
// }
