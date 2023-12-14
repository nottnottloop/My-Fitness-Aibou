package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignUpBody struct {
	FirstName string `json:"firstName"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Dob       string `json:"dob"`
	Gender    string `json:"gender"`
}

// Hello returns a greeting for the named person.
func Signup(c *gin.Context) {
	var signUpDetails SignUpBody

	signUpDetails.FirstName = c.Query("firstname")
	signUpDetails.Surname = c.Query("surname")
	signUpDetails.Email = c.Query("email")
	signUpDetails.Password = c.Query("password")
	signUpDetails.Dob = c.Query("dob")
	signUpDetails.Gender = c.Query("gender")

	// Return a greeting that embeds the name in a message.
	message := "You have succesfully signed up!"

	c.JSON(http.StatusOK, gin.H{
		"message":      message,
		"User Details": signUpDetails,
	})
}
