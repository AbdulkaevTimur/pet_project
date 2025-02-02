package main

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
