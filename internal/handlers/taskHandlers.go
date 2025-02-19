package handlers

import (
	"awesomeProject/internal/taskService"
	"awesomeProject/internal/web/tasks"
	"context"
)

type TaskHandler struct {
	TaskService *taskService.TaskService
}

func NewTaskHandler(taskService *taskService.TaskService) *TaskHandler {
	return &TaskHandler{
		TaskService: taskService,
	}
}

func (h *TaskHandler) GetTasksUserId(ctx context.Context, request tasks.GetTasksUserIdRequestObject) (tasks.GetTasksUserIdResponseObject, error) {
	userID := request.UserId
	userTasks, err := h.TaskService.GetTasksByUserID(userID)
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasksUserId200JSONResponse{}

	for _, tsk := range userTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *TaskHandler) PostTasksUserId(c context.Context, request tasks.PostTasksUserIdRequestObject) (tasks.PostTasksUserIdResponseObject, error) {
	taskRequest := request.Body
	userID := request.UserId
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.TaskService.CreateTaskByUserID(taskToCreate, userID)

	if err != nil {
		return nil, err
	}
	response := tasks.PostTasksUserId201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
		UserId: &createdTask.UserID,
	}
	return response, nil
}

func (h *TaskHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskRequest := request.Body
	id := request.Id
	taskToUpdate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	updatedTask, err := h.TaskService.UpdateTaskByID(id, taskToUpdate)
	if err != nil {
		return nil, err
	}
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}
	return response, nil
}

func (h *TaskHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := request.Id
	err := h.TaskService.DeleteTaskByID(id)
	if err != nil {
		return nil, err
	} else {
		response := tasks.DeleteTasksId204Response{}
		return response, nil
	}
}
