package customevalidator

import (
	"github.com/go-playground/validator/v10"
)

var (
	planType = map[string]struct{}{
		"FREE":  {},
		"BASIC": {},
		"PRO":   {},
	}
)

func GetPlanType() map[string]struct{} {
	return planType
}

var PlanTypeValidator validator.Func = func(fl validator.FieldLevel) bool {
	agent, ok := fl.Field().Interface().(string)
	if ok {
		_, isPresent := planType[agent]
		if !isPresent {
			return false
		}
	}
	return true
}
