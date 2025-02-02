package main

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}
