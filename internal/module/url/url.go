package url

import (
	"errors"
	"log/slog"
	"net/http"
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

	slog.Info("Create Url", "request", req)

	payload, ok := c.Get("payload")
	if !ok {
		errs.Fail(c, errs.ErrTokenInvalid.WithOrigin(errors.New("payload not found")))
		return
	}
	req.UserID = payload.(jwt.Claims).UserId

	var url model.Url
	req.ConvertToUrl(&url)

	// 生成短链接
	shortURL, err := urlService.Create(c.Request.Context(), req.OriginalUrl)
	if err != nil {
		errs.Fail(c, errs.InternalError.WithOrigin(err))
		return
	}

	c.JSON(http.StatusOK, CreateResponse{
		ShortUrl: shortURL,
	})
}

type UrlResponse struct {
	ShortCode   string    `json:"short_code"`
	OriginalUrl string    `json:"original_url"`
	ExpiryTime  time.Time `json:"expiry_time"`
	CreatedAt   time.Time `json:"created_at"`
}

func FetchAll(c *gin.Context) {
	// TODO: implement
}

func Delete(c *gin.Context) {
	shortCode := c.Param("code")
	if shortCode == "" {
		errs.Fail(c, errs.InvalidRequest.WithTips("short code is empty"))
		return
	}

	if err := urlService.Delete(c.Request.Context(), shortCode); err != nil {
		errs.Fail(c, errs.InternalError.WithOrigin(err))
		return
	}

	c.Status(http.StatusNoContent)
}

func Redirect(c *gin.Context) {
	shortCode := c.Param("code")
	if shortCode == "" {
		errs.Fail(c, errs.InvalidRequest.WithTips("short code is empty"))
		return
	}

	longURL, err := urlService.Resolve(c.Request.Context(), shortCode)
	if err != nil {
		errs.Fail(c, errs.NotFound.WithOrigin(err))
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, longURL)
}

func Update(c *gin.Context) {
	// TODO: implement
}
