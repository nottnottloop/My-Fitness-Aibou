package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/Lionel-Wilson/My-Fitness-Aibou/pkg/models"
)

type UserFormData struct {
	UserName  string `json:"userName"`
	About     string `json:"about"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Country   string `json:"country"`
	Password  string `json:"password"`
	Dob       string `json:"dob"`
	Gender    string `json:"gender"`
}

type SignupResponse struct {
	FlashMessage string            `json:"flashMessage"`
	Errors       map[string]string `json:"errors"`
}

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_" + "`" + "{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	var signupData UserFormData

	err = json.Unmarshal(body, &signupData)
	if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Println(signupData)

	errors := make(map[string]string)

	if strings.TrimSpace(signupData.UserName) == "" {
		errors["Username"] = "Username cannot be blank"
	} else if utf8.RuneCountInString(signupData.Email) > 255 {
		errors["Username"] = "Username is too long (maximum is 255 characters)"
	}

	if strings.TrimSpace(signupData.Email) == "" {
		errors["Email"] = "Email cannot be blank"
	} else if utf8.RuneCountInString(signupData.Email) > 255 {
		errors["Email"] = "Email is too long (maximum is 255 characters)"
	} else if !EmailRX.MatchString(signupData.Email) {
		errors["Email"] = "Email is invalid"
	}

	if strings.TrimSpace(signupData.Password) == "" {
		errors["Password"] = "Password cannot be blank"
	} else if utf8.RuneCountInString(signupData.Password) < 10 {
		errors["Password"] = "Password is too short (minimum is 10 characters)"
	}

	if strings.TrimSpace(signupData.FirstName) == "" {
		errors["FirstName"] = "First name cannot be blank"
	}
	if strings.TrimSpace(signupData.LastName) == "" {
		errors["LastName"] = "Last name cannot be blank"
	}
	dob, err := time.Parse("2006-01-02", signupData.Dob)
	if err != nil {
		fmt.Println("Error parsing date of birth:", err)
		errors["DateOfBirth"] = "Error parsing date of birth"
		app.serverError(w, err)
	}
	response := SignupResponse{
		FlashMessage: "",
		Errors:       errors,
	}

	if len(errors) > 0 {
		response.FlashMessage = "You seem to have some incorrect fields. See errors below:"

		resultJson, err := json.Marshal(response)
		if err != nil {
			app.serverError(w, err)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(resultJson)
		return
	}

	err = app.users.Insert(signupData.UserName, signupData.FirstName, signupData.LastName, signupData.Gender, signupData.Country, signupData.Email, signupData.About, signupData.Password, dob)
	if err == models.ErrDuplicateEmail {
		response.FlashMessage = "You seem to have some incorrect fields. See errors below:"
		errors["Email"] = "Email is already in use"
		response.Errors = errors
		resultJson, err := json.Marshal(response)
		if err != nil {
			app.serverError(w, err)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(resultJson)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	response.FlashMessage = "Your signup was successful. Please log in."

	var resultJson []byte
	resultJson, err = json.Marshal(response)
	if err != nil {
		app.serverError(w, err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resultJson)
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authenticate and login the user...")
}
func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout the user...")
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
