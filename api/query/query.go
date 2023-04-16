package query

import (
	"encoding/json"

	"github.com/gin-gonic/gin"

	"github.com/cluster05/linktree/api/model"
)

func User(c *gin.Context) model.JWTPayload {
	reqUser := c.Request.Header.Get("user")

	var user model.JWTPayload
	err := json.Unmarshal([]byte(reqUser), &user)
	if err != nil {
		return model.JWTPayload{}
	}

	return user
}
