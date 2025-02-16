package userService

type UserService struct {
	repo UsersRepository
}

func NewUserService(repo UsersRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name User) (User, error) {
	return s.repo.CreateUser(name)
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) UpdateUserByID(id uint, name User) (User, error) {
	return s.repo.UpdateUserByID(id, name)
}

func (s *UserService) DeleteUserByID(id uint) error {
	return s.repo.DeleteUserByID(id)
}
