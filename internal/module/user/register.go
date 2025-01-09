package user

import (
	"errors"

	"github.com/Cattle0Horse/url-shortener/internal/global/database"
	"github.com/Cattle0Horse/url-shortener/internal/global/errs"
	"github.com/Cattle0Horse/url-shortener/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type CreateRequest struct {
	User
}

func Create(c *gin.Context) {
	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errs.Fail(c, errs.InvalidRequest.WithTips(err.Error()))
		return
	}
	user := &model.User{}
	_ = copier.Copy(user, &req)
	log.Info("Creating User", "user email", user.Email)
	err := database.Query.User.WithContext(c.Request.Context()).Create(user)
	switch {
	case errors.Is(err, gorm.ErrDuplicatedKey):
		log.Info("email exist", "email", user.Email)
		errs.Fail(c, errs.HasExist.WithOrigin(errs.HasExist))
		return
	case err != nil:
		log.Error("database error", "error", err)
		errs.Fail(c, errs.ErrDatabase.WithOrigin(err))
		return
	}

	errs.Success(c, nil)
}
