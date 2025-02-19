package handlers

import (
	"awesomeProject/internal/userService"
	"awesomeProject/internal/web/users"
	"context"
)

type UserHandler struct {
	UserService *userService.UserService
}

func NewUserHandler(userService *userService.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.UserService.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}
	return response, nil
}

func (h *UserHandler) GetUsersUserIdTasks(c context.Context, request users.GetUsersUserIdTasksRequestObject) (users.GetUsersUserIdTasksResponseObject, error) {
	userId := request.UserId
	tasksForUser, err := h.UserService.GetTasksForUser(userId)
	if err != nil {
		return nil, err
	}

	response := users.GetUsersUserIdTasks200JSONResponse{}

	for _, tsk := range tasksForUser {
		task := users.Task{
			Id:     tsk.Id,
			Task:   tsk.Task,
			IsDone: tsk.IsDone,
			UserId: tsk.UserId,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *UserHandler) PostUsers(c context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	userToCreate := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	createdUser, err := h.UserService.CreateUser(userToCreate)

	if err != nil {
		return nil, err
	}
	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}
	return response, nil
}

func (h *UserHandler) PatchUsersUserId(ctx context.Context, request users.PatchUsersUserIdRequestObject) (users.PatchUsersUserIdResponseObject, error) {
	userRequest := request.Body
	userId := request.UserId
	userToUpdate := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	updatedUser, err := h.UserService.UpdateUserByID(userId, userToUpdate)
	if err != nil {
		return nil, err
	}
	response := users.PatchUsersUserId200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}
	return response, nil
}

func (h *UserHandler) DeleteUsersUserId(ctx context.Context, request users.DeleteUsersUserIdRequestObject) (users.DeleteUsersUserIdResponseObject, error) {
	userId := request.UserId
	err := h.UserService.DeleteUserByID(userId)
	if err != nil {
		return nil, err
	} else {
		response := users.DeleteUsersUserId204Response{}
		return response, nil
	}
}
