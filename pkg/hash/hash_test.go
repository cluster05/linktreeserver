package hash_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/cluster05/linktree/pkg/hash"
)

var (
	samplePassword = "sample-password"
	wrongPassword  = "wrong-password"
)

var _ = Describe("Hash", func() {

	Context("Create password hash from plan text", func() {

		It("Create Password Hash", func() {
			hashPassword, err := hash.CreatePasswordHash(samplePassword)
			Expect(hashPassword).ShouldNot(BeEmpty())
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Check Password Hash", func() {
			hashPassword, _ := hash.CreatePasswordHash(samplePassword)

			isValid := hash.CheckPasswordHash(samplePassword, hashPassword)
			Expect(isValid).Should(BeTrue())
		})

		It("Check Password Hash", func() {
			hashPassword, _ := hash.CreatePasswordHash(samplePassword)

			isValid := hash.CheckPasswordHash(wrongPassword, hashPassword)
			Expect(isValid).Should(BeFalse())
		})
	})

})
