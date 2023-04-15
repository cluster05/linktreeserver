package requesthandler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/cluster05/linktree/api/appresponse"
)

type InvalidArgument struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
	Tag   string      `json:"tag"`
	Param string      `json:"param"`
}

func jsonKeyBuilder(key string) string {
	return strings.ToLower(key[:1]) + key[1:]
}

func messageForTag(argument InvalidArgument) string {
	switch argument.Tag {
	case "required":
		return argument.Field + " is required field"
	case "email":
		return argument.Field + " is not valid"
	default:
		return "invalid tag"
	}
}

func BindData(c *gin.Context, req interface{}) bool {
	if c.ContentType() != "application/json" {
		message := fmt.Sprintf("%s only accepts Content-Type application/json", c.FullPath())

		c.JSON(http.StatusOK, appresponse.NewUnsupportedMediaTypeError(message))
		return false
	}
	if err := c.ShouldBind(req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			var invalidArgs []string

			for _, err := range errs {
				msgForTag := messageForTag(InvalidArgument{
					jsonKeyBuilder(err.Field()),
					err.Value(),
					err.Tag(),
					err.Param(),
				})
				invalidArgs = append(invalidArgs, msgForTag)
			}

			c.JSON(http.StatusOK, appresponse.NewBadRequestError(invalidArgs))
			return false
		}
		c.JSON(http.StatusOK, appresponse.NewInternalError("internal server error"))
		return false
	}

	return true
}
