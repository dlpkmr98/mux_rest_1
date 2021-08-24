package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Message struct {
	Msg string
}

func helloWorldJSON(w http.ResponseWriter, r *http.Request) {
	m := Message{"Hello World"}
	response, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
func main() {
	fmt.Println("*******Application Started")
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloWorldJSON).Methods("GET")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
