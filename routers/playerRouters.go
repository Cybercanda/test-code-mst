package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"test-code-mst/controllers"
)

type AppConfig struct {
	Database struct {
		ConnectingString string `mapstructure:"ConnectingString"`
	} `mapstructure:"Database"`
}

var Config AppConfig

func Init() (*gin.Engine, error) {
	viper.SetConfigFile("appsetings.json")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading configuration file: %s", err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling configuration: %s", err)
	}

	fmt.Printf("Database Connection String: %s \n", Config.Database.ConnectingString)

	router := StartServer()

	router.GET("/api/settings", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"connectionString": Config.Database.ConnectingString,
		})
	})

	return router, nil
}

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/api/players", controllers.GetPlayers)
	router.GET("/api/players/:id", controllers.GetPlayersById)
	router.POST("api/players", controllers.CreatePlayer)

	return router
}

func main() {

}
