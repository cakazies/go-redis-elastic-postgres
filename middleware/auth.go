package middleware

import (
	"encoding/base64"
	"os"
	"strings"

	ctr "github.com/cakazies/go-redis-elastic-postgres/controllers"
	"github.com/gin-gonic/gin"
)

// CheckMiddleware function for
func CheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get authorization in request
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
		if len(auth) != 2 || auth[0] != "Basic" {
			respondWithError(c, 401, "Authorization is wrong")
			return
		}

		// decode authorization with basic auth
		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)
		// get username and password env
		username := (os.Getenv("BASIC_USERNAME"))
		password := (os.Getenv("BASIC_PASSWORD"))

		if username != pair[0] || password != password {
			respondWithError(c, 401, "Authorization is wrong")
			return
		}
		c.Next()
	}
}

// respondWithError function for response in middleware if error
func respondWithError(c *gin.Context, code int, message string) {
	var resp ctr.Response
	resp.Code = code
	resp.Message = message
	c.AbortWithStatusJSON(code, gin.H{"response": resp})
}
