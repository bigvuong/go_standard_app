package swagger

import (
	"github.com/gin-gonic/gin"
	_ "standard_app/docs"
   	swaggerFiles "github.com/swaggo/files"
   	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterSwagger đăng ký các route Swagger cho ứng dụng
func RegisterSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
