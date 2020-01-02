package controllers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cakazies/go-redis-elastic-postgres/configs"
	"github.com/go-redis/redis"
)

var (
	dBName = os.Getenv("CONFIGDB_DBNAME")
)

type (
	// Response struct for response client
	Response struct {
		Status  string `json:"status,omitempty"`
		Code    int    `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	}
)

// LogError function for log error in file
func LogError(data string, err error, function string) {
	currentTime := time.Now()
	today := currentTime.Format("2006-01-02")
	location := fmt.Sprintf("utils/logs/%s.log", today)

	f, err := os.OpenFile(location, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(fmt.Sprintf("(%s) %s : %s", function, data, err))
}

func setRedis(key string, value string) error {
	client := configs.DBRedis
	err := client.Set(key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func getRedis(key string) (string, bool, error) {
	client := configs.DBRedis
	result, err := client.Get(key).Result()
	if err == redis.Nil {
		// this redis not error but data is doesn't exist
		return "", true, nil
	} else if err != nil {
		// this redis error
		return "", false, err
	}

	return result, true, nil
}

func checkingRedis(key string) ([]byte, bool) {
	data, ok, err := getRedis(key)
	if err != nil {
		LogError("Error Redis", err, "room/GetRoom")
		return nil, false
	} else if ok && err == nil && data == "" {
		return nil, false
	}
	// if false get from DB , if true get in redis
	return []byte(data), true
}
