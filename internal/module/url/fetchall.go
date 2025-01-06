package url

import (
	"errors"

	"github.com/Cattle0Horse/url-shortener/internal/global/database"
	"github.com/Cattle0Horse/url-shortener/internal/global/errs"
	"github.com/Cattle0Horse/url-shortener/internal/global/jwt"
	"github.com/Cattle0Horse/url-shortener/internal/model"
	"github.com/gin-gonic/gin"
)

type FeechAllRequest struct {
	page int `uri:"page" binding:"min=1"`
	size int `uri:"size" binding:"min=1"`
}

type FetchAllResponse struct {
	Urls []Url `json:"urls"`
}

func (resp *FetchAllResponse) ConvertFromModel(urls []*model.Url) {
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
	userID := payload.(jwt.Claims).UserId

	var req FeechAllRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errs.Fail(c, errs.InvalidPathParams.WithOrigin(err))
		return
	}

	u := database.Query.Url
	// 获取用户的所有短链接
	urls, _, err := u.WithContext(c.Request.Context()).Where(u.UserID.Eq(userID)).FindByPage((req.page-1)*req.size, req.size)
	if err != nil {
		errs.Fail(c, errs.ErrDatabase.WithOrigin(err))
		return
	}

	resp := &FetchAllResponse{}
	resp.ConvertFromModel(urls)

	errs.Success(c, resp)
}
