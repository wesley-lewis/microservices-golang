package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)
	cathandler := http.FileServer(http.Dir("./images"))
	http.Handle("/cat/", http.StripPrefix("/cat/", cathandler))

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "Hello " + request.Name + "!"}

	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}

type validationHandler struct {
	next http.Handler
}

func newValidatorHandler(next http.Handler) validationHandler {
	return validationHandler{next: next}
}

func (h validationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	h.next.ServeHTTP(w, r)
}
