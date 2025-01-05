package user

import (
	"github.com/Cattle0Horse/url-shortener/internal/global/database/mysql"
	"github.com/Cattle0Horse/url-shortener/internal/global/errs"
	"github.com/Cattle0Horse/url-shortener/internal/model"
	"github.com/Cattle0Horse/url-shortener/tools"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type User struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

type CreateRequest struct {
	User
}

func Create(c *gin.Context) {
	var req CreateRequest
	if err := c.BindJSON(&req); err != nil {
		errs.Fail(c, errs.InvalidRequest.WithTips(err.Error()))
		return
	}
	user := &model.User{}
	_ = copier.Copy(user, &req)
	log.Info("Creating User", "user email", user.Email)
	err := mysql.Query.User.WithContext(c.Request.Context()).Create(user)
	switch {
	case tools.IsDuplicateKeyError(err):
		log.Info("email exist", "email", user.Email)
		errs.Fail(c, errs.HasExist.WithOrigin(err))
		return
	case err != nil:
		log.Error("database error", "error", err)
		errs.Fail(c, errs.DatabaseError.WithOrigin(err))
		return
	}

	errs.Success(c, nil)
}
