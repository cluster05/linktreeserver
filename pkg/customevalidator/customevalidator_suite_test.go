package customevalidator_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCustomevalidator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Customevalidator Suite")
}
