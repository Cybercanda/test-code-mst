package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Player struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	BirthPlace string `json:"birthPlace"`
}

var PlayerDatas = []Player{
	{Id: 1, Name: "Cristiano Ronaldo", Age: 38, BirthPlace: "Europe"},
	{Id: 2, Name: "Lionel Messi", Age: 36, BirthPlace: "South America"},
	{Id: 3, Name: "Karim Benzema", Age: 35, BirthPlace: "Europe"},
	{Id: 4, Name: "Erling Haaland", Age: 23, BirthPlace: "Europe"},
	{Id: 5, Name: "Kylian Mbappe", Age: 24, BirthPlace: "Europe"},
}

func GetPlayers(ctx *gin.Context) {
	birthFilter := ctx.Query("BirthPlace")

	var filterPlayers []Player
	if birthFilter != "" {
		for _, player := range PlayerDatas {
			if player.BirthPlace == birthFilter {
				filterPlayers = append(filterPlayers, player)
			}
		}
	} else {
		filterPlayers = PlayerDatas
	}

	ctx.JSON(http.StatusOK, gin.H{
		"players": filterPlayers,
	})
}

func GetPlayersById(ctx *gin.Context) {
	id := ctx.Param("id")
	condition := false
	var playerData Player

	for i, player := range PlayerDatas {
		// konversi id ke integer
		playerID, err := strconv.Atoi(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error_status":  "Bad Request",
				"error_message": "Invalid ID format",
			})
			return
		}
		if playerID == player.Id {
			condition = true
			playerData = PlayerDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Player with ID %V not found", id),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"player": playerData,
	})
}

func CreatePlayer(ctx *gin.Context) {
	var newPlayer Player

	if err := ctx.ShouldBindJSON(&newPlayer); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	maxID := 0
	for _, player := range PlayerDatas {
		if player.Id > maxID {
			maxID = player.Id
		}
	}

	newPlayer.Id = maxID + 1

	PlayerDatas = append(PlayerDatas, newPlayer)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "New Player created successfully",
		"player":  newPlayer,
	})
}

func main() {

}
