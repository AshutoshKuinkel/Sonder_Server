package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func ThrowError(c *gin.Context, msg string) {
	c.Error(errors.New(msg))
	c.Abort()
}
