package internal

import (
	"net/http"
	"strconv"
)

type BmrDTO struct {
	Age    int
	Height int
	Weight int
	Gender string
}

func GetBMR(w http.ResponseWriter, r *http.Request) {
	age, err := strconv.Atoi(r.URL.Query().Get("age"))
	if err != nil {
		http.Error(w, "Invalid age", http.StatusBadRequest)
		return
	}
	height, err := strconv.Atoi(r.URL.Query().Get("height"))
	if err != nil {
		http.Error(w, "Invalid height", http.StatusBadRequest)
		return
	}
	weight, err := strconv.Atoi(r.URL.Query().Get("weight"))
	if err != nil {
		http.Error(w, "Invalid weight", http.StatusBadRequest)
		return
	}
	gender := r.URL.Query().Get("gender")
	if gender != "male" && gender != "female" {
		http.Error(w, "Invalid gender. THERE'S ONLY 2 GENDERS!", http.StatusBadRequest)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	userinfo := BmrDTO{
		Age:    age,
		Height: height,
		Weight: weight,
		Gender: gender,
	}
	bmr := calculateBMR(userinfo)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("BMR : " + strconv.FormatFloat(bmr, 'f', -1, 64)))

}

func calculateBMR(userInfo BmrDTO) float64 {
	//Mifflin-St Jeor Equation:
	if userInfo.Gender == "male" {
		bmr := (10 * float64(userInfo.Weight)) + (6.25 * float64(userInfo.Height)) - (5 * float64(userInfo.Age)) + float64(5)
		return bmr
	}
	bmr := (10 * float64(userInfo.Weight)) + (6.25 * float64(userInfo.Height)) - (5 * float64(userInfo.Age)) - float64(161)
	return bmr
}
