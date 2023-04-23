package middleware_test

import (
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

var (
	errInvalidToken = errors.New("token is invalid")
	errExpiredToken = errors.New("token has expired")

	sampleSecret = "sample-secret"
)

func TestMiddleware(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Middleware Suite")
}
