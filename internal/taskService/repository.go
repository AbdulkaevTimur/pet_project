package taskService

import "gorm.io/gorm"

type MessageRepository interface {
	CreateTask(task Message) (Message, error)
	GetAllTasks() ([]Message, error)
	UpdateTaskByID(id uint, task Message) (Message, error)
	DeleteTaskByID(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Message) (Message, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Message, error) {
	var tasks []Message
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, task Message) (Message, error) {
	result := r.db.Model(&task).Where("id = ?", id).Updates(map[string]interface{}{
		"task":    task.Task,
		"is_done": task.IsDone,
	})
	if result.Error != nil {
		return Message{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	result := r.db.Delete(&Message{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
