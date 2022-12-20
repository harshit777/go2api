package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	. "github.com/go2api/restaurant"
	"github.com/gorilla/mux"
)

var ResponseData []BuildResponse

func main() {
	ResponseData = FindARestaurant("Pizza", "Pune")
	HandleRoutes()
}

var task = []string{"Some Data", "New Data", "old Data"}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am HomePage")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func findRestaurant(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResponseData)
}

func HandleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/getRestaurant", findRestaurant).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))
}
