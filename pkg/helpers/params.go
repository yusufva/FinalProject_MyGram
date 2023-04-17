package helpers

import (
	"final-project/pkg/errrs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetParamsId(c *gin.Context, key string) (int, errrs.MessageErr) {
	value := c.Param(key)

	id, err := strconv.Atoi(value)

	if err != nil {
		return 0, errrs.NewBadRequest("invalid parameter id")
	}

	return id, nil
}
