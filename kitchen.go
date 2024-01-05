package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type BmrDTO struct {
	Age    int
	Height int
	Weight int
	Gender string
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
}

func (app *application) GetBMR(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	age, err := strconv.Atoi(r.URL.Query().Get("age"))
	if err != nil || age < 1 {
		if err != nil {
			app.errorLog.Println(err.Error())
		}
		http.Error(w, "Invalid age", http.StatusBadRequest)
		return
	}
	height, err := strconv.Atoi(r.URL.Query().Get("height"))
	if err != nil || height < 1 {
		if err != nil {
			app.errorLog.Println(err.Error())
		}
		http.Error(w, "Invalid height", http.StatusBadRequest)
		return
	}
	weight, err := strconv.Atoi(r.URL.Query().Get("weight"))
	if err != nil || weight < 1 {
		if err != nil {
			app.errorLog.Println(err.Error())
		}
		http.Error(w, "Invalid weight", http.StatusBadRequest)
		return
	}
	gender := r.URL.Query().Get("gender")
	if gender != "Male" && gender != "Female" {
		http.Error(w, "Invalid gender", http.StatusBadRequest)
		return
	}

	userinfo := BmrDTO{
		Age:    age,
		Height: height,
		Weight: weight,
		Gender: gender,
	}
	bmr := calculateBMR(userinfo)
	bmrJson, err := json.Marshal("BMR : " + strconv.FormatFloat(bmr, 'f', -1, 64))
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write(bmrJson)

}

func calculateBMR(userInfo BmrDTO) float64 {
	//Mifflin-St Jeor Equation:
	bmr := (10 * float64(userInfo.Weight)) + (6.25 * float64(userInfo.Height)) - (5 * float64(userInfo.Age))
	if userInfo.Gender == "Male" {
		bmr += float64(5)
		return bmr
	}
	bmr -= float64(161)
	return bmr
}
