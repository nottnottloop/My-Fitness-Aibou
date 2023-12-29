package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", app.Home)

	//user
	router.HandleFunc("/user/signup", app.Signup)

	//Kitchen Features
	router.HandleFunc("/kitchen/bmr", app.GetBMR)
	return router
}
