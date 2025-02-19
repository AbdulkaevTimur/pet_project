package taskService

import "gorm.io/gorm"

type MessageRepository interface {
	CreateTaskByUserID(task Task) (Task, error)
	GetTasks() ([]Task, error)
	UpdateTaskByID(id uint, task Task) (Task, error)
	DeleteTaskByID(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTaskByUserID(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	result := r.db.Model(&task).Where("id = ?", id).Updates(map[string]interface{}{
		"task":    task.Task,
		"is_done": task.IsDone,
	})
	if result.Error != nil {
		return Task{}, result.Error
	}
	var updatedTask Task
	if err := r.db.First(&updatedTask, id).Error; err != nil {
		return Task{}, err
	}

	return updatedTask, nil
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	result := r.db.Delete(&Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
