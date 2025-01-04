package test

import (
	"bytes"
	"encoding/json"
	"github.com/Cattle0Horse/url-shortener/internal/global/errs"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func DoRequest(t *testing.T, handlerFunc gin.HandlerFunc, request any) (response errs.ResponseBody) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	requestBytes, err := json.Marshal(request)
	require.NoError(t, err)
	c.Request = httptest.NewRequest(http.MethodPost, "/test", bytes.NewReader(requestBytes))
	handlerFunc(c)
	require.NoError(t, json.NewDecoder(w.Body).Decode(&response))
	return
}
