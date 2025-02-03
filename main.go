package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

var messages []Message

func GetMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := DB.Find(&messages).Error; err != nil {
			json.NewEncoder(w).Encode(Response{
				Status:  "Error",
				Message: "Could not get message",
			})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(messages)
	} else {
		fmt.Fprintln(w, "Поддерживается только метод GET")
	}
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		message := Message{}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			json.NewEncoder(w).Encode(Response{
				Status:  "Error",
				Message: "Could not create message",
			})
			return
		}
		json.Unmarshal(body, &message)
		DB.Create(&message)
		json.NewEncoder(w).Encode(message)
	} else {
		fmt.Fprintln(w, "Поддерживается только метод POST")
	}
}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPatch {
		var updatedMessage Message
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		idParam := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idParam)
		if err != nil {
			json.NewEncoder(w).Encode(Response{
				Status:  "Error",
				Message: "Bad ID",
			})
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			json.NewEncoder(w).Encode(Response{
				Status:  "Error",
				Message: "Could not read body",
			})
			return
		}
		json.Unmarshal(body, &updatedMessage)
		if err := DB.Model(&Message{}).Where("id = ?", id).Updates(map[string]interface{}{
			"task":    updatedMessage.Task,
			"is_done": updatedMessage.IsDone,
		}).Error; err != nil {
			json.NewEncoder(w).Encode(Response{
				Status:  "Error",
				Message: "Could not update message",
			})
		} else {
			json.NewEncoder(w).Encode(Response{
				Status:  "Success",
				Message: "Message updated",
			})
		}
	}
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		idParam := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idParam)
		if err != nil {
			json.NewEncoder(w).Encode(Response{
				Status:  "Error",
				Message: "Bad ID",
			})
		}
		if err := DB.Delete(&Message{}, id).Error; err != nil {
			json.NewEncoder(w).Encode(Response{
				Status:  "Error",
				Message: "Could not delete message",
			})
		} else {
			json.NewEncoder(w).Encode(Response{
				Status:  "Success",
				Message: "Message deleted",
			})
		}
	}
}

func main() {
	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	router.HandleFunc("/api/messages/{id}", UpdateMessage).Methods("PATCH")
	router.HandleFunc("/api/messages/{id}", DeleteMessage).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
