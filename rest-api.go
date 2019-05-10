package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test", Test)
	router.HandleFunc("/hello/{name}", Greeting)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Test(w http.ResponseWriter, r *http.Request) {
	response := ExampleResponse{Color: "yellow", Message: "This is a Test", Notify: "false", MessageFormat: "text"}
	json.NewEncoder(w).Encode(response)
}

func Greeting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	response := ExampleResponse{Color: "blue", Message: "Hello " + name, Notify: "false", MessageFormat: "text"}
	json.NewEncoder(w).Encode(response)
}

type ExampleResponse struct {
	Color         string `json:"color"`
	Message       string `json:"message"`
	Notify        string `json:"notify"`
	MessageFormat string `json:"message_format"`
}