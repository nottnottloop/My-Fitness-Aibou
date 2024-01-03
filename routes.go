package main

import "net/http"

func (app *application) routes() http.Handler {
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
	return app.recoverPanic(app.logRequest(secureHeaders(router)))
}
