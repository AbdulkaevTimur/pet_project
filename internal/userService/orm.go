package userService

import (
	"awesomeProject/internal/web/tasks"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string       `json:"email"`
	Password string       `json:"password"`
	Tasks    []tasks.Task `json:"tasks"`
}
