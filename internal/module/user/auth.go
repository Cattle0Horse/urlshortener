package user

import (
	"github.com/Cattle0Horse/url-shortener/internal/global/database/mysql"
	"github.com/Cattle0Horse/url-shortener/internal/global/errs"
	"github.com/Cattle0Horse/url-shortener/internal/global/jwt"
	"github.com/Cattle0Horse/url-shortener/tools"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	var req User
	if err := c.BindJSON(&req); err != nil {
		errs.Fail(c, errs.InvalidRequest.WithTips(err.Error()))
		return
	}
	u := mysql.Query.User
	userInfo, err := u.WithContext(c.Request.Context()).Where(u.Email.Eq(req.Email)).First()
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		errs.Fail(c, errs.NotFound)
		return
	case err != nil:
		errs.Fail(c, errs.DatabaseError.WithOrigin(err))
		return
	}
	if !tools.PasswordCompare(req.Password, userInfo.Password) {
		errs.Fail(c, errs.InvalidPassword)
		return
	}
	token, err := jwt.CreateToken(jwt.Payload{UserId: userInfo.ID})
	if err != nil {
		errs.Fail(c, errs.ErrCreateToken.WithOrigin(err))
		return
	}
	errs.Success(c, map[string]string{"token": token})
}
