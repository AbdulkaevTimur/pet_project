package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var task, message string

type requestBody struct {
	Message string `json:"message"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, message)
	} else {
		fmt.Fprintln(w, "Поддерживается только метод GET")
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		data := requestBody{}
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&data)
		task = data.Message
		message = "hello, " + task
	} else {
		fmt.Fprintln(w, "Поддерживается только метод POST")
	}
}

func main() {
	http.HandleFunc("/get", GetHandler)
	http.HandleFunc("/post", PostHandler)
	http.ListenAndServe("localhost:8080", nil)
}
