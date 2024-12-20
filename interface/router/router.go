package router

import (
	"github.com/gabrielmoura33/neoway-backend-test/interface/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(clientHandler *handler.ClientHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/status", clientHandler.Status)
	r.POST("/clients", clientHandler.CreateClient)
	r.GET("/clients", clientHandler.GetAllClients)
	r.GET("/clients/:document", clientHandler.GetClientByDocument)
	r.GET("/exists", clientHandler.ClientExists)

	return r
}
