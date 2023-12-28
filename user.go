package main

import (
	"net/http"
)

type SignUpBody struct {
	FirstName string `json:"firstName"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Dob       string `json:"dob"`
	Gender    string `json:"gender"`
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("You have reached the my fitness aibou home"))

}
func (app *application) Signup(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("You have signed up successfully"))

}

/*gin version
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
*/
