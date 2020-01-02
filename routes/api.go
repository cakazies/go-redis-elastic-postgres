package routes

import (
	"os"

	"github.com/cakazies/go-redis-elastic-postgres/configs"
	ctr "github.com/cakazies/go-redis-elastic-postgres/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	configs.InitASCII()
}

// Run function for running this API
func Run() {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		room := v1.Group("/room")
		{
			roomCtr := ctr.RoomCtr{}
			room.GET("/detail/:room_id", roomCtr.GetRoom)
			room.POST("/create", roomCtr.InsertRoom)
			room.GET("/", roomCtr.GetRooms)
			room.POST("/update/:room_id", roomCtr.UpdateRoom)
			room.DELETE("/delete/:room_id", roomCtr.DeleteRoom)
		}
	}

	// run this router
	r.Run(os.Getenv("PORT_APPS"))
}
