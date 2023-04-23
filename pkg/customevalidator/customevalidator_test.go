package customevalidator_test

import (
	"github.com/cluster05/linktree/pkg/customevalidator"
	"github.com/go-playground/validator/v10"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CustomValidator", func() {
	var (
		validate *validator.Validate
	)

	BeforeEach(func() {
		validate = validator.New()
		validate.RegisterValidation("plantype", customevalidator.PlanTypeValidator)
		validate.RegisterValidation("subscriptiontype", customevalidator.SubscriptionTypeValidator)
		validate.RegisterValidation("useragent", customevalidator.UserAgentValidator)
	})

	Context("plantype", func() {

		Context("GetPlanType function", func() {
			It("should return a non-nil map with 3 keys", func() {
				pt := customevalidator.GetPlanType()
				Expect(pt).ShouldNot(BeNil())
				Expect(len(pt)).Should(Equal(3))
			})
		})

		Context("PlanTypeValidator function", func() {
			type TestStruct struct {
				Agent string `json:"agent" validate:"required,plantype"`
			}

			It("should validate the plan type correctly", func() {
				// Test case 1: Valid input
				ts := TestStruct{Agent: "FREE"}
				errs := validate.Struct(ts)
				Expect(errs).Should(BeNil())

				// Test case 2: Invalid input
				ts = TestStruct{Agent: "INVALID"}
				errs = validate.Struct(ts)
				Expect(errs).ShouldNot(BeNil())

				// Test case 3: Empty input
				ts = TestStruct{}
				errs = validate.Struct(ts)
				Expect(errs).ShouldNot(BeNil())
			})
		})
	})

	Context("subscriptiontype", func() {

		Context("GetSubscriptionType function", func() {
			It("should return a non-nil map with 2 keys", func() {
				st := customevalidator.GetSubscriptionType()
				Expect(st).ShouldNot(BeNil())
				Expect(len(st)).Should(Equal(2))
			})
		})

		Context("SubscriptionTypeValidator function", func() {
			type TestStruct struct {
				SubscriptionType string `json:"subscriptiontype" validate:"required,subscriptiontype"`
			}

			It("should validate the subscription type correctly", func() {
				// Test case 1: Valid input
				ts := TestStruct{SubscriptionType: "MONTHLY"}
				errs := validate.Struct(ts)
				Expect(errs).Should(BeNil())

				// Test case 2: Invalid input
				ts = TestStruct{SubscriptionType: "INVALID"}
				errs = validate.Struct(ts)
				Expect(errs).ShouldNot(BeNil())

				// Test case 3: Empty input
				ts = TestStruct{}
				errs = validate.Struct(ts)
				Expect(errs).ShouldNot(BeNil())
			})
		})

	})

	Context("useragent", func() {

		Context("GetUserAgents function", func() {
			It("should return a non-nil map with 7 keys", func() {
				ua := customevalidator.GetUserAgents()
				Expect(ua).ShouldNot(BeNil())
				Expect(len(ua)).Should(Equal(7))
			})
		})

		Context("UserAgentValidator function", func() {
			type TestStruct struct {
				UserAgent string `json:"useragent" validate:"required,useragent"`
			}

			It("should validate the user agent correctly", func() {
				// Test case 1: Valid input
				ts := TestStruct{UserAgent: "Mobile"}
				errs := validate.Struct(ts)
				Expect(errs).Should(BeNil())

				// Test case 2: Invalid input
				ts = TestStruct{UserAgent: "INVALID"}
				errs = validate.Struct(ts)
				Expect(errs).ShouldNot(BeNil())

				// Test case 3: Empty input
				ts = TestStruct{}
				errs = validate.Struct(ts)
				Expect(errs).ShouldNot(BeNil())
			})
		})
	})
})
