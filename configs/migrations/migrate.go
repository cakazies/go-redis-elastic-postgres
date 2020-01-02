package main

import (
	"log"

	"github.com/cakazies/go-redis-elastic-postgres/configs"
	"github.com/cakazies/go-redis-elastic-postgres/models"
)

// command go run go-redis-elastic-postgres/configs/migrations/migrate.go
func main() {
	log.Println("Loading ...")
	configs.InitASCII()

	err := configs.GetDB.AutoMigrate(&models.Room{}).Error
	if err != nil {
		log.Println("error : ", err)
		return
	}

	defer configs.GetDB.Close()
	log.Println("Done ...")
}
