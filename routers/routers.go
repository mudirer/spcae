package routers

import (
	"space/routers/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "space/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
)


func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(cors.Default())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//r.Use(middlewares.CORS())
	r.POST("/auth", api.Register)
	r.PUT("/updavatar", api.UpdAvatar)
	r.DELETE("/delete/:id", api.DeleteSpace)
	r.GET("/get/:id", api.GetSpace)
	r.GET("/getlist/:sex", api.GetList)



	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/test",api.Get1)
	return r
}