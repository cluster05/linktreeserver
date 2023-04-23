package requesthandler_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/cluster05/linktree/pkg/requesthandler"
)

var _ = Describe("Requesthandler", func() {

	type TestStruct struct {
		Field1 int     `json:"field1"`
		Field2 string  `json:"field2" binding:"required"`
		Field3 float32 `json:"field3"`
	}

	Context("BindData", func() {

		It("Pass nil JSON object in request", func() {
			w := httptest.NewRecorder()
			request := &http.Request{
				Header: map[string][]string{
					"Content-Type": {""},
				},
			}
			ctx := getTestGinContext(w, request)

			var testStruct TestStruct
			isValid := requesthandler.BindData(ctx, &testStruct)
			Expect(isValid).Should(BeFalse())

			Expect(w.Body).Should(ContainSubstring("\"responseCode\":4000"))
			Expect(w.Body).Should(ContainSubstring("\"data\":\" only accepts Content-Type application/json\""))
		})

		It("Pass empty value for required JSON field", func() {

			bodyReq := TestStruct{
				Field1: 1,
				Field3: float32(10.10),
			}
			reqBody, _ := json.Marshal(bodyReq)

			w := httptest.NewRecorder()
			request := &http.Request{
				Header: map[string][]string{
					"Content-Type": {"application/json"},
				},
				Body: io.NopCloser(bytes.NewBuffer(reqBody)),
			}
			ctx := getTestGinContext(w, request)

			var testStruct TestStruct
			isValid := requesthandler.BindData(ctx, &testStruct)
			Expect(isValid).Should(BeFalse())

			Expect(w.Body).Should(ContainSubstring("\"responseCode\":4000"))
			Expect(w.Body).Should(ContainSubstring("\"data\":[\"field2 is required field\"]"))
		})

		It("Pass valid fields for required JSON field", func() {

			bodyReq := TestStruct{
				Field2: "valid",
			}
			reqBody, _ := json.Marshal(bodyReq)

			w := httptest.NewRecorder()
			request := &http.Request{
				Header: map[string][]string{
					"Content-Type": {"application/json"},
				},
				Body: io.NopCloser(bytes.NewBuffer(reqBody)),
			}
			ctx := getTestGinContext(w, request)

			var testStruct TestStruct
			isValid := requesthandler.BindData(ctx, &testStruct)
			Expect(isValid).Should(BeTrue())
		})

	})

})
