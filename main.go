package main

import (
	"log"
	"net/http"

	"github.com/Lionel-Wilson/My-Fitness-Aibou/internal"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", internal.Home)

	//user
	router.HandleFunc("/user/signup", internal.Signup)

	//Kitchen Features
	router.HandleFunc("/kitchen/bmr", internal.GetBMR)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)

}

/*gin version
func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	user := router.Group("/user")
	{
		user.POST("signup", internal.Signup)
	}

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}
*/
