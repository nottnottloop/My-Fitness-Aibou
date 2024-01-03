package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	msssql "github.com/Lionel-Wilson/My-Fitness-Aibou/pkg/models/mssql"
	"github.com/joho/godotenv"

	//_ "github.com/denisenkom/go-mssqldb"
	//_ "github.com/microsoft/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog    *log.Logger
	infoLog     *log.Logger
	workoutLogs *msssql.WorkoutLogModel
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func openDB(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func buildConnectionString() string {
	//server := os.Getenv("DEV_SERVER")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	database := os.Getenv("DATABASE")

	connectionString := fmt.Sprintf(`%s:%s@/%s?parseTime=true`, user, password, database)
	return connectionString
}

func main() {
	//Configuration
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	addr := os.Getenv("DEV_ADDRESS")
	connectionString := buildConnectionString()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(connectionString)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		errorLog:    errorLog,
		infoLog:     infoLog,
		workoutLogs: &msssql.WorkoutLogModel{DB: db},
	}

	srv := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}
