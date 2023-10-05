package routers

import (
	"github.com/gin-gonic/gin"
	"test-code-mst/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/api/players", controllers.GetPlayers)
	router.GET("/api/players/:id", controllers.GetPlayersById)
	router.POST("api/players", controllers.CreatePlayer)

	return router
}

func main() {

}
