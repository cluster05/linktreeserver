package appresponse_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAppresponse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Appresponse Suite")
}
