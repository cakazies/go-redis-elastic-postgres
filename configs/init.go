package configs

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

// InitEnv function for first load ENV file
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file", err)
	}
}

// InitASCII this function for load in first time, for beauty command line, and link for make it
func InitASCII() {
	fmt.Println(`
	_____ _____       ______ ___________ _____ _____ 
	|  __ \  _  |      | ___ \  ___|  _  \_   _/  ___|
	| |  \/ | | |______| |_/ / |__ | | | | | | \  --. 
	| | __| | | |______|    /|  __|| | | | | |   --. \
	| |_\ \ \_/ /      | |\ \| |___| |/ / _| |_/\__/ /
	 \____/\___/       \_| \_\____/|___/  \___/\____/ `)
	// http://patorjk.com/software/taag/#p=display&f=Doom&t=GO-REDIS
}
