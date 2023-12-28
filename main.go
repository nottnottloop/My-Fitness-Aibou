package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	//Configuration

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	addr := os.Getenv("DEV_ADDRESS")
	/* command line - flag version
	addr := flag.String("addr", ":8080", "Http network address")
	flag.Parse()
	*/

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	router := http.NewServeMux()
	router.HandleFunc("/", app.Home)

	//user
	router.HandleFunc("/user/signup", app.Signup)

	//Kitchen Features
	router.HandleFunc("/kitchen/bmr", app.GetBMR)

	srv := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  router,
	}

	infoLog.Printf("Starting server on %s", addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

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
