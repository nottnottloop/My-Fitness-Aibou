package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	//Alice package allows use to chain middleware in a much more understandable and clearer way.
	// recoverPanic -> logrequest -> secureHeaders ->...
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	//router := http.NewServeMux()
	// using bmizernay/pat package to implement RESTful routes
	router := pat.New()
	router.Get("/", http.HandlerFunc(app.Home))

	//user
	router.Post("/user/signup", http.HandlerFunc(app.Signup))

	//Kitchen Features
	router.Post("/kitchen/bmr", http.HandlerFunc(app.GetBMR))

	//Workout Area
	router.Get("/workout/getworkoutlog", http.HandlerFunc(app.getWorkoutLog))
	router.Get("/workout/getallworkoutlogs", http.HandlerFunc(app.getAllWorkoutLogs))
	router.Post("/workout/addworkoutlog", http.HandlerFunc(app.addNewWorkoutLog))

	//...router -> secureHeaders -> logrequest -> recoverPanic -> client
	return standardMiddleware.Then(router)
}
