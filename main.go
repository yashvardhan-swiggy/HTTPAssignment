package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	Name    string `json:"name"`
	EmailId string `json:"emailId"`
}

func main() {
	fmt.Println("Starting Server...")
	http.HandleFunc("/", RequestHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userDetailsSent := Data{
		Name:    "Yash Vardhan Thakur",
		EmailId: "yash.vardhan@swiggy.in",
	}
	userDetailsReceived := Data{}
	switch r.Method {
	case "GET":
		err := json.NewEncoder(w).Encode(userDetailsSent)
		if err != nil {
			fmt.Errorf("error in encoding data")
		}
	case "POST":
		err := json.NewDecoder(r.Body).Decode(&userDetailsReceived)
		if err != nil {
			fmt.Errorf("error in decoding data")
		}
		fmt.Fprintf(w, "User Details Received : %v\n", userDetailsReceived)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
