package testhelper

import (
	"encoding/json"
	"fmt"
	"github.com/cluster05/linktree/api/appresponse"
	"github.com/cluster05/linktree/api/config"
	"github.com/cluster05/linktree/api/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
)

func GenerateTestJWT(payload model.JWTPayload) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(config.AppConfig.JWTSecret))
}

type GinTestBuidler struct {
	ctx *gin.Context
	w   *httptest.ResponseRecorder
}

func NewGinTestBuidler() *GinTestBuidler {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	request := &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
		Method: "POST",
	}
	ctx.Request = request

	gb := &GinTestBuidler{
		ctx: ctx,
		w:   w,
	}
	return gb
}

func (gb *GinTestBuidler) WithRequest(req *http.Request) *GinTestBuidler {
	gb.ctx.Request = req
	return gb
}

func (gb *GinTestBuidler) WithToken(token string) *GinTestBuidler {
	gb.ctx.Request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	return gb
}

func (gb *GinTestBuidler) GetGinContext() *gin.Context {
	return gb.ctx
}

func (gb *GinTestBuidler) GetResponseRecorder() *httptest.ResponseRecorder {
	return gb.w
}

func (gb *GinTestBuidler) GetBody() appresponse.Response {

	bytes, err := io.ReadAll(gb.w.Body)
	if err != nil {
		panic(err)
	}

	response := appresponse.Response{}

	err = json.Unmarshal(bytes, &response)
	if err != nil {
		panic(err)
	}
	return response
}
