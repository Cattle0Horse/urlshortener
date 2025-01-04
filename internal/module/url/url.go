package url

import (
	"errors"
	"time"

	"github.com/Cattle0Horse/url-shortener/config"
	"github.com/Cattle0Horse/url-shortener/internal/global/errs"
	"github.com/Cattle0Horse/url-shortener/internal/global/jwt"
	"github.com/Cattle0Horse/url-shortener/internal/model"
	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	OriginalUrl string `json:"original_url" binding:"required,url"`
	Duration    *int   `json:"duration" binding:"omitempty,min=1,max=100"`
	UserID      uint   `json:"user_id" binding:"-"`
}

func (r *CreateRequest) ConvertToUrl(url *model.Url) {
	url.OriginalUrl = r.OriginalUrl
	if r.Duration == nil {
		url.ExpiryTime = time.Now().Add(config.Get().Url.DefaultDuration)
	} else {
		url.ExpiryTime = time.Now().Add(time.Hour * time.Duration(*r.Duration))
	}
	url.UserID = r.UserID
}

type CreateResponse struct {
	ShortUrl string `json:"short_url"`
}

func Create(c *gin.Context) {
	var req CreateRequest
	if err := c.BindJSON(&req); err != nil {
		errs.Fail(c, errs.InvalidRequest.WithTips(err.Error()))
		return
	}
	log.Info("Create Url", "request", req) // "msg="Create Url", request={...}"
	payload, ok := c.Get("payload")
	if !ok {
		errs.Fail(c, errs.ErrTokenInvalid.WithOrigin(errors.New("payload not found")))
		return
	}
	req.UserID = payload.(jwt.Claims).UserId
}

func FetchAll(c *gin.Context) {
	// TODO: implement
}

func Delete(c *gin.Context) {
	// TODO: implement
}

func Redirect(c *gin.Context) {
	// TODO: implement
}

func Update(c *gin.Context) {
	// TODO: implement
}
