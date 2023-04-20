package customevalidator

import (
	"github.com/go-playground/validator/v10"
)

var (
	subscriptionType = map[string]struct{}{
		"MONTHLY": {},
		"YEARLY":  {},
	}
)

func GetSubscriptionType() map[string]struct{} {
	return subscriptionType
}

var SubscriptionTypeValidator validator.Func = func(fl validator.FieldLevel) bool {
	agent, ok := fl.Field().Interface().(string)
	if ok {
		_, isPresent := subscriptionType[agent]
		if !isPresent {
			return false
		}
	}
	return true
}
