package main

import (
	"log"
	"net/http"
	"os"

	handlers "github.com/Lionel-Wilson/My-Fitness-Aibou/handlers"
	"github.com/joho/godotenv"
)

func main() {
	//Configuration

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	addr, exists := os.LookupEnv("DEV_ADDRESS")
	if !exists {
		log.Fatal(exists)
	}
	/* command line - flag version
	addr := flag.String("addr", ":8080", "Http network address")
	flag.Parse()
	*/
	router := http.NewServeMux()
	router.HandleFunc("/", handlers.Home)

	//user
	router.HandleFunc("/user/signup", handlers.Signup)

	//Kitchen Features
	router.HandleFunc("/kitchen/bmr", handlers.GetBMR)

	log.Printf("Starting server on %s", addr)
	err = http.ListenAndServe(addr, router)
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
