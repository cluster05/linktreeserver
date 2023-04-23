package appresponse_test

import (
	"fmt"
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/cluster05/linktree/api/appresponse"
)

var _ = Describe("Appresponse", func() {

	Context("Error generation validation", func() {

		It("Check for NewAuthorizationError", func() {
			errortype := http.StatusUnauthorized
			statustext := http.StatusText(errortype)
			res := appresponse.NewAuthorizationError(statustext)

			Expect(res.ResponseCode).Should(Equal(appresponse.ErrorCode))
			Expect(res.HttpCode).Should(Equal(errortype))
			Expect(res.Data).Should(Equal(statustext))
		})

		It("Check for NewBadRequestError", func() {
			errortype := http.StatusBadRequest
			statustext := http.StatusText(errortype)
			res := appresponse.NewBadRequestError(statustext)

			Expect(res.ResponseCode).Should(Equal(appresponse.ErrorCode))
			Expect(res.HttpCode).Should(Equal(errortype))
			Expect(res.Data).Should(Equal(statustext))
		})

		It("Check for NewConflictError", func() {
			errortype := http.StatusConflict
			statustext := http.StatusText(errortype)
			res := appresponse.NewConflictError(statustext)

			Expect(res.ResponseCode).Should(Equal(appresponse.ErrorCode))
			Expect(res.HttpCode).Should(Equal(errortype))
			Expect(res.Data).Should(Equal(statustext))
		})

		It("Check for NewInternalError", func() {
			errortype := http.StatusInternalServerError
			statustext := http.StatusText(errortype)
			res := appresponse.NewInternalError(statustext)

			Expect(res.ResponseCode).Should(Equal(appresponse.ErrorCode))
			Expect(res.HttpCode).Should(Equal(errortype))
			Expect(res.Data).Should(Equal(statustext))
		})

		It("Check for NewNotFoundError", func() {
			errortype := http.StatusNotFound
			statustext := http.StatusText(errortype)
			res := appresponse.NewNotFoundError(statustext)

			Expect(res.ResponseCode).Should(Equal(appresponse.ErrorCode))
			Expect(res.HttpCode).Should(Equal(errortype))
			Expect(res.Data).Should(Equal(statustext))
		})

		It("Check for NewPayloadTooLargeError", func() {
			errortype := http.StatusRequestEntityTooLarge
			maxBodySize := int64(10)
			contentLength := int64(100)
			res := appresponse.NewPayloadTooLargeError(maxBodySize, contentLength)

			Expect(res.ResponseCode).Should(Equal(appresponse.ErrorCode))
			Expect(res.HttpCode).Should(Equal(errortype))
			Expect(res.Data).Should(Equal(fmt.Sprintf("Max payload size of %v exceeded. Actual payload size: %v", maxBodySize, contentLength)))
		})

		It("Check for NewServiceUnavailableError", func() {
			errortype := http.StatusServiceUnavailable
			res := appresponse.NewServiceUnavailableError()

			Expect(res.ResponseCode).Should(Equal(appresponse.ErrorCode))
			Expect(res.HttpCode).Should(Equal(errortype))
			Expect(res.Data).Should(Equal("Service unavailable or timed out"))
		})

		It("Check for NewUnsupportedMediaTypeError", func() {
			errortype := http.StatusUnsupportedMediaType
			statustext := http.StatusText(errortype)
			res := appresponse.NewUnsupportedMediaTypeError(statustext)

			Expect(res.ResponseCode).Should(Equal(appresponse.ErrorCode))
			Expect(res.HttpCode).Should(Equal(errortype))
			Expect(res.Data).Should(Equal(statustext))
		})

	})

	Context("Success generation validation", func() {

		It("Check for Success Response", func() {
			successtype := http.StatusOK
			statustext := http.StatusText(successtype)
			res := appresponse.NewSuccess(statustext)

			Expect(res.ResponseCode).Should(Equal(appresponse.SuccessCode))
			Expect(res.HttpCode).Should(Equal(successtype))
			Expect(res.Data).Should(Equal(statustext))
		})

	})

})
