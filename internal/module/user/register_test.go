package user

import (
	"context"
	"testing"

	"github.com/Cattle0Horse/url-shortener/internal/global/database"
	"github.com/Cattle0Horse/url-shortener/internal/global/errs"
	"github.com/Cattle0Horse/url-shortener/pkg/tools"
	"github.com/Cattle0Horse/url-shortener/test"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	test.SetupEnvironment(t)
	selfInit()
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		req := CreateRequest{
			User: User{
				Email:    "test@test.com",
				Password: "123456",
			},
		}
		resp := test.DoRequest(t, Create, req)
		test.NoError(t, resp)
		u := database.Query.User
		userInfo, err := u.WithContext(ctx).Where(u.Email.Eq(req.Email)).First()
		require.NoError(t, err)
		require.Equal(t, true, tools.PasswordCompare(req.Password, userInfo.Password))
	})

	t.Run("PassWordTooShort", func(t *testing.T) {
		req := CreateRequest{
			User: User{
				Email:    "test@test.com",
				Password: "123",
			},
		}
		resp := test.DoRequest(t, Create, req)
		test.ErrorEqual(t, errs.InvalidRequest.WithTips(
			`Key: 'CreateRequest.User.Password' Error:Field validation for 'Password' failed on the 'min' tag`,
		), resp)
	})

	t.Run("EmailInvalid", func(t *testing.T) {
		req := CreateRequest{
			User: User{
				Email:    "test",
				Password: "123456",
			},
		}
		resp := test.DoRequest(t, Create, req)
		test.ErrorEqual(t, errs.InvalidRequest.WithTips(
			`Key: 'CreateRequest.User.Email' Error:Field validation for 'Email' failed on the 'email' tag`,
		), resp)
	})

	t.Run("EmailExist", func(t *testing.T) {
		req := CreateRequest{
			User: User{
				Email:    "test@test.com",
				Password: "123456",
			},
		}
		u := database.Query.User
		_, err := u.WithContext(ctx).Unscoped().Where(u.Email.Eq(req.Email)).Delete()
		require.NoError(t, err)
		resp := test.DoRequest(t, Create, req)
		test.NoError(t, resp)
		resp = test.DoRequest(t, Create, req)
		test.ErrorEqual(t, errs.HasExist, resp)
	})
}
