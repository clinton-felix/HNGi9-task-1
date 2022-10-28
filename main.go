package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type ZuriDetails struct {
	SlackUsername string  `json:"slackUsername"`
	Backend       bool    `json:"backend"`
	Age           int64 `json:"age"`
	Bio           string  `json:"bio"`
}

var myZuriDetail ZuriDetails

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
 }

// setting up a function to get my JSON encoded details with CORS
func getDetails(w http.ResponseWriter, r *http.Request)  {
	setupCorsResponse(&w, r)
		if (*r).Method == "OPTIONS" {
		   return
		}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(myZuriDetail)
}


func main()  {
	// instantiating the routers with gorilla mux
	s := mux.NewRouter()

	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
	}
	port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

	// populating the myZuriDetail Slice with my Details
	myZuriDetail = ZuriDetails{
		SlackUsername: "ClintElix", 
		Backend: true, 
		Age: 25, 
		Bio: "Backend Dev, Proficient with Golang, Rust, Nodejs and web3 Solidity",
	}

	// setting the handler function...
	s.HandleFunc("/", getDetails).Methods("GET")

	// running the Server...
	fmt.Printf("Listening on port %v...", port)
	log.Fatal(http.ListenAndServe(":"+port, s))
	http.Handle("/", s)

}