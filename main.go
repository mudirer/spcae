package main

import (
	"github.com/gin-gonic/gin"
	"space/res"
	"space/routers"
)

func main() {


	gin.SetMode(res.Config.Server.RunMode)

	router := routers.InitRouter()


	router.Run(":5050")

}
