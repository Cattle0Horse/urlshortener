package url

import (
	"errors"

	"github.com/Cattle0Horse/url-shortener/internal/global/database"
	"github.com/Cattle0Horse/url-shortener/internal/global/errs"
	"github.com/Cattle0Horse/url-shortener/internal/global/jwt"
	"github.com/Cattle0Horse/url-shortener/internal/model"
	"github.com/gin-gonic/gin"
)

type FetchAllRequest struct {
	Page int `form:"page" binding:"required,min=1"`
	Size int `form:"size" binding:"required,min=1,max=100"`
}

type FetchAllResponse struct {
	Total int64 `json:"total"`
	Urls  []Url `json:"urls"`
}

func (resp *FetchAllResponse) ConvertFromModel(urls []*model.Url, total int64) {
	resp.Total = total
	resp.Urls = make([]Url, len(urls))
	for i, url := range urls {
		resp.Urls[i].ConvertFromModel(url)
	}
}

func FetchAll(c *gin.Context) {
	payload, ok := c.Get("payload")
	if !ok {
		errs.Fail(c, errs.InvaildToken.WithOrigin(errors.New("payload not found")))
		return
	}
	userID := payload.(*jwt.Claims).UserId

	var req FetchAllRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Info("fetch all url", "err", err)
		errs.Fail(c, errs.InvalidRequest.WithOrigin(err))
		return
	}

	u := database.Query.Url
	// 获取用户的所有短链接
	urls, total, err := u.WithContext(c.Request.Context()).Where(u.DeletedAt.IsNull(), u.UserID.Eq(userID)).FindByPage((req.Page-1)*req.Size, req.Size)
	if err != nil {
		log.Info("fetch all url", "err", err)
		errs.Fail(c, errs.ErrDatabase.WithOrigin(err))
		return
	}

	resp := &FetchAllResponse{}
	resp.ConvertFromModel(urls, total)

	errs.Success(c, resp)
}
