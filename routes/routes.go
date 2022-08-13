package routes

import (
	"github.com/gin-gonic/gin"
	"service/controllers"
)

func Routes(g *gin.RouterGroup) {
	g.POST("/add", controllers.Add())
}
