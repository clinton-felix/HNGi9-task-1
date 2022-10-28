package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

type ZuriDetails struct {
	SlackUserName string  `json:"slackUserName"`
	Backend       bool    `json:"backend"`
	Age           int64 `json:"age"`
	Bio           string  `json:"bio"`
}

var myZuriDetail ZuriDetails

// setting up a function to get my JSON encoded details
func getDetails(w http.ResponseWriter, r *http.Request)  {
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
		SlackUserName: "ClintElix", 
		Backend: true, 
		Age: 25, 
		Bio: "Backend Dev, Proficient with Golang, Rust, Nodejs and web3 Solidity",
	}
	

	// setting the handler function...
	s.HandleFunc("/", getDetails).Methods("GET")


	// handling cors
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
		AllowedMethods: []string{"GET", "DELETE", "POST", "PUT"},
	})

	// running the Server...
	fmt.Printf("Listening on port %v...", port)
	handler := c.Handler(s)
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":"+port, handler))

}