package middleware_test

import (
	"github.com/cluster05/linktree/api/config"
	"github.com/cluster05/linktree/api/middleware"
	"github.com/cluster05/linktree/api/model"
	"github.com/cluster05/linktree/pkg/testhelper"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Middleware", func() {

	gin.SetMode(gin.TestMode)

	Context("Auth Middleware", func() {

		It("validation empty header key", func() {
			tg := testhelper.NewGinTestBuidler()
			middleware.Auth(tg.GetGinContext())

			response := tg.GetBody()
			Expect(response.Data).Should(ContainSubstring(errInvalidToken.Error()))
			Expect(response.ResponseCode).Should(Equal(4000))
		})

		It("validation empty header value", func() {
			tg := testhelper.NewGinTestBuidler()
			middleware.Auth(tg.GetGinContext())

			response := tg.GetBody()
			Expect(response.Data).Should(ContainSubstring(errInvalidToken.Error()))
			Expect(response.ResponseCode).Should(Equal(4000))
		})

		It("validation random token value", func() {
			tg := testhelper.NewGinTestBuidler().WithToken("invalid-token-value")
			middleware.Auth(tg.GetGinContext())

			response := tg.GetBody()
			Expect(response.Data).Should(ContainSubstring(errInvalidToken.Error()))
			Expect(response.ResponseCode).Should(Equal(4000))
		})

		It("validation expired token value", func() {

			config.AppConfig.JWTSecret = sampleSecret

			jwtPayload := model.JWTPayload{
				IssuedAt:  time.Now().Unix(),
				ExpiredAt: time.Now().Unix(),
			}
			token, _ := testhelper.GenerateTestJWT(jwtPayload)
			By("Waiting for jwt to be expired")
			time.Sleep(time.Second * 2)

			tg := testhelper.NewGinTestBuidler().WithToken(token)
			middleware.Auth(tg.GetGinContext())

			response := tg.GetBody()
			Expect(response.Data).Should(ContainSubstring(errExpiredToken.Error()))
			Expect(response.ResponseCode).Should(Equal(4000))
		})

		It("validation expired token value", func() {

			config.AppConfig.JWTSecret = sampleSecret

			jwtPayload := model.JWTPayload{
				IssuedAt:  time.Now().Unix(),
				ExpiredAt: time.Now().Add(time.Minute).Unix(),
			}
			token, _ := testhelper.GenerateTestJWT(jwtPayload)

			tg := testhelper.NewGinTestBuidler().WithToken(token)
			middleware.Auth(tg.GetGinContext())

			user := tg.GetGinContext().Request.Header.Get("user")
			Expect(user).ShouldNot(BeEmpty())
		})

	})

})
