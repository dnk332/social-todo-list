package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social-todo-list/common"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					c.AbortWithStatusJSON(http.StatusInternalServerError, common.ErrInterval(err))
				}
				panic(r)
			}
		}()
		c.Next()
	}
}
