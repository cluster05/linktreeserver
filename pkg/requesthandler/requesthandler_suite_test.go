package requesthandler_test

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRequesthandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Requesthandler Suite")
}

var _ = BeforeSuite(func() {
})

var _ = AfterSuite(func() {

})

func getTestGinContext(w *httptest.ResponseRecorder, request *http.Request) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = request
	return ctx
}
