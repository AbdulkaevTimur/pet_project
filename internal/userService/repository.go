package userService

import (
	"awesomeProject/internal/web/tasks"
	"awesomeProject/pkg/utils"
	"gorm.io/gorm"
)

type UsersRepository interface {
	CreateUser(name User) (User, error)
	GetTasksForUser(userID uint) ([]tasks.Task, error)
	GetAllUsers() ([]User, error)
	UpdateUserByID(id uint, name User) (User, error)
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(name User) (User, error) {
	password := name.Password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return User{}, err
	}
	name.Password = hashedPassword
	result := r.db.Create(&name)
	if result.Error != nil {
		return User{}, result.Error
	}
	return name, nil
}

func (r *userRepository) GetTasksForUser(id uint) ([]tasks.Task, error) {
	var tasksForUser []tasks.Task
	err := r.db.Where("user_id = ?", id).Find(&tasksForUser).Error
	return tasksForUser, err
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var names []User
	err := r.db.Find(&names).Error
	return names, err
}

func (r *userRepository) UpdateUserByID(id uint, name User) (User, error) {
	password := name.Password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return User{}, err
	}
	result := r.db.Model(&name).Where("id = ?", id).Updates(map[string]interface{}{
		"email":    name.Email,
		"password": hashedPassword,
	})
	if result.Error != nil {
		return User{}, result.Error
	}
	var updatedUser User
	if err := r.db.First(&updatedUser, id).Error; err != nil {
		return User{}, err
	}

	return updatedUser, nil
}

func (r *userRepository) DeleteUserByID(id uint) error {
	result := r.db.Unscoped().Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
