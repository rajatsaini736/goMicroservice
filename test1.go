package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Roll struct {
	ID			string `json:"id"`
	ImageNumber string `json:"imageNumber"`
	Name		string `json:"name"`
	Ingredients string `json:"ingredients"`
}

// init rolls var as a SLICE
var rolls []Roll

func getRolls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rolls)
}

func getRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range rolls {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newRoll Roll
	json.NewDecoder(r.Body).Decode(&newroll)
	newRoll.ID = strconv.Itoa(len(rolls) + 1)
}

func updateRoll(w http.ResponseWriter, r *http.Request) {

}

func deleteRoll(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// newRoll := Roll{ID: "123", ImageNumber: "imageNum", Name: "Thomos", Ingredients: "Shelby"}
	// rolls = append(rolls, newRoll)
	// fmt.Println(rolls)

	// initialize router
	router := mux.NewRouter()

	//endpoints
	router.HandleFunc("/sushi", getRolls).Methods("GET")
	router.HandleFunc("/sushi/{id}", getRoll).Methods("GET")
	router.HandleFunc("/sushi", createRoll).Methods("POSTw")
	router.HandleFunc("/shushi/{id}", updateRoll).Methods("POST")
	router.HandleFunc("/shushi/{id}", deleteRoll).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))

}