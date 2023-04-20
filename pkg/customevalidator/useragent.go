package customevalidator

import (
	"github.com/go-playground/validator/v10"
)

var (
	userAgents = map[string]struct{}{
		"Mobile":    {},
		"Tablet":    {},
		"Desktop":   {},
		"SmartTV":   {},
		"Raspberry": {},
		"Bot":       {},
		"Other":     {},
	}
)

func GetUserAgents() map[string]struct{} {
	return userAgents
}

var UserAgentValidator validator.Func = func(fl validator.FieldLevel) bool {
	agent, ok := fl.Field().Interface().(string)
	if ok {
		_, isPresent := planType[agent]
		if !isPresent {
			return false
		}
	}
	return true
}
