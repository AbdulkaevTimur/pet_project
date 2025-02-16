package userService

import "gorm.io/gorm"

type UsersRepository interface {
	CreateUser(task User) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUserByID(id uint, task User) (User, error)
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(name User) (User, error) {
	result := r.db.Create(&name)
	if result.Error != nil {
		return User{}, result.Error
	}
	return name, nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var names []User
	err := r.db.Find(&names).Error
	return names, err
}

func (r *userRepository) UpdateUserByID(id uint, name User) (User, error) {
	result := r.db.Model(&name).Where("id = ?", id).Updates(map[string]interface{}{
		"email":    name.Email,
		"password": name.Password,
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
	result := r.db.Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
