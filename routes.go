package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	//Alice package allows use to chain middleware in a much more understandable and clearer way.
	// recoverPanic -> logrequest -> secureHeaders ->...
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	router := http.NewServeMux()
	router.HandleFunc("/", app.Home)

	//user
	router.HandleFunc("/user/signup", app.Signup)

	//Kitchen Features
	router.HandleFunc("/kitchen/bmr", app.GetBMR)

	//Workout Area
	router.HandleFunc("/workout/addworkoutlog", app.addNewWorkoutLog)
	router.HandleFunc("/workout/getworkoutlog", app.getWorkoutLog)
	router.HandleFunc("/workout/getallworkoutlogs", app.getAllWorkoutLogs)

	//...router -> secureHeaders -> logrequest -> recoverPanic -> client
	return standardMiddleware.Then(router)
}
