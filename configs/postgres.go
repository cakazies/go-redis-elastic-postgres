package configs

import (
	"fmt"
	"log"
	"os"

	// repo for connect mysql
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// GetDB variable for saving DB
	GetDB *gorm.DB
)

func init() {
	host := os.Getenv("POSTGRESQL_HOST")
	port := os.Getenv("POSTGRESQL_PORT")
	user := os.Getenv("POSTGRESQL_USER")
	password := os.Getenv("POSTGRESQL_PASSWORD")
	dbname := os.Getenv("POSTGRESQL_DBNAME")

	// connect database in here
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		log.Println(err)
	}
	GetDB = db
}
