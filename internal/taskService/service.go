package taskService

type TaskService struct {
	repo MessageRepository
}

func NewTaskService(repo MessageRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTaskByUserID(task Task) (Task, error) {
	return s.repo.CreateTaskByUserID(task)
}

func (s *TaskService) GetTasks() ([]Task, error) {
	return s.repo.GetTasks()
}

func (s *TaskService) UpdateTaskByID(id uint, task Task) (Task, error) {
	return s.repo.UpdateTaskByID(id, task)
}

func (s *TaskService) DeleteTaskByID(id uint) error {
	return s.repo.DeleteTaskByID(id)
}
