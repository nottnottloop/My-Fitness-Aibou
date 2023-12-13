package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Hello returns a greeting for the named person.
func Signup(c *gin.Context) {
	// Return a greeting that embeds the name in a message.
	message := "You have succesfully signed up!"

	c.String(http.StatusOK, message)
}
