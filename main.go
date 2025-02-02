package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func getMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var messages []Message
		DB.Find(&messages)
		json.Marshal(messages)
		fmt.Fprintln(w, messages)
	} else {
		fmt.Fprintln(w, "Поддерживается только метод GET")
	}
}

func createMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		data := Message{}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintln(w, "Ошибка чтения данных")
			return
		}
		json.Unmarshal(body, &data)
		DB.Create(&data)
		fmt.Fprintln(w, data)
	} else {
		fmt.Fprintln(w, "Поддерживается только метод POST")
	}
}

func main() {
	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", createMessage).Methods("POST")
	router.HandleFunc("/api/messages", getMessages).Methods("GET")
	http.ListenAndServe(":8080", router)
}
