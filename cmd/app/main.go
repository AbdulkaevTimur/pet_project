package main

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/handlers"
	"awesomeProject/internal/taskService"
	"awesomeProject/internal/userService"
	"awesomeProject/internal/web/tasks"
	"awesomeProject/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()

	tasksRepo := taskService.NewTaskRepository(database.DB)
	usersRepo := userService.NewUserRepository(database.DB)

	tasksService := taskService.NewTaskService(tasksRepo)
	userService := userService.NewUserService(usersRepo)

	tasksHandler := handlers.NewTaskHandler(tasksService)
	usersHandler := handlers.NewUserHandler(userService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictTasksHandler := tasks.NewStrictHandler(tasksHandler, nil)
	strictUsersHandler := users.NewStrictHandler(usersHandler, nil)

	tasks.RegisterHandlers(e, strictTasksHandler)
	users.RegisterHandlers(e, strictUsersHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
