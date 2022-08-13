package main

import (
	"github.com/gin-gonic/gin"
	"log"
	m "service/middleware"
	"service/routes"
)

var (
	logger = m.Logger{}
)

func init() {
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("%v", err)
		}
	}()
	router := gin.Default()

	api := router.Group("/api/")
	routes.Routes(api)

	logger.LogInfo("Start service")
	logger.LogErr(router.Run(":80"))
}

