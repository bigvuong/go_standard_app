package utils

import "github.com/gin-gonic/gin"

func JSONResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{"data": data})
}
