package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ZuriDetails struct {
	SlackUserName string  `json:"slackUserName"`
	Backend       bool    `json:"backend"`
	Age           int64 `json:"age"`
	Bio           string  `json:"bio"`
}

// declaring a slice for my details
var myZuriDetail []ZuriDetails


// setting up a function to get my JSON encoded details
func getDetails(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(myZuriDetail)
}


func main()  {
	// instantiating the routers with gorilla mux
	s := mux.NewRouter()

	// populating the myZuriDetail Slice with my Details
	myZuriDetail = append(myZuriDetail, 
		ZuriDetails{
			SlackUserName: "ClintElix", 
			Backend: true, 
			Age: 25, 
			Bio: "Backend Dev, Proficient with Golang, Rust, Nodejs and web3 Solidity"},
	)

	// setting the handler function...
	s.HandleFunc("/", getDetails).Methods("GET")


	http.Handle("/", s)
	// running the Server...
	fmt.Printf("Starting Server...")
	log.Fatal(http.ListenAndServe(":80", s))
}