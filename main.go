package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func main() {
	http.HandleFunc("/time", timeHandler)
	log.Println("Server started on http://localhost:8795")
	log.Fatal(http.ListenAndServe(":8795", nil))
}

func timeHandler(writer http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC3339)

	response := TimeResponse{
		Time: currentTime,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonResponse)
}
